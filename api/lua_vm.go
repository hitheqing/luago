package api

type LuaVM interface {
	LuaState
	// PC get program counter, test only
	PC() int
	// AddPC 用于实现jump指令，n可以为负数
	AddPC(n int)
	// Fetch get pc, pc to next
	Fetch() uint32
	// GetConst push const value to stack
	GetConst(idx int)
	// GetRK push register const to stack
	GetRK(rk int)
}
