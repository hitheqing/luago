package binchunk

const (
	LUA_SIGNATURE = "\x1bLua"            //magic bytes
	LUAC_VERSION  = 0x53                 //version
	LUAC_FORMAT   = 0                    // 格式号，官方为 0
	LUAC_DATA     = "\x19\x93\r\n\x1a\n" // 文件损坏不加载
	// 长度
	CINT_SIZE        = 4 //
	CSIZET_SIZE      = 4
	INSTRUCTION_SIZE = 4
	LUA_INTEGER_SIZE = 8
	LUA_NUMBER_SIZE  = 8
	// 检测二进制大小端 和 本机 是否匹配
	LUAC_INT = 0x5678
	//检查浮点数格式是否匹配
	LUAC_NUM = 370.5
)

const (
	TAG_NIL       = 0x00
	TAG_BOOLEAN   = 0x01
	TAG_NUMBER    = 0x03
	TAG_SHORT_STR = 0x04

	TAG_INTEGER  = 0x13
	TAG_LONG_STR = 0x14
)

type binaryChunk struct {
	header
	sizeUpvalues byte // ?
	mainFunc     *Prototype
}

type header struct {
	signature       [4]byte
	version         byte
	format          byte
	luacData        [6]byte
	cintSize        byte
	sizetSize       byte
	instructionSize byte
	luaIntegerSize  byte
	luaNumberSize   byte
	luacInt         int64
	luacNum         float64
}

// function prototype
type Prototype struct {
	Source          string // debug
	LineDefined     uint32
	LastLineDefined uint32
	NumParams       byte
	IsVararg        byte
	//寄存器数量（栈大小）
	MaxStackSize byte
	//指令表
	Code []uint32
	//常量表
	Constants []interface{}
	//闭包值表
	Upvalues []Upvalue
	//子函数表
	Protos   []*Prototype
	LineInfo []uint32 // debug
	//local vars
	LocVars      []LocVar // debug
	UpvalueNames []string // debug
}

type Upvalue struct {
	Instack byte
	Idx     byte
}

type LocVar struct {
	VarName string
	StartPC uint32
	EndPC   uint32
}

// 反序列化，解析lua编译后的二进制文件
func Undump(data []byte) *Prototype {
	reader := &reader{data}
	reader.checkHeader()
	reader.readByte() // size_upvalues
	return reader.readProto("")
}
