package models

import (
	"encoding/binary"
	"math"
)

type TabletLocPayload struct {
	fields           []string
	_type            uint8
	_subtype         uint8
	_length          uint8
	payload          []byte
	data             map[string]interface{}
}

func NewTabletLocPayload(payload []byte) *TabletLocPayload {
	return &TabletLocPayload{
		fields:    []string{"latitudeTablet", "longitudeTablet"},
		_type:     0xc1,
		_subtype:  0x2b,
		_length:   0,
		payload:   payload,
		data:      make(map[string]interface{}),
	}
}

func (t *TabletLocPayload) parse() {
	data := t.data

	data["longitudeTablet"] = binary.LittleEndian.Uint64(t.payload[155:163])
	data["latitudeTablet"] = binary.LittleEndian.Uint64(t.payload[163:171])

	if data["latitudeTablet"] == 0 || data["longitudeTablet"] == 0 || math.Abs(data["latitudeTablet"].(float64)) <= 0.0175 || math.Abs(data["longitudeTablet"].(float64)) <= 0.0175 || math.Abs(data["latitudeTablet"].(float64)) >= 181.0 || math.Abs(data["longitudeTablet"].(float64)) >= 181.0 {
		data["longitudeTablet"] = ""
		data["latitudeTablet"] = ""
	}
}
