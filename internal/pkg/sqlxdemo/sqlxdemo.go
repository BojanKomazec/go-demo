package sqlxdemo

import (
	"fmt"
	"log"

	"github.com/BojanKomazec/go-demo/internal/pkg/config"
	"github.com/jmoiron/sqlx"

	// Pure Go Postgres driver for database/sql
	"github.com/lib/pq"
	// _ "github.com/lib/pq"
)

// Resource struct
type Resource struct {
	URL  string
	Desc string
}

// Contact struct
// pg array types have to be used for variables into which array data is to be read from Postgres DB:
//    pq.StringArray instead of []string
//    pq.Int64Array instead of []int
type Contact struct {
	ID           int
	Name         string
	Phones       pq.StringArray
	MagicNumbers pq.Int64Array  `db:"magic_numbers"`
	Resources    pq.StringArray // []Resource
}

// ShowDemo function
func ShowDemo(conf *config.Config) error {
	db, err := sqlx.Connect(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			conf.DB.ConnParams().Host(),
			conf.DB.ConnParams().Port(),
			conf.DB.ConnParams().User(),
			conf.DB.ConnParams().Password(),
			conf.DB.ConnParams().DbName()))
	if err != nil {
		log.Fatalln(err)
		return err
	}

	log.Println("connected")

	defer func() error {
		err = db.Close()
		if err != nil {
			log.Println(err)
			return err
		}
		return nil
	}()

	rows, err := db.Queryx(`
SELECT id, name, phones, magic_numbers, resources
FROM contacts`)

	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		var c Contact
		err = rows.StructScan(&c)
		log.Printf("\n\n%v %v %v %v %v\n", c.ID, c.Name, c.Phones, c.MagicNumbers, c.Resources)
		for _, phone := range c.Phones {
			log.Printf("%v\n", phone)
		}
		for _, number := range c.MagicNumbers {
			log.Printf("%v\n", number)
		}
		for _, resource := range c.Resources {
			log.Printf("%v\n", resource)

			// @todo here convert string to json data struct
		}
	}

	return nil
}
