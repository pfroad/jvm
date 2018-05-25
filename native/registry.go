package native

import "jvm/runtime"

type NativeMethod func(frame *runtime.Frame)

var registry = map[string]NativeMethod{}

func Register(className, methodName, descriptor string, method NativeMethod) {
	key := nativeMethodKey(className, methodName, descriptor)
	registry[key] = method
}

func findNativeMethod(className, methodName, descriptor string) NativeMethod {
	key := nativeMethodKey(className, methodName, descriptor)
	if nativeMethod, ok := registry[key]; ok {
		return nativeMethod
	}

	if descriptor == "()V" && methodName == "registerNative" {
		return emptyNativeMethod
	}

	return nil
}

func emptyNativeMethod(frame *runtime.Frame) {
	// todo
}

func nativeMethodKey(className, methodName, descriptor string) string {
	return className + "~" + methodName + "~" + descriptor
}
