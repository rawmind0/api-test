package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/mattn/go-colorable"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

// VERSION gets overridden at build time using -X main.VERSION=$VERSION
var (
	NAME    = "dev"
	VERSION = "dev"
)

func main() {
	log.SetOutput(colorable.NewColorableStdout())

	if err := mainErr(); err != nil {
		log.Fatal(err)
	}
}

func mainErr() error {
	app := cli.NewApp()
	app.Name = "api-test"
	app.Version = VERSION
	app.Usage = "api-test [OPTIONS]"
	app.Author = "Raul Sanchez"
	app.Email = "rawmind@gmail.com"
	app.Action = runAPITest
	app.Before = func(ctx *cli.Context) error {
		if ctx.GlobalBool("debug") {
			log.SetLevel(log.DebugLevel)
		}
		NAME = ctx.GlobalString("name")
		return nil
	}
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug,d",
			Usage: "Debug logging",
		},
		cli.StringFlag{
			Name:  "name",
			Usage: "service name",
			Value: "dev",
		},
		cli.StringFlag{
			Name:  "port",
			Usage: "service port",
			Value: "8080",
		},
	}

	return app.Run(os.Args)
}

func runAPITest(ctx *cli.Context) error {
	fmt.Printf("Starting api-test %s version %s at port %s\n", NAME, VERSION, ctx.GlobalString("port"))

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", apiHandler)
	return http.ListenAndServe(":"+ctx.GlobalString("port"), router)
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"service":     NAME,
		"version":     VERSION,
		"description": fmt.Sprintf("Microservice %s version %s", NAME, VERSION),
	}
	json.NewEncoder(w).Encode(response)
}
