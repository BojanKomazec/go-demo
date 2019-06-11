package pgclientdemo

import (
	"fmt"
	"reflect"

	"github.com/BojanKomazec/go-demo/internal/pkg/config"
	"github.com/BojanKomazec/go-demo/internal/pkg/dbclient"
	"github.com/BojanKomazec/go-demo/internal/pkg/onerr"
	"github.com/BojanKomazec/go-demo/internal/pkg/pgclient"
)

// ShowDemo func
func ShowDemo(conf *config.Config) {
	fmt.Println("pgclientdemo.ShowDemo()")

	dbConnParams := dbclient.NewConnParams(
		conf.DB.ConnParams().Host(),
		conf.DB.ConnParams().Port(),
		conf.DB.ConnParams().DbName(),
		conf.DB.ConnParams().User(),
		conf.DB.ConnParams().Password())

	dbClient, err := pgclient.New(dbConnParams)
	onerr.Panic(err)

	defer dbClient.Close()

	err = dbClient.Ping()
	onerr.Panic(err)

	fmt.Println("Successfully connected!")

	tableNames, err := dbClient.GetTables()
	onerr.Panic(err)
	fmt.Println("Table names:", tableNames)

	columnInfo, err := dbClient.GetColumnInfo(tableNames[0], "public")
	onerr.Panic(err)
	fmt.Println("Column info:", columnInfo)

	allRows, err := dbClient.GetAllRows(tableNames[0])
	onerr.Panic(err)
	fmt.Println("All rows:", allRows)

	for i := range allRows {
		fmt.Println("Row #", i, allRows[i])
		for j := range allRows[i] {
			colValType := reflect.TypeOf(allRows[i][j])
			fmt.Printf("Column #%d\n\ttype: %s\n\tvalue: ", j, colValType)

			switch allRows[i][j].(type) {
			case int64:
				fmt.Printf("%d\n", allRows[i][j])
			case string:
				fmt.Printf("%s\n", allRows[i][j])
			case []uint8:
				fmt.Printf("%v\n", allRows[i][j])
			default:
				fmt.Println("Type not supported")
			}
		}
	}
}
