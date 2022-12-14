package state

// 实现接口方法可以利用ide的功能，生成接口方法，然后放到不同的文件中

func (self *luaState) PC() int {
	return self.pc
}

func (self *luaState) AddPC(n int) {
	self.pc += n
}

func (self *luaState) Fetch() uint32 {
	i := self.proto.Code[self.pc]
	self.pc++
	return i
}

func (self *luaState) GetConst(idx int) {
	c := self.proto.Constants[idx]
	self.stack.push(c)
}

// iABC OpArgK类型参数。 9bit 如果最高位1，则存放的常量表索引。取后面的位
// 否则存放的寄存器索引。 该索引从0开始，而api从1开始。需要+1
func (self *luaState) GetRK(rk int) {
	if rk > 0xff {
		self.GetConst(rk & 0xff)
	} else {
		self.PushValue(rk + 1)
	}
}
