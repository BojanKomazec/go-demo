package pgclient

import (
	"database/sql"
	"fmt"

	// Pure Go Postgres driver for database/sql
	_ "github.com/lib/pq"

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

// Ping function verifies DB connection
func (pgClient pgClient) Ping() error {
	err := pgClient.db.Ping()
	return err
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

// Close function closes DB connetion
func (pgClient pgClient) Close() {
	pgClient.db.Close()
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
