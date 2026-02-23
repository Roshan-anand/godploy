package lib

// returns the address of the value passed in
func GetValAddrs[T any](val T) *T {
	return &val
}
