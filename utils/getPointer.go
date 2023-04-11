package utils

type any interface{}

func Pointer[T any](in T) (out *T) {
	return &in
}
