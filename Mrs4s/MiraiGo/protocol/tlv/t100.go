package tlv

import "Mrs4s/MiraiGo/binary"

func T100() []byte {
	return binary.NewWriterF(func(w *binary.Writer) {
		w.WriteUInt16(0x100)
		w.WriteTlv(binary.NewWriterF(func(w *binary.Writer) {
			w.WriteUInt16(1)
			w.WriteUInt32(5)
			w.WriteUInt32(16)
			w.WriteUInt32(537062409) // Sub app id
			w.WriteUInt32(0)         // App client version
			w.WriteUInt32(34869472)
		}))
	})
}
