package vm

import . "luago/api"

// 加载类指令

func loadNil(i Instruction, vm LuaVM) {
	a, b, _ := i.ABC()
	a++
	// 这里不能这么写，因为push的位置不一定是在栈顶
	//for i := a; i <= a+b; i++ {
	//	vm.PushNil()
	//}
	vm.PushNil()
	for i := a; i <= a+b; i++ {
		vm.Copy(-1, i)
	}
	vm.Pop(1)
}

func loadBool(i Instruction, vm LuaVM) {
	a, b, c := i.ABC()
	a++
	vm.PushBoolean(b != 0)
	vm.Replace(a)
	if c != 0 {
		vm.AddPC(1)
	}
}

// bx 18bit 可表示262143.通常不会超过这个数。但是对于某些配置表可能超出，用loadkx来弥补
func loadK(i Instruction, vm LuaVM) {
	a, b := i.ABx()
	a++
	vm.GetConst(b)
	vm.Replace(a)
}

// 解决上述问题，额外读取下一条指令，取ax操作数，
func loadKx(i Instruction, vm LuaVM) {
	a, _ := i.ABx()
	a++
	ax := Instruction(vm.Fetch()).Ax()
	vm.GetConst(ax)
	vm.Replace(a)
}
