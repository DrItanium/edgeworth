// stack operations
package edgeworth

func (c *Core) PushOntoStack(value Word) {
	sp := c.GetStackPointer() + 1
	c.StackMemory[sp] = value
	c.SetStackPointer(sp)
}

func (c *Core) PopOffStack() Word {
	sp := c.GetStackPointer()
	val := c.StackMemory[sp]
	c.SetStackPointer(sp - 1)
	return val
}

func (c *Core) InitStack() {
	c.SetStackPointer(0xFFFF)
}
