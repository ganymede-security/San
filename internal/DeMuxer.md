# The DeMuxer (De-Multiplexer)
*Look into starting in BPF for extra efficiency*
The DeMuxer is the entrypoint into the application. It performs setup and 
listens for events to occur. 

When an event occurs, it will inform the proper Driver and hand off the 
data. 

It will also notify an Driver when an event stream is 'Ready' for I/O

It's possible for multiple functions to read from the same channel 
with the fan-out pattern goroutines enable, but it's unclear if this is 
something that will be possible for multiple Drivers to take advantage of

# eBPF
eBPF allows the DeMuxer to interop directly with the kernel, allowing
insight and control over Networking, Observability and Security that
would not otherwise be possible.

Networking applications that don't take utilize eBPF are limited to
the user space which doesn't have any way to take advantage of these
utilities. 

Some examples include:
Since eBPF can get closer to the NIC than any application otherwise
would, it is able to filter a colossal amount of traffic at an
extremely fine grained level, preventing DDOS attacks or general
traffic filtering.

It can be used to create extremely smart performant firewalls, 
ensuring firewalls are being complied with, detecting 
vulnerability patterns, or auditing network traffic. 

It can create application level load balancers

## Startup
Each Driver will register a minimum amount of information with the 
DeMuxer so that it can perform the following syscalls:

### socket() - net.Listen:

### bind() - net.Listen:

### listen() - net.Listen:

### accept() - Listener.Accept interface:

## Performing Operations

After the DeMuxer has successfully bound to all the ports required by each
Driver, it can begin performing operations. It doesn't care about the type 
of data received, it simply passes it on to the appropriate Driver.

It is able to do this concurrently using a goroutine to listen on the TCP/
UDP port requested by each engine. When an event is received by a certain 
port it will forward the event to the DeMuxer which will then pass it to 
the Driver that requested the port be opened.

### Handling Multiple Services on the same connection
The demuxer also has the ability to handle multiple services on the same
port, if for example you had a gRPC and a JSON service both could run on 
port 443 at the same time.

## Closing the connection
After the client closes the connection, the DeMuxer passes a close event 
to the Driver, freeing the thread and the resources used by the Driver and 
Routines.

# Security
Because the DeMuxer's job is only to forward traffic to the appropriate
Engine, ensuring that data is of the correct type is the responsibility of 
the Driver. 

# Performance 
 - Notify Engine of 'Ready' status
 - Release Engine and drivers depending on if there are more operations in the queue 