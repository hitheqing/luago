package state

// luaState 目前仅有一个栈可供操作
type luaState struct {
	stack *luaStack
}

func New() *luaState {
	return &luaState{
		stack: newLuaStack(20),
	}
}
