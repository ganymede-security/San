package ethernet

// Driver is the interface for the Engine.
// A Driver is fulfilled when it is supplied with a
// Handler method to handle the request, a context to handle the
// request with, as well as a standard log entry for emitting logs.
type EthernetDriver interface {
	// Handler is an undefined function containing the
	// action to be performed by the driver.
	Handle(ctx context.Context, request []byte, response []byte)
	// Defines the information to be included in the
	// standard log Entry format for the driver
	//Log (*DriverEntry)
	// return/add to context
	//Context(context.Context)
}

type EthernetHandler struct {
}

// Creates a New Driver with all middleware included
func NewDriver(EthernetHandler) *engine.Driver {
	eng := engine.NewEngine()

	return &engine.Driver{
		Handler: 
	}
}

func (d Driver) ReadByte() {

}

// Handle function defines how the Ethernet driver handles server events.
func (e engine.Engine) Handle(ctx context.Context) {
	
}