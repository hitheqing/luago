package state

import "luago/binchunk"

type luaState struct {
	// luaState 目前仅有一个栈可供操作
	stack *luaStack
	// 增加函数原型和pc
	proto *binchunk.Prototype
	pc    int
}

func New(size int, proto *binchunk.Prototype) *luaState {
	return &luaState{
		stack: newLuaStack(size),
		proto: proto,
		pc:    0,
	}
}
