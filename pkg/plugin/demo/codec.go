package main

import (
	"context"

	"github.com/fdingiit/mpl/pkg/plugin/demo/codec"
	"mosn.io/api"
)

type Codec struct {
	exampleStatusMapping codec.StatusMapping

	exampleMatcher codec.Matcher

	proto codec.Proto
}

func (r Codec) NewXProtocol(ctx context.Context) api.XProtocol {
	return &codec.Proto{}
}

func (r Codec) ProtocolName() api.ProtocolName {
	return r.proto.Name()
}

func (r Codec) XProtocol() api.XProtocol {
	return &r.proto
}

func (r Codec) ProtocolMatch() api.ProtocolMatch {
	return r.exampleMatcher.ExampleMatcher
}

func (r Codec) HTTPMapping() api.HTTPMapping {
	return &r.exampleStatusMapping
}

//loader_func_name that go-Plugin use,LoadCodec is default name
func LoadCodec() api.XProtocolCodec {
	return &Codec{}
}
