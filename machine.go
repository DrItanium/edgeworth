// generic machine interface
package edgeworth

import (
	"fmt"
)

var registrations map[string]MachineRegistration

func RegisterMachine(name string, gen MachineRegistration) error {
	if registrations == nil {
		registrations := make(map[string]MachineRegistration)
	}
	if _, ok := registrations[name]; ok {
		return fmt.Errorf("Key %s is already registered!", name)
	} else {
		registrations[name] = gen
		return nil
	}
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
	Run() error
}
