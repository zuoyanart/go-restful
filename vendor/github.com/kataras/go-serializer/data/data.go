package data

const (
	// ContentType the custom key for the serializer, when used inside iris, Q web frameworks or simply net/http
	ContentType = "application/octet-stream"
)

// Serializer the serializer which renders a binary, data
type Serializer struct {
	config Config
}

// New returns a new data serializer
func New(cfg ...Config) *Serializer {
	c := Config{} // I know it's just empty for now
	if len(cfg) > 0 {
		c = cfg[0]
	}
	return &Serializer{config: c}
}

// Serialize accepts the 'object' value and converts it to bytes in order to be 'renderable'
// implements the go-serializer.Serializer interface
func (e *Serializer) Serialize(val interface{}, options ...map[string]interface{}) ([]byte, error) {
	return val.([]byte), nil
}
