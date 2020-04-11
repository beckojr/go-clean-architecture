package adapter

// Context to be implemented by the web framework
type Context interface {
	Bind(i interface{}) error
	JSON(code int, i interface{}) error
	String(code int, s string) error
	Param(s string) string
	NoContent(code int) error
}
