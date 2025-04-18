package json_ext

type KeyOption struct {
	TaggedKeyPrefix   *bool          // `json:"key,prefix:app_"`
	NoTaggedKeyNaming NamingStrategy // no jsonTag give a name
}

type NamingStrategy interface {
	MarshalToJSONKey(fieldName string) (jsonKey string)
	UnmarshalToFieldName(jsonKey string) (fieldName string)
}

type CustomizeNaming struct {
	FieldNameToJSONKey func(string) string
	JSONKeyToFieldName func(string) string
}

func (s CustomizeNaming) MarshalToJSONKey(fieldName string) string {
	return s.FieldNameToJSONKey(fieldName)
}

func (s CustomizeNaming) UnmarshalToFieldName(jsonKey string) string {
	return s.JSONKeyToFieldName(jsonKey)
}

type NormalizedNaming struct {
	CamelCase          *bool // camelCase
	SnakeCase          *bool // snake_case
	KebabCase          *bool // kebab-case
	ScreamingSnakeCase *bool // SCREAMING_SNAKE_CASE
	PascalCase         *bool // GolangExportedFieldName, ByDefault
}
