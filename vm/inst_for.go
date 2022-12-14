package vm

import "luago/api"

//for循环指令

func forPrep(i Instruction, vm api.LuaVM) {
	a, sBx := i.AsBx()
	a++
	//ra -= r(a+2)
	vm.PushValue(a)
	vm.PushValue(a + 2)
	vm.Arith(api.LUA_OPSUB)
	vm.Replace(a)
	//pc+=sbx
	vm.AddPC(sBx)
}

func forLoop(i Instruction, vm api.LuaVM) {
	a, sBx := i.AsBx()
	a++
	//ra += r(a+2)
	vm.PushValue(a)
	vm.PushValue(a + 2)
	vm.Arith(api.LUA_OPADD)
	vm.Replace(a)
	// 步长是否大于0
	positive := vm.ToNumber(a+2) >= 0
	if positive && vm.Compare(a, a+1, api.LUA_OPLE) ||
		!positive && vm.Compare(a+1, a, api.LUA_OPLE) {
		vm.AddPC(sBx)
		vm.Copy(a, a+3)
	}
}
