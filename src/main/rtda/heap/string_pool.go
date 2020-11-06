package heap

import (
	"unicode/utf16"
)

var internedStrings = map[string]*Object{}

func JString(loader *ClassLoader, goStr string) *Object {
	if internedStr, ok := internedStrings[goStr]; ok {
		return internedStr
	}

	chars := stringToUtf16(goStr)
	jChars := &Object{
		loader.LoadClass("[C"),
		chars,
		nil,
	}

	jStr := loader.LoadClass("java/lang/String").NewObject()
	jStr.setRefVar("value", "[C", jChars)

	internedStrings[goStr] = jStr
	return jStr
}

func stringToUtf16(s string) []uint16 {
	// 转成utf32
	runes := []rune(s)
	// 编码称utf16
	return utf16.Encode(runes)
}

func GoString(jStr *Object) string {
	charArr := jStr.GetRefVar("value", "[C")
	return utf16ToString(charArr.Chars())
}

func utf16ToString(chars []uint16) string {
	runes := utf16.Decode(chars)
	return string(runes)
}
