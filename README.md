# Background
The Drone API project has come about as a result of wanting to offer REST APIs that can control a drone. By exposing the means to control the drone via REST it becomes possible for people for people to build solutions with just about any technology of choice. From simple scripts using CURL to clever web apps.

The development of the platform wraps and extends the [Gobot](https://gobot.io/) framework so eventually the framework could be utilized to provide APIs on any of the Gobot supported drones. Today the focus is entirely on the use of the [Parrot Bebop 2 drone](https://www.parrot.com/uk/drones/parrot-bebop-2)

# How We Have Developed the APIs
An [API First approasch](http://www.api-first.com/) - as a result the API Blueprint has been developed and is periodically committed to this GitHub repo. By adopting an API 1st approach it means that anyone can start building against the APIs whilst we build out the layer around Gobot

The [API Blueprint](https://app.apiary.io/dronedevmeetup/) can be seen [here](https://app.apiary.io/dronedevmeetup/).

## Implementation Details

The API layer is obviously also written in [Go](http://golang.org).  Go is a well supported and mature language with strong typing like Java, but has characteristics of C by being explicit about the use of pointers, there is no class inheritance, data structures are defined using struc definitions and methods can be part of an interface, as a result polymorphic behaviour is easy to achieve. Go also offers the more relaxed approach to the need for semicolons, type definitions you see in [Groovy](http://groovy-lang.org). 

Go also offers the portability of Java. So this functionality can be run on any platform supporting Go which includes Windows, Linux OS X.

The implementation completely wraps and abstracts the Gobot framework so no understanding of Gobot is required. Masny of the URIs are very simple and can be realized using a browser.

The Web server / URI handling has been implemented using Gorrilla Mux.

## A Simulation Mode
In addition to driving a real drone. It is possible to also cooperate the backend with a simulation display.  Each of the action calls are forwarded onto a JET Application called drone-dash which will display messages of what the drone is being asked to do, and provide a visual representation.  Drone-dash is available [here](https://github.com/oracledeveloperslondon/dronedash).

# Getting Started

To get started and use this framework you need to:
  1. Install Go - instructions [here](https://golang.org/doc/install)
  2. Retrieve Gorilla/Mux - instructions [here](http://www.gorillatoolkit.org)
  2b. Retrieve Gorilla/Websocket - [here](https://github.com/gorilla/websocket)
  3. Download this library [here](https://gobot.io/documentation/getting-started/)
  4. Compile everything using the [Go build command](https://golang.org/pkg/go/build/)
  5. Start up the server - this is down the filesystem naming, but the preceding steps should give a binary executable

# How the Code Hangs Together
Whilst we have worked hard to ensure the code is self documenting the following provides details on which files do what so you can start to traverse the code, or add to it without compromising the abstractions and code structure ...

## Packages
there are two packages, **main** which contains the REST handling layer and **dronecore** which provides the wrapping layer.

### main
   - **webhandler.go** this includes the declarations for the Gorrilla Multiplexer (MUX) providing the declarations for the handlers. The URIs defined are presented in a manner that makes it as close as possible to the Blueprint representation. additionally the URIs are declared with both a trailing **/** as it impacts the MUX ability to handle the call. We have also included alias paths for example __yaw__ and __rotation__
   - **performActions.go** The handler functions have been grouped together in the main tier. All the handler files incorporate the task of taking the REST calls and translate the web call to the relevant calls in the **dronecore** layer. __performactions__ takes the actions like takeoff, land etc. The ohers are __performEmergency__ which takes an Emergency Stop which then goes through all the known drones and lands and stops them regardless of what is happening.  __PerformRegistrations__ takes the registration call and locates drone(s) or creates dummy ones and registers them  __performSiompleNavigations__ contains the change yaw, gaz etc
   - **dronelogging** makes use of the basic logging framework to handle simple Log4J style logging and exposes 4 objects Trace, Info, Warning, Fatal -- these entities have no linkage to the __DroneDash__ feature
   
## dronecore
  - **droneobj.go** This provides the means to wrap the __gobot__ library and provides the primary struc which holds, a number of key elements relating to a drone, this includes the __driver__ and __connection__ object, plus a handle to a __reporter__ object which provides the means to record what the drone is being asked to perform
  - **drone-services.go** This provides the handler for non drone instance specific logic, such as translating a URI to get the name of the drone and locating the cached drone object.
  - **droneReport.go**
  - **droneactions.go**

# API LifeCycle

To control the drone it first has to be registered through the RegisterDrone API call - then any command can be sent, but IF the drone is in the wrong state a **BadRequest** response will be provided.  APIs not yet implemented will return **NotImplemented**
