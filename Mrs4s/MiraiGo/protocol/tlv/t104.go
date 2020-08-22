package tlv

import "Mrs4s/MiraiGo/binary"

func T104(data []byte) []byte {
	return binary.NewWriterF(func(w *binary.Writer) {
		w.WriteUInt16(0x104)
		w.WriteTlv(data)
	})
}
