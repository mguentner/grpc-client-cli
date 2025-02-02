package caller

import (
	"github.com/golang/protobuf/proto"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/desc/builder"
	"github.com/jhump/protoreflect/dynamic"
)

var customAnyDescr *desc.MessageDescriptor

func init() {
	md, err := builder.NewMessage("any").
		AddField(builder.NewField("value", builder.FieldTypeBytes())).
		AddField(builder.NewField("err", builder.FieldTypeString())).
		Build()
	if err != nil {
		panic(err)
	}

	customAnyDescr = md
}

// anyResolver resolves types specified in typeURL/@type field of google.protobuf.Any message
// or falls back to the one that just represents Any message with an error field
type anyResolver struct{}

func (a *anyResolver) Resolve(typeURL string) (proto.Message, error) {
	baseResolver := dynamic.AnyResolver(nil)
	m, err := baseResolver.Resolve(typeURL)
	if err == nil {
		return m, nil
	}
	return &anyWrapper{dynamic.NewMessage(customAnyDescr)}, nil
}

type anyWrapper struct {
	*dynamic.Message
}

func (a *anyWrapper) Unmarshal(b []byte) error {
	a.Message.SetFieldByName("value", b)
	a.Message.SetFieldByName("err", "type not found")
	return nil
}

var _ proto.Message = (*anyWrapper)(nil)
