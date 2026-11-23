package wsframe

type Frame struct {
	Buff *[]byte
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
