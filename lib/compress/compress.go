package compress

type Encoder interface {
	Encode(data string) []byte
}
type Decoder interface {
	Decode(data []byte) string
}
