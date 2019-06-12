package pgclientdemo

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

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
				// if DB column type is string array then convert
				// this value into a string and parse it into a
				// string array
				if columnInfo[j].RegType == "text[]" {
					if byteArr, ok := allRows[i][j].([]byte); ok {
						columnStringifiedValue := string(byteArr)
						columnStringifiedValue = strings.TrimPrefix(columnStringifiedValue, "{")
						columnStringifiedValue = strings.TrimSuffix(columnStringifiedValue, "}")
						stringElements := strings.Split(columnStringifiedValue, ",")

						fmt.Printf("%s\n", string(byteArr))
						fmt.Println("Array elements (strings):", stringElements)
					} else {
						fmt.Println("Conversion failed")
					}
				} else if columnInfo[j].RegType == "integer[]" {
					if byteArr, ok := allRows[i][j].([]byte); ok {
						columnStringifiedValue := string(byteArr)
						columnStringifiedValue = strings.TrimPrefix(columnStringifiedValue, "{")
						columnStringifiedValue = strings.TrimSuffix(columnStringifiedValue, "}")
						stringifiedNumbers := strings.Split(columnStringifiedValue, ",")
						numbers := make([]int, 0)
						for _, v := range stringifiedNumbers {
							number, err := strconv.Atoi(v)
							if err == nil {
								numbers = append(numbers, number)
							} else {
								fmt.Println("Failed to convert string to number")
							}
						}
						fmt.Printf("%s\n", columnStringifiedValue)
						fmt.Println("Array elements (integers):", numbers)
					} else {
						fmt.Println("Conversion to []byte failed")
					}
				} else {

					fmt.Printf("%v\n", allRows[i][j])
				}
			default:
				fmt.Println("Type not supported")
			}
		}
	}
}
