package proto

import (
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
	"io"
)

type Encoder struct {
	w io.Writer
}

func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{
		w: w,
	}
}

func (enc *Encoder) Encode(v interface{}) error {
	m, ok := v.(proto.Message)
	if !ok {
		return errors.New("parameter v is not a protobuf message")
	}
	b, err := proto.Marshal(m)
	if err != nil {
		return err
	}
	enc.w.Write(b)
	return nil
}

type Decoder struct {
	r io.Reader
}

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{
		r: r,
	}
}

func (dec *Decoder) Decode(v interface{}) error {
	//proto.Unmarshal()
	return nil
}
