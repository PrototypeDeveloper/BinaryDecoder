package decode

import (
	"BinaryDecoder/src/common/util"
	"encoding/hex"
	"fmt"
)

func DecodeHandler(protocol string, data string) (map[string]interface{}, error) {

	var err error
	var decoded map[string]interface{}

	// Remove space
	data = util.RemoveSpace(data)

	// Convert bytes
	msg, _ := hex.DecodeString(data)

	switch protocol {
	case "MessagePack":
		decoded, err = messagePackDecode(msg)
	case "CBOR":
		decoded, err = cborDecode(msg)
	case "BSON":
		decoded, err = bsonDecode(msg)
	case "JSON":
		decoded, err = jsonDecode(msg)
	default:
		err = fmt.Errorf("unknown binary")
	}

	return decoded, err
}
