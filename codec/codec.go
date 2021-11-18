package codec

import "io"

type Header struct {
	ServiceMethod string // 服务名.方法名
	Seq           uint64 // 请求序号
	Error         string
}

type Codec interface {
	io.Closer
	ReadHeader(*Header) error
	ReadBody(interface{}) error
	Write(*Header, interface{}) error
}
