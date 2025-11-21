package decode

import (
	"BinaryDecoder/src/common/util"
	"encoding/json"
	"log"

	"github.com/fxamacker/cbor/v2"
	"github.com/vmihailenco/msgpack/v5"
	"go.mongodb.org/mongo-driver/bson"
)

func messagePackDecode(msg []byte) (map[string]interface{}, error) {

	// MessagePack
	var obj map[string]interface{}
	if err := msgpack.Unmarshal(msg, &obj); err != nil {
		log.Fatalf("MessagePack parse error: %v", err)
		return obj, err
	}
	return obj, nil
}

func cborDecode(msg []byte) (map[string]interface{}, error) {

	// CBOR
	var obj map[string]interface{}
	var result interface{}
	if err := cbor.Unmarshal(msg, &result); err != nil {
		log.Fatalf("CBOR parse error: %v", err)
		return obj, err
	}
	result = util.Convert(result)
	j, _ := json.Marshal(&result)
	json.Unmarshal(j, &obj)
	return obj, nil
}

func bsonDecode(msg []byte) (map[string]interface{}, error) {

	// BSON
	var obj map[string]interface{}
	if err := bson.Unmarshal(msg, &obj); err != nil {
		log.Fatalf("BSON parse error: %v", err)
		return obj, err
	}
	return obj, nil
}

func jsonDecode(msg []byte) (map[string]interface{}, error) {

	// JSON
	var obj map[string]interface{}
	if err := json.Unmarshal(msg, &obj); err != nil {
		log.Fatalf("JSON parse error: %v", err)
		return obj, err
	}

	return obj, nil
}
