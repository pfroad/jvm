package main

import (
	"jvm/runtime"
	"fmt"
	"jvm/instructions/common"
	"jvm/instructions"
	"jvm/runtime/data"
)

func interpret(method *data.Method) {
	thread := runtime.NewThread()
	frame := runtime.NewFrame(thread, method)
	thread.PushFrame(frame)

	defer catchError(frame)
	loop(thread, method.Code())
}

func loop(thread *runtime.Thread, code []byte) {
	frame := thread.PopFrame()
	reader := &common.BytecodeReader{}

	for {
		pc := thread.PC()
		//thread.SetPC(pc)
		reader.Reset(code, pc)

		opcode := reader.ReadUint8()
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader, frame)

		// execute
		fmt.Printf("pc:%2d inst:%T %v\n", pc, inst, inst)
		inst.Execute(frame)
	}
}

func catchError(frame *runtime.Frame) {
	if r := recover(); r != nil {
		fmt.Printf("Variables:%v\n", frame.LocalVars())
		fmt.Printf("OperandStack: %v\n", frame.OperandStack())
		panic(r)
	}
}