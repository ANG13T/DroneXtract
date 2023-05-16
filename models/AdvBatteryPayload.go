package models

import (
	"fmt"
	"encoding/binary"
)

type AdvBatteryPayload struct {
	fields             []string
	_type              byte
	_subtype           byte
	_length            byte
	payload            []byte
	data               map[string]interface{}
}

func NewAdvBatteryPayload(payload []byte) *AdvBatteryPayload {
	p := &AdvBatteryPayload{
		fields:    []string{"current", "volt1", "volt2", "volt3", "volt4", "volt5", "volt6", "totalVolts", "voltSpread", "Watts", "batteryTemp(C)", "ratedCapacity", "remainingCapacity", "percentageCapacity"},
		_type:     0x44,
		_subtype:  0x11,
		_length:   0x39,
		payload:   payload,
		data:      make(map[string]interface{}),
	}

	p.parse()

	return p
}

func (p *AdvBatteryPayload) parse() {
	data := p.data
	payload := p.payload

	data["ratedCapacity"] = int16(binary.LittleEndian.Uint16(payload[2:4]))
	data["remainingCapacity"] = int16(binary.LittleEndian.Uint16(payload[4:6]))
	data["totalVolts"] = float32(int16(binary.LittleEndian.Uint16(payload[6:8]))) / 1000.0
	data["current"] = -float32(binary.LittleEndian.Uint16(payload[8:10])-65536) / 1000.0
	data["percentageCapacity"] = payload[11]
	data["batteryTemp(C)"] = payload[12]
	data["volt1"] = float32(int16(binary.LittleEndian.Uint16(payload[18:20]))) / 1000.0
	data["volt2"] = float32(int16(binary.LittleEndian.Uint16(payload[20:22]))) / 1000.0
	data["volt3"] = float32(int16(binary.LittleEndian.Uint16(payload[22:24]))) / 1000.0
	data["volt4"] = float32(int16(binary.LittleEndian.Uint16(payload[24:26]))) / 1000.0

	// Only DJI Inspire has 6 cell battery, comment out for DJI Phantom 3
	//data["volt5"] = float32(int16(binary.LittleEndian.Uint16(payload[26:28]))) / 1000.0
	//data["volt6"] = float32(int16(binary.LittleEndian.Uint16(payload[28:30]))) / 1000.0

	voltMax := max(data["volt1"].(float32), max(data["volt2"].(float32), max(data["volt3"].(float32), data["volt4"].(float32))))
	voltMin := min(data["volt1"].(float32), min(data["volt2"].(float32), min(data["volt3"].(float32), data["volt4"].(float32))))
	data["voltSpread"] = voltMax - voltMin

	data["Watts"] = data["totalVolts"].(float32) * data["current"].(float32)
}

func max(a, b float32) float32 {
	if a > b {
		return a
	}
	return b
}

func min(a, b float32) float32 {
	if a < b {
		return a
	}
	return b
}