package tlv

import "Mrs4s/MiraiGo/binary"

func T191(k byte) []byte {
	return binary.NewWriterF(func(w *binary.Writer) {
		w.WriteUInt16(0x191)
		w.WriteTlv(binary.NewWriterF(func(w *binary.Writer) {
			w.WriteByte(k)
		}))
	})
}
