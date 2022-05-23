package codec

import "mosn.io/api"

type CodecSimple struct {
}

func LoadCodec() api.XProtocolCodec {
	return &CodecSimple{}
}
