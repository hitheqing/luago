package state

// [-0, +1, e]
// http://www.lua.org/manual/5.3/manual.html#lua_len
func (self *luaState) Len(idx int) {
	val := self.stack.get(idx)

	if s, ok := val.(string); ok {
		self.stack.push(int64(len(s)))
	} else {
		panic("length error!")
	}
}

// [-n, +1, e]
// http://www.lua.org/manual/5.3/manual.html#lua_concat
// 依次连接栈顶n个元素
func (self *luaState) Concat(n int) {
	if n == 0 {
		self.stack.push("")
	} else if n >= 2 {
		//todo 重复pop push了，这里应该有更好的实现..
		for i := 1; i < n; i++ {
			if self.IsString(-1) && self.IsString(-2) {
				s2 := self.ToString(-1)
				s1 := self.ToString(-2)
				self.stack.pop()
				self.stack.pop()
				self.stack.push(s1 + s2)
				continue
			}

			panic("concatenation error!")
		}
	}
	// n == 1, do nothing
}
