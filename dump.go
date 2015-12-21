// dumper interface
package edgeworth

type Dumper interface {
	Dump(output chan<- byte) error
}
