package codec

import (
	"github.com/fdingiit/mpl/pkg/simple"
	"mosn.io/api"
	"mosn.io/pkg/header"
)

type RequestFrame struct {
	simple.Request

	header.CommonHeader

	// todo, lab3 taskc
}

func (r *RequestFrame) Get(key string) (string, bool) {
	// todo, lab3 taskc
	panic("implement me")
}

func (r *RequestFrame) Set(key, value string) {
	// todo, lab3 taskc
	panic("implement me")
}

func (r *RequestFrame) Add(key, value string) {
	// todo, lab3 taskc
	panic("implement me")
}

func (r *RequestFrame) Del(key string) {
	// todo, lab3 taskc
	panic("implement me")
}

func (r *RequestFrame) Range(f func(key string, value string) bool) {
	// todo, lab3 taskc
	panic("implement me")
}

func (r *RequestFrame) Clone() api.HeaderMap {
	// todo, lab3 taskc
	panic("implement me")
}

func (r *RequestFrame) ByteSize() uint64 {
	// todo, lab3 taskc
	panic("implement me")
}

func (r *RequestFrame) GetRequestId() uint64 {
	panic("implement me later")
}

func (r *RequestFrame) SetRequestId(id uint64) {
	panic("implement me later")
}

func (r *RequestFrame) IsHeartbeatFrame() bool {
	panic("implement me later")
}

func (r *RequestFrame) GetTimeout() int32 {
	// todo, lab3 taskc
	panic("implement me")
}

func (r *RequestFrame) GetStreamType() api.StreamType {
	// todo, lab3 taskc
	panic("implement me")
}

func (r *RequestFrame) GetHeader() api.HeaderMap {
	// todo, lab3 taskc
	panic("implement me")
}

func (r *RequestFrame) GetData() api.IoBuffer {
	// todo, lab3 taskc
	panic("implement me")
}

func (r *RequestFrame) SetData(data api.IoBuffer) {
	// todo, lab3 taskc
	panic("implement me")
}

type ResponseFrame struct {
	// todo, lab3 taskc
	simple.Response

	header.CommonHeader
}

func (r *ResponseFrame) Get(key string) (string, bool) {
	// todo, lab3 taskc
	panic("implement me")
}

func (r *ResponseFrame) Set(key, value string) {
	// todo, lab3 taskc
	panic("implement me")
}

func (r *ResponseFrame) Add(key, value string) {
	// todo, lab3 taskc
	panic("implement me")
}

func (r *ResponseFrame) Del(key string) {
	// todo, lab3 taskc
	panic("implement me")
}

func (r *ResponseFrame) Range(f func(key string, value string) bool) {
	// todo, lab3 taskc
	panic("implement me")
}

func (r *ResponseFrame) Clone() api.HeaderMap {
	// todo, lab3 taskc
	panic("implement me")
}

func (r *ResponseFrame) ByteSize() uint64 {
	// todo, lab3 taskc
	panic("implement me")
}

func (r *ResponseFrame) GetRequestId() uint64 {
	panic("implement me later")
}

func (r *ResponseFrame) SetRequestId(id uint64) {
	panic("implement me later")
}

func (r *ResponseFrame) IsHeartbeatFrame() bool {
	panic("implement me later")
}

func (r *ResponseFrame) GetTimeout() int32 {
	// todo, lab3 taskc
	panic("implement me")
}

func (r *ResponseFrame) GetStreamType() api.StreamType {
	// todo, lab3 taskc
	panic("implement me")
}

func (r *ResponseFrame) GetHeader() api.HeaderMap {
	// todo, lab3 taskc
	panic("implement me")
}

func (r *ResponseFrame) GetData() api.IoBuffer {
	// todo, lab3 taskc
	panic("implement me")
}

func (r *ResponseFrame) SetData(data api.IoBuffer) {
	// todo, lab3 taskc
	panic("implement me")
}

func (r *ResponseFrame) GetStatusCode() uint32 {
	// todo, lab3 taskc
	panic("implement me")
}

type XProtocolSimple struct {
}
