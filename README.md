# San

Information flows from the Server to the Router then finally to the Engine. The base context
is defined in the Server, which passes it to the router to hand the data to the appropriate
engine. 

Request --> Server ---> Router ---> Engine ---> Response

San creates and manages much of the middleware so that starting a new servers is seamless. 
Logging, context, and server baselines are all handled from the start to enable near 
instant setup of a new server or group of servers.

## Server
Defines the configuration and setup of server runtime environment.

Baseline Server Config:
 - Port the server should listen on
 - Defines and configures middleware
 - Allows configuration of servers for different use cases such as testing, development, or production

## Router
Routers consume the context defined in the server to efficiently determine which engine
or engines to send the data to.

## Engine
Engines are different server types that can be used. There are engines for handling multiple 
types of data including HTTP servers, API Gateways, DNS Engines, etc. They manage formatting
and deserializing data into an easily handled form.

### Drivers
Drivers specify how data should be handled, all the while standardizing logging, context, and
baselines so the focus can be on development.