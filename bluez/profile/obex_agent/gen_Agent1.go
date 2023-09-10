// Code generated by go-bluetooth generator DO NOT EDIT.

package obex_agent

import (
	"sync"

	"github.com/godbus/dbus/v5"
	"github.com/muka/go-bluetooth/bluez"
	"github.com/muka/go-bluetooth/props"
	"github.com/muka/go-bluetooth/util"
)

var Agent1Interface = "org.bluez.obex.Agent1"

// NewAgent1 create a new instance of Agent1
//
// Args:
// - servicePath: unique name
// - objectPath: freely definable
func NewAgent1(servicePath string, objectPath dbus.ObjectPath) (*Agent1, error) {
	a := new(Agent1)
	a.client = bluez.NewClient(
		&bluez.Config{
			Name:  servicePath,
			Iface: Agent1Interface,
			Path:  dbus.ObjectPath(objectPath),
			Bus:   bluez.SystemBus,
		},
	)
	a.Properties = new(Agent1Properties)

	_, err := a.GetProperties()
	if err != nil {
		return nil, err
	}
	return a, nil
}

/*
Agent1 Agent hierarchy
*/
type Agent1 struct {
	client                 *bluez.Client
	propertiesSignal       chan *dbus.Signal
	objectManagerSignal    chan *dbus.Signal
	objectManager          *bluez.ObjectManager
	Properties             *Agent1Properties
	watchPropertiesChannel chan *dbus.Signal
}

// Agent1Properties contains the exposed properties of an interface
type Agent1Properties struct {
	lock sync.RWMutex `dbus:"ignore"`
}

// Lock access to properties
func (p *Agent1Properties) Lock() {
	p.lock.Lock()
}

// Unlock access to properties
func (p *Agent1Properties) Unlock() {
	p.lock.Unlock()
}

// Close the connection
func (a *Agent1) Close() {
	a.unregisterPropertiesSignal()
	a.client.Disconnect()
}

// Path return Agent1 object path
func (a *Agent1) Path() dbus.ObjectPath {
	return a.client.Config.Path
}

// Client return Agent1 dbus client
func (a *Agent1) Client() *bluez.Client {
	return a.client
}

// Interface return Agent1 interface
func (a *Agent1) Interface() string {
	return a.client.Config.Iface
}

// GetObjectManagerSignal return a channel for receiving updates from the ObjectManager
func (a *Agent1) GetObjectManagerSignal() (chan *dbus.Signal, func(), error) {

	if a.objectManagerSignal == nil {
		if a.objectManager == nil {
			om, err := bluez.GetObjectManager()
			if err != nil {
				return nil, nil, err
			}
			a.objectManager = om
		}

		s, err := a.objectManager.Register()
		if err != nil {
			return nil, nil, err
		}
		a.objectManagerSignal = s
	}

	cancel := func() {
		if a.objectManagerSignal == nil {
			return
		}
		a.objectManagerSignal <- nil
		a.objectManager.Unregister(a.objectManagerSignal)
		a.objectManagerSignal = nil
	}

	return a.objectManagerSignal, cancel, nil
}

// ToMap convert a Agent1Properties to map
func (a *Agent1Properties) ToMap() (map[string]interface{}, error) {
	return props.ToMap(a), nil
}

// FromMap convert a map to an Agent1Properties
func (a *Agent1Properties) FromMap(props map[string]interface{}) (*Agent1Properties, error) {
	props1 := map[string]dbus.Variant{}
	for k, val := range props {
		props1[k] = dbus.MakeVariant(val)
	}
	return a.FromDBusMap(props1)
}

// FromDBusMap convert a map to an Agent1Properties
func (a *Agent1Properties) FromDBusMap(props map[string]dbus.Variant) (*Agent1Properties, error) {
	s := new(Agent1Properties)
	err := util.MapToStruct(s, props)
	return s, err
}

// ToProps return the properties interface
func (a *Agent1) ToProps() bluez.Properties {
	return a.Properties
}

// GetWatchPropertiesChannel return the dbus channel to receive properties interface
func (a *Agent1) GetWatchPropertiesChannel() chan *dbus.Signal {
	return a.watchPropertiesChannel
}

// SetWatchPropertiesChannel set the dbus channel to receive properties interface
func (a *Agent1) SetWatchPropertiesChannel(c chan *dbus.Signal) {
	a.watchPropertiesChannel = c
}

// GetProperties load all available properties
func (a *Agent1) GetProperties() (*Agent1Properties, error) {
	a.Properties.Lock()
	err := a.client.GetProperties(a.Properties)
	a.Properties.Unlock()
	return a.Properties, err
}

// SetProperty set a property
func (a *Agent1) SetProperty(name string, value interface{}) error {
	return a.client.SetProperty(name, value)
}

// GetProperty get a property
func (a *Agent1) GetProperty(name string) (dbus.Variant, error) {
	return a.client.GetProperty(name)
}

// GetPropertiesSignal return a channel for receiving udpdates on property changes
func (a *Agent1) GetPropertiesSignal() (chan *dbus.Signal, error) {

	if a.propertiesSignal == nil {
		s, err := a.client.Register(a.client.Config.Path, bluez.PropertiesInterface)
		if err != nil {
			return nil, err
		}
		a.propertiesSignal = s
	}

	return a.propertiesSignal, nil
}

// Unregister for changes signalling
func (a *Agent1) unregisterPropertiesSignal() {
	if a.propertiesSignal != nil {
		a.propertiesSignal <- nil
		a.propertiesSignal = nil
	}
}

// WatchProperties updates on property changes
func (a *Agent1) WatchProperties() (chan *bluez.PropertyChanged, error) {
	return bluez.WatchProperties(a)
}

func (a *Agent1) UnwatchProperties(ch chan *bluez.PropertyChanged) error {
	return bluez.UnwatchProperties(a, ch)
}

/*
Release 			This method gets called when the service daemon

	unregisters the agent. An agent can use it to do
	cleanup tasks. There is no need to unregister the
	agent, because when this method gets called it has
	already been unregistered.
*/
func (a *Agent1) Release() error {
	return a.client.Call("Release", 0).Store()
}

/*
AuthorizePush 			This method gets called when the service daemon

	needs to accept/reject a Bluetooth object push request.
	Returns the full path (including the filename) or the
	folder name suffixed with '/' where the object shall
	be stored. The transfer object will contain a Filename
	property that contains the default location and name
	that can be returned.
	Possible errors: org.bluez.obex.Error.Rejected
	                 org.bluez.obex.Error.Canceled
*/
func (a *Agent1) AuthorizePush(transfer dbus.ObjectPath) (string, error) {
	var val0 string
	err := a.client.Call("AuthorizePush", 0, transfer).Store(&val0)
	return val0, err
}

/*
Cancel 			This method gets called to indicate that the agent

	request failed before a reply was returned. It cancels
	the previous request.
*/
func (a *Agent1) Cancel() error {
	return a.client.Call("Cancel", 0).Store()
}
