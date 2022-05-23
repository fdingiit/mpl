package test

import (
	"context"
	"testing"

	"github.com/fdingiit/mpl/pkg/plugin/simple/codec"
	"github.com/fdingiit/mpl/pkg/simple"
	"github.com/stretchr/testify/assert"
	"mosn.io/api"
	"mosn.io/pkg/buffer"
	"mosn.io/pkg/header"
)

func Test_Lab2_TaskA_Interface(t *testing.T) {
	var c api.XProtocolCodec
	var x api.XProtocol

	c = codec.CodecSimple{}
	x = c.NewXProtocol(context.TODO())

	t.Logf("codec: %+v", c)
	t.Logf("xprotocol: %+v", x)
}

func Test_Lab2_TaskA_LoadCodec(t *testing.T) {
	var c api.XProtocolCodec

	c = codec.LoadCodec()
	t.Logf("codec: %+v", c)
}

func Test_Lab2_TaskB_ProtocolName(t *testing.T) {
	tests := []struct {
		name string
		want api.ProtocolName
	}{
		{
			name: "",
			want: codec.ProtocolName,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			codec := codec.LoadCodec()
			xp := codec.NewXProtocol(context.TODO())

			if !assert.Equal(t, tt.want, xp.Name()) {
				t.Errorf("[failed] incorrect protocol name: %+v", xp.Name())
				t.FailNow()
			}

			if !assert.Equal(t, tt.want, codec.ProtocolName()) {
				t.Errorf("[failed] incorrect protocol name: %+v", codec.ProtocolName())
				t.FailNow()
			}
		})
	}
}

func Test_Lab2_TaskB_ProtocolMatch(t *testing.T) {
	tests := []struct {
		name string
		data []byte
		want api.MatchResult
	}{
		{
			name: "",
			data: []byte{},
			want: api.MatchAgain,
		},
		{
			name: "notenough",
			data: []byte{},
			want: api.MatchAgain,
		},
		{
			name: "",
			data: []byte("12345shangshandalaohulaohumeidazhaodadaoxiaosongshu"),
			want: api.MatchFailed,
		},
		{
			name: "",
			data: []byte("00000328RQ0tPK6UhVeIHb2hrsedxXMJHw         010005010<timestamp>1648811583</timestamp><serial_no>12345</serial_no><currency>2</currency><amount>100</amount><unit>0</unit><out_bank_id>2</out_bank_id><out_account_id>1234567899321</out_account_id><in_bank_id>2</in_bank_id><in_account_id>3211541298661</in_account_id><notes></notes>"),
			want: api.MatchSuccess,
		},
		{
			name: "",
			data: []byte("00000156RS0665db818fa5ef08e9f10ec77d76b9a0e010005010<timestamp>1648811583</timestamp><serial_no>12345</serial_no><err_code>0</err_code><message>ok</message>"),
			want: api.MatchSuccess,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := codec.CodecSimple{}
			got := c.ProtocolMatch()(tt.data)
			if !assert.Equal(t, tt.want, got) {
				t.Errorf("[failed] incorrect protocol match")
				t.FailNow()
			}
		})
	}
}

func Test_Lab2_TaskC_XFrame(t *testing.T) {
	buf := []byte("this is magic")

	var reqF interface{}
	reqF = &codec.RequestFrame{}
	req, ok := reqF.(api.XFrame)
	if !assert.True(t, ok) {
		t.Errorf("[failed] should implement api xframe")
		t.FailNow()
	}
	if !assert.Equal(t, api.Request, req.GetStreamType()) {
		t.Errorf("[failed] value mismatch, got: %+v, wanted: %+v", req.GetStreamType(), api.Request)
		t.FailNow()
	}
	_, ok = req.GetHeader().(api.XFrame)
	if !assert.True(t, ok) {
		t.Errorf("[failed] should implement api xframe")
		t.FailNow()
	}
	req.SetData(buffer.NewIoBufferBytes(buf))
	if !assert.Equal(t, buf, req.GetData().Bytes()) {
		t.Errorf("[failed] value mismatch, got: %+v, wanted: %+v", req.GetData().String(), string(buf))
		t.FailNow()
	}

	var rspF interface{}
	rspF = &codec.ResponseFrame{}
	rsp, ok := rspF.(api.XRespFrame)
	if !assert.True(t, ok) {
		t.Errorf("[failed] should implement api xframe-resp")
		t.FailNow()
	}
	if !assert.Equal(t, api.Response, rsp.GetStreamType()) {
		t.Errorf("[failed] value mismatch, got: %+v, wanted: %+v", req.GetStreamType(), api.Request)
		t.FailNow()
	}
	_, ok = rsp.GetHeader().(api.XFrame)
	if !assert.True(t, ok) {
		t.Errorf("[failed] should implement api xframe")
		t.FailNow()
	}
	rsp.SetData(buffer.NewIoBufferBytes(buf))
	if !assert.Equal(t, buf, rsp.GetData().Bytes()) {
		t.Errorf("[failed] value mismatch, got: %+v, wanted: %+v", req.GetData().String(), string(buf))
		t.FailNow()
	}
}

func Test_Lab2_TaskC_Codec_Encode(t *testing.T) {
	type args struct {
		ctx   context.Context
		model interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    api.IoBuffer
		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx: context.TODO(),
				model: &codec.RequestFrame{
					Request: simple.Request{
						Header: simple.Header{
							TotalLength: 328,
							Type:        "RQ",
							PageMark:    0,
							Checksum:    "tPK6UhVeIHb2hrsedxXMJHw         ",
							ServiceCode: 1000501,
							Reserved:    0,
						},
						UnixTimestamp: 1648811583,
						SerialNo:      12345,
						Currency:      2,
						Amount:        100,
						Unit:          0,
						OutBankId:     2,
						OutAccountId:  1234567899321,
						InBankId:      2,
						InAccountId:   3211541298661,
						Notes:         "",
					},
					CommonHeader: header.CommonHeader{
						"total_length": "328",
						"type":         "RQ",
						"page_mark":    "0",
						"checksum":     "tPK6UhVeIHb2hrsedxXMJHw         ",
						"service_code": "1000501",
						"reserved":     "0",
					},
				},
			},
			want: buffer.NewIoBufferBytes([]byte("00000328RQ0tPK6UhVeIHb2hrsedxXMJHw         010005010<timestamp>1648811583</timestamp><serial_no>12345</serial_no><currency>2</currency><amount>100</amount><unit>0</unit><out_bank_id>2</out_bank_id><out_account_id>1234567899321</out_account_id><in_bank_id>2</in_bank_id><in_account_id>3211541298661</in_account_id><notes></notes>")),
		},
		{
			name: "",
			args: args{
				ctx: context.TODO(),
				model: &codec.ResponseFrame{
					Response: simple.Response{
						Header: simple.Header{
							TotalLength: 156,
							Type:        "RS",
							PageMark:    0,
							Checksum:    "665db818fa5ef08e9f10ec77d76b9a0e",
							ServiceCode: 1000501,
							Reserved:    0,
						},
						UnixTimestamp: 1648811583,
						SerialNo:      12345,
						ErrCode:       0,
						Message:       "ok",
					},
					CommonHeader: header.CommonHeader{
						"total_length": "156",
						"type":         "RS",
						"page_mark":    "0",
						"checksum":     "665db818fa5ef08e9f10ec77d76b9a0e",
						"service_code": "1000501",
						"reserved":     "0",
					},
				},
			},
			want: buffer.NewIoBufferBytes([]byte("00000156RS0665db818fa5ef08e9f10ec77d76b9a0e010005010<timestamp>1648811583</timestamp><serial_no>12345</serial_no><err_code>0</err_code><message>ok</message>")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			X := codec.XProtocolSimple{}
			got, err := X.Encode(tt.args.ctx, tt.args.model)
			if !assert.Nil(t, err) {
				t.Errorf("[failed] Encode() error = %v", err)
				t.FailNow()
			}
			if !assert.Equal(t, tt.want.String(), got.String()) {
				t.Errorf("[failed] value mismatch, got: %+v, wanted: %+v", tt.want.String(), got.String())
				t.FailNow()
			}
		})
	}
}

func Test_Lab2_TaskC_Codec_Decode(t *testing.T) {
	ending := []byte("this is magic")
	buf := buffer.NewIoBufferBytes([]byte("00000328RQ0tPK6UhVeIHb2hrsedxXMJHw         010005010<timestamp>1648811583</timestamp><serial_no>12345</serial_no><currency>2</currency><amount>100</amount><unit>0</unit><out_bank_id>2</out_bank_id><out_account_id>1234567899321</out_account_id><in_bank_id>2</in_bank_id><in_account_id>3211541298661</in_account_id><notes></notes>"))
	buf.Append([]byte("00000156RS0665db818fa5ef08e9f10ec77d76b9a0e010005010<timestamp>1648811583</timestamp><serial_no>12345</serial_no><err_code>0</err_code><message>ok</message>"))
	buf.Append(ending)

	type args struct {
		ctx  context.Context
		data api.IoBuffer
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "req",
			args: args{
				ctx:  context.TODO(),
				data: buf,
			},
			want: &codec.RequestFrame{
				Request: simple.Request{
					Header: simple.Header{
						TotalLength: 328,
						Type:        "RQ",
						PageMark:    0,
						Checksum:    "tPK6UhVeIHb2hrsedxXMJHw         ",
						ServiceCode: 1000501,
						Reserved:    0,
					},
					UnixTimestamp: 1648811583,
					SerialNo:      12345,
					Currency:      2,
					Amount:        100,
					Unit:          0,
					OutBankId:     2,
					OutAccountId:  1234567899321,
					InBankId:      2,
					InAccountId:   3211541298661,
					Notes:         "",
				},
				CommonHeader: header.CommonHeader{
					"total_length": "328",
					"type":         "RQ",
					"page_mark":    "0",
					"checksum":     "tPK6UhVeIHb2hrsedxXMJHw         ",
					"service_code": "1000501",
					"reserved":     "0",
				},
			},
			wantErr: false,
		},
		{
			name: "rsp",
			args: args{
				ctx:  context.TODO(),
				data: buf,
			},
			want: &codec.ResponseFrame{
				Response: simple.Response{
					Header: simple.Header{
						TotalLength: 156,
						Type:        "RS",
						PageMark:    0,
						Checksum:    "665db818fa5ef08e9f10ec77d76b9a0e",
						ServiceCode: 1000501,
						Reserved:    0,
					},
					UnixTimestamp: 1648811583,
					SerialNo:      12345,
					ErrCode:       0,
					Message:       "ok",
				},
				CommonHeader: header.CommonHeader{
					"total_length": "156",
					"type":         "RS",
					"page_mark":    "0",
					"checksum":     "665db818fa5ef08e9f10ec77d76b9a0e",
					"service_code": "1000501",
					"reserved":     "0",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			X := codec.XProtocolSimple{}
			got, err := X.Decode(tt.args.ctx, tt.args.data)
			if !assert.Nil(t, err) {
				t.Errorf("[failed] Decode() error = %v", err)
				t.FailNow()
			}

			gf, _ := got.(api.XFrame)
			wf, _ := tt.want.(api.XFrame)

			// check headers
			if !assert.NotNil(t, gf.GetHeader()) {
				t.Errorf("[failed] null header")
				t.FailNow()
			}
			wf.GetHeader().Range(func(key string, value string) bool {
				if v, ok := gf.GetHeader().Get(key); !ok {
					t.Errorf("[failed] cannot get key %+v in header", key)
					t.FailNow()
				} else if v != value {
					t.Errorf("[failed] value mismatch for key: %+v, got: %+v, wanted: %+v", key, v, value)
					t.FailNow()
				}
				return true
			})

			// check body
			if !assert.NotNil(t, gf.GetData()) {
				t.Errorf("[failed] null body")
				t.FailNow()
			}
			if !assert.Equal(t, tt.args.data.String(), gf.GetData().String()) {
				t.Errorf("[failed] body mismatch, got: %+v, wanted: %+v", gf.GetData().String(), wf.GetData().String())
				t.FailNow()
			}
		})
	}

	if !assert.Equal(t, len(ending), buf.Len()) {
		t.Errorf("[failed] should read all data except ending, current buf: %+v", buf.String())
		t.FailNow()
	}
}
