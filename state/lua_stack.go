package state

// lua栈 ： 可以用索引访问的栈
type luaStack struct {
	// 槽位数组
	slots []luaValue
	// 栈顶index， 默认0， push第一个值以后变成1，用 slot[top-1]来表示当前栈顶元素
	top int
}

// 创建lua栈
func newLuaStack(size int) *luaStack {
	return &luaStack{
		slots: make([]luaValue, size),
		top:   0,
	}
}

// 确保栈容量有n
func (self *luaStack) check(n int) {
	free := len(self.slots) - self.top
	for i := free; i < n; i++ {
		self.slots = append(self.slots, nil)
	}
}

// 往栈里push luaValue
func (self *luaStack) push(val luaValue) {
	if self.top == len(self.slots) {
		panic("stack overflow!")
	}
	self.slots[self.top] = val
	self.top++
}

// 从栈顶取出 luaValue
func (self *luaStack) pop() luaValue {
	if self.top < 1 {
		panic("stack underflow!")
	}
	val := self.slots[self.top-1]
	self.slots[self.top-1] = nil
	self.top -= 1
	return val
}

// 转换index为绝对index。
func (self *luaStack) absIndex(idx int) int {
	if idx >= 0 {
		return idx
	}
	return idx + self.top + 1
}

// 是否valid index  [1,top]之内为合法
func (self *luaStack) isValid(idx int) bool {
	absIndex := self.absIndex(idx)
	return absIndex > 0 && absIndex <= self.top
}

// get, [1,top]之内为合法
func (self *luaStack) get(idx int) luaValue {
	absIndex := self.absIndex(idx)
	if absIndex > 0 && absIndex <= self.top {
		return self.slots[absIndex-1]
	}
	return nil
}

// set,[1,top]之内为合法
func (self *luaStack) set(idx int, val luaValue) {
	absIndex := self.absIndex(idx)
	if absIndex > 0 && absIndex <= self.top {
		self.slots[absIndex-1] = val
		return
	}
	panic("invalid index")
}

func (self *luaStack) reverse(from, to int) {
	slots := self.slots
	for from < to {
		slots[from], slots[to] = slots[to], slots[from]
		from++
		to--
	}
}
