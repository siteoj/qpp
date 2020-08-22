package tlv

import (
	"Mrs4s/MiraiGo/binary"
	"math/rand"
	"time"
)

func T1(uin uint32, ip []byte) []byte {
	if len(ip) != 4 {
		panic("invalid ip")
	}
	return binary.NewWriterF(func(w *binary.Writer) {
		w.WriteUInt16(0x01)
		w.WriteTlv(binary.NewWriterF(func(w *binary.Writer) {
			w.WriteUInt16(1)
			w.WriteUInt32(rand.Uint32())
			w.WriteUInt32(uin)
			w.WriteUInt32(uint32(time.Now().UnixNano() / 1e6))
			w.Write(ip)
			w.WriteUInt16(0)
		}))
	})
}
