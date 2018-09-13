// the code in this package should not embody any understanding of the drones
// the purpose here is to receive the REST calls and then use the dronecore
//layer to command the drone(s)
package main

import (
	"dronecore"
	"encoding/json"
	"fmt"
	"net/http"
)

type DroneRegistrationResponseType struct {
	DroneId          string
	DroneDescription string
}

type RegistrationResponseMsgType struct {
	//	drones []droneRegistrationResponseType
	//// TODO: make the sresponse structure an array
	Drones DroneRegistrationResponseType
}

// this function haandles the registration of the drone
func droneRegistrationHandler(response http.ResponseWriter, request *http.Request) {
	Trace.Println("RegistrationHandler Called")
	var newDrone = dronecore.FindNewDrone()
	myResponse := new(RegistrationResponseMsgType)
	var newElement DroneRegistrationResponseType

	if newDrone == nil {
		newDrone = dronecore.CreateMockDrone()
		Info.Println("Creating a mock drone")
	}

	if newDrone != nil {
		Trace.Println("setting limits and response object")

		//// TODO: replace these hardcoded values
		newDrone.SetLimits(360, 360, 360, 2)
		drones.RegisterDrone(newDrone.GetName(), newDrone)

		//// TODO: at a future point we need to handle multiple drones here
		//for idx, el := range drones.GetDroneList() {

		newElement.DroneId = newDrone.GetName()
		newElement.DroneDescription = newDrone.GetDescription()
		//// TODO: next line needs to become an array handler in due course
		myResponse.Drones = newElement
	} else {
		Error.Panicln("Could NOT create a drone")
	}

	Trace.Println("about to respond with:\n" + fmt.Sprintf("%v", myResponse))

	body, err := json.Marshal(myResponse)

	if err == nil {
		response.WriteHeader(http.StatusOK)
		response.Write(body)
	} else {
		response.WriteHeader(http.StatusInternalServerError)
		Error.Panicln(fmt.Sprintf("registrationResponseMsgType - error when marshalling %v", err))
	}

}
