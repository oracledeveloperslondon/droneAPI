// this provides all the output for the drone - directing it to the correct
// output - be that logging on the display UI.
// future enhancement is to make the implementation polymophic
//// TODO: connect to the real Drone Reporter URI
package dronecore

import (
	"log"

	"io"
	//	"io/ioutil"
	"os"
)

const (
	PREPOSTFIX      string = "****"
	BIGPREPOSTFIX   string = "*******"
	PANICPREPOSTFIX string = "!!!!"

	Log           = 0 // use the logging mechanism only - by setting to 0 we default to log
	Display       = 1 // there isn't a real drone available use the REST based UI to display what is happening
	LogAndDisplay = 2
	Silent        = 10
)

var (
	trace   *log.Logger
	info    *log.Logger
	warning *log.Logger
	error   *log.Logger
)

func Init(
	traceHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {

	trace = log.New(traceHandle,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	warning = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func init() {
	//// TODO: make this configuration driven by sys properties
	//Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	Init(os.Stdout, os.Stdout, os.Stdout, os.Stderr)

	//Trace.Println("I have something standard to say")
	//Info.Println("Special Information")
	//Warning.Println("There is something you need to know about")
	//Error.Println("Something has failed")
	trace.Println("Drone Entity Logging initialized")
}

type ReportSvc struct {
	address string
	mode    uint8
}

func (reporter *ReportSvc) ReportSimpleMessage(message string) {

	switch reporter.mode {
	case Display, Log, LogAndDisplay:

		info.Println(PREPOSTFIX + message + PREPOSTFIX)

	case Silent: // do nothing
	}
}

func (reporter *ReportSvc) ReportStructuredMessage(messageId uint, message string) {
	switch reporter.mode {
	case Display, Log, LogAndDisplay:
		info.Println(PREPOSTFIX+"messaged:%v \n message:"+message+"\n"+PREPOSTFIX, messageId)
	case Silent: // do nothing
	}
}

func (reporter *ReportSvc) ReportMovement(yaw int8, roll int8, pitch int8, altitude int8) {
	switch reporter.mode {
	case Display, Log, LogAndDisplay:
		info.Println(BIGPREPOSTFIX+"yaw:%v roll:%v pitch:%v altitude:%v"+BIGPREPOSTFIX, yaw, roll, pitch, altitude)
	case Silent: // do nothing
	}
}

func (reporter *ReportSvc) ReportStatus(status int, statusMap map[int]string) {
	switch reporter.mode {
	case Display, Log, LogAndDisplay:
		info.Println(PREPOSTFIX+"Status:%s (%v)"+PREPOSTFIX, status, statusMap[status])
	case Silent: // do nothing
	}
}

func (reporter *ReportSvc) ReportWarning(message string) {
	switch reporter.mode {
	case Display, Log, LogAndDisplay:
		info.Println(PANICPREPOSTFIX + "messaged: message:" + message + "\n" + PANICPREPOSTFIX)
	case Silent: // do nothing
	}
}

func (reporter *ReportSvc) ReportTrace(message string) {
	switch reporter.mode {
	case Display, Log, LogAndDisplay:
		info.Println(PREPOSTFIX + "messaged: message:" + message + "\n" + PREPOSTFIX)
	case Silent: // do nothing
	}
}

func ReportTrace(message string) {
	trace.Println(PREPOSTFIX + "ERR :" + message + "\n" + PREPOSTFIX)

}

func ReportWarning(message string) {
	warning.Println(PREPOSTFIX + "message:" + message + "\n" + PREPOSTFIX)

}

func ReportError(message string) {
	error.Println(PREPOSTFIX + "ERR message:" + message + "\n" + PREPOSTFIX)

}
