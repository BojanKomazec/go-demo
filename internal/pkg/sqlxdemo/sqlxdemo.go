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
// fields must be exported (capitalized) in order for sqlx to be able to write into them
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

// ContactSimple struct
type ContactSimple struct {
	ID    int
	Name  string
	Phone string `db:"phones"` // phones[n] in query will return a single value here
}

// ContactSimpleAlias struct
type ContactSimpleAlias struct {
	ID    int
	Name  string
	Phone string `db:"main_phone"` // main_phone is an alias for phones[1] where "phones" is the column name
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

	log.Printf("\n\nDemo using placeholder parameters:")
	// It is not possible to substitute identifier names (like table name), only values.
	// source: https://www.postgresql.org/docs/9.1/xfunc-sql.html

	var c Contact

	// note that $1, $2, ...notation is used (at least for Postgres)
	err = db.QueryRowx(`
SELECT *
FROM contacts
WHERE id = $1`, 1).StructScan(&c)
	if err != nil {
		return err
	}
	log.Printf("\n\n%v %v %v %v %v\n", c.ID, c.Name, c.Phones, c.MagicNumbers, c.Resources)

	log.Printf("\n\nDemo selecting single element from column which is an array (placeholder is int):")

	var contactSimple ContactSimple
	err = db.QueryRowx(`
SELECT id, name, phones[1]
FROM contacts
WHERE id = $1`, 1).StructScan(&contactSimple)

	if err != nil {
		return err
	}

	log.Printf("\n\n%v %v %v\n", contactSimple.ID, contactSimple.Name, contactSimple.Phone)

	log.Printf("\n\nDemo selecting single element from column which is an array (placeholder is string):")

	err = db.QueryRowx(`
SELECT id, name, phones[1]
FROM contacts
WHERE name = $1`, "Jeremy Clarckson").StructScan(&contactSimple)

	if err != nil {
		return err
	}

	log.Printf("\n\n%v %v %v\n", contactSimple.ID, contactSimple.Name, contactSimple.Phone)

	log.Printf("\n\nDemo selecting column with alias name in the query:")

	var contactSimpleAlias ContactSimpleAlias
	err = db.QueryRowx(`
SELECT id, name, phones[1] as main_phone
FROM contacts
WHERE id = $1`, 1).StructScan(&contactSimpleAlias)
	if err != nil {
		return err
	}
	log.Printf("\n\n%v %v %v\n", contactSimpleAlias.ID, contactSimpleAlias.Name, contactSimpleAlias.Phone)

	return nil
}
