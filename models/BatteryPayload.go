package models

import (
	"encoding/binary"
)

type BatteryPayload struct {
	fields             []string
	_type              byte
	_subtype           byte
	_length            byte
	payload            []byte
	data               map[string]interface{}
}

func NewBatteryPayload(payload []byte) *BatteryPayload {
	p := &BatteryPayload{
		fields:    []string{"batteryUsefulTime", "voltagePercent"},
		_type:     0x1e,
		_subtype:  0x12,
		_length:   0x59,
		payload:   payload,
		data:      make(map[string]interface{}),
	}

	p.parse()

	return p
}

func (p *BatteryPayload) parse() {
	data := p.data
	payload := p.payload

	data["batteryUsefulTime"] = int16(binary.LittleEndian.Uint16(payload[0:2]))
	data["voltagePercent"] = payload[72]
}

