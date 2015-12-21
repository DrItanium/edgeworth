// registration mechanisms for machine target parsing
package parser

var parsers map[string]Registration

type Entry struct {
	Line  string
	Index int
}
type Registration interface {
	New(args ...interface{}) (Parser, error)
}
type Parser interface {
	Dumper
	Parse(lines <-chan Entry) error
}

func Register(name string, reg Registration) error {
	if parsers == nil {
		parsers = make(map[string]Registration)
	}

}
