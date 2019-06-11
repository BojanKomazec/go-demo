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
		Ping() error
		GetTables() ([]string, error)
		GetColumnInfo(tableName string, schemaName string) ([]ColumnInfo, error)
		// string argument specifies the name of the table
		// Return value is a slice of slices.
		GetAllRows(string) ([]([]interface{}), error)
		Close()
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
