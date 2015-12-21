// generic machine interface
package edgeworth

import (
	"fmt"
)

var registrations map[string]MachineRegistration

func RegisterMachine(name string, gen MachineRegistration) error {
	if registrations == nil {
		registrations = make(map[string]MachineRegistration)
	}
	if _, ok := registrations[name]; ok {
		return fmt.Errorf("Key %s is already registered!", name)
	} else {
		registrations[name] = gen
		return nil
	}
}
func RegisteredMachines() []string {
	var names []string
	if registrations != nil {
		for name, _ := range registrations {
			names = append(names, name)
		}
	}
	return names
}

func NewMachine(name string, args ...interface{}) (Machine, error) {
	if registrations == nil {
		return nil, fmt.Errorf("No machines registered!")
	}
	if gen, ok := registrations[name]; ok {
		return gen.New(args)
	} else {
		return nil, fmt.Errorf("%s does not refer to a registered machine!", name)
	}
}

type MachineRegistration interface {
	New(args ...interface{}) (Machine, error)
}

type Machine interface {
	GetDebugStatus() bool
	SetDebug(value bool)
	InstallProgram(input <-chan byte) error
	Dump(output chan<- byte) error
	Startup() error
	Shutdown() error
	Run() error
}

// Dummy function used to force inclusion of this library
func Activate() {}
