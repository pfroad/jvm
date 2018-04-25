package classfile

const (
	CONSTANT_VALUE = "ConstantValue"
	CODE = "Code"
	STACK_MAP_TABLE = "StackMapTable"
	EXCEPTIONS = "Exceptions"
	INNER_CLASSES = "InnerClasses"
	ENCLOSING_METHOD = "EnclosingMethod"
	SYNTHETIC = "Synthetic"
	SIGNATURE = "Signature"
	SOURCE_FILE = "SourceFile"
	SOURCE_DEBUG_EXTENSION = "SourceDebugExtension"
	LINE_NUMBER_TABLE = "LineNumberTable"
	LOCAL_VARIABLE_TABLE = "LocalVariableTable"
	LOCAL_VARIABLE_TYPE = "LocalVariableTypeTable"
	DEPRECATED = "Deprecated"
	RUNTIME_VISIBALE_ANNOTATIONS = "RuntimeVisibleAnnotations"
	RUNTIME_INVISIBLE_ANNOTATIONS = "RuntimeInvisibleAnnotations"
	RUNTIME_VISIBLE_PARAMETER_ANNOTATIONS = "RuntimeVisibleParameterAnnotations"
	RUNTIME_INVISIBLE_PARAMETER_ANNOTATIONS = "RuntimeInvisibleParameterAnnotations"
	RUNTIME_VISIBLE_TYPE_ANNOTATIONS = "RuntimeVisibleTypeAnnotations"
	RUNTIME_INVISIBLE_TYPE_ANNOTATIONS = "RuntimeInvisibleTypeAnnotations"
	ANNOATATION_DEFAULT = "AnnotationDefault"
	BOOTSTRAP_METHODS = "BootstrapMethods"
	METHOD_PARAMETERS = "MethodParameters"
)

type AttributeInfo interface {
	readerInfo(reader *ClassReader)
}

func readAttributes(reader *ClassReader, cp ConstantPool) []AttributeInfo {
	attributeCount := reader.readUint16()
	attributes := make([]AttributeInfo, attributeCount)
	for i := range attributes {
		attributes[i] = readAttribute(reader, cp)
	}
	return attributes
}

func readAttribute(reader *ClassReader, cp ConstantPool) AttributeInfo {
	attrName := cp.getUtf8(reader.readUint16())
	attrLength := reader.readUint32()
	attributeInfo := newAttributeInfo(attrName, attrLength, cp)
	attributeInfo.readerInfo(reader)
	return attributeInfo
}

func newAttributeInfo(attrName string, attrLen uint32, cp ConstantPool) AttributeInfo {
	switch attrName {
	case CODE:
		return CodeAttribute{cp}
	case CONSTANT_VALUE:
		return ConstantValueAttribute{}
	case DEPRECATED:
		return DeprecatedAttribute{}
	case EXCEPTIONS:
		return ExceptionsAttribute{}
	case LINE_NUMBER_TABLE:
		return LineNumberTableAttribute{}
	case LOCAL_VARIABLE_TABLE:
		return LocalVariableTableAttribute{}
	case SOURCE_FILE:
		return SourceFileAttribute{cp}
	case SYNTHETIC:
		return SyntheticAttribute{}
	default:
		return UnparsedAttribute{}
	}
}
