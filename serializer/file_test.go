package serializer_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/techschool/pcbook/pb"
	"gitlab.com/techschool/pcbook/sample"
	"gitlab.com/techschool/pcbook/serializer"
	"google.golang.org/protobuf/proto"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()

	binaryFile := "../tmp/laptop.bin"
	jsonFile := "../tmp/laptop.json"

	laptop1 := sample.NewLaptop()

	err := serializer.WriteProtobufToBinaryFile(laptop1, binaryFile)
	require.NoError(t, err)

	err = serializer.WriteProtobufToJSONFile(laptop1, jsonFile)
	require.NoError(t, err)

	laptop2 := &pb.Laptop{}
	err = serializer.ReadProtobufFromBinaryFile(binaryFile, laptop2)
	require.NoError(t, err)

	require.True(t, proto.Equal(laptop1, laptop2))
}

func TestWriteProtobufToBinaryFile_Error(t *testing.T) {
	t.Parallel()

	// Test with invalid file path
	laptop := sample.NewLaptop()
	err := serializer.WriteProtobufToBinaryFile(laptop, "/invalid-path/laptop.bin")
	require.Error(t, err)
}

func TestReadProtobufFromBinaryFile_Error(t *testing.T) {
	t.Parallel()

	// Test with invalid file path
	laptop := &pb.Laptop{}
	err := serializer.ReadProtobufFromBinaryFile("/invalid-path/laptop.bin", laptop)
	require.Error(t, err)

	// Test with corrupted data
	err = os.WriteFile("../tmp/corrupted.bin", []byte("corrupted data"), 0644)
	require.NoError(t, err)
	err = serializer.ReadProtobufFromBinaryFile("../tmp/corrupted.bin", laptop)
	require.Error(t, err)
}

func TestWriteProtobufToJSONFile_Error(t *testing.T) {
	t.Parallel()

	// Test with invalid file path
	laptop := sample.NewLaptop()
	err := serializer.WriteProtobufToJSONFile(laptop, "/invalid-path/laptop.json")
	require.Error(t, err)
}
