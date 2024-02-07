package ref

// Ptr は v のポインタを返す
func Ptr[T any](v T) *T {
	return &v
}
