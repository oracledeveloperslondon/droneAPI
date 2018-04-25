package main

import (
	"dronecore"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

const (
	SVR_ADDR = "localhost:80"
	ENCODING = "Content-Encoding"
	TEXTENC  = "text/plain"
)

var drones *dronecore.DroneService

//var srv *http.Server

func init() {
	drones = dronecore.InitialiseDroneService()
}

func catchAllHandler(w http.ResponseWriter, r *http.Request) {
	Info.Println("Default handler Call " + fmt.Sprint(r.Header))
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "-- your body -- %s!", r.URL.Path[1:])
}

// walk through the multiplexer configuration and display the relevant details
func walkMux(r *mux.Router) {
	err := r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		pathTemplate, err := route.GetPathTemplate()
		if err == nil {
			fmt.Println("ROUTE:", pathTemplate)
		}
		pathRegexp, err := route.GetPathRegexp()
		if err == nil {
			fmt.Println("Path regexp:", pathRegexp)
		}
		queriesTemplates, err := route.GetQueriesTemplates()
		if err == nil {
			fmt.Println("Queries templates:", strings.Join(queriesTemplates, ","))
		}
		queriesRegexps, err := route.GetQueriesRegexp()
		if err == nil {
			fmt.Println("Queries regexps:", strings.Join(queriesRegexps, ","))
		}
		methods, err := route.GetMethods()
		if err == nil {
			fmt.Println("Methods:", strings.Join(methods, ","))
		}
		fmt.Println()
		return nil
	})

	if err != nil {
		Error.Panicln(err)
	}
}

func main() {
	Trace.Println("Starting ...")

	// going to use Gorrilla MUX - can be obtained by go get -u github.com/gorilla/mux
	// http://www.gorillatoolkit.org/pkg/mux
	r := mux.NewRouter()

	r.HandleFunc("/drone/STOP", droneAllEmergencyStopHandler)
	r.HandleFunc("/drone/STOP/", droneAllEmergencyStopHandler)

	r.HandleFunc("/drone/{droneId}/STOP", droneEmergencyStopHandler)
	r.HandleFunc("/drone/{droneId}/STOP/", droneEmergencyStopHandler)

	r.HandleFunc("/drone/register", droneRegistrationHandler)
	r.HandleFunc("/drone/register/", droneRegistrationHandler)

	r.HandleFunc("/drone/{droneId}/nav/action/{action}", droneActionHandler)
	r.HandleFunc("/drone/{droneId}/nav/action/{action}/", droneActionHandler)

	r.HandleFunc("/drone/{droneId}/simplenav/pitch/{pitch}", droneSimpleNavPitchHandler)
	r.HandleFunc("/drone/{droneId}/simplenav/pitch/{pitch}/", droneSimpleNavPitchHandler)

	r.HandleFunc("/drone/{droneId}/simplenav/altitude/{altitude}", droneSimpleNavGazHandler)
	r.HandleFunc("/drone/{droneId}/simplenav/altitude/{altitude}/", droneSimpleNavGazHandler)

	r.HandleFunc("/drone/{droneId}/simplenav/gaz/{gaz}", droneSimpleNavGazHandler)
	r.HandleFunc("/drone/{droneId}/simplenav/gaz/{gaz}/", droneSimpleNavGazHandler)

	r.HandleFunc("/drone/{droneId}/simplenav/roll/{roll}", droneSimpleNavRollHandler)
	r.HandleFunc("/drone/{droneId}/simplenav/roll/{roll}/", droneSimpleNavRollHandler)

	r.HandleFunc("/drone/{droneId}/simplenav/rotation/{yaw}", droneSimpleNavYawHandler)
	r.HandleFunc("/drone/{droneId}/simplenav/rotation/{yaw}/", droneSimpleNavYawHandler)
	r.HandleFunc("/drone/{droneId}/simplenav/yaw/{yaw}", droneSimpleNavYawHandler)
	r.HandleFunc("/drone/{droneId}/simplenav/yaw/{yaw}/", droneSimpleNavYawHandler)

	r.HandleFunc("/", catchAllHandler)
	r.HandleFunc("/drone/", catchAllHandler)

	Trace.Println("Multiplexer configured. starting server....")

	if IsLogging("Trace") {
		walkMux(r)
	}

	log.Fatal(http.ListenAndServe(SVR_ADDR, r))

	Trace.Println("The End")
}
