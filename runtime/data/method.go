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
//func (method *Method) Name() string {
//	return method.Name()
//}
