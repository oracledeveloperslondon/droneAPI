package main

import (
	"dronecore"
	"net/http"
)

// locates the drone and commands it to perform specific simple axtions such as takeoff and land
func droneActionHandler(w http.ResponseWriter, request *http.Request) {
	status := http.StatusOK
	myDrone := drones.GetDrone(dronecore.GetDroneNameFromURI(request))
	action := myDrone.GetAction(request)

	if myDrone != nil {
		Trace.Println("droneActionHandler drone is " + myDrone.GetName() + " action is " + action)

		switch action {
		case dronecore.TAKEOFF:
			if myDrone.GetStatus() != dronecore.Flying {
				go sendAction(action)
				myDrone.Takeoff()
			} else {
				status = http.StatusBadRequest
			}
		case dronecore.LAND:
			if myDrone.GetStatus() == dronecore.Flying {
				go sendAction(action)
				myDrone.Land()
			} else {
				status = http.StatusBadRequest
			}

		case dronecore.HOVER:
			if myDrone.GetStatus() == dronecore.Flying {
				myDrone.SetGaz(0)
				myDrone.SetRoll(0)
				myDrone.SetPitch(0)
				myDrone.SetYaw(0)
			} else {
				status = http.StatusBadRequest
			}

			//// TODO: Implement other axtion types
			//		case "cutpower",
			//			"return",
			//			"startFollowing", "stopFollowing":
			//			Warning.Println("Request not implemented " + action)
			//			status = http.StatusNotImplemented
		default:
			Warning.Println("Unknown action requested " + action)
			status = http.StatusNotImplemented
		}
	} else {
		Warning.Println("No drone identified so cant process request")
		status = http.StatusBadRequest
	}

	w.WriteHeader(status)
}
