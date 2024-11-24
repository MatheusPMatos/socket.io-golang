package socket_protocol

import (
	"encoding/json"
	"io"
	"strconv"

	"github.com/doquangtan/gofiber-socket.io/engineio"
)

type PacketType int

const (
	CONNECT PacketType = iota
	DISCONNECT
	EVENT
	ACK
	CONNECT_ERROR
	BINARY_EVENT
	BINARY_ACK
)

func (id PacketType) String() string {
	return strconv.Itoa(int(id))
}

type writer struct {
	t   PacketType
	nps string
	i   int64
	w   io.Writer
}

func (w *writer) Write(p []byte) (int, error) {
	paserData := append([]byte(w.t.String()+w.nps), p...)
	return engineio.WriteByte(w.w, engineio.MESSAGE, paserData)
}

func WriteTo(w io.Writer, t PacketType, nps string, arg ...interface{}) (int64, error) {
	writer := writer{
		t:   t,
		nps: nps,
		w:   w,
	}
	if len(arg) > 0 {
		err := json.NewEncoder(&writer).Encode(arg[0])
		return writer.i, err
	} else {
		_, err := writer.Write([]byte{})
		return writer.i, err
	}
}
