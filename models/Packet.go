package models

type Packet struct {
	pktlen     int
	header     []byte
	pkttype    byte
	pktsubtype byte
	label      string
	msg        byte
	tickNo     uint32
	payload    interface{}
}

func NewPacket(pktlen int, header []byte, payload []byte) *Packet {
	pkt := &Packet{
		pktlen: pktlen,
		header: header,
		pkttype: header[0] & 0xFF,
		pktsubtype: header[1] & 0xFF,
		msg: header[2],
		tickNo: uint32(header[3]) | uint32(header[4])<<8 | uint32(header[5])<<16 | uint32(header[6])<<24,
	}

	pkt.processPayload(payload)

	return pkt
}