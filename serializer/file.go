package serializer

import (
	"fmt"
	"os"

	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/runtime/protoiface"
)

func WritesProtoBufToBinary(message proto.Message, filename string) error {
	data, err := proto.Marshal(message)
	if err != nil {
		return fmt.Errorf("marshaling error: %w", err)
	}

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("creating file error: %w", err)
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("writing to file error: %w", err)
	}

	fmt.Println("Data has been written to", filename)
	return nil
}

func ReadProtobufFromBinaryFile(filename string, message proto.Message) error {
	// data, err := ioutil.ReadFile(filename)
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("cannot read file: %w", err)
	}

	err = proto.Unmarshal(data, message)
	if err != nil {
		return fmt.Errorf("cannot unmarshal binary data: %w", err)
	}

	return nil
}

// WriteProtobufToJSONFile writes protocol buffer message to JSON file
func WriteProtobufToJSONFile(messageOk proto.Message, filename string) error {
	// Marshal the message to JSON format
	data, err := ProtobufToJson(messageOk)
	if err != nil {
		return fmt.Errorf("cannot marshal proto message to JSON: %w", err)
	}

	err = os.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		return fmt.Errorf("cannot write JSON data to file: %w", err)
	}

	return nil
}

func ProtobufToJson(message proto.Message) (string, error) {

	protoMsg := message.ProtoReflect().Interface()

	// Convert protoMsg to protoiface.MessageV1
	protoMessage, ok := protoMsg.(protoiface.MessageV1)
	if !ok {
		return "", fmt.Errorf("cannot convert proto message to protoiface.MessageV1")
	}

	marshaler := jsonpb.Marshaler{
		EnumsAsInts:  false,
		EmitDefaults: true,
		Indent:       "  ",
		OrigName:     true,
	}
	json, err := marshaler.MarshalToString(protoMessage)
	if err != nil {
		return "", err
	}
	return json, nil
}
