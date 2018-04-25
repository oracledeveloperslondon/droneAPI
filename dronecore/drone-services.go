package dronecore

// this part of the dronecore package handles the list of drones

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const (
	droneIdLabel = "droneId"

	action = "action"
)

var NODRONE *Drone
var validActions []string

// defines the storage structure
type DroneService struct{ drones (map[string]*Drone) }

func init() {
	validActions = append(validActions, TAKEOFF)
	validActions = append(validActions, LAND)

}

func InitialiseDroneService() (instance *DroneService) {
	if instance == nil {
		instance = new(DroneService)
	}
	instance.drones = make(map[string]*Drone)

	return instance
}

// retuirns an array of drone references from the DroneService
func (service *DroneService) GetDroneList() []*Drone {
	list := make([]*Drone, 0, len(service.drones))

	for _, value := range service.drones {
		list = append(list, value)
	}

	return list
}

// add the details of a drone to the the list of drones
func (service *DroneService) RegisterDrone(droneName string, drone *Drone) bool {

	var registered bool = false

	if service == nil {
		service = InitialiseDroneService()
		ReportTrace("Creating drone service  object")
	}

	if service.drones[droneName] == nil {
		service.drones[droneName] = drone
		registered = true
	}

	return registered
}

func (service *DroneService) GetDrone(droneName string) *Drone {
	var result *Drone
	if (service == nil) || (len(droneName) == 0) {
		ReportError("Requested drone " + droneName + " but NO droneList")
		result = NODRONE
	} else {
		result = service.drones[droneName]
	}
	return result
}

// removes the drone from our own list of drones
func (service *DroneService) RemoveDrone(droneName string) {
	var drone *Drone

	drone = (service.drones)[droneName]

	if drone != nil {
		drone.Land()
		(service.drones)[droneName] = nil
	}
}

func (service *DroneService) LogStatus() {
	var drone *Drone

	drone = (service.drones)["droneName"]

	if drone != nil {
		drone.reporter.ReportSimpleMessage("status check")
		(service.drones)["droneName"] = nil
	}
}

func GetDroneNameFromURI(request *http.Request) string {
	vars := mux.Vars(request)
	droneName := vars[droneIdLabel]

	ReportTrace("GetDroneNameFromURI returning:" + droneName)
	return (droneName)

}

func GetNavValue(request *http.Request, navType string) int {
	vars := mux.Vars(request)
	val, err := strconv.Atoi(vars[navType])

	if err == nil {
		if (val <= MaxAngle) && (val >= MinAngle) {
			return val
		}
	}
	return 0
}
