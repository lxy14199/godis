package protocol

type UnknownErrReply struct{}

var unkonwnErrBytes = []byte("-Err unknown\r\n")

func (r *UnknownErrReply) ToBytes() []byte {
	return unkonwnErrBytes
}

func (r *UnknownErrReply) Error() string {
	return "Err unknown"
}

type ArgNumErrReply struct {
	Cmd string
}

func (r *ArgNumErrReply) ToBytes() []byte {
	return []byte("-ERR wrong number of argument for '" + r.Cmd + "'command\r\n")
}

func (r *ArgNumErrReply) Error() string {
	return "ERR wrong number of argument for '" + r.Cmd + "'command"
}

func MakeArgNumErrReply(cmd string) *ArgNumErrReply {
	return &ArgNumErrReply{
		Cmd: cmd,
	}
}

type SyntaxErrReply struct{}

var syntaxErrBytes = []byte("-Err syntax error\r\n")
var theSyntaxErrReply = &SyntaxErrReply{}

func MakeSyntaxErrReply() *SyntaxErrReply {
	return theSyntaxErrReply
}

func (r *SyntaxErrReply) ToBytes() []byte {
	return syntaxErrBytes
}

func (r *SyntaxErrReply) Error() string {
	return "Err syntax error"
}

type WrongTypeErrReply struct{}

var wrongTypeErrBytes = []byte("-WRONGTYPE Operation against a key hoding the wrong king of value\r\n")

func (r *WrongTypeErrReply) ToBytes() []byte {
	return wrongTypeErrBytes
}

func (r *WrongTypeErrReply) Error() string {
	return "WRONGTYPE Operation against a key hoding the wrong king of value"
}

type ProtoclErrReply struct {
	Msg string
}

func (r *ProtoclErrReply) ToBytes() []byte {
	return []byte("-ERR Protocl error: '" + r.Msg + "'\r\n")
}

func (r *ProtoclErrReply) Error() string {
	return "ERR Protocl error: '" + r.Msg + "'"
}
