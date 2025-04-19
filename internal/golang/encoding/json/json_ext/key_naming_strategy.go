package json_ext

import "github.com/iancoleman/strcase"

type KeyOption struct {
	// `json:"key,prefix:app_"`
	// only marshal set tagged jsonKey a prefix
	TaggedKeyPrefix *bool

	// to give no-jsonTag-field a name(embedType will be ignored)
	// Unmarshal Set fieldName From jsonKey
	// Marshal Set fieldName To jsonKey
	NoTaggedKeyNaming NamingStrategy
}

type NamingStrategy interface {
	MarshalToJSONKey(fieldName string) (jsonKey string)
	UnmarshalToFieldName(jsonKey string) (fieldName string)
}

type CustomizeNaming struct {
	FieldNameToJSONKey func(string) string
	JSONKeyToFieldName func(string) string
}

func (n CustomizeNaming) MarshalToJSONKey(fieldName string) string {
	return n.FieldNameToJSONKey(fieldName)
}

func (n CustomizeNaming) UnmarshalToFieldName(jsonKey string) string {
	return n.JSONKeyToFieldName(jsonKey)
}

type normalizedNaming struct {
	CamelCase          *bool // camelCase
	SnakeCase          *bool // snake_case
	KebabCase          *bool // kebab-case
	ScreamingSnakeCase *bool // SCREAMING_SNAKE_CASE
	PascalCase         *bool // GolangExportedFieldName, ByDefault
}

type CamelCaseNaming struct{}

func (n CamelCaseNaming) MarshalToJSONKey(fieldName string) string {
	return strcase.ToLowerCamel(fieldName)
}

func (n CamelCaseNaming) UnmarshalToFieldName(jsonKey string) string {
	return strcase.ToCamel(jsonKey)
}
