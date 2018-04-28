# Background
The Drone API project has come about as a result of wanting to offer REST APIs that can control a drone. By exposing the means to control the drone via REST it becomes possible for people for people to build solutions with just about any technology of choice. From simple scripts using CURL to clever web apps.

The developement of the platform wraps and extends the [Gobot](https://gobot.io/) framework so eventually the framework could be utilized to provide APIs on any of the Gobot supported drones. Today the focus is entirely on the use of the [Parrot Bebop 2 drone](https://www.parrot.com/uk/drones/parrot-bebop-2)

# How We Have Developed the APIs
An [API First approasch](http://www.api-first.com/) - as a result the API Blueprint has been developed and is periodically committed to this GitHub repo. By adopting an API 1st approach it means that anyonew can start building against the APIs whilst we build out the layer around Gobot

The API Blueprint can be seen [here](https://app.apiary.io/dronedevmeetup/).

## Implementation Details

The API layer is obviously also written in [Go](http://golang.org).  Go is a well supported and mature language with strong typinging Like Java, but has characteristics of C by being explicit about the use of pointers, there is no class inheritance, data stratuctures are defined using struc definitions and methods can be part of an interface, as a result polymorphic behaviour is easy to achieve. Go also offers the more relaxed approach to the need for semicolons, type definitions you see in [Groovy](http://groovy-lang.org). 

Go also offers the portability of Java. So this functionality can be run on any platform supporting Go which includes Windows, Linux OS X.

The implementation completely wraps and abstracts the Gobot framework so no understanding of Gobot is required. Masny of the URIs are very simple and can be realized using a browser.

The Web server / URI handling has been implemented using Gorrilla Mux.

## A Simulation Mode
In addition to driving a real drone. It is possible to also coperate the backend with a simulation display.  Each of the action calls are forwarded onto a JET Application called drone-dash which will display messages of what the drone is being asked to do, and provide a visual representation.  Drooe-dash is avaialable [here](https://github.com/oracledeveloperslondon/dronedash).

#Getting Started

To get started and use this framework you need to:
  1. Install Go
  2. Retrive Gorilla/Mux
  3. Download this library
  4. Complie everything
  5. Start up the server

# How the Code Hangs Together
Whilst we have worked hard to ensure the code is self documenting the following provides details on which files do what so you can start to traverse the code, or add to it without compromising the abstratcions and code structure ...

TBD
