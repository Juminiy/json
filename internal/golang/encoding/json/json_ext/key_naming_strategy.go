package json_ext

type KeyOption struct {
	TaggedKeyPrefix   *bool          // `json:"key,prefix:app_"`
	NoTaggedKeyNaming NamingStrategy // no jsonTag give a named
}

type NamingStrategy interface {
	MarshalToJSONKey(fieldName string) (jsonKey string)
	UnmarshalToFieldName(jsonKey string) (fieldName string)
}

type CustomizeNamingStrategy struct {
	FieldNameToJSONKey func(string) string
	JSONKeyToFieldName func(string) string
}

func (s CustomizeNamingStrategy) MarshalToJSONKey(fieldName string) string {
	return s.FieldNameToJSONKey(fieldName)
}

func (s CustomizeNamingStrategy) UnmarshalToFieldName(jsonKey string) string {
	return s.JSONKeyToFieldName(jsonKey)
}

type NormalizedNamingStrategy struct {
	CamelCase          *bool // camelCase
	SnakeCase          *bool // snake_case
	KebabCase          *bool // kebab-case
	ScreamingSnakeCase *bool // SCREAMING_SNAKE_CASE
	PascalCase         *bool // GolangExportedFieldName, byDefault
}
