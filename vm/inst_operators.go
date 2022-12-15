package vm

import . "luago/api"

//运算符相关指令

func binaryop(i Instruction, vm LuaVM, op ArithOp) {
	a, b, c := i.ABC()
	a++
	vm.GetRK(b)
	vm.GetRK(c)
	vm.Arith(op)
	vm.Replace(a)
}

// 对操作数b的寄存器的值一元运算，然后写入操作数a中
func unaryop(i Instruction, vm LuaVM, op ArithOp) {
	a, b, _ := i.ABC()
	a++
	b++
	vm.PushValue(b)
	vm.Arith(op)
	vm.Replace(a)
}

func add(i Instruction, vm LuaVM)  { binaryop(i, vm, LUA_OPADD) }  // +
func sub(i Instruction, vm LuaVM)  { binaryop(i, vm, LUA_OPSUB) }  // -
func mul(i Instruction, vm LuaVM)  { binaryop(i, vm, LUA_OPMUL) }  // *
func mod(i Instruction, vm LuaVM)  { binaryop(i, vm, LUA_OPMOD) }  // %
func pow(i Instruction, vm LuaVM)  { binaryop(i, vm, LUA_OPPOW) }  // ^
func div(i Instruction, vm LuaVM)  { binaryop(i, vm, LUA_OPDIV) }  // /
func idiv(i Instruction, vm LuaVM) { binaryop(i, vm, LUA_OPIDIV) } // //
func band(i Instruction, vm LuaVM) { binaryop(i, vm, LUA_OPBAND) } // &
func bor(i Instruction, vm LuaVM)  { binaryop(i, vm, LUA_OPBOR) }  // |
func bxor(i Instruction, vm LuaVM) { binaryop(i, vm, LUA_OPBXOR) } // ~ 异或
func shl(i Instruction, vm LuaVM)  { binaryop(i, vm, LUA_OPSHL) }  // <<
func shr(i Instruction, vm LuaVM)  { binaryop(i, vm, LUA_OPSHR) }  // >>
func unm(i Instruction, vm LuaVM)  { unaryop(i, vm, LUA_OPUNM) }   // -
func bnot(i Instruction, vm LuaVM) { unaryop(i, vm, LUA_OPBNOT) }  // ~ 按位取反

// ra = length of rb
func _len(i Instruction, vm LuaVM) {
	a, b, _ := i.ABC()
	a++
	b++
	vm.Len(b)
	vm.Replace(a)
}

// ra = rb.. ... ..rc
func concat(i Instruction, vm LuaVM) {
	a, b, c := i.ABC()
	a++
	b++
	c++
	n := c - b + 1
	vm.CheckStack(n)
	for i := b; i <= c; i++ {
		vm.PushValue(i)
	}
	vm.Concat(n)
	vm.Replace(a)
}

func _compare(i Instruction, vm LuaVM, op CompareOp) {
	a, b, c := i.ABC()
	vm.GetRK(b)
	vm.GetRK(c)
	if vm.Compare(-2, -1, op) != (a != 0) {
		vm.AddPC(1)
	}
	vm.Pop(2)
}

func eq(i Instruction, vm LuaVM) { _compare(i, vm, LUA_OPEQ) } // ==
func lt(i Instruction, vm LuaVM) { _compare(i, vm, LUA_OPLT) } // <
func le(i Instruction, vm LuaVM) { _compare(i, vm, LUA_OPLE) } // <=

func not(i Instruction, vm LuaVM) {
	a, b, _ := i.ABC()
	a++
	b++
	vm.PushBoolean(!vm.ToBoolean(b))
	vm.Replace(a)
}

// b = d and e
func testSet(i Instruction, vm LuaVM) {
	a, b, c := i.ABC()
	a++
	b++
	if vm.ToBoolean(b) == (c != 0) {
		vm.Copy(b, a)
	} else {
		vm.AddPC(1)
	}
}

// b = b and e
func test(i Instruction, vm LuaVM) {
	a, _, c := i.ABC()
	a++
	if vm.ToBoolean(a) != (c != 0) {
		vm.AddPC(1)
	}
}
