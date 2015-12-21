// registration mechanisms for machine target parsing
package parser

import (
	"fmt"
	"github.com/DrItanium/edgeworth"
)

var parsers map[string]Registration

type Entry struct {
	Line  string
	Index int
}
type Registration interface {
	New(args ...interface{}) (Parser, error)
}
type Parser interface {
	edgeworth.Dumper
	Parse(lines <-chan Entry) error
}

func Register(name string, reg Registration) error {
	if parsers == nil {
		parsers = make(map[string]Registration)
	}
	if _, ok := parsers[name]; ok {
		return fmt.Errorf("Parser %s is already registered!", name)
	} else {
		parsers[name] = reg
		return nil
	}

}

func GetRegistered() []string {
	var names []string
	if parsers != nil {
		for name, _ := range parsers {
			names = append(names, name)
		}
	}
	return names
}

func New(name string, args ...interface{}) (Parser, error) {
	if parsers == nil {
		return nil, fmt.Errorf("No parsers registered!")
	}
	if gen, ok := parsers[name]; ok {
		return gen.New(args)
	} else {
		return nil, fmt.Errorf("%s does not refer to a registered parser!", name)
	}
}
func IsRegistered(name string) bool {
	if parsers == nil {
		return false
	} else {
		_, ok := parsers[name]
		return ok
	}
}

func Activate() {}
