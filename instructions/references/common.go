package references

import (
	"jvm/runtime/data"
	"jvm/runtime"
)

func InvokeMethod(method *data.Method, invokerFrame *runtime.Frame) {
	thread := invokerFrame.Thread()
	newFrame := runtime.NewFrame(thread, method)
	thread.PushFrame(newFrame)
	argCount := int(method.ArgCount())
	for i := 0; i < argCount; i++ {
		slot := invokerFrame.OperandStack().Pop()
		newFrame.LocalVars().SetSlot(uint(i), slot)
	}
}
