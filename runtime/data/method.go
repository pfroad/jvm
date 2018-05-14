package data

import "jvm/classfile"

type Method struct {
	ClassMember
	maxStack  uint
	maxLocals uint
	code      []byte
}

func newMethods(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))
	for i, cfMethod := range cfMethods {
		method := &Method{}
		method.class = class
		method.copyFromMember(cfMethod)
		method.copyFromAttributes(cfMethod)
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

func (method *Method) Class() *Class {
	return method.class
}