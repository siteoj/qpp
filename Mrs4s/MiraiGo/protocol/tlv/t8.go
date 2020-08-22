package tlv

import "Mrs4s/MiraiGo/binary"

func T8(localId uint32) []byte {
	return binary.NewWriterF(func(w *binary.Writer) {
		w.WriteUInt16(0x8)
		w.WriteTlv(binary.NewWriterF(func(w *binary.Writer) {
			w.WriteUInt16(0)
			w.WriteUInt32(localId)
			w.WriteUInt16(0)
		}))
	})
}
