package main

import (
	"io"
	//	"io/ioutil"
	"log"
	"os"
)

var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func Init(
	traceHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {

	Trace = log.New(traceHandle,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func init() {
	//Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	Init(os.Stdout, os.Stdout, os.Stdout, os.Stderr)

	//Trace.Println("I have something standard to say")
	//Info.Println("Special Information")
	//Warning.Println("There is something you need to know about")
	//Error.Println("Something has failed")
	Trace.Println("Drone Logging initialized")
}

func IsLogging(logType string) bool {
	//// TODO: implement proper switch
	return false
}
