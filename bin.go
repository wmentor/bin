package bin

import (
	"bytes"
	"encoding/binary"
)

func Encode(list ...interface{}) []byte {
	b := new(bytes.Buffer)
	for _, item := range list {

		switch v := item.(type) {
		case string:
			binary.Write(b, binary.BigEndian, []byte(v))
		default:
			binary.Write(b, binary.BigEndian, item)
		}
	}
	return b.Bytes()
}

func Decode(data []byte, list ...interface{}) error {
	reader := bytes.NewReader(data)
	for _, item := range list {
		if err := binary.Read(reader, binary.BigEndian, item); err != nil {
			return err
		}
	}

	return nil
}
