package main

import (
	"context"
	"flag"
	"golang-echo-layout/bootstrap"
)

var (
	env      string
	showHelp bool
)

func init() {
	flag.StringVar(&env, "env", "dev", "environment for server:[dev|prod]")
	flag.BoolVar(&showHelp, "h", false, "show help")
	flag.Parse()
}

func main() {
	if showHelp {
		flag.PrintDefaults()
		return
	}

	ctx := context.WithValue(context.Background(), "env", env)
	bootstrap.StartServer(ctx)

}
