package json_ext

// ValueOption
// encoding/json support `json:",string"`
// ValueOption support
// `json:",int"`
// `json:",bool"`
// `json:",time"`
// `json:",time:DateTime"` OR `json:",time:DateTime"` OR `json:",time:RFC3339"` OR ...
type ValueOption struct {
	AnyToInt          *bool
	AnyToBool         *bool
	AnyToTime         *bool
	AnyToTimeByFormat *bool // support Format in time/format.go,eg: time.DateTime, time.ANSIC
}
