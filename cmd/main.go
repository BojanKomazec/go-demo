package main

import (
	"log"

	"github.com/BojanKomazec/go-demo/internal/pkg/config"
	"github.com/BojanKomazec/go-demo/internal/pkg/cryptodemo"
	"github.com/BojanKomazec/go-demo/internal/pkg/datatypesdemo"
	"github.com/BojanKomazec/go-demo/internal/pkg/goroutinedemo"
	"github.com/BojanKomazec/go-demo/internal/pkg/httpdemo"
	"github.com/BojanKomazec/go-demo/internal/pkg/jsondemo"
	"github.com/BojanKomazec/go-demo/internal/pkg/mapdemo"
	"github.com/BojanKomazec/go-demo/internal/pkg/onerr"
	"github.com/BojanKomazec/go-demo/internal/pkg/osdemo"
	"github.com/BojanKomazec/go-demo/internal/pkg/randdemo"
	"github.com/BojanKomazec/go-demo/internal/pkg/regexdemo"
	"github.com/BojanKomazec/go-demo/internal/pkg/runtimedemo"
	"github.com/BojanKomazec/go-demo/internal/pkg/stringdemo"
	"github.com/BojanKomazec/go-demo/internal/pkg/types"
	"github.com/joho/godotenv"
	// "github.com/BojanKomazec/go-demo/internal/pkg/bufiodemo"
	// "github.com/BojanKomazec/go-demo/internal/pkg/fmtdemo"
	// "github.com/BojanKomazec/go-demo/internal/pkg/function"
)

func init() {
	err := godotenv.Load()
	if err != nil { // not critical
		log.Println("WARNING: loading .env file failed")
	}
}

func main() {
	go runtimedemo.GoRoutineCountBackgroundMonitor()

	conf, err := config.New()
	if err != nil {
		onerr.Panic(err)
	}

	// err = pgclientdemo.ShowDemo(conf)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// function.VariadicFunction(1, 'a', true, "bcdef")
	// fmtdemo.ReadIntegersFromLine()
	// bufiodemo.ReadIntegersLineDemo()
	cryptodemo.ShowDemo()
	datatypesdemo.ShowDemo()
	goroutinedemo.ShowDemo()
	httpdemo.ShowDemo(conf.OutputDir)
	jsondemo.ShowDemo()
	mapdemo.ShowDemo()
	osdemo.ShowDemo()
	randdemo.ShowDemo()
	regexdemo.ShowDemo()
	runtimedemo.ShowDemo()
	stringdemo.ShowDemo()
	types.EnumDemo()
	types.IotaDemo()
}
