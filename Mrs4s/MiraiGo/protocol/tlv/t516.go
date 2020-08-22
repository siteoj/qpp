package tlv

import "Mrs4s/MiraiGo/binary"

func T516() []byte {
	return binary.NewWriterF(func(w *binary.Writer) {
		w.WriteUInt16(0x516)
		w.WriteTlv(binary.NewWriterF(func(w *binary.Writer) {
			w.WriteUInt32(0)
		}))
	})
}
