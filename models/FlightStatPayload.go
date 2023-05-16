package models

import (
	"fmt"
	"encoding/binary"
	"math"
)

type FlightStatPayload struct {
	fields           []string
	NGPEErrs         map[byte]string
	FLCSStates       map[byte]string
	flyc_to_dw       map[byte]byte
	_type            byte
	_subtype         byte
	_length          byte
	payload          []byte
	data             map[string]interface{}
}

func NewFlightStatPayload(payload []byte) *FlightStatPayload {
	p := &FlightStatPayload{
		fields:     []string{"FSlongitude", "FSlatitude", "height", "FSpitch", "FSroll", "FSyaw", "flyc_state", "flycStateStr", "connectedToRC", "failure", "nonGPSError", "nonGPSErrStr", "time(millisecond)", "DWflyCState"},
		NGPEErrs:   map[byte]string{1: "FORBIN", 2: "GPSNUM_NONENOUGH", 3: "GPS_HDOP_LARGE", 4: "GPS_POSITION_NON_MATCH", 5: "SPEED_ERROR_LARGE", 6: "YAW_ERROR_LARGE", 7: "COMPASS_ERROR_LARGE"},
		FLCSStates: map[byte]string{0: "MANUAL", 1: "ATTI", 2: "ATTI_CL", 3: "ATTI_HOVER", 4: "HOVER", 5: "GSP_BLAKE", 6: "GPS_ATTI", 7: "GPS_CL", 8: "GPS_HOME_LOCK", 9: "GPS_HOT_POINT", 10: "ASSISTED_TAKEOFF", 11: "AUTO_TAKEOFF", 12: "AUTO_LANDING", 13: "ATTI_LANDING", 14: "NAVI_GO", 15: "GO_HOME", 16: "CLICK_GO", 17: "JOYSTICK", 23: "ATTI_LIMITED", 24: "GPS_ATTI_LIMITED", 25: "FOLLOW_ME", 100: "OTHER"},
		flyc_to_dw: map[byte]byte{0: 1, 1: 2, 2: 3, 3: 4, 4: 5, 5: 6, 6: 7, 7: 8, 8: 9, 9: 20, 10: 30, 11: 40, 12: 50, 13: 60, 14: 70, 15: 80, 16: 90, 17: 200, 23: 300, 24: 400},
		_type:      0x2A,
		_subtype:   0x0C,
		_length:    0x3E,
		payload:    payload,
		data:       make(map[string]interface{}),
	}

	p.parse()

	return p
}

func (p *FlightStatPayload) parse() {
	data := p.data
	payload := p.payload

	data["FSlongitude"] = p.convertpos(binary.LittleEndian.Uint64(payload[0:8]))
	data["FSlatitude"] = p.convertpos(binary.LittleEndian.Uint64(payload[8:16]))
	data["height"] = int16(binary.LittleEndian.Uint16(payload[16:18])) / 10.0

	data["FSpitch"] = int16(binary.LittleEndian.Uint16(payload[24:26])) / 10.0
	data["FSroll"] = int16(binary.LittleEndian.Uint16(payload[26:28])) / 10.0
	data["FSyaw"] = int16(binary.LittleEndian.Uint16(payload[28:30])) / 10.0

	flycState := payload[30] & 0x7f
	if _, ok := p.FLCSStates[flycState]; !ok {
		flycState = 0
	}
	data["flyc_state"] = flycState
	data["flycStateStr"] = p.FLCSStates[flycState]

	data["connectedToRC"] = 0
	if payload[30]&0x80 == 0 {
		data["connectedToRC"] = 1
	}

	data["failure"] = payload[38]
	data["nonGPSError"] = payload[39] & 0x7
	data["nonGPSErrStr"] = p.NGPEErrs[data["nonGPSError"].(byte)]

	data["time(millisecond)"] = int16(binary.LittleEndian.Uint16(payload[42:44])) * 100

	data["DWflyCState"] = p.flyc_to_dw[data["flyc_state"].(byte)]
}

func (p *FlightStatPayload) convertpos(pos uint64) float64 {
	// Convert position from radians to degrees
	return math.degrees(float64(pos))
}

