package dbclient

type (
	// ConnParams struct contains all options necessary to create DB connection string
	ConnParams struct {
		host     string
		port     int
		dbName   string
		user     string
		password string
	}

	// ColumnInfo describes a column in a table: its name and type of the data stored in it
	ColumnInfo struct {
		Name     string
		DataType string
		RegType  string
	}

	// DbClient interface defines a common DB client behaviour
	DbClient interface {
		Close()
		// string argument specifies the name of the table
		// Return value is a slice of slices.
		GetAllRows(string) ([]([]interface{}), error)
		GetColumnInfo(tableName string, schemaName string) ([]ColumnInfo, error)
		GetTables() ([]string, error)
		ListTables() error
		Ping() error
		// Executes query and returns entire query result at once.
		Query(query string) (result [][]interface{}, err error)
		// Executes query which returns a single columns of text array type.
		// Returns entire query result at once.
		// Returns error if:
		//  - query result is not a single column of this type
		//  - RDBM does not implement text array data type. PostgreSQL and
		//    Oracle have support for arrays but MySQL does not.
		//  - query execution fails
		//  - data type conversion fails
		ReadColumnTextArray(query string) (result [][]string, err error)
	}

	// New is a DbClient factory method
	New func(dbConnectionParams ConnParams) (DbClient, error)
)

// Host is a getter method which supports ConnParams immutability.
func (params ConnParams) Host() string {
	return params.host
}

// Port is a getter method which supports ConnParams immutability.
func (params ConnParams) Port() int {
	return params.port
}

// DbName is a getter method which supports ConnParams immutability.
func (params ConnParams) DbName() string {
	return params.dbName
}

// User is a getter method which supports ConnParams immutability.
func (params ConnParams) User() string {
	return params.user
}

// Password is a getter method which supports ConnParams immutability.
func (params ConnParams) Password() string {
	return params.password
}

// NewConnParams creates and returns a new ConnParams instance
func NewConnParams(host string, port int, dbName string, user string, password string) ConnParams {
	return ConnParams{
		host,
		port,
		dbName,
		user,
		password,
	}
}
