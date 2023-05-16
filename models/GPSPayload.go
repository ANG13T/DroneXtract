package models

import (
	"encoding/binary"
	"math"
)

type GPSPayload struct {
	Fields   []string
	Type     uint8
	Subtype  uint8
	Length   uint8
	Payload  []byte
	Data     map[string]interface{}
}

func NewGPSPayload(payload []byte) *GPSPayload {
	g := &GPSPayload{
		Fields:  []string{"latitude", "longitude", "altitude", "accelX", "accelY", "accelZ", "gyroX", "gyroY", "gyroZ", "baroAlt", "quatW", "quatX", "quatY", "quatZ", "errorX", "errorY", "errorZ", "velN", "velE", "velD", "x4", "x5", "x6", "magX", "magY", "magZ", "imuTemp", "i2", "i3", "i4", "i5", "satnum", "vel", "velH", "error", "accel", "magMod", "gyro", "roll", "pitch", "yaw", "yaw360", "magYawX", "velGPS-velH", "totalGyroZ", "distanceTravelled", "directionOfTravel", "directionOfTravelTrue"},
		Type:    0xcf,
		Subtype: 0x01,
		Length:  0x84,
		Payload: payload,
		Data:    make(map[string]interface{}),
	}
	g.parse()
	return g
}

func (g *GPSPayload) convertpos(pos float64) float64 {
	// convert position from radians to degrees
	return pos * 180.0 / math.Pi
}

func (g *GPSPayload) mtoft(meter float32) float32 {
	// convert meters to feet
	return meter * 3.2808
}

func (g *GPSPayload) parse() {
	data := make(map[string]interface{})

	longitude := math.Float64frombits(binary.LittleEndian.Uint64(g.Payload[0:8]))
	data["longitude"] = g.convertpos(longitude)

	latitude := math.Float64frombits(binary.LittleEndian.Uint64(g.Payload[8:16]))
	data["latitude"] = g.convertpos(latitude)

	altitude := math.Float32frombits(binary.LittleEndian.Uint32(g.Payload[16:20]))
	data["altitude"] = g.mtoft(altitude)

	data["accelX"] = math.Float32frombits(binary.LittleEndian.Uint32(g.Payload[20:24]))
	data["accelY"] = math.Float32frombits(binary.LittleEndian.Uint32(g.Payload[24:28]))
	data["accelZ"] = math.Float32frombits(binary.LittleEndian.Uint32(g.Payload[28:32]))
	data["accel"] = float32(math.Sqrt(float64(data["accelX"].(float32))*float64(data["accelX"].(float32)) + float64(data["accelY"].(float32))*float64(data["accelY"].(float32)) + float64(data["accelZ"].(float32))*float64(data["accelZ"].(float32))))

	data["gyroX"] = math.Float32frombits(binary.LittleEndian.Uint32(g.Payload[32:36]))
	data["gyroY"] = math.Float32frombits(binary.LittleEndian.Uint32(g.Payload[36:40]))
	data["gyroZ"] = math.Float32frombits(binary.LittleEndian.Uint32(g.Payload[40:44]))
	data["gyro"] = float32(math.Sqrt(float64(data["gyroX"].(float32))*float64(data["gyroX"].(float32)) + float64(data["gyroY"].(float32))*float64(data["gyroY"].(float32)) + float64(data["gyroZ"].(float32))*float64(data["gyroZ"].(float32))))

	data["baroAlt"] = g.mtoft(math.Float32frombits(binary.LittleEndian.Uint32(g.Payload[44:48])))

	data["quatW"] = math.Float32frombits(binary.LittleEndian.Uint32(g.Payload[48:52]))
	data["quatX"] = math.Float32frombits(binary.LittleEndian.Uint32(g.Payload[52:56]))
	data["quatY"] = math.Float32frombits(binary.LittleEndian.Uint32(g.Payload[56:60]))
	data["quatZ"] = math.Float32frombits(binary.LittleEndian.Uint32(g.Payload[60:64]))

	q := NewQuaternion(data["quatX"].(float32), data["quatY"].(float32), data["quatZ"].(float32), data["quatW"].(float32))
	eAngs := q.ToEuler()
	data["pitch"] = g.convertpos(float64(eAngs[0]))
	data["roll"] = g.convertpos(float64(eAngs[1]))
	data["yaw"] = g.convertpos(float64(eAngs[2]))
	data["yaw360"] = (data["yaw"].(float64) + 360.0) % 360.0

	data["errorX"] = math.Float32frombits(binary.LittleEndian.Uint32(g.Payload[64:68]))
	data["errorY"] = math.Float32frombits(binary.LittleEndian.Uint32(g.Payload[68:72]))
	data["errorZ"] = math.Float32frombits(binary.LittleEndian.Uint32(g.Payload[72:76]))
	data["error"] = float32(math.Sqrt(float64(data["errorX"].(float32))*float64(data["errorX"].(float32)) + float64(data["errorY"].(float32))*float64(data["errorY"].(float32)) + float64(data["errorZ"].(float32))*float64(data["errorZ"].(float32))))

	data["velN"] = math.Float32frombits(binary.LittleEndian.Uint32(g.Payload[76:80]))
	data["velE"] = math.Float32frombits(binary.LittleEndian.Uint32(g.Payload[80:84]))
	data["velD"] = math.Float32frombits(binary.LittleEndian.Uint32(g.Payload[84:88]))
	data["vel"] = float32(math.Sqrt(float64(data["velN"].(float32))*float64(data["velN"].(float32)) + float64(data["velE"].(float32))*float64(data["velE"].(float32)) + float64(data["velD"].(float32))*float64(data["velD"].(float32))))
	data["velH"] = float32(math.Sqrt(float64(data["velN"].(float32))*float64(data["velN"].(float32)) + float64(data["velE"].(float32))*float64(data["velE"].(float32))))

	x4 := math.Float32frombits(binary.LittleEndian.Uint32(g.Payload[88:92]))
	x5 := math.Float32frombits(binary.LittleEndian.Uint32(g.Payload[92:96]))
	x6 := math.Float32frombits(binary.LittleEndian.Uint32(g.Payload[96:100]))

	data["magX"] = int16(binary.LittleEndian.Uint16(g.Payload[100:102]))
	data["magY"] = int16(binary.LittleEndian.Uint16(g.Payload[102:104]))
	data["magZ"] = int16(binary.LittleEndian.Uint16(g.Payload[104:106]))
	data["magMod"] = float32(math.Sqrt(float64(data["magX"].(int16))*float64(data["magX"].(int16)) + float64(data["magY"].(int16))*float64(data["magY"].(int16)) + float64(data["magZ"].(int16))*float64(data["magZ"].(int16))))

	qAcc := NewQuaternion(eAngs[0].(float32), eAngs[1].(float32), 0.0, 0.0)
	qMag := NewQuaternion(data["magX"].(float32), data["magY"].(float32), data["magZ"].(float32), 0.0)
	magXYPlane := qAcc.Times(qMag).Times(qAcc.Conjugate())
	x := magXYPlane.X
	y := magXYPlane.Y
	data["magYawX"] = g.convertpos(-math.Atan2(float64(y), float64(x)))

	data["imuTemp"] = int16(binary.LittleEndian.Uint16(g.Payload[106:108]))

	i2 := int16(binary.LittleEndian.Uint16(g.Payload[108:110]))
	i3 := int16(binary.LittleEndian.Uint16(g.Payload[110:112]))
	i4 := int16(binary.LittleEndian.Uint16(g.Payload[112:114]))
	i5 := int16(binary.LittleEndian.Uint16(g.Payload[114:116]))

	data["satnum"] = g.Payload[116]

	// only output legitimate location data
	if data["latitude"].(float64) == 0 || data["longitude"].(float64) == 0 || math.Abs(data["latitude"].(float64)) <= 0.0175 || math.Abs(data["longitude"].(float64)) <= 0.0175 || data["satnum"].(uint8) <= 2 || data["satnum"].(uint8) >= 32 {
		data["longitude"] = ""
		data["latitude"] = ""
	}

	g.Data = data
}


