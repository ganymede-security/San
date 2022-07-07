# San

# The San Model

## Design Goals:
San is a framework for setting up high efficiency network applications capable 
of handling many different protocols from the Data Link layer up to the 
Application layer.

It aims to allow network applications to work in concert with 
each other extremely efficiently, using a combination of de-multiplexing for 
concurrent event handling as well as context aware code that shares 
information that may be useful to another server. 

Most network applications handle events at the Application layer, meaning
the time they have and the scope of their awareness is relatively limited
to what happens at that layer.

San is able to start that process earlier and faster than other applications 
by utilizing eBPF to analyze network events at the Data Link layer as 
opposed to other applications that handle events at the application layer.

A DeMuxer is highly efficient at determining what should be done with 
data, whether it should be passed on to an Driver(s)
or determining if it should be passed on to another network location.

Routines are how Driver process data. While Drivers handle the setup of 
a server and manage context, they don't have any built in functions to
handle data. That's where Routines come in. Routines are custom built to 
handle data, following the pattern set by Drivers to ensure information 
is logged correctly, and providing additional context to the DeMuxer to
ensure efficient routing.

# Mapping The Flow of Data

In a system where data only needs to use one, the data flow is 
straightforward.

```
Request --> DeMuxer --> Driver --> Routine --> Response
```

# Reactor Pattern
A reactor design pattern works in concert with the concept of multiplexing. It 
handles requests coming from one or multiple clients concurrently. 
It takes these concurrent events and uses a de-multiplexer to handle them
synchronously.

Benefits: 
 - Application level code is completely separated from the reactor, which allows it to be divided into modular components without fear of disrupting a separate application's code.
 - Enables Separation of Concerns. An Application Handler is implemented independently from the other components, enabled greater modularization.
 - Improves portability: The Driver interface exists independently of the OS being used
 - Provides concurrent performance without the additional programmer effort required by multithreading

Detractors:
 - Can add difficulties when trying to debug. Pattern is based on single threaded operations, can switch to a Proactor pattern if async becomes a requirement
 - Rotuines are disincentivized from using blocking I/O, as this impedes the thread from processing requests in the queue. If long periods of blocking are required Routines should consider multithreading

### Axioms:
 - Scalability is limited both by the calling of request handlers synchronously, but also by the de-multiplexer, therefore performance should be optimized around these first.
 - Lookups and Updates are trivial and should be prioritized 
 - Reads and Writes 

## The DeMuxer(De-Multiplexer)
### What does it do: 
 - The de-multiplexer is responsible for waiting for events to occur
 - When an event occurs, it informs the Initiation Dispatcher (Driver) to deserialize the information and hand off to a driver for handling
 - Notifies the Driver when an event source becomes 'Ready'

### Interfacing:
 - Handle

### Mapping the syscalls:
 - socket()
 - bind()
 - listen()
 - accept()?

### Separation of Concerns:
 - The demuxer should only be concerned with forwarding data, not analyzing it

## The Driver (Initiation Dispatcher)
### What Does it do:
 - Defines an interface for adding, removing, and dispatching event handlers
 - Requests each Routine to pass back its internal handle (The method that the driver performs)
 - When the demuxer receives an event, it forwards to the appropriate Routine which then signals the driver
 - It uses the handle methods of the driver as 'Keys' to locate and dispatch the hook method.
 - Maintains a table of Rotuines and Subroutines following the 'Chain of Responsibility Pattern' as well as methods to Register and Remove Routines
 - Provide an entry point into the blocking event loop with a use_driver method

### Interfacing:
 - Read? io.Reader interface
 - Request/ResponseWriter

### Mapping the syscalls:
 - 

### Separation of Concerns:
 - The Driver Should only be concerned with handling a specific type of data

## Routines:
### What does it do:
 - Follows the Chain or Responsibility Pattern for routines and sub-routines
 - Specifies an interface consisting of a hook method that abstractly represents the dispatching operation for service specific events. Must be implemented by a Driver as it performs application specific services.
 - Implements the hook method as well as methods to process data in an application specific manner.
 - Composed of an interface, the 'Routine' and a function 'Handle'
### Interfacing:
 - Serve?
 - Handle?

### Mapping the syscalls:

### Separation of Concerns:
 - The driver should not be concerned with receiving data of a different type than expected

## Helpful Reading
 - [The Reactor Pattern](http://www.dre.vanderbilt.edu/~schmidt/PDF/reactor-siemens.pdf)
