package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"golang-labbaika_gaji-fasthttp/handler/routing"
	"golang-labbaika_gaji-fasthttp/model"
	"golang-labbaika_gaji-fasthttp/util"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

func main(){
	config := flag.String("config", "config.json", "config is used to load config file")
	if *config == "" {
		*config = "config.json"
	}
	configFile, err := os.Open(*config)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	properties := loadProperties(configFile)
	util.SetupLogging(properties.LogPath)
	timeOut := parseTimeOut(properties.TimeOut)

	//init
	fmt.Println(strings.ToUpper(properties.ServiceName), "SERVICE")
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal(err.Error())
	}
	wg := sync.WaitGroup{}
	wg.Add(1)
	runApp := routing.InitRouting(properties, timeOut)
	go runApp.Routing()
	wg.Wait()
}

func loadProperties(configFile *os.File) model.Properties {
	properties := model.Properties{}
	fmt.Println("Init configuration....")
	jsonParser := json.NewDecoder(configFile)
	if err := jsonParser.Decode(&properties); err != nil {
		fmt.Println("Error init configuration !!")
		os.Exit(1)
	}
	return properties
}

func parseTimeOut(sTimeout string) time.Duration {
	timeout, err := time.ParseDuration(sTimeout)
	if err != nil {
		fmt.Println("Error while parsing timeout")
		os.Exit(1)
	}
	return timeout
}
