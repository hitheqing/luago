package api

type LuaVM interface {
	LuaState
	PC() int          // get program counter, test only
	AddPC(n int)      // used for 'JUMP'
	Fetch() uint32    // get pc, pc to next
	GetConst(idx int) // push const value to stack
	GetRK(rk int)     // push regist const to stack
}
