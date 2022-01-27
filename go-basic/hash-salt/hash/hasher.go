package hash

type Hasher interface {
	EncodeWithSalt(value string) (string, error)
	Encode(value string) (string, error)
}
