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

func (pkt *Packet) decode(payload []byte) []byte {
	xorKey := int(pkt.tickNo % 256)
	decodedPld := make([]byte, len(payload))
	for i, byteVal := range payload {
		decodedPld[i] = byte(byteVal) ^ byte(xorKey)
	}
	return decodedPld
}

func (pkt *Packet) processPayload(payload []byte) interface{} {
    if pkt.pkttype == 0xcf && pkt.pktsubtype == 0x01 { // GPS Packet
        pkt.label = "GPS"
        // fmt.Println(fmt.Sprintf("%d - GPS pkt len: %d", tickNo, pktlen))
        payload = pkt.decode(payload)
        pldObj := NewGPSPayload(payload)
        if len(pldObj.Data) > 0 {
            return pldObj
        }
    } else if pkt.pkttype == 0xda && pkt.pktsubtype == 0xf1 { // Motor Packet
        pkt.label = "MOTOR"
        // fmt.Println(fmt.Sprintf("%d - Motor pkt len: %d", tickNo, pktlen))
        payload = pkt.decode(payload)
        pldObj := NewMotorPayload(payload)
        if len(pldObj.Data) > 0 {
            return pldObj
        }
    } else if pkt.pkttype == 0xc6 && pkt.pktsubtype == 0x0d { // Home Point Packet
        pkt.label = "HP"
        // fmt.Println(fmt.Sprintf("%d - HP pkt len: %d", tickNo, pktlen))
        payload = pkt.decode(payload)
        pldObj := NewHPPayload(payload)
        if len(pldObj.Data) > 0 {
            return pldObj
        }
    } else if pkt.pkttype == 0x98 && pkt.pktsubtype == 0x00 { // Remote Control Packet
        pkt.label = "RC"
        // fmt.Println(fmt.Sprintf("%d - RC pkt len: %d", tickNo, pktlen))
        payload = pkt.decode(payload)
        pldObj := NewRCPayload(payload)
        if len(pldObj.data) > 0 {
            return pldObj
        }
    } else if pkt.pkttype == 0xc1 && pkt.pktsubtype == 0x2b { // Tablet Location Packet
        pkt.label = "TABLET"
        // fmt.Println(fmt.Sprintf("%d - TABLET pkt len: %d", tickNo, pktlen))
        payload = pkt.decode(payload)
        pldObj := NewTabletLocPayload(payload)
        if len(pldObj.data) > 0 {
            return pldObj
        }
    } else if pkt.pkttype == 0x1e && pkt.pktsubtype == 0x12 { // Battery Packet
        pkt.label = "BATTERY"
        // fmt.Println(fmt.Sprintf("%d - BATTERY pkt len: %d", tickNo, pktlen))
        payload = pkt.decode(payload)
        pldObj := NewBatteryPayload(payload)
        if len(pldObj.data) > 0 {
            return pldObj
        }
    } else if pkt.pkttype == 0x2c && pkt.pktsubtype == 0x34 { // Gimbal Packet
        pkt.label = "GIMBAL"
        // fmt.Println(fmt.Sprintf("%d - GIMBAL pkt len: %d", tickNo, pktlen))
        payload = pkt.decode(payload)
        pldObj := NewGimbalPayload(payload)
        if len(pldObj.Data) > 0 {
            return pldObj
        }
    } else if pkt.pkttype == 0x2A && pkt.pktsubtype == 0x0C { // Flight Status Packet
        pkt.label = "FLIGHT STAT"
        // fmt.Println(fmt.Sprintf("%d - FLIGHT STAT pkt len: %d", tickNo, pktlen))
        payload = pkt.decode(payload)
        pldObj := NewFlightStatPayload(payload)
        if len(pldObj.data) > 0 {
            return pldObj
        }
    } else if pkt.pkttype == 0x44 && pkt.pktsubtype == 0x11 { // Advanced Battery Packet
        pkt.label = "ADV BATTERY"
        // fmt.Println(fmt.Sprintf("%d - ADV BATTERY pkt len: %d", tickNo, pktlen))
        payload = pkt.decode(payload)
        pldObj := NewAdvBatteryPayload(payload)
        if len(pldObj.data) > 0 {
            return pldObj
        }
    }
    return nil
}

