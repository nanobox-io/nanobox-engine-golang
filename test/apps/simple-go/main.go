package main

import (
	"./api"
	"./backends"
	"./models"
	_ "net/http/pprof"
	"github.com/jcelliott/lumber"
	"flag"
	"runtime"
)

func main() {

	setLogLevel()

	runtime.GOMAXPROCS(runtime.NumCPU())

	var be models.Storage

	switch selectedBackend {
	case "local":
		be = backends.NewLocalStorage(backendCredentials)
	default:
		be = backends.NewLocalStorage(backendCredentials)
	}

	err := models.Initialize(dbCredentials, be)
	if err != nil {
		panic(err)
	}

	models.CleanEmptyObjects()

	err = api.Start(port)
	lumber.Error(err.Error())

}

func setLogLevel() {
	switch level {
	case "TRACE":
		lumber.Level(0)
	case "DEBUG":
		lumber.Level(1)
	case"INFO":
		lumber.Level(2)
	case"WARN":
		lumber.Level(3)
	case"ERROR":
		lumber.Level(4)
	case"FATAL":
		lumber.Level(5)
	default:
		lumber.Info("the log level provided ("+level+") is not available, defaulting to INFO")
	}
}

var port string
var dbCredentials string
var selectedBackend string
var backendCredentials string
var level string

func init() {
	flag.StringVar(&level, "level", "INFO", "Log level (available options are TRACE, DEBUG, INFO, WARN, ERROR, FATAL")
	flag.StringVar(&port, "port", "8080", "Port to listen on")
	flag.StringVar(&dbCredentials, "dbCredentials", "dbname=live sslmode=disable", "Connection string for database connection")
	flag.StringVar(&selectedBackend, "backend", "local", "Backend data storage")
	flag.StringVar(&backendCredentials, "backendcredentials", "/tmp/data", "Backend data credentials")
	flag.Parse()
}
