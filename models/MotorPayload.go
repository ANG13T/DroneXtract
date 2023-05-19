package models

import (
	"encoding/binary"
	"fmt"
	"math"
)

type MotorPayload struct {
	Fields   []string
	Type     uint8
	Subtype  uint8
	Length   uint8
	Payload  []byte
	Data     map[string]interface{}
}

func NewMotorPayload(payload []byte) *MotorPayload {
	mp := &MotorPayload{
		Fields:  []string{"rFrontSpeed", "lFrontSpeed", "lBackSpeed", "rBackSpeed", "rFrontLoad", "lFrontLoad", "lBackLoad", "rBackLoad", "thrustAngle"},
		Type:    0xda,
		Subtype: 0xf1,
		Payload: payload,
		Data:    make(map[string]interface{}),
	}
	mp.parse()
	return mp
}

func (mp *MotorPayload) parse() {
	data := make(map[string]interface{})

	data["rFrontLoad"] = int16(binary.LittleEndian.Uint16(mp.Payload[1:3]))
	data["rFrontSpeed"] = int16(binary.LittleEndian.Uint16(mp.Payload[3:5]))
	data["lFrontLoad"] = int16(binary.LittleEndian.Uint16(mp.Payload[20:22]))
	data["lFrontSpeed"] = int16(binary.LittleEndian.Uint16(mp.Payload[22:24]))
	data["lBackLoad"] = int16(binary.LittleEndian.Uint16(mp.Payload[39:41]))
	data["lBackSpeed"] = int16(binary.LittleEndian.Uint16(mp.Payload[41:43]))
	data["rBackLoad"] = int16(binary.LittleEndian.Uint16(mp.Payload[58:60]))
	data["rBackSpeed"] = int16(binary.LittleEndian.Uint16(mp.Payload[60:62]))

	lbrfdiff := data["lBackSpeed"].(int16) - data["rFrontSpeed"].(int16)
	rblfdiff := data["rBackSpeed"].(int16) - data["lFrontSpeed"].(int16)

	thrust1 := math.Atan2(float64(lbrfdiff), float64(rblfdiff))
	thrust2 := (math.Mod(convert_rad_to_degrees(thrust1)+315.0, 360.0))
	data["thrustAngle"] = thrust2
	if thrust2 > 180.0 {
		data["thrustAngle"] = thrust2 - 360.0
	}

	mp.Data = data
}