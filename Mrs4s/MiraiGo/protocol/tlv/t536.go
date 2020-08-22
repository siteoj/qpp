package tlv

import "Mrs4s/MiraiGo/binary"

func T536(loginExtraData []byte) []byte {
	return binary.NewWriterF(func(w *binary.Writer) {
		w.WriteUInt16(0x536)
		w.WriteTlv(binary.NewWriterF(func(w *binary.Writer) {
			w.Write(loginExtraData)
		}))
	})
}
