package pgclientdemo

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/BojanKomazec/go-demo/internal/pkg/config"
	"github.com/BojanKomazec/go-demo/internal/pkg/dbclient"
	"github.com/BojanKomazec/go-demo/internal/pkg/pgclient"
)

func readFromDBDemo(conf *config.Config) error {
	dbConnParams := dbclient.NewConnParams(
		conf.DB.ConnParams().Host(),
		conf.DB.ConnParams().Port(),
		conf.DB.ConnParams().DbName(),
		conf.DB.ConnParams().User(),
		conf.DB.ConnParams().Password())

	dbClient, err := pgclient.New(dbConnParams)
	if err != nil {
		return err
	}

	defer dbClient.Close()

	err = dbClient.Ping()
	if err != nil {
		return err
	}

	fmt.Println("Successfully connected!")

	tableNames, err := dbClient.GetTables()
	if err != nil {
		return err
	}

	fmt.Println("Table names:", tableNames)

	columnInfo, err := dbClient.GetColumnInfo(tableNames[0], "public")
	if err != nil {
		return err
	}

	fmt.Println("Column info:", columnInfo)

	allRows, err := dbClient.GetAllRows(tableNames[0])
	if err != nil {
		return err
	}

	fmt.Println("All rows:", allRows)

	for i := range allRows {
		fmt.Println("\n\nRow #", i, allRows[i])

		// use this approach to discover data type for each column
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
				} else if columnInfo[j].RegType == "json[]" {
					stringifiedJSONArrValue := string(allRows[i][j].([]uint8))
					fmt.Printf("%s\n", stringifiedJSONArrValue)
				} else {
					fmt.Printf("%v\n", allRows[i][j])
				}
			default:
				fmt.Println("Type not supported")
			}
		}

		// if we know in advance type (and/or name) of each column we can manually map column value to data
		// name, err := getString(row[1]) // row[1] is of type TEXT; function returns (string, error)
		// if err != nil {
		// 	log.Println(err)
		// 	continue
		// }

		// phones, err := getStrings(row[2]) // row[2] is of type TEXT[]; function returns ([]string, error)
		// if err != nil {
		// 	log.Println(err)
		// 	continue
		// }

		// magicNumbers, err := getSIntegers(row[3]) // row[3] is of type integer[]; function returns ([]integer, error)
		// if err != nil {
		// 	log.Println(err)
		// 	continue
		// }

		// attachments, err := getResources(row[4]) // row[4] is of type json[]; function returns ([]resource, error) where resource is struct that matches json
		// if err != nil {
		// 	log.Println(err)
		// 	continue
		// }

		// (!) approach above is very brittle - as soon as db schema changes, this code will break!
		// It is better to use some tool which can automatically map columns to variables.
	}

	return nil
}

func writeAndReadDemo(conf *config.Config) error {
	dbConnParams := dbclient.NewConnParams(
		conf.DB.ConnParams().Host(),
		conf.DB.ConnParams().Port(),
		conf.DB.ConnParams().DbName(),
		conf.DB.ConnParams().User(),
		conf.DB.ConnParams().Password())

	dbClient, err := pgclient.New(dbConnParams)
	if err != nil {
		return err
	}

	defer dbClient.Close()

	err = dbClient.Ping()
	if err != nil {
		return err
	}

	fmt.Println("Successfully connected!")

	// @todo
	// write an array of structs into DB
	// read table rows
	// extract data from each row

	return nil
}

// ShowDemo func
func ShowDemo(conf *config.Config) error {
	fmt.Println("pgclientdemo.ShowDemo()")
	err := readFromDBDemo(conf)
	return err
}
