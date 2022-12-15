package state

import (
	"luago/number"
	"math"
)

type luaTable struct {
	arr  []luaValue
	_map map[luaValue]luaValue
}

func newLuaTable(nArr int, nRecord int) *luaTable {
	t := &luaTable{}
	if nArr > 0 {
		t.arr = make([]luaValue, 0, nArr)
	}
	if nRecord > 0 {
		t._map = make(map[luaValue]luaValue, nRecord)
	}
	return t
}

func (self *luaTable) get(key luaValue) luaValue {
	key = _floatToInteger(key)
	if idx, ok := key.(int64); ok {
		//整数且位于数组索引之内，取数组部分
		if idx > 0 && idx <= int64(len(self.arr)) {
			return self.arr[idx-1]
		}
	}
	return self._map[key]
}

func (self *luaTable) put(key, value luaValue) {
	if key == nil {
		panic("table index key is nil!")
	}
	if f, ok := key.(float64); ok && math.IsNaN(f) {
		panic("table index key is NaN!")
	}
	key = _floatToInteger(key)
	// 数组部分
	if idx, ok := key.(int64); ok && idx >= 1 {
		arrLen := int64(len(self.arr))
		// 已有的slot
		if idx <= arrLen {
			self.arr[idx-1] = value
			// 最后一个元素nil， 数组缩小
			if value == nil && idx == arrLen {
				self._shrinkArray()
			}
			return
		}
		// 刚好是数组的长度+1，视为数组扩容。 从原有的map中移除，转移到数组中
		if idx == arrLen+1 {
			delete(self._map, key)
			if value != nil {
				self.arr = append(self.arr, value)
				self._expandArray()
			}
			return
		}
	}

	// 其他情况， 视为对map的操作
	if value != nil {
		if self._map == nil {
			self._map = make(map[luaValue]luaValue, 4)
		}
		self._map[key] = value
	} else {
		delete(self._map, key)
	}
}

func (self *luaTable) _shrinkArray() {
	for i := len(self.arr) - 1; i >= 0; i-- {
		if self.arr[i] == nil {
			self.arr = self.arr[0:i]
		}
	}
}

func (self *luaTable) _expandArray() {
	for i := int64(len(self.arr)); true; i++ {
		if val, found := self._map[i]; found {
			delete(self._map, i)
			self.arr = append(self.arr, val)
		} else {
			break
		}
	}
}

func (self *luaTable) len() int {
	return len(self.arr)
}

// 浮点数的key，先转换成整数，
func _floatToInteger(key luaValue) luaValue {
	if f, ok := key.(float64); ok {
		if i, ok := number.FloatToInteger(f); ok {
			return i
		}
	}
	return key
}
