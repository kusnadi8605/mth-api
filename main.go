package main

import (
	"flag"
	"fmt"
	"log"
	conf "mth-api/config"
	hdlr "mth-api/handler"
	mdw "mth-api/middleware"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// open mysql connection
	configFile := flag.String("conf", "config/config.yml", "main configuration file")
	flag.Parse()
	conf.LoadConfigFromFile(configFile)

	logDate := time.Now().Format("20060102")
	conf.SetFilename(conf.Param.Log.FileName + logDate + ".txt")

	//redis connection
	conf.RedisDbInit(conf.Param.RedisURL)

	// mysqlConnection
	conn, err := conf.NewConn(conf.Param.DBType, conf.Param.DBUrl)
	conf.Logf("Load Database Conf: %s ", conf.Param.DBType)
	conf.Logf("running App on port: %s ", conf.Param.ListenPort)

	if err != nil {
		conf.Logf("Load Database Conf: %s ", err)
		log.Fatal(err)
	}

	// handler
	http.HandleFunc("/api/token", mdw.Chain(hdlr.TokenHandler(conn), mdw.ContentType("application/json"), mdw.Method("POST")))
	http.HandleFunc("/api/create", mdw.Chain(hdlr.CreateHandler(conn), mdw.ContentType("application/json"), mdw.Method("POST"), mdw.CheckToken()))
	http.HandleFunc("/api/update", mdw.Chain(hdlr.UpdateHandler(conn), mdw.ContentType("application/json"), mdw.Method("POST"), mdw.CheckToken()))
	http.HandleFunc("/api/detail", mdw.Chain(hdlr.DetailHandler(conn), mdw.ContentType("application/json"), mdw.Method("POST"), mdw.CheckToken()))
	http.HandleFunc("/api/delete", mdw.Chain(hdlr.DeleteHandler(conn), mdw.ContentType("application/json"), mdw.Method("POST"), mdw.CheckToken()))
	http.HandleFunc("/api/list", mdw.Chain(hdlr.ListHandler(conn), mdw.ContentType("application/json"), mdw.Method("POST"), mdw.CheckToken()))

	var errors error
	errors = http.ListenAndServe(conf.Param.ListenPort, nil)

	if errors != nil {
		fmt.Println("error", errors)
		conf.Logf("Unable to start the server: %s ", conf.Param.ListenPort)
		os.Exit(1)
	}
}
