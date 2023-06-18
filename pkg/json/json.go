package json

import encodingJSON "encoding/json"

type JSON interface {
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
}
type json struct{}

var instance JSON

func Instance() JSON {
	if instance == nil {
		instance = &json{}
	}
	return instance
}

func (u *json) Marshal(v interface{}) ([]byte, error) {
	return encodingJSON.Marshal(v)
}

func (u *json) Unmarshal(data []byte, v interface{}) error {
	return encodingJSON.Unmarshal(data, v)
}
