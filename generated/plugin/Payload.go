// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package plugin

import (
	flatbuffers "github.com/google/flatbuffers/go"
	"strconv"

	generated__plugin__client "github.com/siemens/wfx/generated/plugin/client"
)

type Payload byte

const (
	PayloadNONE                             Payload = 0
	Payloadgenerated_plugin_client_Request  Payload = 1
	Payloadgenerated_plugin_client_Response Payload = 2
)

var EnumNamesPayload = map[Payload]string{
	PayloadNONE:                             "NONE",
	Payloadgenerated_plugin_client_Request:  "generated_plugin_client_Request",
	Payloadgenerated_plugin_client_Response: "generated_plugin_client_Response",
}

var EnumValuesPayload = map[string]Payload{
	"NONE":                             PayloadNONE,
	"generated_plugin_client_Request":  Payloadgenerated_plugin_client_Request,
	"generated_plugin_client_Response": Payloadgenerated_plugin_client_Response,
}

func (v Payload) String() string {
	if s, ok := EnumNamesPayload[v]; ok {
		return s
	}
	return "Payload(" + strconv.FormatInt(int64(v), 10) + ")"
}

type PayloadT struct {
	Type  Payload
	Value interface{}
}

func (t *PayloadT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	switch t.Type {
	case Payloadgenerated_plugin_client_Request:
		return t.Value.(*generated__plugin__client.RequestT).Pack(builder)
	case Payloadgenerated_plugin_client_Response:
		return t.Value.(*generated__plugin__client.ResponseT).Pack(builder)
	}
	return 0
}

func (rcv Payload) UnPack(table flatbuffers.Table) *PayloadT {
	switch rcv {
	case Payloadgenerated_plugin_client_Request:
		var x generated__plugin__client.Request
		x.Init(table.Bytes, table.Pos)
		return &PayloadT{Type: Payloadgenerated_plugin_client_Request, Value: x.UnPack()}
	case Payloadgenerated_plugin_client_Response:
		var x generated__plugin__client.Response
		x.Init(table.Bytes, table.Pos)
		return &PayloadT{Type: Payloadgenerated_plugin_client_Response, Value: x.UnPack()}
	}
	return nil
}
