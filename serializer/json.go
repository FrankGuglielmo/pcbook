package serializer

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func ProtobufToJSON(message proto.Message) (string, error) {
	marshaler := protojson.MarshalOptions{
		Indent:            "  ",
		UseProtoNames:     true,
		UseEnumNumbers:    false,
		EmitDefaultValues: true,
	}

	jsonBytes, err := marshaler.Marshal(message)
	if err != nil {
		return "Cannot Marshal message to JSON: ", err
	}

	return string(jsonBytes), nil
}
