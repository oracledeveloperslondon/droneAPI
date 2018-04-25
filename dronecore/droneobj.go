// the drone struct and associated functions in this file provide the  drones
// which wrap the gobot library so we can replace / extend etc when we start
// incorporating the enhanced features
package dronecore

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	bebop "gobot.io/x/gobot/platforms/parrot/bebop"
)

const (
	MaxAngle = 100
	MinAngle = -100

	// drone states
	Inactive  = 0
	Started   = 1
	Stopped   = 2
	Landed    = 10
	Flying    = 11
	Undefined = 99
)

var bebopStatus map[int]string

type Drone struct {
	name        string // although duplicating the driver - means we're tied to having a real drone
	description string
	adaptor     *bebop.Adaptor
	driver      *bebop.Driver
	//robot       *Robot
	reporter *ReportSvc
	live     bool

	status uint8

	pvSafeYaw    int
	nvSafeYaw    int
	currentYaw   int
	pvSafePitch  int
	nvSafePitch  int
	currentPitch int
	pvSafeRoll   int
	nvSafeRoll   int
	currentRoll  int

	pvSafeGaz  int
	nvSafeGaz  int
	currentGaz int

	// https://emissarydrones.com/what-is-roll-pitch-and-yaw
}

func init() {

	//state StatusType = new (StatusType{ 1: ""})
	//ebopStatus =

	bebopStatus = make(map[int]string)
	bebopStatus[Inactive] = "Inactive"
	bebopStatus[Landed] = "Landed"
	bebopStatus[Flying] = "Flying"
	bebopStatus[Started] = "Started"
	bebopStatus[Stopped] = "Stopped"
}

// populates the drone object - will set the drone live flag if the
// adaptor or driver provided are null objects
func (drone *Drone) initialise(adaptor *bebop.Adaptor, driver *bebop.Driver) {
	drone.pvSafeYaw = MaxAngle
	drone.nvSafeYaw = MinAngle
	drone.pvSafePitch = MaxAngle
	drone.nvSafePitch = MinAngle
	drone.pvSafeRoll = MaxAngle
	drone.nvSafeRoll = MinAngle
	drone.pvSafeGaz = MaxAngle
	drone.nvSafeGaz = MinAngle

	drone.status = Inactive

	drone.reporter = new(ReportSvc)

	drone.adaptor = adaptor
	drone.driver = driver
	drone.reporter.ReportSimpleMessage("Drone controls alive " + strconv.FormatBool(adaptor == nil) + " " + strconv.FormatBool(driver == nil))
	if (adaptor == nil) || (driver == nil) {
		drone.reporter.ReportSimpleMessage("Set alive flag for drone ")
		drone.live = false
	} else {
		drone.live = true
		drone.reporter.ReportSimpleMessage("Caught a live drone")
	}

}

func (drone *Drone) GetStatus() uint8 {
	var status uint8

	if drone.live {
		//TODO if the drone is real then interrogate the drone
		status = drone.status
	} else {
		status = drone.status
	}

	return status
}

// sets the drone name - if the provided name is empty after removing whitespace
// then take the name provided by the drone object
func (drone *Drone) SetName(name string) {
	var nameSet bool

	name = strings.TrimSpace(name)
	if len(name) > 0 {
		drone.name = name
		if drone.adaptor != nil {
			drone.adaptor.SetName(name)
		}
		nameSet = true
	}

	// if we dont have a usable name then pull it from the adaptor if we're live
	if (!nameSet) && (drone.live) {
		drone.name = drone.adaptor.Name()
	}
}

func (drone *Drone) GetName() string {
	if drone != nil {
		return drone.name
	} else {
		ReportWarning("Name request on a null object drone")
		return ""
	}
}

func (drone *Drone) GetDescription() string { return drone.description }

// if we've accidentally connected to a drone then shutdown and drop connection
func (drone *Drone) SetToTest() {
	drone.live = false

	if drone.driver != nil {
		drone.driver.Land()
		drone.driver.Stop()
		drone.driver.StopRecording()
	}
	drone.driver = nil
	if drone.adaptor != nil {
		drone.adaptor.Finalize()
	}
	drone.adaptor = nil
}

// links the drone to an output object.
func (drone *Drone) SetDisplay(ReportSvc *ReportSvc) {
	drone.reporter = ReportSvc
}

// sets the maximum values for the different values - expect a positive value 0 - MaxAngle
func (drone *Drone) SetLimits(safeYaw int, safePitch int, safeRoll int, safeGaz int) {
	if (safeYaw <= MaxAngle) && (safeYaw >= 0) {
		drone.pvSafeYaw = safeYaw
		drone.nvSafeYaw = 0 - safeYaw
	} else {
		drone.reporter.ReportWarning("Yaw value out of allow range " + strconv.Itoa(MinAngle) + " - " + strconv.Itoa(MaxAngle))
	}

	if (safeRoll <= MaxAngle) && (safeRoll >= 0) {
		drone.pvSafePitch = safeYaw
		drone.nvSafePitch = 0 - safeYaw
	} else {
		drone.reporter.ReportWarning("Roll value out of allow range " + strconv.Itoa(MinAngle) + " - " + strconv.Itoa(MaxAngle))
	}
	if (safePitch <= MaxAngle) && (safePitch >= 0) {
		drone.pvSafeRoll = safeRoll
		drone.nvSafeRoll = 0 - safeRoll
	} else {
		drone.reporter.ReportWarning("Pitch value out of allow range " + strconv.Itoa(MinAngle) + " - " + strconv.Itoa(MaxAngle))
	}

}

//// TODO: do we need this factory?
type DroneFactory func(name string) Drone

func (drone *Drone) Takeoff() {
	drone.reporter.ReportTrace(drone.name + " received takeoff request")

	if drone.status != Flying {
		// if we have operating mode that means we actually signal the drone ...
		if drone.live {
			drone.driver.TakeOff()
		}
		drone.status = Flying
		drone.reporter.ReportStatus(Flying, bebopStatus)
	} else {
		drone.reporter.ReportWarning("Request to take off being ignored - already flying")
	}
	drone.reporter.ReportStatus(Flying, bebopStatus)

}

func (drone *Drone) Land() {

	if drone.status == Flying {
		if drone.live {
			drone.driver.Land()
		}
		drone.status = Landed
		drone.reporter.ReportStatus(Landed, bebopStatus)
	} else {
		drone.reporter.ReportWarning("Request to land being ignored - as not flying")
	}

}

func (drone *Drone) Start() {
	// if we have operating mode that means we actually signal the drone ...
	if drone.live {

		drone.driver.Start()
	}

	drone.status = Started
	drone.reporter.ReportStatus(Started, bebopStatus)
}

func (drone *Drone) Stop() {
	// if we have operating mode that means we actually signal the drone ...
	if drone.live {
		drone.driver.Stop()
	}
	drone.status = Stopped
	drone.reporter.ReportStatus(Stopped, bebopStatus)
}

func (drone *Drone) SetYaw(yaw int) {
	if (yaw >= drone.nvSafeYaw) && (yaw <= drone.pvSafeYaw) {
		if yaw >= 0 {
			if drone.live {
				drone.driver.Clockwise(int(yaw))
			}
		} else {
			if drone.live {
				drone.driver.CounterClockwise(int(0 - yaw))
			}
		}
		drone.currentYaw = yaw
	} else {
		drone.reporter.ReportWarning("Requested Yaw out of limits " + strconv.Itoa(int(drone.nvSafeYaw)) + "-" + strconv.Itoa(int(drone.pvSafeYaw)) + " ignoring request")
	}

}

func (drone *Drone) SetRoll(roll int) {
	if (roll >= drone.nvSafeRoll) && (roll <= drone.pvSafeRoll) {
		if roll >= 0 {
			if drone.live {
				drone.driver.Left(int(roll))
			}
		} else {
			if drone.live {
				drone.driver.Right(int(0 - roll))
			}
		}
		drone.currentRoll = roll
	} else {
		drone.reporter.ReportWarning("Requested Roll out of limits " + strconv.Itoa(int(drone.nvSafeRoll)) + "-" + strconv.Itoa(int(drone.pvSafeRoll)) + " ignoring request")
	}
}

func (drone *Drone) SetPitch(pitch int) {
	if (pitch >= drone.nvSafePitch) && (pitch <= drone.pvSafePitch) {
		if pitch >= 0 {
			if drone.live {
				drone.driver.Backward(int(pitch))
			}
		} else {
			if drone.live {
				drone.driver.Forward(int(0 - pitch))
			}
		}
		drone.currentPitch = pitch
	} else {
		drone.reporter.ReportWarning("Requested Pitch out of limits " + strconv.Itoa(int(drone.nvSafePitch)) + "-" + strconv.Itoa(int(drone.pvSafePitch)) + " ignoring request")
	}

}

func (drone *Drone) SetGaz(gaz int) {
	if (gaz >= drone.nvSafeGaz) && (gaz <= drone.pvSafeGaz) {
		if gaz >= 0 {
			if drone.live {
				drone.driver.Backward(int(gaz))
			}
		} else {
			if drone.live {
				drone.driver.Forward(int(0 - gaz))
			}
		}
		drone.currentGaz = gaz
	} else {
		drone.reporter.ReportWarning("Requested Gaz out of limits " + strconv.Itoa(int(drone.nvSafePitch)) + "-" + strconv.Itoa(int(drone.pvSafePitch)) + " ignoring request")
	}

}

func CreateMockDrone() *Drone {
	drone := new(Drone)
	drone.initialise(nil, nil)
	drone.description = "Dummy drone"

	//// TODO: make the drone name pseudo random e.g. grab milliseconds from epoch
	drone.SetName("dummyDrone")
	return drone
}

// creates a new drone object - currently we're allowing the framework to
// default so it can only handle 1 drone
//// TODO: revise to support multiple drones
// enhance to perform networtk scanning for drones not already identified
func FindNewDrone() *Drone {
	ReportTrace("Looking for the drone")
	var drone Drone
	var dronePtr *Drone
	dronePtr = &drone

	// create the bebopAdaptor and establish a connection
	bebopAdaptor := bebop.NewAdaptor()
	driver := bebop.NewDriver(bebopAdaptor)

	dronePtr.initialise(bebopAdaptor, driver)

	drone.reporter.ReportTrace("Is live:" + strconv.FormatBool(drone.live))

	if drone.live {
		drone.name = driver.Name()

		if len(drone.name) > 5 {
			drone.name = "realDrone"
			//// TODO: fix this
		}
		drone.description = driver.Name()

		err := bebopAdaptor.Connect()

		//driver.Subscribe()
		//// TODO: Set event listener up
		//// TODO: setup hull HullProtection
		//// TODO: setup indoor/outdoor mode
		driver.Start()

		if err != nil {
			drone.reporter.ReportWarning("Error Starting up switching live off " + fmt.Sprint(err))
			drone.live = false
			driver.Stop()
			// TODO: improve this message
		}
	}
	return dronePtr
}

// the drone can take actions such as take off and land, this retrieves the value and validates the action
// by linking this to the drone we can overload the validation for other drones
func (drone *Drone) GetAction(request *http.Request) string {
	vars := mux.Vars(request)
	var thisAction string = vars[action]
	var isValid bool = false
	thisAction = strings.ToLower(thisAction)

	for _, act := range validActions {
		if act == thisAction {
			isValid = true
			ReportTrace("validated action " + thisAction)
		}
	}
	if !isValid {
		thisAction = ""
	}
	return thisAction
}

func (drone *Drone) EMERGENCYSTOP() {
	if drone != nil {
		ReportWarning("Emergency Stop requested for " + drone.GetName() + " disregarding everything issueing land command")

		if drone.live {
			// we dealing with a real drone here
			drone.driver.Land()
			err := drone.driver.Halt()

			if err != nil {
				ReportError("EMERGENCYSTOP error during process : " + fmt.Sprint(err))
			}
		}
		drone.currentGaz = 0
		drone.currentYaw = 0
		drone.currentRoll = 0
		drone.currentPitch = 0
		drone.status = Stopped

	} else {
		ReportError("Emergency Stop requested for NULL drone object - CANT ACTION")
	}
}
