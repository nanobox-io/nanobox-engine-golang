package main

import (
	"flag"
	"os"
)

var host string
var key string
var id string
var bucketid string
var objectid string
var needHelp bool
var public bool

func init() {
	flag.BoolVar(&needHelp, "h", false, "Show Help")
	flag.BoolVar(&needHelp, "help", false, "Show Help")
	flag.BoolVar(&public, "public", false, "Public object on create")
	flag.StringVar(&host, "location", "", "Server host and port in the format of 'host:port'")
	flag.StringVar(&host, "host", "", "Server host and port in the format of 'host:port'")
	flag.StringVar(&host, "server", "", "Server host and port in the format of 'host:port'")
	flag.StringVar(&key, "key", "", "Access key for your user")
	flag.StringVar(&key, "user-key", "", "Access key for your user")
	flag.StringVar(&id, "id", "", "Access id for your user")
	flag.StringVar(&id, "user-id", "", "Access id for your user")
	flag.StringVar(&bucketid, "bucketid", "", "The UUID of the bucket you intend to use")
	flag.StringVar(&bucketid, "bucketname", "", "The Name of the bucket you intend to use")
	flag.StringVar(&objectid, "objectid", "", "The UUID of the object you intend to use")
	flag.StringVar(&objectid, "objectalias", "", "The Alias of the object you intend to use")

	flag.Parse()
	if needHelp {
		help()
		os.Exit(0)
	}
	getHost()
	userKey()
	userId()
	bucketId()
	objectId()
}

func getHost() {
	val := host
	if val == "" {
		val = os.Getenv("HOST")
	}
	if val == "" {
		val = os.Getenv("LOCATION")
	}
	if val == "" {
		val = os.Getenv("SERVER")
	}
	host = val
}

func userKey() {
	val := key
	if val == "" {
		val = os.Getenv("USERKEY")
	}
	if val == "" {
		val = os.Getenv("KEY")
	}
	key = val
}

func userId() {
	val := id
	if val == "" {
		val = os.Getenv("USERID")
	}
	if val == "" {
		val = os.Getenv("ID")
	}
	id = val
}

func bucketId() {
	val := bucketid
	if val == "" {
		val = os.Getenv("BUCKETID")
	}
	if val == "" {
		val = os.Getenv("BUCKETNAME")
	}
	bucketid = val
}

func objectId() {
	val := objectid
	if val == "" {
		val = os.Getenv("OBJECTID")
	}
	if val == "" {
		val = os.Getenv("OBJECTALIAS")
	}
	objectid = val
}
