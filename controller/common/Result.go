package common

type H map[string]interface{}

type Result[T any] struct {
	Code int
	Data T
	Msg string
}

func Success[T any]() Result[T] {
	return Result[T]{
		Code: 1,
		Msg: "success",
		Data: *new(T),
	}
}

func SuccessWithData[T any](data T) Result[T] {
	return Result[T]{
		Code: 1,
		Msg: "success",
		Data: data,
	}
}

func Error[T any](msg string) Result[T] {
	return Result[T]{
		Code: 0,
		Msg: msg,
		Data: *new(T),
	}
}
