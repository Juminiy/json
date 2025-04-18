package json

import "sigs.k8s.io/json/internal/golang/encoding/json/json_ext"

func DecodeExtensionConfig(d *decodeState, config json_ext.Config) {
	d.extensionConfig = &config
}

func EncodeExtensionConfig(e *encodeState, config json_ext.Config) {
	e.extensionConfig = &config
}
