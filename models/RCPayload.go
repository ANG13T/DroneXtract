package models

import (
	"fmt"
	"encoding/binary"
)

type RCPayload struct {
	fields      []string
	_type       uint8
	_subtype    uint8
	_length     uint8
	payload     []byte
	data        map[string]interface{}
}

func NewRCPayload(payload []byte) *RCPayload {
	return &RCPayload{
		fields:   []string{"aileron", "elevator", "throttle", "rudder", "modeSwitch", "gpsHealth"},
		_type:    0x98,
		_subtype: 0x00,
		_length:  0x37,
		payload:  payload,
		data:     make(map[string]interface{}),
	}
}

func (r *RCPayload) parse() {
	data := r.data

	data["aileron"] = int16(binary.LittleEndian.Uint16(r.payload[4:6]))
	data["elevator"] = int16(binary.LittleEndian.Uint16(r.payload[6:8]))
	data["throttle"] = int16(binary.LittleEndian.Uint16(r.payload[8:10]))
	data["rudder"] = int16(binary.LittleEndian.Uint16(r.payload[10:12]))
	data["modeSwitch"] = r.payload[31]
	data["gpsHealth"] = r.payload[41]
}