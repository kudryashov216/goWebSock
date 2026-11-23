package wsframe

import (
	"errors"
	"net"
)

func NewCloser() CloserFrame {

	return &Closer{}

}

func NewFrame(buff *[]byte) WSFrame {

	return &Frame{Buff: buff, CloserFrame: NewCloser()}

}

func NewReader(frame WSFrame) ReaderFrame {

	return &Reader{WSFrame: frame}

}

func NewBuilder(frame WSFrame) BuilderFrame {

	return &Builder{WSFrame: frame}

}

func NewReaderBuilder(frame WSFrame) ReaderBuilder {

	return ReaderBuilder{
		ReaderFrame:  NewReader(frame),
		BuilderFrame: NewBuilder(frame),
	}

}

func getOpcodeToFrame(buff *[]byte) byte {

	if len((*buff)) == 0 {
		return 0xFF
	}

	opcode := (*buff)[0] & 0x0F

	return opcode

}

func (r *Reader) Read(conn net.Conn, buff *[]byte) (code byte, n int, err error) {

	if r.IsClosed() {
		return 0xFF, 0, errors.New("It was not possible to read the frame because the reader is closed.")
	}

	n, err = conn.Read(*buff)

	opcode := getOpcodeToFrame(buff)

	switch opcode {
	case 0x08:
		frame := NewFrame(buff)
		builder := NewBuilder(frame)
		defer builder.Close()

		data, _ := builder.Build(opcode, buff)
		conn.Write(*data)
		return 0xFF, 0, nil
	}

	return opcode, n, err

}

func (b *Builder) Build(opcode byte, buff *[]byte) (data *[]byte, err error) {

	if b.IsClosed() {
		return nil, errors.New("The frame could not be built because the builder is closed.")
	}

	var (
		header, payload []byte
		length          uint64
	)

	payload = *buff

	header = append(header, FIN|opcode)

	switch {

	case len(payload) <= 125:

		length = uint64(len(payload))
		header = append(header, byte(length))

	case len(payload) <= 65535:

		length = uint64(len(payload))
		header = append(header, 126, byte(length>>8), byte(length))

	default:

		length = uint64(len(payload))
		header = append(header, 127)

		for i := 7; i >= 0; i-- {
			header = append(header, byte(length>>(8*i)))
		}

	}

	header = append(header, payload...)

	return &header, nil

}

func (c *Closer) Close() {

	c.closed = true

}
func (c *Closer) IsClosed() bool {

	return c.closed == true

}
