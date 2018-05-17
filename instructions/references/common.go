package references

import (
	"jvm/runtime/data"
	"jvm/runtime"
	"fmt"
)

func InvokeMethod(method *data.Method, invokerFrame *runtime.Frame) {
	thread := invokerFrame.Thread()
	newFrame := thread.NewFrame(method)
	thread.PushFrame(newFrame)
	argCount := int(method.ArgCount())
	for i := argCount - 1; i >= 0; i-- {
		slot := invokerFrame.OperandStack().Pop()
		newFrame.LocalVars().SetSlot(uint(i), slot)
	}

	// hack!
	if method.IsNative() {
		if method.Name() == "registerNatives" {
			thread.PopFrame()
		} else {
			panic(fmt.Sprintf("native method: %v.%v%v\n",
				method.Class().ClassName(), method.Name(), method.Descriptor()))
		}
	}
}
