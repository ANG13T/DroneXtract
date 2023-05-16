package models

import (
	"encoding/binary"
	"fmt"
	"math"
)

type GimbalPayload struct {
	Fields   []string
	Type     uint8
	Subtype  uint8
	Length   uint8
	Payload  []byte
	Data     map[string]interface{}
}

func NewGimbalPayload(payload []byte) *GimbalPayload {
	g := &GimbalPayload{
		Fields:  []string{"quatW", "quatX", "quatY", "quatZ", "Gimbal:roll", "Gimbal:pitch", "Gimbal:yaw", "rFront", "lFront", "lBack", "rBack", "Gimbal:Xroll", "Gimbal:Xpitch", "Gimbal:Xyaw"},
		Type:    0x2c,
		Subtype: 0x34,
		Length:  0xF7,
		Payload: payload,
		Data:    make(map[string]interface{}),
	}
	g.parse()
	return g
}

func (g *GimbalPayload) parse() {
	data := make(map[string]interface{})

	quatW := math.Float32frombits(binary.LittleEndian.Uint32(g.Payload[78:82]))
	quatX := math.Float32frombits(binary.LittleEndian.Uint32(g.Payload[82:86]))
	quatY := math.Float32frombits(binary.LittleEndian.Uint32(g.Payload[86:90]))
	quatZ := math.Float32frombits(binary.LittleEndian.Uint32(g.Payload[90:94]))

	qGimbal := Quaternion{W: quatW, X: quatX, Y: quatY, Z: quatZ}
	rpy := qGimbal.toEuler()
	data["Gimbal:Xpitch"] = math.Degrees(rpy[0])
	data["Gimbal:Xroll"] = math.Degrees(rpy[1])
	data["Gimbal:Xyaw"] = math.Degrees(rpy[2])

	data["Gimbal:yaw"] = math.Degrees(math.Float32frombits(binary.LittleEndian.Uint32(g.Payload[94:98])))
	data["Gimbal:roll"] = math.Degrees(math.Float32frombits(binary.LittleEndian.Uint32(g.Payload[98:102])))
	data["Gimbal:pitch"] = math.Degrees(math.Float32frombits(binary.LittleEndian.Uint32(g.Payload[102:106])))

	data["rFront"] = int16(binary.LittleEndian.Uint16(g.Payload[219:221]))
	data["lFront"] = int16(binary.LittleEndian.Uint16(g.Payload[221:223]))
	data["lBack"] = int16(binary.LittleEndian.Uint16(g.Payload[223:225]))
	data["rBack"] = int16(binary.LittleEndian.Uint16(g.Payload[225:227]))

	g.Data = data
}