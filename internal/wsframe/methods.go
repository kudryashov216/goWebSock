package wsframe

import (
	"encoding/binary"
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

func (f *Frame) Unmask(data *[]byte) (*[]byte, error) {

	if len(*data) < 2 {
		return nil, errors.New("Insufficient bytes for the header.")
	}

	masked := (*data)[1]&0x80 != 0
	payloadLen := int((*data)[1] & 0x7F)
	headerLen := 2

	switch payloadLen {
	case 126:

		if len((*data)) < 4 {
			return nil, errors.New("The extended length byte is missing.")
		}

		payloadLen = int(binary.BigEndian.Uint16((*data)[2:4]))

		headerLen = 4

	case 127:

		if len((*data)) < 10 {
			return nil, errors.New("The byte length of 64 bits is missing.")
		}

		payloadLen = int(binary.BigEndian.Uint64((*data)[2:10]))

		headerLen = 10

	}

	payloadStart := headerLen
	if masked {

		if len((*data)) < headerLen+4 {
			return nil, errors.New("The byte of the masking key is missing.")
		}

		maskKey := (*data)[headerLen : headerLen+4]
		payloadStart += 4

		if len((*data)) < payloadStart+payloadLen {
			return nil, errors.New("The payload is incomplete: read another byte from the socket.")
		}

		for i := 0; i < payloadLen; i++ {
			(*data)[payloadStart+i] ^= maskKey[i%4]
		}

	} else if len((*data)) < payloadStart+payloadLen {
		return nil, errors.New("The payload is incomplete: read another byte from the socket.")
	}

	payload := (*data)[payloadStart : payloadStart+payloadLen]

	return &payload, nil

}

func (r *Reader) Read(conn net.Conn, buff *[]byte) (code byte, n int, err error) {

	if r.IsClosed() {
		return 0xFF, 0, errors.New("It was not possible to read the frame because the reader is closed.")
	}

	n, err = conn.Read(*buff)

	opcode := getOpcodeToFrame(buff)

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
