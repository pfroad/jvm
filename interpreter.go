package main

import (
	"fmt"
	"jvm/instructions"
	"jvm/instructions/common"
	"jvm/runtime"
	"jvm/runtime/data"
)

func interpret(method *data.Method, logInst bool) {
	thread := runtime.NewThread()
	frame := thread.NewFrame(method)
	//frame := runtime.NewFrame(thread, method)
	thread.PushFrame(frame)

	defer catchError(thread)
	// loop(thread, method.Code())
	loop(thread, logInst)
}

// func loop(thread *runtime.Thread, code []byte) {
// 	frame := thread.PopFrame()
// 	reader := &common.CodeReader{}

// 	for {
// 		pc := thread.PC()
// 		//thread.SetPC(pc)
// 		reader.Reset(code, pc)

// 		opcode := reader.ReadUint8()
// 		inst := instructions.NewInstruction(opcode)
// 		inst.FetchOperands(reader, frame)

// 		// execute
// 		fmt.Printf("pc:%2d inst:%T %v\n", pc, inst, inst)
// 		inst.Execute(frame)
// 	}
// }

func loop(thread *runtime.Thread, logInst bool) {
	// frame := thread.PopFrame()	// return instruction will Pop frame
	// reader := &common.BytecodeReader{}
	for {
		frame := thread.TopFrame()
		thread.SetPC(frame.PC())
		//thread.SetPC(pc)
		// reader.Reset(frame.Method().Code(), pc)
		reader := frame.CodeReader()
		opcode := reader.ReadUint8()
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)

		if logInst {
			logInstruction(frame, inst)
		}

		// execute
		// fmt.Printf("pc:%2d inst:%T %v\n", pc, inst, inst)
		inst.Execute(frame)

		if thread.StackIsEmpty() {
			break
		}
	}
}

func catchError(thread *runtime.Thread) {
	if r := recover(); r != nil {
		// fmt.Printf("Variables:%v\n", frame.LocalVars())
		// fmt.Printf("OperandStack: %v\n", frame.OperandStack())
		logFrames(thread)
		panic(r)
	}
}

func logFrames(thread *runtime.Thread) {
	for !thread.StackIsEmpty() {
		frame := thread.PopFrame()
		method := frame.Method()
		className := method.Class().ClassName()
		fmt.Printf(">> pc: %4d, %v.%v%v \n",
			thread.PC(), className, method.Name(), method.Descriptor())
	}
}

func logInstruction(frame *runtime.Frame, inst common.Instruction) {
	method := frame.Method()
	className := method.Class().ClassName()
	pc := frame.Thread().PC()
	fmt.Printf("%v.%v() #%2d %T %v\n", className, method.Name(), pc, inst, inst)
}
