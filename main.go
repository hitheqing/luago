package main

import "fmt"
import . "luago/api"
import _ "luago/binchunk"

func main() {
	//ls := state.New()
	//ls.PushInteger(1)
	//ls.PushString("2.0")
	//ls.PushString("3.0")
	//ls.PushNumber(4.0)
	//printStack(ls)
	////[1]["2.0"]["3.0"][4]
	//
	//ls.Arith(LUA_OPADD)
	//printStack(ls)
	////[1]["2.0"][7]
	//
	//ls.Arith(LUA_OPBNOT)
	//printStack(ls)
	////[1]["2.0"][-8]
	//ls.Len(2)
	//printStack(ls)
	////[1]["2.0"][-8][3]
	//ls.Concat(3)
	//printStack(ls)
	////[1]["2.0-83"]
	//ls.PushBoolean(ls.Compare(1, 2, LUA_OPEQ))
	//printStack(ls)
	//[1]["2.0-83"][false]
}

func printStack(ls LuaState) {
	top := ls.GetTop()
	for i := 1; i <= top; i++ {
		t := ls.Type(i)
		switch t {
		case LUA_TBOOLEAN:
			fmt.Printf("[%t]", ls.ToBoolean(i))
		case LUA_TNUMBER:
			fmt.Printf("[%g]", ls.ToNumber(i))
		case LUA_TSTRING:
			fmt.Printf("[%q]", ls.ToString(i))
		default: // other values
			fmt.Printf("[%s]", ls.TypeName(t))
		}
	}
	fmt.Println()
}
