package state

import . "luago/api"

// [-0, +1, m]
// http://www.lua.org/manual/5.3/manual.html#lua_newtable
func (self *luaState) NewTable() {
	self.CreateTable(0, 0)
}

// [-0, +1, m]
// http://www.lua.org/manual/5.3/manual.html#lua_createtable
func (self *luaState) CreateTable(nArr, nRecord int) {
	t := newLuaTable(nArr, nRecord)
	self.stack.push(t)
}

// [-1, +1, e]
// http://www.lua.org/manual/5.3/manual.html#lua_gettable
// 栈顶元素作为key，-1，然后取t[key], +1
func (self *luaState) GetTable(idx int) LuaType {
	t := self.stack.get(idx)
	k := self.stack.pop()
	return self.getTable(t, k)

}

// [-0, +1, e]
// http://www.lua.org/manual/5.3/manual.html#lua_getfield
func (self *luaState) GetField(idx int, key string) LuaType {
	t := self.stack.get(idx)
	return self.getTable(t, key)
}

// [-0, +1, e]
// http://www.lua.org/manual/5.3/manual.html#lua_geti
// 为数组设计的
func (self *luaState) GetI(idx int, i int64) LuaType {
	t := self.stack.get(idx)
	v := self.getTable(t, i)
	return v
}

// push(t[k])
func (self *luaState) getTable(t, k luaValue) LuaType {
	if tbl, ok := t.(*luaTable); ok {
		v := tbl.get(k)
		self.stack.push(v)
		return typeOf(v)
	}

	panic("not a table!") // todo
}

// [-2, +0, e]
// http://www.lua.org/manual/5.3/manual.html#lua_settable
func (self *luaState) SetTable(idx int) {
	t := self.stack.get(idx)
	v := self.stack.pop()
	k := self.stack.pop()
	self.setTable(t, k, v)
}

// [-1, +0, e]
// http://www.lua.org/manual/5.3/manual.html#lua_setfield
func (self *luaState) SettField(idx int, key string) {
	t := self.stack.get(idx)
	v := self.stack.pop()
	k := key
	self.setTable(t, k, v)
}

// [-1, +0, e]
// http://www.lua.org/manual/5.3/manual.html#lua_seti
func (self *luaState) SetI(idx int, i int64) {
	t := self.stack.get(idx)
	v := self.stack.pop()
	//k := self.stack.pop()
	self.setTable(t, i, v)
}

// t[k]=v
func (self *luaState) setTable(t, k, v luaValue) {
	if tbl, ok := t.(*luaTable); ok {
		tbl.put(k, v)
		return
	}

	panic("not a table!")
}
