# San
  - [Design Goals:](#design-goals)
- [San](#san)
  - [Design Goals:](#design-goals)
- [Mapping The Flow of Data](#mapping-the-flow-of-data)
  - [The DeMuxer(De-Multiplexer)](#the-demuxerde-multiplexer)
    - [What does it do:](#what-does-it-do)
    - [Interfacing:](#interfacing)
    - [Mapping the syscalls:](#mapping-the-syscalls)
    - [Separation of Concerns:](#separation-of-concerns)
  - [The Driver (Initiation Dispatcher)](#the-driver-initiation-dispatcher)
    - [What Does it do:](#what-does-it-do-1)
    - [Interfacing:](#interfacing-1)
    - [Mapping the syscalls:](#mapping-the-syscalls-1)
    - [Separation of Concerns:](#separation-of-concerns-1)
  - [Routines:](#routines)
    - [What does it do:](#what-does-it-do-2)
    - [Interfacing:](#interfacing-2)
    - [Mapping the syscalls:](#mapping-the-syscalls-2)
    - [Separation of Concerns:](#separation-of-concerns-2)
## Design Goals:
San is a framework for setting up high efficiency network applications capable 
of handling many different protocols from the Data Link layer up to the 
Application layer.

It aims to allow network applications to work in concert with 
each other extremely efficiently, using a combination of de-multiplexing for 
concurrent event handling as well as context aware code that shares 
information that may be useful to another server or Driver. 

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
a network application and manage context, they don't have any built in 
functions to handle data. That's where Routines come in. 

Routines are custom built to  handle data, following the pattern set by 
Drivers to ensure information is logged correctly, and providing 
additional context to the DeMuxer to ensure efficient dispatching.

# Mapping The Flow of Data

In a system where data only needs to use one Driver, the data flow is 
straightforward.

```
Request --> DeMuxer --> Driver --> Routine --> Response
```

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
 - Maintains a table of Routines and Subroutines following the 'Chain of Responsibility Pattern' as well as methods to Register and Remove Routines
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