package main

import (
	"log"

	"github.com/BojanKomazec/go-demo/internal/pkg/config"
	"github.com/BojanKomazec/go-demo/internal/pkg/cryptodemo"
	"github.com/BojanKomazec/go-demo/internal/pkg/datatypesdemo"
	"github.com/BojanKomazec/go-demo/internal/pkg/errordemo"
	"github.com/BojanKomazec/go-demo/internal/pkg/function"
	"github.com/BojanKomazec/go-demo/internal/pkg/httpdemo"
	"github.com/BojanKomazec/go-demo/internal/pkg/iodemo"
	"github.com/BojanKomazec/go-demo/internal/pkg/jsondemo"
	"github.com/BojanKomazec/go-demo/internal/pkg/mapdemo"
	"github.com/BojanKomazec/go-demo/internal/pkg/onerr"
	"github.com/BojanKomazec/go-demo/internal/pkg/osdemo"
	"github.com/BojanKomazec/go-demo/internal/pkg/pathdemo"
	"github.com/BojanKomazec/go-demo/internal/pkg/randdemo"
	"github.com/BojanKomazec/go-demo/internal/pkg/regexdemo"
	"github.com/BojanKomazec/go-demo/internal/pkg/runtimedemo"
	"github.com/BojanKomazec/go-demo/internal/pkg/stringdemo"
	"github.com/BojanKomazec/go-demo/internal/pkg/structdemo"
	"github.com/BojanKomazec/go-demo/internal/pkg/types"
	"github.com/joho/godotenv"
	// "github.com/BojanKomazec/go-demo/internal/pkg/bufiodemo"
	// "github.com/BojanKomazec/go-demo/internal/pkg/fmtdemo"
	// "github.com/BojanKomazec/go-demo/internal/pkg/goroutinedemo"
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

	// fmtdemo.ReadIntegersFromLine()
	// bufiodemo.ReadIntegersLineDemo()
	cryptodemo.ShowDemo()
	datatypesdemo.ShowDemo()
	// goroutinedemo.ShowDemo()
	errordemo.ShowDemo()
	function.ShowDemo()
	httpdemo.ShowDemo(conf.OutputDir)
	iodemo.ShowDemo()
	jsondemo.ShowDemo()
	mapdemo.ShowDemo()
	osdemo.ShowDemo()
	pathdemo.ShowDemo()
	randdemo.ShowDemo()
	regexdemo.ShowDemo()
	runtimedemo.ShowDemo()
	stringdemo.ShowDemo()
	structdemo.ShowDemo()
	types.EnumDemo()
	types.IotaDemo()
}
