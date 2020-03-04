package message

type Message interface {
	ToByte() ([]byte, error)
}
