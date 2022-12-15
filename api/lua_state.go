package api

type LuaType = int
type ArithOp = int
type CompareOp = int

/*
目前为止luastate包含的功能
读写栈/栈上算术运算/封装栈上的按类型访问helper
*/
type LuaState interface {
	/* basic stack manipulation */
	GetTop() int
	AbsIndex(idx int) int
	CheckStack(n int) bool
	Pop(n int)
	Copy(fromIdx, toIdx int)
	PushValue(idx int)
	// 把idx的值替换为栈顶的值
	Replace(idx int)
	Insert(idx int)
	Remove(idx int)
	Rotate(idx, n int)
	SetTop(idx int)
	/* access functions (stack -> Go) */
	TypeName(tp LuaType) string
	Type(idx int) LuaType
	IsNone(idx int) bool
	IsNil(idx int) bool
	IsNoneOrNil(idx int) bool
	IsBoolean(idx int) bool
	IsInteger(idx int) bool
	IsNumber(idx int) bool
	IsString(idx int) bool
	IsTable(idx int) bool
	IsThread(idx int) bool
	IsFunction(idx int) bool
	ToBoolean(idx int) bool
	ToInteger(idx int) int64
	ToIntegerX(idx int) (int64, bool)
	ToNumber(idx int) float64
	ToNumberX(idx int) (float64, bool)
	ToString(idx int) string
	ToStringX(idx int) (string, bool)
	/* push functions (Go -> stack) */
	PushNil()
	PushBoolean(b bool)
	PushInteger(n int64)
	PushNumber(n float64)
	PushString(s string)
	/* 算术运算 */
	Arith(op ArithOp)
	Compare(idx1, idx2 int, op CompareOp) bool
	Len(idx int)
	Concat(n int)
	/* table 操作 get */
	NewTable()
	CreateTable(nArr, nRecord int)
	GetTable(idx int) LuaType
	GetField(idx int, key string) LuaType
	GetI(idx int, i int64) LuaType
	/* table 操作 set */
	SetTable(idx int)
	SettField(idx int, key string)
	SetI(idx int, i int64)
}
