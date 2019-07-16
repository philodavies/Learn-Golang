package main

import "encoding/binary"

func byteToInt(byt []byte) (num uint16) {
	num = 256*uint16(byt[0]) + uint16(byt[1])
	return num
}

func byteDecode(byt []byte) (num uint16) {
	num = binary.BigEndian.Uint16(byt)
	return num
}

func main() {
	test_byte := make([]byte, 2)
	test_byte[0] = 215
	test_byte[1] = 48
	byteToInt(test_byte)
}
