package models

import (
	"encoding/binary"
	"math"
)

type HPPayload struct {
	Fields   []string
	Type     uint8
	Subtype  uint8
	Length   uint8
	Payload  []byte
	Data     map[string]interface{}
}

func NewHPPayload(payload []byte) *HPPayload {
	hp := &HPPayload{
		Fields:  []string{"longitudeHP", "latitudeHP"},
		Type:    0xc6,
		Subtype: 0x0d,
		Length:  0x2E,
		Payload: payload,
		Data:    make(map[string]interface{}),
	}
	hp.parse()
	return hp
}

func (hp *HPPayload) convertpos(pos float64) float64 {
	// convert position from radians to degrees
	return pos * 180.0 / math.Pi
}

func (hp *HPPayload) parse() {
	data := make(map[string]interface{})

	longitude := math.Float64frombits(binary.LittleEndian.Uint64(hp.Payload[0:8]))
	data["longitudeHP"] = hp.convertpos(longitude)

	latitude := math.Float64frombits(binary.LittleEndian.Uint64(hp.Payload[8:16]))
	data["latitudeHP"] = hp.convertpos(latitude)

	// only output legitimate location data
	if data["latitudeHP"].(float64) == 0 || data["longitudeHP"].(float64) == 0 || math.Abs(data["latitudeHP"].(float64)) <= 0.0175 || math.Abs(data["longitudeHP"].(float64)) <= 0.0175 || math.Abs(data["latitudeHP"].(float64)) >= 181.0 || math.Abs(data["longitudeHP"].(float64)) >= 181.0 {
		data["longitudeHP"] = ""
		data["latitudeHP"] = ""
	}

	hp.Data = data
}