package classfile

import (
	"fmt"
	"unicode/utf16"
)

type ConstantUtf8Info struct {
	val string
}

func (self *ConstantUtf8Info) readInfo(reader *ClassReader) {
	length := uint32(reader.readUint16())
	bytes := reader.readBytes(length)
	self.val = decodeMUTF8(bytes)
}

func decodeMUTF8(bytes []byte) string {
	utfLen := len(bytes)
	chars := make([]uint16, utfLen)

	var char1, char2, char3 uint16
	count := 0
	charsCount := 0

	for count < utfLen {
		char1 = uint16(bytes[count])
		if char1 > 127 {
			break
		}
		count++
		chars[charsCount] = char1
		charsCount++
	}
	for count < utfLen {
		char1 = uint16(bytes[count])
		switch char1 >> 4 {
		case 0, 1, 2, 3, 4, 5, 6, 7:
			/* 0xxxxxxx*/
			count++
			chars[charsCount] = char1
			charsCount++
		case 12, 13:
			/* 110x xxxx   10xx xxxx*/
			count += 2
			if count > utfLen {
				panic("malformed input: partial character at end")
			}
			char2 = uint16(bytes[count-1])
			if char2&0xC0 != 0x80 {
				panic(fmt.Errorf("malformed input around byte %v", count))
			}
			chars[charsCount] = char1&0x1F<<6 | char2&0x3F
			charsCount++
		case 14:
			/* 1110 xxxx  10xx xxxx  10xx xxxx*/
			count += 3
			if count > utfLen {
				panic("malformed input: partial character at end")
			}
			char2 = uint16(bytes[count-2])
			char3 = uint16(bytes[count-1])
			if char2&0xC0 != 0x80 || char3&0xC0 != 0x80 {
				panic(fmt.Errorf("malformed input around byte %v", (count - 1)))
			}
			chars[charsCount] = char1&0x0F<<12 | char2&0x3F<<6 | char3&0x3F<<0
			charsCount++
		default:
			/* 10xx xxxx,  1111 xxxx */
			panic(fmt.Errorf("malformed input around byte %v", count))
		}
	}
	chars = chars[0:charsCount]
	runes := utf16.Decode(chars)
	return string(runes)
}
