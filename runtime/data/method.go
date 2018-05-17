package data

import "jvm/classfile"

type Method struct {
	ClassMember
	maxStack  uint
	maxLocals uint
	code      []byte
	argCount  uint
}

func newMethods(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))
	for i, cfMethod := range cfMethods {
		method := &Method{}
		method.class = class
		method.copyFromMember(cfMethod)
		method.copyFromAttributes(cfMethod)
		method.calcArgCount()
		methods[i] = method
	}
	return methods
}

func (method *Method) copyFromAttributes(cfMethod *classfile.MemberInfo) {
	if code := cfMethod.CodeAttribute(); code != nil {
		method.maxStack = uint(code.MaxStack())
		method.maxLocals = uint(code.MaxLocals())
		method.code = code.Code()
	}
}

func (method *Method) MaxStack() uint {
	return method.maxStack
}

func (method *Method) MaxLocals() uint {
	return method.maxLocals
}

func (method *Method) IsStatic() bool {
	return method.accessFlags.IsStatic()
}

func (method *Method) IsFinal() bool {
	return method.accessFlags.IsFinal()
}

func (method *Method) IsPublic() bool {
	return method.accessFlags.IsPublic()
}

func (method *Method) Code() []byte {
	return method.code
}

func (method *Method) ArgCount() uint {
	return method.argCount
}

func (method *Method) calcArgCount() {
	parser := &MethodDescriptorParser{}
	methodDescriptor := parser.parse(method.descriptor)

	for _, pt := range methodDescriptor.parameterTypes {
		method.argCount++
		if pt == "J" || pt == "D" {
			method.argCount++
		}
	}

	if !method.IsStatic() {
		method.argCount++	// this
	}
}

func (method *Method) IsAbstract() bool {
	return method.accessFlags.IsAbstract()
}

func (method *Method) IsNative() bool {
	return method.accessFlags.IsNative()
}
//func (method *Method) Name() string {
//	return method.Name()
//}
