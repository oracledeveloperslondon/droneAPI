package main

import (
	"dronecore"
	"net/http"
	"strconv"
	//"github.com/gorilla/mux"
	//"dronecore"
	
)

// This handles the Roll request, utilizing dronecore to facilitate then
// call interpreetation and validation
func droneSimpleNavRollHandler(w http.ResponseWriter, r *http.Request) {
	droneName := dronecore.GetDroneNameFromURI(r)
	roll := dronecore.GetNavValue(r, dronecore.ROLL)
	drone := drones.GetDrone(droneName)

	if drone != nil {
	
		drone.SetRoll(roll)
		
		go sendMessage("0", strconv.Itoa(roll), "0")

		msg := "Roll received " + strconv.Itoa(roll) + " for " + droneName
		w.WriteHeader(http.StatusOK)
		w.Header().Set(ENCODING, TEXTENC)
		w.Write([]byte(msg))
		Trace.Println(msg)
	} else {
		msg := "Roll received " + strconv.Itoa(roll) + " for " + droneName + " No drone to command"
		w.Header().Set(ENCODING, TEXTENC)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(msg))
		Warning.Println(msg)
	}
}

func droneSimpleNavYawHandler(w http.ResponseWriter, r *http.Request) {
	droneName := dronecore.GetDroneNameFromURI(r)
	yaw := dronecore.GetNavValue(r, dronecore.YAW)
	drone := drones.GetDrone(droneName)

	if drone != nil {
	
		drone.SetYaw(yaw)
		
		go sendMessage(strconv.Itoa(yaw), "0", "0")

		msg := "Yaw received " + strconv.Itoa(yaw) + " for " + droneName
		w.WriteHeader(http.StatusOK)
		w.Header().Set(ENCODING, TEXTENC)
		w.Write([]byte(msg))
		
		Trace.Println(msg)
	} else {
		msg := "Yaw received " + strconv.Itoa(yaw) + " for " + droneName + " No drone to command"
		w.Header().Set(ENCODING, TEXTENC)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(msg))
		Warning.Println(msg)
	}
}

func droneSimpleNavGazHandler(w http.ResponseWriter, r *http.Request) {
	droneName := dronecore.GetDroneNameFromURI(r)
	gaz := dronecore.GetNavValue(r, dronecore.GAZ)
	drone := drones.GetDrone(droneName)

	if drone != nil {
	
		drone.SetGaz(gaz)
		
		go sendMessage("0", "0", strconv.Itoa(gaz))

		msg := "Gaz received " + strconv.Itoa(gaz) + " for " + droneName
		w.WriteHeader(http.StatusOK)
		w.Header().Set(ENCODING, TEXTENC)
		w.Write([]byte(msg))
		Trace.Println(msg)
	} else {
		msg := "Gaz received " + strconv.Itoa(gaz) + " for " + droneName + " No drone to command"
		w.Header().Set(ENCODING, TEXTENC)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(msg))
		Warning.Println(msg)
	}

}

func droneSimpleNavPitchHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
