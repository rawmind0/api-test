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
	NAME     = "dev"
	VERSION  = "dev"
	notFound = fmt.Sprintf("%d", http.StatusNotFound)
	statusOK = fmt.Sprintf("%d", http.StatusOK)
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
			Name:  "name,n",
			Usage: "service name",
			Value: "dev",
		},
		cli.StringFlag{
			Name:  "port,p",
			Usage: "service port",
			Value: "8080",
		},
	}

	return app.Run(os.Args)
}

func runAPITest(ctx *cli.Context) error {
	log.Infof("Starting api-test %s version %s at port %s\n", NAME, VERSION, ctx.GlobalString("port"))

	router := mux.NewRouter().StrictSlash(true)
	router.NotFoundHandler = handle404()
	router.HandleFunc("/", apiHandler)
	return http.ListenAndServe(":"+ctx.GlobalString("port"), router)
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	log.Debugf("%s-%s-%s-%s%s-%s", statusOK, r.Method, r.Proto, r.Host, r.RequestURI, r.RemoteAddr)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusOK)
	response := map[string]string{
		"name":    NAME,
		"version": VERSION,
		"message": fmt.Sprintf("Microservice %s version %s", NAME, VERSION),
		"status":  statusOK,
	}
	json.NewEncoder(w).Encode(response)
}

func handle404() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Debugf("%s-%s-%s-%s%s-%s", notFound, r.Method, r.Proto, r.Host, r.RequestURI, r.RemoteAddr)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(http.StatusNotFound)
		response := map[string]string{
			"name":    NAME,
			"version": VERSION,
			"message": http.StatusText(http.StatusNotFound),
			"status":  notFound,
		}
		json.NewEncoder(w).Encode(response)
	})
}
