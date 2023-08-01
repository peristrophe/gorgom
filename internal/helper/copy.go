package helper

import (
	"bytes"
	"encoding/gob"
)

func DeepCopy(src interface{}, dst interface{}) (err error) {
	b := new(bytes.Buffer)
	enc := gob.NewEncoder(b)
	err = enc.Encode(src)
	if err != nil {
		return err
	}
	dec := gob.NewDecoder(b)
	err = dec.Decode(dst)
	if err != nil {
		return err
	}
	return nil
}
