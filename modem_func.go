package main

import (
	"fmt"
	"log"

	"github.com/tarm/serial"
)

// Make conection to RaspberryPi USB serial port
func ConnectModem() *serial.Port {
	c := &serial.Config{Name: "/dev/ttyUSB0", Baud: 115200}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	return s
}

// Set status Modem ready to send SMS
func InitModem(m *serial.Port) {
	// Sure that modem is in AT Mode
	RunAT(m)
	// Set text mode to send SMS
	SetTextMode(m)
	// Set GSM char encoding
	SetCharCoding(m)
}

// Sure that modem is in AT Mode
func RunAT(m *serial.Port) {
	_, err := m.Write([]byte("AT\n\r"))
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, 128)
	b, _ := m.Read(buf)
	log.Printf("%q", buf[:b])
}

// Set text mode to send SMS
func SetTextMode(m *serial.Port) {
	_, err := m.Write([]byte("AT+CMGF=1\n\r"))
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, 128)
	b, _ := m.Read(buf)
	log.Printf("%q", buf[:b])
}

// Set GSM char encoding
func SetCharCoding(m *serial.Port) {
	init := fmt.Sprintf("AT+CSCS=%q", "GSM")
	m.Write([]byte(init))
	m.Write([]byte("\n\r"))
	buf := make([]byte, 128)
	b, _ := m.Read(buf)
	log.Printf("%q", buf[:b])
}

// Send message
func WriteSMS(m *serial.Port, msg string, nr string) ([]byte, error) {
	SetRecipient(m, nr)
	b, err := WriteMsg(m, msg)
	return b, err
}

// Set SMS recipient
func SetRecipient(m *serial.Port, nr string) {
	number := fmt.Sprintf("48%s", nr)
	init := fmt.Sprintf("AT+CMGS=%q", number)
	m.Write([]byte(init))
	m.Write([]byte("\n\r"))
	buf := make([]byte, 128)
	b, _ := m.Read(buf)
	log.Printf("%q", buf[:b])
}

// Write SMS
func WriteMsg(m *serial.Port, msg string) ([]byte, error) {
	_, err := m.Write([]byte(fmt.Sprintf("Veryfication code: %s", msg)))
	if err != nil {
		log.Fatal(err)
	}
	_, err = m.Write([]byte{0x1A})
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, 128)
	b, err := m.Read(buf)
	return buf[:b], err
}
