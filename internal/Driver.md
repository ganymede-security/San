# The Driver (Initiation Dispatcher)
The Driver is the interface between the network Callers and the Routine
that handles the call. 

It defines methods for adding, removing and dispatching the 
Routines that process network Calls. It maintains a table of Routines and 
their Subroutines using the 'Method' that Routines define as the 
'Keys' for registering and removing Routines from the table.

Drivers are the entry point into the blocking event loop when the 
Routine implements their Method.

Drivers also enforce standards for logging, creating Contexts, and 
implement any middleware that is required for the application.

## Working with a DeMuxer
The demuxer is the entrypoint for the application, so has to be able to
pass data to the Driver. It requires configuration settings from the 
Driver in order to be able to open a port on its behalf.

This means that configuration of the "environment" de facto 
happens at the Driver.

Configuration Settings that can be passed to the Demuxer include:
 - Listen Port
 - ReadBuffer - Optional
 - WriteBuffer - Optional
 - DefaultContext - Optional (Because there will be a default)
 - CallContext
 - EventBus - Optional (Because there will be a default)
 - Logger

Why are things like a Logger and CallContext required? The goal of San
is not only to make servers performant, but also to make setup and
maintenance features baked in to the development experience. 

In order to make servers run better in the long run, San strongly 
encourages (and tries to make easy) following best practices from the 
get go. Therefore, Context and Logging get first class treatment.

There will also be methods that a Driver interface specifies:

*Interfaces should NOT be a list of useful methods that some subsystem you wrote exports. (I bring this up because the article suggests such an architecture.) Such an interface is very difficult to implement, which leads to there only being one implementation or to dramatically increasing the maintenance burden of the codebase. If there’s only one implementation, the codebase would be easier to understand with concrete types. If there are multiple implementations, you should ask yourself why.*

 - InterruptHandler()?
 - Init()?
 - Open()?
 - Close()? - Shutdown()
 - Read()?
 - Write()?

## The Dispatcher
Dispatches calls/callers to the appropriate Routine which handles them
in a trie list of routines and subroutines.