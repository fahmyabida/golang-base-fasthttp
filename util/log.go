package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/jeanphorn/log4go"
)

type config struct {
	Console console `json:"console"`
	Files   []files `json:"files"`
}

type console struct {
	Enable bool `json:"enable"`
}

type files struct {
	Enable   bool   `json:"enable"`
	Level    string `json:"level"`
	Filename string `json:"filename"`
	Category string `json:"category"`
	Pattern  string `json:"pattern"`
	Rotate   bool   `json:"rotate"`
	Maxsize  string `json:"maxsize"`
	Maxlines string `json:"maxlines"`
	Daily    bool   `json:"daily"`
	Sanitize bool   `json:"sanitize"`
}

//Error is public function to create logging
func Error(text ...string) {
	log.error(text...)
}

//Event is public function to create logging
func Event(text ...string) {
	log.event(text...)
}

//Message is public function to create logging
func Message(text ...string) {
	log.message(text...)

}

//Fatal is public function to create logging
func Fatal(text ...string) {
	log.fatal(text...)

}

var (
	log iLog
)

// ILog is used to create interface logging
type iLog interface {
	error(text ...string)
	event(text ...string)
	message(text ...string)
	fatal(text ...string)
}

//Logger is used to create object logging
type logger struct {
	errorLog   *log4go.Filter
	eventLog   *log4go.Filter
	messageLog *log4go.Filter
}

// Error is use to write error log
func (l *logger) error(text ...string) {
	logText := strings.Join(text, "][")
	timenow := "[" + time.Now().Format("2006-01-02 15:04:05.000") + "]"
	l.errorLog.Error(timenow + logText)
}

// Event is use to write event log
func (l *logger) event(text ...string) {
	logText := strings.Join(text, "][")
	timenow := "[" + time.Now().Format("2006-01-02 15:04:05.000") + "]"
	l.eventLog.Info(timenow + logText)

}

// Message is use to write Message log
func (l *logger) message(text ...string) {
	logText := "[" + strings.Join(text, "][") + "]"
	timenow := "[" + time.Now().Format("2006-01-02 15:04:05.000") + "]"
	l.messageLog.Info(timenow + logText)
}

// Fatal is use to write Message log
func (l *logger) fatal(text ...string) {
	log := "[" + strings.Join(text, "][") + "]"
	timenow := "[" + time.Now().Format("2006-01-02 15:04:05.000") + "]"
	l.errorLog.Info(timenow + log)
	time.Sleep(3 * time.Second)
	os.Exit(1)
}

// SetupLogging is used to set up logging system
func SetupLogging(path string) {

	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, 0775)
	}

	if _, err := os.Stat(path + "/error"); os.IsNotExist(err) {
		os.MkdirAll(path+"/error", 0775)
	}

	if _, err := os.Stat(path + "/event"); os.IsNotExist(err) {
		os.MkdirAll(path+"/event", 0775)
	}

	if _, err := os.Stat(path + "/message"); os.IsNotExist(err) {
		os.MkdirAll(path+"/message", 0775)
	}

	data := config{
		Console: console{
			Enable: false,
		},
		Files: []files{
			files{
				Enable:   true,
				Level:    "INFO",
				Filename: path + "/error/error.log",
				Category: "error",
				Pattern:  "%M",
				Rotate:   true,
				Maxsize:  "500M",
				Maxlines: "10K",
				Daily:    true,
				Sanitize: true,
			},
			files{
				Enable:   true,
				Level:    "INFO",
				Filename: path + "/event/event.log",
				Category: "event",
				Pattern:  "%M",
				Rotate:   true,
				Maxsize:  "500M",
				Maxlines: "10K",
				Daily:    true,
				Sanitize: true,
			},
			files{
				Enable:   true,
				Level:    "INFO",
				Filename: path + "/message/message.log",
				Category: "message",
				Pattern:  "%M",
				Rotate:   true,
				Maxsize:  "500M",
				Maxlines: "10K",
				Daily:    true,
				Sanitize: true,
			},
		},
	}

	ConfigFileLog, errorParsingConfig := json.Marshal(data)
	if errorParsingConfig != nil {
		fmt.Println(errorParsingConfig.Error())
		os.Exit(2)
	}
	_, errorCheckExistence := os.Stat("config-log.json")
	if os.IsNotExist(errorCheckExistence) {
		ioutil.WriteFile("config-log.json", ConfigFileLog, 0775)
	}

	log4go.LoadConfiguration("config-log.json")
	log = &logger{
		errorLog:   log4go.LOGGER("error"),
		eventLog:   log4go.LOGGER("event"),
		messageLog: log4go.LOGGER("message"),
	}

}

