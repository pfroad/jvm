package data

import "unicode/utf16"

var internedStrings = map[string]*Object {}

func JString(loader *ClassLoader, goString string) *Object {
	if internedString, ok := internedStrings[goString]; ok {
		return internedString
	}

	chars := stringToUtf16(goString)
	jChars := &Object{loader.LoadClass("[C"), chars}
	jString := loader.LoadClass("java/lang.String").NewObject()
	jString.SetRefVar("value", "[C", jChars)
	internedStrings[goString] = jString
	return jString
}

func stringToUtf16(str string) []uint16 {
	runes := []rune(str)
	return utf16.Encode(runes)
}

func GoString(jStr *Object) string {
	jChars := jStr.GetRefVar("value", "[C")
	return Utf16ToString(jChars.Chars())
}

func Utf16ToString(jChars []uint16) string {
	return string(utf16.Decode(jChars))
}