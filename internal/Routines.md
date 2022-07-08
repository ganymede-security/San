# Routines
Routines are the 'Event' that callers are requesting. They are a
list of functions in a chain of a parent 'Routine' and child 'Subroutines' that a Routine can call as necessary, with the final subroutine breaking 
the event loop and freeing the resources used. 

They contain a hook method (the Routine/Subroutine name) that 
represents the operation and also is the 'Key' to call them from the 
event dispatcher. 

In Go terminology it would be called a Pipeline. 

## Routine/Subroutine Model

Dispatcher --> Routine --> Subroutine -- Subroutine --> End

## Adding Middleware
Middleware is added in a closure function wrapping the Routine 
with a function that needs to be performed on each instance of
the Routine being used.