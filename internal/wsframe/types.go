package wsframe

import "context"

type Frame struct {
	Context context.Context
	GUID    string
	Data    *[]byte
	NBytes  int
	CloserFrame
}

type ReaderBuilder struct {
	ReaderFrame
	BuilderFrame
}

type Reader struct {
	WSFrame
}

type Builder struct {
	WSFrame
}

type Closer struct {
	closed bool
}
