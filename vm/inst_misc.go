package vm

import . "luago/api"

//其他类型 MOVE JMP等

// R(a)=R(b) 操作数只有8bit，局部变量不能大于2^8 255个。实际上限制在200
func move(i Instruction, vm LuaVM) {
	a, b, _ := i.ABC()
	a++
	b++
	vm.Copy(b, a)
}

func jmp(i Instruction, vm LuaVM) {
	a, sbx := i.AsBx()
	vm.AddPC(sbx)
	if a != 0 { // a 和upvalue有关 先不处理
		panic("todo!")
	}
}
