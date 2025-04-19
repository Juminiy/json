package json_ext

// Config will never fail and never throw-or-return any error
// if error occurred, it is
// - `encoding/json` definedError or encodeError/decodeError
// - sigs.k8s.io/json definedError
type Config struct {
	KeyOption
	ValueOption
	StreamOption
}
