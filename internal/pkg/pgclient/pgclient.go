package pgclient

import (
	"database/sql"
	"fmt"
	"log"

	// Pure Go Postgres driver for database/sql
	"github.com/lib/pq"

	"github.com/BojanKomazec/go-demo/internal/pkg/dbclient"
)

// pgClient structure represents a PostgreSQL client
type pgClient struct {
	driverName     string
	dataSourceName string
	db             *sql.DB
}

// New function creates an instance of PostgreSQL client
func New(params dbclient.ConnParams) (dbclient.DbClient, error) {
	return newImpl(params)
}

// Close function closes DB connetion
func (pgClient pgClient) Close() {
	pgClient.db.Close()
}

// If a DB column has type (udt_name) text[] (data_type is ARRAY) then
// when Rows.Scan() populates matching destination variable, this variable
// will have Go type []uint8 (which is essentially byte array). Each byte
// is decimal ASCII code so entire variable can be converted into a string.
// Postgres defines that this string comes in form:
//    {element1,element2,...elementN}
// (there are no spaces between curly braces or between elements)
func (pgClient pgClient) GetAllRows(tableName string) ([]([]interface{}), error) {
	query := fmt.Sprintf("SELECT * FROM %s;", tableName)
	rows, err := pgClient.db.Query(query)
	if err != nil {
		return nil, err
	}

	columnsNames, err := rows.Columns()
	columnsCount := len(columnsNames)
	fmt.Println("columns count =", columnsCount)
	fmt.Println("columns names =", columnsNames)

	allRowsValues := make([]([]interface{}), 0)

	defer rows.Close()
	for rows.Next() {
		rowValues := make([]interface{}, columnsCount)
		rowValuesPtrs := make([]interface{}, columnsCount)
		for i := range rowValuesPtrs {
			rowValuesPtrs[i] = &rowValues[i]
		}
		err := rows.Scan(rowValuesPtrs...)
		if err != nil {
			return nil, err
		}

		allRowsValues = append(allRowsValues, rowValues)
	}

	return allRowsValues, nil
}

func (pgClient pgClient) GetColumnInfo(tableName string, schemaName string) ([]dbclient.ColumnInfo, error) {
	query := fmt.Sprintf(
		"SELECT column_name, data_type, udt_name::regtype "+
			"FROM information_schema.columns "+
			"WHERE table_name = '%s' AND table_schema = '%s';", tableName, schemaName)
	fmt.Println("Query =", query)
	rows, err := pgClient.db.Query(query)
	if err != nil {
		fmt.Println("db.Query failed")
		return nil, err
	}

	columnsInfo := make([]dbclient.ColumnInfo, 0)
	defer rows.Close()

	for rows.Next() {
		var colInf dbclient.ColumnInfo
		err := rows.Scan(&colInf.Name, &colInf.DataType, &colInf.RegType)
		if err != nil {
			fmt.Println("rows.Scan failed")
			return nil, err
		}

		columnsInfo = append(columnsInfo, colInf)
	}

	return columnsInfo, nil
}

func (pgClient pgClient) GetTables() ([]string, error) {
	query := "SELECT table_name FROM information_schema.tables WHERE table_schema='public' ORDER BY table_name;"
	rows, err := pgClient.db.Query(query)
	if err != nil {
		return nil, err
	}

	tableNames := make([]string, 0)
	var tableName string
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&tableName)
		if err != nil {
			return nil, err
		}
		tableNames = append(tableNames, tableName)
	}

	err = rows.Err()
	return tableNames, err
}

func (pgClient pgClient) ListTables() error {
	query := "SELECT table_name FROM information_schema.tables WHERE table_schema='public' ORDER BY table_name;"
	rows, err := pgClient.db.Query(query)
	if err != nil {
		return err
	}

	var tableName string
	var logOutput string
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&tableName)
		if err != nil {
			return err
		}
		logOutput += tableName + "\n"
	}

	log.Printf("Database tables:\n%s\n", logOutput)

	err = rows.Err()
	return err
}

// Ping function verifies DB connection
func (pgClient pgClient) Ping() error {
	err := pgClient.db.Ping()
	return err
}

// Function which executes an arbitrary SQL query and returns raw, untyped data.
//
// If a DB column has type (udt_name) text[] (data_type is ARRAY) then
// when Rows.Scan() populates matching destination variable, this variable
// will have Go type []uint8 (which is essentially byte array). Each byte
// is decimal ASCII code so entire variable can be converted into a string.
// Postgres defines that this string comes in form:
//    {element1,element2,...elementN}
// (there are no spaces between curly braces or between elements)
func (pgClient pgClient) Query(query string) ([][]interface{}, error) {
	rows, err := pgClient.db.Query(query)
	if err != nil {
		return nil, err
	}

	columnsNames, err := rows.Columns()
	columnsCount := len(columnsNames)

	allRowsValues := make([]([]interface{}), 0)

	defer rows.Close()
	for rows.Next() {
		rowValues := make([]interface{}, columnsCount)
		rowValuesPtrs := make([]interface{}, columnsCount)
		for i := range rowValuesPtrs {
			rowValuesPtrs[i] = &rowValues[i]
		}
		err := rows.Scan(rowValuesPtrs...)
		if err != nil {
			return nil, err
		}

		allRowsValues = append(allRowsValues, rowValues)
	}

	return allRowsValues, nil
}

func (pgClient pgClient) ReadColumnTextArray(query string) (result [][]string, err error) {
	rows, err := pgClient.db.Query(query)
	if err != nil {
		return nil, err
	}

	allRowsValues := make([][]string, 0)

	defer rows.Close()
	for rows.Next() {
		rowValues := make([]string, 0)
		err := rows.Scan(pq.Array(&rowValues))
		if err != nil {
			return nil, err
		}

		allRowsValues = append(allRowsValues, rowValues)
	}

	return allRowsValues, nil
}

// todo: evaluate if it's necessary to parameterize sslmode value
func newImpl(dbConnectionParams dbclient.ConnParams) (pgClient, error) {
	pgClient := pgClient{
		driverName: "postgres",
		dataSourceName: fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			dbConnectionParams.Host(), dbConnectionParams.Port(), dbConnectionParams.User(),
			dbConnectionParams.Password(), dbConnectionParams.DbName()),
	}
	var err error
	pgClient.db, err = sql.Open(pgClient.driverName, pgClient.dataSourceName)
	return pgClient, err
}
