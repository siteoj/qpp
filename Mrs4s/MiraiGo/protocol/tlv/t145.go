package tlv

import "Mrs4s/MiraiGo/binary"

func T145(guid []byte) []byte {
	return binary.NewWriterF(func(w *binary.Writer) {
		w.WriteUInt16(0x145)
		w.WriteTlv(binary.NewWriterF(func(w *binary.Writer) {
			w.Write(guid)
		}))
	})
}
