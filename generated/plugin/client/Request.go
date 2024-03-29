// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package client

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type RequestT struct {
	Cookie      uint64       `json:"cookie"`
	Destination string       `json:"destination"`
	Action      Action       `json:"action"`
	Envelope    []*EnvelopeT `json:"envelope"`
	Content     []byte       `json:"content"`
}

func (t *RequestT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	destinationOffset := flatbuffers.UOffsetT(0)
	if t.Destination != "" {
		destinationOffset = builder.CreateString(t.Destination)
	}
	envelopeOffset := flatbuffers.UOffsetT(0)
	if t.Envelope != nil {
		envelopeLength := len(t.Envelope)
		envelopeOffsets := make([]flatbuffers.UOffsetT, envelopeLength)
		for j := 0; j < envelopeLength; j++ {
			envelopeOffsets[j] = t.Envelope[j].Pack(builder)
		}
		RequestStartEnvelopeVector(builder, envelopeLength)
		for j := envelopeLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(envelopeOffsets[j])
		}
		envelopeOffset = builder.EndVector(envelopeLength)
	}
	contentOffset := flatbuffers.UOffsetT(0)
	if t.Content != nil {
		contentOffset = builder.CreateByteString(t.Content)
	}
	RequestStart(builder)
	RequestAddCookie(builder, t.Cookie)
	RequestAddDestination(builder, destinationOffset)
	RequestAddAction(builder, t.Action)
	RequestAddEnvelope(builder, envelopeOffset)
	RequestAddContent(builder, contentOffset)
	return RequestEnd(builder)
}

func (rcv *Request) UnPackTo(t *RequestT) {
	t.Cookie = rcv.Cookie()
	t.Destination = string(rcv.Destination())
	t.Action = rcv.Action()
	envelopeLength := rcv.EnvelopeLength()
	t.Envelope = make([]*EnvelopeT, envelopeLength)
	for j := 0; j < envelopeLength; j++ {
		x := Envelope{}
		rcv.Envelope(&x, j)
		t.Envelope[j] = x.UnPack()
	}
	t.Content = rcv.ContentBytes()
}

func (rcv *Request) UnPack() *RequestT {
	if rcv == nil {
		return nil
	}
	t := &RequestT{}
	rcv.UnPackTo(t)
	return t
}

type Request struct {
	_tab flatbuffers.Table
}

func GetRootAsRequest(buf []byte, offset flatbuffers.UOffsetT) *Request {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Request{}
	x.Init(buf, n+offset)
	return x
}

func FinishRequestBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsRequest(buf []byte, offset flatbuffers.UOffsetT) *Request {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Request{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedRequestBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *Request) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Request) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Request) Cookie() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Request) MutateCookie(n uint64) bool {
	return rcv._tab.MutateUint64Slot(4, n)
}

func (rcv *Request) Destination() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Request) Action() Action {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return Action(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *Request) MutateAction(n Action) bool {
	return rcv._tab.MutateInt8Slot(8, int8(n))
}

func (rcv *Request) Envelope(obj *Envelope, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *Request) EnvelopeLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Request) Content(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *Request) ContentLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Request) ContentBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Request) MutateContent(j int, n byte) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func RequestStart(builder *flatbuffers.Builder) {
	builder.StartObject(5)
}
func RequestAddCookie(builder *flatbuffers.Builder, cookie uint64) {
	builder.PrependUint64Slot(0, cookie, 0)
}
func RequestAddDestination(builder *flatbuffers.Builder, destination flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(destination), 0)
}
func RequestAddAction(builder *flatbuffers.Builder, action Action) {
	builder.PrependInt8Slot(2, int8(action), 0)
}
func RequestAddEnvelope(builder *flatbuffers.Builder, envelope flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(envelope), 0)
}
func RequestStartEnvelopeVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func RequestAddContent(builder *flatbuffers.Builder, content flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(content), 0)
}
func RequestStartContentVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func RequestEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
