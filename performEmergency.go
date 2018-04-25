package main

import (
	"dronecore"
	"net/http"
	"strconv"
)

// goes through all the drones and tells them in turn to execute their
// emergency operations
func droneAllEmergencyStopHandler(w http.ResponseWriter, request *http.Request) {
	var stopCount int

	Warning.Println("Received EMERGENCY STOP for ALL")
	droneList := drones.GetDroneList()

	for index, drone := range droneList {
		if drone == nil {
			Warning.Println("CANT Perform EMERGENCY STOP for drone :" + strconv.Itoa(index))
		} else {
			drone.EMERGENCYSTOP()
			drones.RemoveDrone(drone.GetName())
			Info.Println("EMERGENCY STOP called")
			stopCount++
		}
	}
	w.Header().Set("Content-Encoding", "text/plain")
	w.Write([]byte("Stopped " + strconv.Itoa(stopCount)))
}

// disregards all lifecycle and related rules and calls the drone to land and halt immediately
func droneEmergencyStopHandler(w http.ResponseWriter, request *http.Request) {
	droneName := dronecore.GetDroneNameFromURI(request)
	Warning.Println("Received EMERGENCY STOP for drone :" + droneName)

	drone := drones.GetDrone(droneName)

	if drone == nil {
		Warning.Println("CANT Perform EMERGENCY STOP for drone :" + droneName)
		w.Header().Set("Content-Encoding", "text/plain")
		w.Write([]byte("Emergency stop for " + droneName + " FAILED"))
	} else {
		drone.EMERGENCYSTOP()
		drones.RemoveDrone(drone.GetName())
		Info.Println("EMERGENCY STOP called")
		w.Header().Set("Content-Encoding", "text/plain")
		w.Write([]byte("Stopped " + droneName))
	}

}
