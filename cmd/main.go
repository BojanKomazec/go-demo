package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/BojanKomazec/go-demo/internal/pkg/config"
	"github.com/BojanKomazec/go-demo/internal/pkg/cryptodemo"
	"github.com/BojanKomazec/go-demo/internal/pkg/datatypesdemo"
	"github.com/BojanKomazec/go-demo/internal/pkg/errordemo"
	"github.com/BojanKomazec/go-demo/internal/pkg/function"
	"github.com/BojanKomazec/go-demo/internal/pkg/htmltemplatedemo"
	"github.com/BojanKomazec/go-demo/internal/pkg/httpdemo"
	"github.com/BojanKomazec/go-demo/internal/pkg/iodemo"
	"github.com/BojanKomazec/go-demo/internal/pkg/jsondemo"
	"github.com/BojanKomazec/go-demo/internal/pkg/kingpindemo"
	"github.com/BojanKomazec/go-demo/internal/pkg/mapdemo"
	"github.com/BojanKomazec/go-demo/internal/pkg/onerr"
	"github.com/BojanKomazec/go-demo/internal/pkg/osdemo"
	"github.com/BojanKomazec/go-demo/internal/pkg/pathdemo"
	"github.com/BojanKomazec/go-demo/internal/pkg/pgclientdemo"
	"github.com/BojanKomazec/go-demo/internal/pkg/randdemo"
	"github.com/BojanKomazec/go-demo/internal/pkg/regexdemo"
	"github.com/BojanKomazec/go-demo/internal/pkg/runtimedemo"
	"github.com/BojanKomazec/go-demo/internal/pkg/sqlxdemo"
	"github.com/BojanKomazec/go-demo/internal/pkg/statementsdemo"
	"github.com/BojanKomazec/go-demo/internal/pkg/stringdemo"
	"github.com/BojanKomazec/go-demo/internal/pkg/structdemo"
	"github.com/BojanKomazec/go-demo/internal/pkg/texttemplatedemo"
	"github.com/BojanKomazec/go-demo/internal/pkg/types"
	"github.com/joho/godotenv"
	// "github.com/BojanKomazec/go-demo/internal/pkg/bufiodemo"
	// "github.com/BojanKomazec/go-demo/internal/pkg/fmtdemo"
	// "github.com/BojanKomazec/go-demo/internal/pkg/goroutinedemo"
)

var (
	printHelp       bool
	runPostgresDemo bool
	verboseMode     bool
	name            string
)

// https://stackoverflow.com/questions/24790175/when-is-the-init-function-run
// init() is guaranteed to run before main() is called.
// https://stackoverflow.com/questions/19761963/flag-command-line-parsing-in-golang
// https://grokbase.com/t/gg/golang-nuts/134rcnkas9/go-nuts-why-not-flag-parse-in-init
func init() {
	flag.BoolVar(&printHelp, "help", false, "print this help")
	flag.BoolVar(&runPostgresDemo, "postgres", false, "[true|false] - run Postgres Client demo (requires PostgresDB running prior to this application)")

	// the following flags are added to match those defined in kingpindemo package
	flag.StringVar(&name, "name", "", "Name of user.")
	flag.BoolVar(&verboseMode, "verbose", false, "Verbose mode")
	flag.BoolVar(&verboseMode, "v", false, "Verbose mode")

	flag.Parse()

	if printHelp {
		flag.PrintDefaults()
		os.Exit(0)
	}

	fmt.Println("runPostgresDemo =", runPostgresDemo)

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

	if runPostgresDemo {
		err = pgclientdemo.ShowDemo(conf)
		if err != nil {
			fmt.Println(err)
		}

		err = sqlxdemo.ShowDemo(conf)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		// fmtdemo.ReadIntegersFromLine()
		// bufiodemo.ReadIntegersLineDemo()
		cryptodemo.ShowDemo()
		datatypesdemo.ShowDemo()
		// goroutinedemo.ShowDemo()
		errordemo.ShowDemo()
		function.ShowDemo()
		htmltemplatedemo.ShowDemo()
		httpdemo.ShowDemo(conf.OutputDir)
		iodemo.ShowDemo()
		jsondemo.ShowDemo()
		kingpindemo.ShowDemo()
		mapdemo.ShowDemo()
		osdemo.ShowDemo()
		pathdemo.ShowDemo()
		randdemo.ShowDemo()
		regexdemo.ShowDemo()
		runtimedemo.ShowDemo()
		statementsdemo.ShowDemo()
		stringdemo.ShowDemo()
		structdemo.ShowDemo()
		texttemplatedemo.ShowDemo()
		types.EnumDemo()
		types.IotaDemo()
	}
}
