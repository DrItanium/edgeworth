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

func (c *Core) PeekAtStack() Word {
	return c.StackMemory[c.GetStackPointer()]
}

func (c *Core) InitStack() {
	c.SetStackPointer(0xFFFF)
}

/* Swap the top two elements on the stack */
func (c *Core) Swap2Stack() {
	sp := c.GetStackPointer()
	f := c.StackMemory[sp]
	s := c.StackMemory[sp-1]
	c.StackMemory[sp] = s
	c.StackMemory[sp-1] = f
}

func (c *Core) DupStack() {
	c.PushOntoStack(c.PeekAtStack())
}
