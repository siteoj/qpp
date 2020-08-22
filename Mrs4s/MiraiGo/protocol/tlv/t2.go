package tlv

import "Mrs4s/MiraiGo/binary"

func T2(result string, sign []byte) []byte {
	return binary.NewWriterF(func(w *binary.Writer) {
		w.WriteUInt16(0x02)
		w.WriteTlv(binary.NewWriterF(func(w *binary.Writer) {
			w.WriteUInt16(0)
			w.WriteStringShort(result)
			w.WriteTlv(sign)
		}))
	})
}
