package cache

// Output struct
type Output struct {
	Result interface{}
	Error  error
}

// Cache interface
type Cache interface {
	Get(key string, value interface{}) Output
	Set(key string, data interface{}) Output
	Del(key string) Output
	Flush() Output
}
