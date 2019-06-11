package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/BojanKomazec/go-demo/internal/pkg/array"
	"github.com/BojanKomazec/go-demo/internal/pkg/pgclientdemo"

	// "github.com/BojanKomazec/go-demo/internal/pkg/bufiodemo"
	// "github.com/BojanKomazec/go-demo/internal/pkg/fmtdemo"
	// "github.com/BojanKomazec/go-demo/internal/pkg/function"
	"github.com/BojanKomazec/go-demo/internal/pkg/config"
	"github.com/BojanKomazec/go-demo/internal/pkg/jsondemo"
	"github.com/BojanKomazec/go-demo/internal/pkg/mapdemo"
	"github.com/BojanKomazec/go-demo/internal/pkg/onerr"
	"github.com/BojanKomazec/go-demo/internal/pkg/randdemo"
	"github.com/BojanKomazec/go-demo/internal/pkg/types"
)

func init() {
	err := godotenv.Load()
	if err != nil { // not critical
		log.Println("WARNING: loading .env file failed")
	}
}

func main() {
	conf, err := config.New()
	onerr.Panic(err)

	pgclientdemo.ShowDemo(conf)
	// function.VariadicFunction(1, 'a', true, "bcdef")
	// fmtdemo.ReadIntegersFromLine()
	// bufiodemo.ReadIntegersLineDemo()
	array.DemoDeclaration()
	types.IotaDemo()
	types.EnumDemo()
	mapdemo.ShowDemo()
	jsondemo.ShowDemo()
	randdemo.ShowDemo()
}
