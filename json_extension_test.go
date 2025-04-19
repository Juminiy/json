package json

import (
	internaljson "sigs.k8s.io/json/internal/golang/encoding/json"
	internaljsonext "sigs.k8s.io/json/internal/golang/encoding/json/json_ext"
	"testing"
)

var _extConfig = internaljsonext.Config{
	KeyOption: internaljsonext.KeyOption{
		TaggedKeyPrefix:   func(bv bool) *bool { return &bv }(true),
		NoTaggedKeyNaming: internaljsonext.CamelCaseNaming{},
	},
	ValueOption: internaljsonext.ValueOption{
		AnyToBool:         func(bv bool) *bool { return &bv }(true),
		AnyToInt:          func(bv bool) *bool { return &bv }(true),
		AnyToTime:         func(bv bool) *bool { return &bv }(true),
		AnyToTimeByFormat: func(bv bool) *bool { return &bv }(true),
	},
	StreamOption: internaljsonext.StreamOption{},
}

func ExtMarshal(v any) ([]byte, error) {
	return internaljson.MarshalWithExtension(v, _extConfig)
}

func ExtUnmarshal(b []byte, v any) error {
	return internaljson.UnmarshalWithExtension(b, v, _extConfig)
}

func TestJSONExtension(t *testing.T) {
	type tbsTyp struct {
		ValuePub     int64
		AValueName   string
		BValueICanDo string
		GPUWeBought  int64
		UUID         string
	}

	tbs := []byte(`
{
	"ValuePub": 888,
	"aValueSame": "rich-mode-sort",
	"b_value_i_can_do": "how_??>",
	"GPUWeBought": 114564,
	"UUID": "9099999-xx2a"
}
`)
	t.Run("only sigs.k8s.io/json compatible with stdlib", func(tt *testing.T) {
		var tbsVal tbsTyp
		if err := internaljson.UnmarshalStrict(tbs, &tbsVal); err != nil {
			tt.Logf("some error occurs: %s", err.Error())
			return
		}
		t.Logf("%+v", tbsVal)
	})

	t.Run("Extension Unmarshal jsonKey -> FieldName", func(tt *testing.T) {
		var tbsVal tbsTyp
		if err := ExtUnmarshal(tbs, &tbsVal); err != nil {
			tt.Logf("some error occurs: %s", err.Error())
			return
		}
		t.Logf("%+v", tbsVal)
	})
}
