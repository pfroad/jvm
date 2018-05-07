package main

import (
	"jvm/classfile"
	"jvm/runtime"
	"fmt"
	"jvm/instructions/common"
	"jvm/instructions"
)

func interpret(methodInfo *classfile.MemberInfo) {
	codeArr := methodInfo.CodeAttribute()
	maxLocals := uint(codeArr.MaxLocals())
	maxStack := uint(codeArr.MaxStack())
	code := codeArr.Code()

	thread := runtime.NewThread()
	frame := runtime.NewFrame(thread, maxLocals, maxStack)
	thread.PushFrame(frame)

	defer catchError(frame)
	loop(thread, code)
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
		fmt.Printf("LocalVars:%v\n", frame.LocalVars())
		fmt.Printf("OperandStack: %v\n", frame.OperandStack())
		panic(r)
	}
}