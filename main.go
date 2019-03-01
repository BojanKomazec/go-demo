package main

import (
	"github.com/BojanKomazec/go-demo/array"
	// "github.com/BojanKomazec/go-demo/bufiodemo"
	// "github.com/BojanKomazec/go-demo/fmtdemo"
	// "github.com/BojanKomazec/go-demo/function"
	"github.com/BojanKomazec/go-demo/jsondemo"
	"github.com/BojanKomazec/go-demo/mapdemo"
	"github.com/BojanKomazec/go-demo/types"
)

func main() {
	// function.VariadicFunction(1, 'a', true, "bcdef")
	// fmtdemo.ReadIntegersFromLine()
	// bufiodemo.ReadIntegersLineDemo()
	array.DemoDeclaration()
	types.IotaDemo()
	types.EnumDemo()
	mapdemo.ShowDemo()
	jsondemo.ShowDemo()
}
