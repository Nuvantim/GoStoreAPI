package response

type Response[T any] struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

type Failed struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

func Error(msg, err string) Failed {
	return Failed{
		Success: false,
		Message: msg,
		Error:   err,
	}
}
func Pass[T any](msg string, result T) Response[T] {
	return Response[T]{
		Success: true,
		Message: msg,
		Data:    result,
	}
}
