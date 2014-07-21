// a series of memory related operations
package edgeworth

func (c *Core) SetRegister(index, value Word) {
	c.Registers[index] = value
}
