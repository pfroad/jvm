package references

import (
	"fmt"
	"jvm/runtime"
	"jvm/runtime/data"
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

// InitClass exec clinit to init class
func InitClass(thread *runtime.Thread, class *data.Class) {
	class.StartInit()
	scheduleClinit(thread, class)
	initSuperClass(thread, class)
}

func scheduleClinit(thread *runtime.Thread, class *data.Class) {
	clinit := class.GetClinitMethod()
	if clinit != nil {
		// exec clinit
		newFrame := thread.NewFrame(clinit)
		thread.PushFrame(newFrame)
	}
}

func initSuperClass(thread *runtime.Thread, class *data.Class) {
	if !class.IsInterface() {
		superClass := class.SuperClass()
		if superClass != nil && !superClass.InitStarted() {
			InitClass(thread, superClass)
		}
	}
}
