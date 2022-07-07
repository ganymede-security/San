# Design Philosophy

## Reactor Pattern
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

## Helpful Reading
 - [The Reactor Pattern](http://www.dre.vanderbilt.edu/~schmidt/PDF/reactor-siemens.pdf)