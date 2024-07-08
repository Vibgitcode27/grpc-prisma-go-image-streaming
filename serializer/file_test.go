package serializer_test

import (
	"great/psm"
	"great/sample"
	"great/serializer"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestWriteBinary(t *testing.T) {

	t.Parallel()

	binaryFile := "../test_dump/laptop.bin"

	laptop := sample.NewLaptop()

	err := serializer.WritesProtoBufToBinary(laptop, binaryFile)
	require.NoError(t, err)

	laptop2 := &psm.Laptop{}
	err = serializer.ReadProtobufFromBinaryFile(binaryFile, laptop2)
	require.NoError(t, err)

	err = serializer.WriteProtobufToJSONFile(laptop, "../test_dump/laptop.json")
	require.NoError(t, err)

	require.True(t, proto.Equal(laptop, laptop2), "Test Successful! ")
}
