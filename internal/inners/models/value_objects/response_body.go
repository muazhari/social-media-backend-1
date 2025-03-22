package value_objects

type ResponseBody[T any] struct {
	Message string `json:"message"`
	Data    T      `json:"data"`
	Errors  any    `json:"errors"`
}
