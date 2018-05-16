package data

import "strings"

/*
java code:
public Integer test(Long[][] a, Long b, int[] c)
descriptor: ([[Ljava/lang/Long;Ljava/lang/Long;[I)Ljava/lang/Integer;
*/
type MethodDescriptor struct {
	parameterTypes []string
	returnType     string
}

func (md *MethodDescriptor) addParameterType(parameterType string) {
	pLen := len(md.parameterTypes)

	if pLen >= cap(md.parameterTypes) {
		newArr := make([]string, pLen, pLen+4)
		copy(newArr, md.parameterTypes)
		md.parameterTypes = newArr
	}

	md.parameterTypes = append(md.parameterTypes, parameterType)
}

type MethodDescriptorParser struct {
	descriptor string
	offset     uint
	parsedMD   *MethodDescriptor
}

func (parser *MethodDescriptorParser) parse(descriptor string) *MethodDescriptor {
	parser.parsedMD = &MethodDescriptor{}
	parser.descriptor = descriptor
	parser.parseStart()
	parser.parseParameterTypes()
	parser.parseEnd()
	parser.parseReturnType()
	parser.finished()
	return parser.parsedMD
}

func (parser *MethodDescriptorParser) parseStart() {
	startTag := parser.descriptor[parser.offset]
	parser.offset++

	if startTag != '(' {
		parser.causePanic()
	}
}

func (parser *MethodDescriptorParser) parseParameterTypes() {
	for p := parser.parseType(); p != ""; p = parser.parseType() {
		parser.parsedMD.addParameterType(p)
	}
}

func (parser *MethodDescriptorParser) parseReturnType() {
	parser.parsedMD.returnType = parser.parseType()
}

func (parser *MethodDescriptorParser) parseEnd() {
	endTag := parser.descriptor[parser.offset]
	parser.offset++

	if endTag != ')' {
		parser.causePanic()
	}
}

func (parser *MethodDescriptorParser) causePanic() {
	panic("BAD descriptor: " + parser.descriptor)
}

func (parser *MethodDescriptorParser) parseType() string {
	tag := parser.descriptor[parser.offset]
	parser.offset++
	switch tag {
	case 'B':
		return "B"
	case 'C':
		return "C"
	case 'D':
		return "D"
	case 'F':
		return "F"
	case 'I':
		return "I"
	case 'J':
		return "J"
	case 'S':
		return "S"
	case 'Z':
		return "Z"
	case 'L':
		return parser.parseObjectType()
	case '[':
		return parser.parseArrayType()
	default:
		parser.offset--
		return ""
	}
}

func (parser *MethodDescriptorParser) parseObjectType() string {
	unParseDescriptor := parser.descriptor[parser.offset-1:]
	endIndex := strings.Index(unParseDescriptor, ";")

	if endIndex < 0 {
		parser.causePanic()
	}

	parser.offset += uint(endIndex)
	return unParseDescriptor[:endIndex + 1]
}

func (parser *MethodDescriptorParser) parseArrayType() string {
	startIndex := parser.offset - 1
	elementType := parser.parseType()

	if elementType == "" {
		parser.causePanic()
	}

	return parser.descriptor[startIndex : parser.offset]
}

func (parser *MethodDescriptorParser) finished() {
	if int(parser.offset) > len(parser.descriptor) {
		parser.causePanic()
	}
}