package item
import "bytes"

type File interface {
	Name() string
	Type() string
	Data() bytes.Buffer
}