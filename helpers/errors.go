package helpers

import (
	"fmt"
	"io"
)

type NotDATFileError struct {
	filename string
}

func (e NotDATFileError) Error() string {
	return fmt.Sprintf("*** ERROR: %s is not a recognized DJI DAT file. ***", e.filename)
}

type NoNewPacketError struct {
	byteValue byte
	filePos   int64
}

func (e NoNewPacketError) Error() string {
	return fmt.Sprintf("No new packet found. Byte: 0x%x, File Position: %d", e.byteValue, e.filePos)
}

// CorruptPacketError represents a custom error for a corrupt packet.
type CorruptPacketError struct {
	message string
}

func (e CorruptPacketError) Error() string {
	return e.message
}

// Message represents the message structure.
type Message struct {
	// Define the fields of the Message structure as needed
}

func (m *Message) setTickNo(tickNo int) {
	// Implement the setTickNo method as needed
}

func (m *Message) writeRow(writer io.Writer, tickNo int) {
	// Implement the writeRow method as needed
}

func (m *Message) addPacket(pktLen int, header, payload []byte) bool {
	// Implement the addPacket method as needed
	return false
}

func (m *Message) getRow() []string {
	// Implement the getRow method as needed
	return nil
}

func (m *Message) writeKml(row []string) {
	// Implement the writeKml method as needed
}

func (m *Message) finalizeKml() {
	// Implement the finalizeKml method as needed
}
