package json

import "sigs.k8s.io/json/internal/golang/encoding/json/json_ext"

/*func DecodeExtensionConfig(d *decodeState, config json_ext.Config) {
	d.extensionConfig = &config
}

func EncodeExtensionConfig(e *encodeState, config json_ext.Config) {
	e.extensionConfig = &config
}*/

func newEncodeStateExtensionConfig(extCfg json_ext.Config) *encodeState {
	e := newEncodeState()
	e.extensionConfig = &extCfg
	return e
}

// MarshalWithExtension
// copy from Marshal add json_ext.Config only
func MarshalWithExtension(v any, extCfg json_ext.Config) ([]byte, error) {
	e := newEncodeStateExtensionConfig(extCfg)
	defer encodeStatePool.Put(e)

	err := e.marshal(v, encOpts{escapeHTML: true})
	if err != nil {
		return nil, err
	}
	buf := append([]byte(nil), e.Bytes()...)

	return buf, nil
}

// TODO:
func NewEncoderWithExtension() {

}

func UnmarshalStrict(b []byte, v any) error {
	return Unmarshal(b, v,
		// options matching UnmarshalCaseSensitivePreserveInts
		CaseSensitive,
		PreserveInts,
		// all strict options
		DisallowDuplicateFields,
		DisallowUnknownFields,
	)
}

func UnmarshalWithExtension(b []byte, v any, extCfg json_ext.Config) error {
	return Unmarshal(b, v,
		// options matching UnmarshalCaseSensitivePreserveInts
		CaseSensitive,
		PreserveInts,
		// all strict options
		DisallowDuplicateFields,
		DisallowUnknownFields,
		func(d *decodeState) {
			d.extensionConfig = &extCfg
		},
	)
}

// TODO:
func NewDecoderWithExtension() {

}
