package instructions

import (
	"fmt"
	"jvm/instructions/common"
	"jvm/instructions/comparisons"
	"jvm/instructions/constants"
	"jvm/instructions/control"
	"jvm/instructions/conversions"
	"jvm/instructions/loads"
	"jvm/instructions/math"
	"jvm/instructions/stack"
	"jvm/instructions/stores"
)

/*
nop = 0x00
aconst_null  = 0x01
iconst_m1 = 0x02
iconst_0 = 0x03
iconst_1 = 0x04
iconst_2 = 0x05
iconst_3 = 0x06
iconst_4 = 0x07
iconst_5 = 0x08
lconst_0 = 0x09
lconst_1 = 0x0A
fconst_0 = 0x0B
fconst_1 = 0x0C
fconst_2 = 0x0D
dconst_0 = 0x0E
dconst_1 = 0x0F
bipush = 0x10
sipush = 0x11
ldc = 0x12
ldc_w = 0x13
ldc2_w = 0x14
iload = 0x15
lload = 0x16
fload = 0x17
dload = 0x18
aload = 0x19
iload_0 = 0x1A
iload_1 = 0x1B
iload_2 = 0x1C
iload_3 = 0x1D
lload_0 = 0x1E
lload_1 = 0x1F
lload_2 = 0x20
lload_3 = 0x21
fload_0 = 0x22
fload_1 = 0x23
fload_2 = 0x24
fload_3 = 0x25
dload_0 = 0x26
dload_1 = 0x27
dload_2 = 0x28
dload_3 = 0x29
aload_0 = 0x2A
aload_1 = 0x2B
aload_2 = 0x2C
aload_3 = 0x2D
iaload = 0x2E
laload = 0x2F
faload = 0x30
daload = 0x31
aaload = 0x32
baload = 0x33
caload = 0x34
saload = 0x35
istore = 0x36
lstore = 0x37
fstore = 0x38
dstore = 0x39
astore = 0x3A
istore_0 = 0x3B
istore_1 = 0x3C
istore_2 = 0x3D
istore_3 = 0x3E
lstore_0 = 0x3F
lstore_1 = 0x40
lstore_2 = 0x41
lstore_3 = 0x42
fstore_0 = 0x43
fstore_1 = 0x44
fstore_2 = 0x45
fstore_3 = 0x46
dstore_0 = 0x47
dstore_1 = 0x48
dstore_2 = 0x49
dstore_3 = 0x4A
astore_0 = 0x4B
astore_1 = 0x4C
astore_2 = 0x4D
astore_3 = 0x4E
iastore  = 0x4F
lastore = 0x50
fastore = 0x51
dastore = 0x52
aastore = 0x53
bastore = 0x54
castore = 0x55
sastore = 0x56
pop = 0x57
pop2 = 0x58
dup = 0x59
dup_x1 = 0x5A
dup_x2 = 0x5B
dup2 = 0x5C
dup2_x1 = 0x5D
dup2_x2 = 0x5E
swap = 0x5F
iadd = 0x60
ladd = 0x61
fadd = 0x62
dadd = 0x63
isub = 0x64
lsub = 0x65
fsub = 0x66
dsub = 0x67
imul = 0x68
lmul = 0x69
fmul = 0x6A
dmul = 0x6B
idiv = 0x6C
ldiv = 0x6D
fdiv = 0x6E
ddiv = 0x6F
irem = 0x70
lrem = 0x71
frem = 0x72
drem = 0x73
ineg = 0x74
lneg = 0x75
fneg = 0x76
dneg = 0x77
ishl = 0x78
lshl = 0x79
ishr = 0x7A
lshr = 0x7B
iushr = 0x7C
lushr = 0x7D
iand = 0x7E
land = 0x7F
ior = 0x80
lor = 0x81
ixor = 0x82
lxor = 0x83
iinc = 0x84
i2l = 0x85
i2f = 0x86
i2d = 0x87
l2i = 0x88
l2f = 0x89
l2d = 0x8A
f2i = 0x8B
f2l = 0x8C
f2d = 0x8D
d2i = 0x8E
d2l = 0x8F
d2f = 0x90
i2b = 0x91
i2c = 0x92
i2s = 0x93
lcmp = 0x94
fcmpl = 0x95
fcmpg = 0x96
dcmpl = 0x97
dcmpg = 0x98
ifeq = 0x99
ifne = 0x9A
iflt = 0x9B
ifge = 0x9C
ifgt = 0x9D
ifle = 0x9E
if_icmpeq = 0x9F
if_icmpne = 0xA0
if_icmplt = 0xA1
if_icmpge = 0xA2
if_icmpgt = 0xA3
if_icmple = 0xA4
if_acmpeq = 0xA5
if_acmpne = 0xA6
goto = 0xA7
jsr = 0xA8
ret = 0xA9
tableswitch = 0xAA
lookupswitch = 0xAB
ireturn = 0xAC
lreturn = 0xAD
freturn = 0xAE
dreturn = 0xAF
areturn = 0xB0
return = 0xB1
getstatic = 0xB2
putstatic = 0xB3
getfield = 0xB4
putfield = 0xB5
invokevirtual = 0xB6
invokespecial = 0xB7
invokestatic = 0xB8
invokeinterface = 0xb9
// --- = 0xBA
new = 0xBB
newarray = 0xBC
anewarray = 0xBD
arraylength = 0xBE
athrow = 0xBF
checkcast = 0xC0
instanceof = 0xc1
monitorenter = 0xC2
monitorexit = 0xC3
wide = 0xC4
multianewarray = 0xC5
ifnull = 0xC6
ifnonnull = 0xC7
goto_w = 0xC8
jsr_w = 0xC9
breakpoint = 0xCA
impdep1 = 0xFE
impdep = 0xFF
*/

const (
	NOP             = 0X00
	ACONST_NULL     = 0X01
	ICONST_M1       = 0X02
	ICONST_0        = 0X03
	ICONST_1        = 0X04
	ICONST_2        = 0X05
	ICONST_3        = 0X06
	ICONST_4        = 0X07
	ICONST_5        = 0X08
	LCONST_0        = 0X09
	LCONST_1        = 0X0A
	FCONST_0        = 0X0B
	FCONST_1        = 0X0C
	FCONST_2        = 0X0D
	DCONST_0        = 0X0E
	DCONST_1        = 0X0F
	BIPUSH          = 0X10
	SIPUSH          = 0X11
	LDC             = 0X12
	LDC_W           = 0X13
	LDC2_W          = 0X14
	ILOAD           = 0X15
	LLOAD           = 0X16
	FLOAD           = 0X17
	DLOAD           = 0X18
	ALOAD           = 0X19
	ILOAD_0         = 0X1A
	ILOAD_1         = 0X1B
	ILOAD_2         = 0X1C
	ILOAD_3         = 0X1D
	LLOAD_0         = 0X1E
	LLOAD_1         = 0X1F
	LLOAD_2         = 0X20
	LLOAD_3         = 0X21
	FLOAD_0         = 0X22
	FLOAD_1         = 0X23
	FLOAD_2         = 0X24
	FLOAD_3         = 0X25
	DLOAD_0         = 0X26
	DLOAD_1         = 0X27
	DLOAD_2         = 0X28
	DLOAD_3         = 0X29
	ALOAD_0         = 0X2A
	ALOAD_1         = 0X2B
	ALOAD_2         = 0X2C
	ALOAD_3         = 0X2D
	IALOAD          = 0X2E
	LALOAD          = 0X2F
	FALOAD          = 0X30
	DALOAD          = 0X31
	AALOAD          = 0X32
	BALOAD          = 0X33
	CALOAD          = 0X34
	SALOAD          = 0X35
	ISTORE          = 0X36
	LSTORE          = 0X37
	FSTORE          = 0X38
	DSTORE          = 0X39
	ASTORE          = 0X3A
	ISTORE_0        = 0X3B
	ISTORE_1        = 0X3C
	ISTORE_2        = 0X3D
	ISTORE_3        = 0X3E
	LSTORE_0        = 0X3F
	LSTORE_1        = 0X40
	LSTORE_2        = 0X41
	LSTORE_3        = 0X42
	FSTORE_0        = 0X43
	FSTORE_1        = 0X44
	FSTORE_2        = 0X45
	FSTORE_3        = 0X46
	DSTORE_0        = 0X47
	DSTORE_1        = 0X48
	DSTORE_2        = 0X49
	DSTORE_3        = 0X4A
	ASTORE_0        = 0X4B
	ASTORE_1        = 0X4C
	ASTORE_2        = 0X4D
	ASTORE_3        = 0X4E
	IASTORE         = 0X4F
	LASTORE         = 0X50
	FASTORE         = 0X51
	DASTORE         = 0X52
	AASTORE         = 0X53
	BASTORE         = 0X54
	CASTORE         = 0X55
	SASTORE         = 0X56
	POP             = 0X57
	POP2            = 0X58
	DUP             = 0X59
	DUP_X1          = 0X5A
	DUP_X2          = 0X5B
	DUP2            = 0X5C
	DUP2_X1         = 0X5D
	DUP2_X2         = 0X5E
	SWAP            = 0X5F
	IADD            = 0X60
	LADD            = 0X61
	FADD            = 0X62
	DADD            = 0X63
	ISUB            = 0X64
	LSUB            = 0X65
	FSUB            = 0X66
	DSUB            = 0X67
	IMUL            = 0X68
	LMUL            = 0X69
	FMUL            = 0X6A
	DMUL            = 0X6B
	IDIV            = 0X6C
	LDIV            = 0X6D
	FDIV            = 0X6E
	DDIV            = 0X6F
	IREM            = 0X70
	LREM            = 0X71
	FREM            = 0X72
	DREM            = 0X73
	INEG            = 0X74
	LNEG            = 0X75
	FNEG            = 0X76
	DNEG            = 0X77
	ISHL            = 0X78
	LSHL            = 0X79
	ISHR            = 0X7A
	LSHR            = 0X7B
	IUSHR           = 0X7C
	LUSHR           = 0X7D
	IAND            = 0X7E
	LAND            = 0X7F
	IOR             = 0X80
	LOR             = 0X81
	IXOR            = 0X82
	LXOR            = 0X83
	IINC            = 0X84
	I2L             = 0X85
	I2F             = 0X86
	I2D             = 0X87
	L2I             = 0X88
	L2F             = 0X89
	L2D             = 0X8A
	F2I             = 0X8B
	F2L             = 0X8C
	F2D             = 0X8D
	D2I             = 0X8E
	D2L             = 0X8F
	D2F             = 0X90
	I2B             = 0X91
	I2C             = 0X92
	I2S             = 0X93
	LCMP            = 0X94
	FCMPL           = 0X95
	FCMPG           = 0X96
	DCMPL           = 0X97
	DCMPG           = 0X98
	IFEQ            = 0X99
	IFNE            = 0X9A
	IFLT            = 0X9B
	IFGE            = 0X9C
	IFGT            = 0X9D
	IFLE            = 0X9E
	IF_ICMPEQ       = 0X9F
	IF_ICMPNE       = 0XA0
	IF_ICMPLT       = 0XA1
	IF_ICMPGE       = 0XA2
	IF_ICMPGT       = 0XA3
	IF_ICMPLE       = 0XA4
	IF_ACMPEQ       = 0XA5
	IF_ACMPNE       = 0XA6
	GOTO            = 0XA7
	JSR             = 0XA8
	RET             = 0XA9
	TABLESWITCH     = 0XAA
	LOOKUPSWITCH    = 0XAB
	IRETURN         = 0XAC
	LRETURN         = 0XAD
	FRETURN         = 0XAE
	DRETURN         = 0XAF
	ARETURN         = 0XB0
	RETURN          = 0XB1
	GETSTATIC       = 0XB2
	PUTSTATIC       = 0XB3
	GETFIELD        = 0XB4
	PUTFIELD        = 0XB5
	INVOKEVIRTUAL   = 0XB6
	INVOKESPECIAL   = 0XB7
	INVOKESTATIC    = 0XB8
	INVOKEINTERFACE = 0XB9
	// --- = 0XBA
	NEW            = 0XBB
	NEWARRAY       = 0XBC
	ANEWARRAY      = 0XBD
	ARRAYLENGTH    = 0XBE
	ATHROW         = 0XBF
	CHECKCAST      = 0XC0
	INSTANCEOF     = 0XC1
	MONITORENTER   = 0XC2
	MONITOREXIT    = 0XC3
	WIDE           = 0XC4
	MULTIANEWARRAY = 0XC5
	IFNULL         = 0XC6
	IFNONNULL      = 0XC7
	GOTO_W         = 0XC8
	JSR_W          = 0XC9
	BREAKPOINT     = 0XCA
	IMPDEP1        = 0XFE
	IMPDEP         = 0XFF
)

func NewInstruction(opcode uint8) common.Instruction {
	switch opcode {
	case NOP:
		return &constants.NOP{}
	case ACONST_NULL:
		return &constants.AConstNull{}
	case ICONST_M1:
		return &constants.IConstM1{}
	case ICONST_0:
		return &constants.IConst0{}
	case ICONST_1:
		return &constants.IConst1{}
	case ICONST_2:
		return &constants.IConst2{}
	case ICONST_3:
		return &constants.IConst3{}
	case ICONST_4:
		return &constants.IConst4{}
	case ICONST_5:
		return &constants.IConst5{}
	case LCONST_0:
		return &constants.LConst0{}
	case LCONST_1:
		return &constants.LConst1{}
	case FCONST_0:
		return &constants.FConst0{}
	case FCONST_1:
		return &constants.FConst1{}
	case FCONST_2:
		return &constants.FConst2{}
	case DCONST_0:
		return &constants.DConst0{}
	case DCONST_1:
		return &constants.DConst1{}
	case BIPUSH:
		return &constants.BIPush{}
	case SIPUSH:
		return &constants.SIPush{}
	//case LDC:
	//
	//case LDC_W:
	//
	//case LDC2_W:

	case ILOAD:
		return &loads.ILoad{}
	case LLOAD:
		return &loads.LLoad{}
	case FLOAD:
		return &loads.FLoad{}
	case DLOAD:
		return &loads.DLoad{}
	case ALOAD:
		return &loads.ALoad{}
	case ILOAD_0:
		return &loads.ILoad0{}
	case ILOAD_1:
		return &loads.ILoad1{}
	case ILOAD_2:
		return &loads.ILoad2{}
	case ILOAD_3:
		return &loads.ILoad3{}
	case LLOAD_0:
		return &loads.LLoad0{}
	case LLOAD_1:
		return &loads.LLoad1{}
	case LLOAD_2:
		return &loads.LLoad2{}
	case LLOAD_3:
		return &loads.LLoad3{}
	case FLOAD_0:
		return &loads.FLoad0{}
	case FLOAD_1:
		return &loads.FLoad1{}
	case FLOAD_2:
		return &loads.FLoad2{}
	case FLOAD_3:
		return &loads.FLoad3{}
	case DLOAD_0:
		return &loads.DLoad0{}
	case DLOAD_1:
		return &loads.DLoad1{}
	case DLOAD_2:
		return &loads.DLoad2{}
	case DLOAD_3:
		return &loads.DLoad3{}
	case ALOAD_0:
		return &loads.ALoad0{}
	case ALOAD_1:
		return &loads.ALoad1{}
	case ALOAD_2:
		return &loads.ALoad2{}
	case ALOAD_3:
		return &loads.ALoad3{}
	//case IALOAD:
	//
	//case LALOAD:
	//
	//case FALOAD:
	//
	//case DALOAD:
	//
	//case AALOAD:
	//
	//case BALOAD:
	//
	//case CALOAD:
	//
	//case SALOAD:

	case ISTORE:
		return &stores.IStore{}
	case LSTORE:
		return &stores.LStore{}
	case FSTORE:
		return &stores.FStore{}
	case DSTORE:
		return &stores.DStore{}
	case ASTORE:
		return &stores.AStore{}
	case ISTORE_0:
		return &stores.IStore0{}
	case ISTORE_1:
		return &stores.IStore1{}
	case ISTORE_2:
		return &stores.IStore2{}
	case ISTORE_3:
		return &stores.IStore3{}
	case LSTORE_0:
		return &stores.LStore0{}
	case LSTORE_1:
		return &stores.LStore1{}
	case LSTORE_2:
		return &stores.LStore2{}
	case LSTORE_3:
		return &stores.LStore3{}
	case FSTORE_0:
		return &stores.FStore0{}
	case FSTORE_1:
		return &stores.FStore1{}
	case FSTORE_2:
		return &stores.FStore2{}
	case FSTORE_3:
		return &stores.FStore3{}
	case DSTORE_0:
		return &stores.DStore0{}
	case DSTORE_1:
		return &stores.DStore1{}
	case DSTORE_2:
		return &stores.DStore2{}
	case DSTORE_3:
		return &stores.DStore3{}
	case ASTORE_0:
		return &stores.AStore0{}
	case ASTORE_1:
		return &stores.AStore1{}
	case ASTORE_2:
		return &stores.AStore2{}
	case ASTORE_3:
		return &stores.AStore3{}
	//case IASTORE:
	//
	//case LASTORE:
	//
	//case FASTORE:
	//
	//case DASTORE:
	//
	//case AASTORE:
	//
	//case BASTORE:
	//
	//case CASTORE:
	//
	//case SASTORE:

	case POP:
		return &stack.Pop{}
	case POP2:
		return &stack.Pop2{}
	case DUP:
		return &stack.Dup{}
	case DUP_X1:
		return &stack.DupX1{}
	case DUP_X2:
		return &stack.DupX2{}
	case DUP2:
		return &stack.Dup2{}
	case DUP2_X1:
		return &stack.Dup2X1{}
	case DUP2_X2:
		return &stack.Dup2X2{}
	case SWAP:
		return &stack.Swap{}
	case IADD:
		return &math.IAdd{}
	case LADD:
		return &math.LAdd{}
	case FADD:
		return &math.FAdd{}
	case DADD:
		return &math.DAdd{}
	case ISUB:
		return &math.ISub{}
	case LSUB:
		return &math.LSub{}
	case FSUB:
		return &math.FSub{}
	case DSUB:
		return &math.DSub{}
	case IMUL:
		return &math.IMul{}
	case LMUL:
		return &math.LMul{}
	case FMUL:
		return &math.FMul{}
	case DMUL:
		return &math.DMul{}
	case IDIV:
		return &math.IDiv{}
	case LDIV:
		return &math.LDiv{}
	case FDIV:
		return &math.FDiv{}
	case DDIV:
		return &math.DDiv{}
	case IREM:
		return &math.IRem{}
	case LREM:
		return &math.LRem{}
	case FREM:
		return &math.FRem{}
	case DREM:
		return &math.DRem{}
	case INEG:
		return &math.INeg{}
	case LNEG:
		return &math.LNeg{}
	case FNEG:
		return &math.FNeg{}
	case DNEG:
		return &math.DNeg{}
	case ISHL:
		return &math.IShL{}
	case LSHL:
		return &math.IShL{}
	case ISHR:
		return &math.IShR{}
	case LSHR:
		return &math.LShR{}
	case IUSHR:
		return &math.IUShR{}
	case LUSHR:
		return &math.LUShR{}
	case IAND:
		return &math.IAnd{}
	case LAND:
		return &math.LAnd{}
	case IOR:
		return &math.IOR{}
	case LOR:
		return &math.LOR{}
	case IXOR:
		return &math.IXOR{}
	case LXOR:
		return &math.LXOR{}
	case IINC:
		return &math.IINC{}
	case I2L:
		return &conversions.I2L{}
	case I2F:
		return &conversions.I2F{}
	case I2D:
		return &conversions.I2D{}
	case L2I:
		return &conversions.L2I{}
	case L2F:
		return &conversions.L2F{}
	case L2D:
		return &conversions.L2D{}
	case F2I:
		return &conversions.F2I{}
	case F2L:
		return &conversions.F2L{}
	case F2D:
		return &conversions.F2D{}
	case D2I:
		return &conversions.D2I{}
	case D2L:
		return &conversions.D2L{}
	case D2F:
		return &conversions.D2F{}
	case I2B:
		return &conversions.I2B{}
	case I2C:
		return &conversions.I2C{}
	case I2S:
		return &conversions.I2S{}
	case LCMP:
		return &comparisons.LCMP{}
	case FCMPL:
		return &comparisons.FCMPL{}
	case FCMPG:
		return &comparisons.FCMPG{}
	case DCMPL:
		return &comparisons.DCMPL{}
	case DCMPG:
		return &comparisons.DCMPG{}
	case IFEQ:
		return &control.IFEQ{}
	case IFNE:
		return &control.IFNE{}
	case IFLT:
		return &control.IFLT{}
	case IFGE:
		return &control.IFGE{}
	case IFGT:
		return &control.IFGT{}
	case IFLE:
		return &control.IFLE{}
	case IF_ICMPEQ:
		return &control.IFICMPEQ{}
	case IF_ICMPNE:
		return &control.IFICMPNE{}
	case IF_ICMPLT:
		return &control.IFICMPLT{}
	case IF_ICMPGE:
		return &control.IFICMPGE{}
	case IF_ICMPGT:
		return &control.IFICMPGT{}
	case IF_ICMPLE:
		return &control.IFICMPLE{}
	case IF_ACMPEQ:
		return &control.IFACMPEQ{}
	case IF_ACMPNE:
		return &control.IFACMPNE{}
	case GOTO:
		return &control.GoTo{}
	//case JSR:
	//
	//case RET:

	case TABLESWITCH:
		return &control.TableSwitch{}
	case LOOKUPSWITCH:
		return &control.LookupSwitch{}
	//case IRETURN:
	//
	//case LRETURN:
	//
	//case FRETURN:
	//
	//case DRETURN:
	//
	//case ARETURN:
	//
	//case RETURN:
	//
	//case GETSTATIC:
	//
	//case PUTSTATIC:
	//
	//case GETFIELD:
	//
	//case PUTFIELD:
	//
	//case INVOKEVIRTUAL:
	//
	//case INVOKESPECIAL:
	//
	//case INVOKESTATIC:
	//
	//case INVOKEINTERFACE:
	//
	//case NEW:
	//
	//case NEWARRAY:
	//
	//case ANEWARRAY:
	//
	//case ARRAYLENGTH:
	//
	//case ATHROW:
	//
	//case CHECKCAST:
	//
	//case INSTANCEOF:
	//
	//case MONITORENTER:
	//
	//case MONITOREXIT:
	//
	//case WIDE:
	//
	//case MULTIANEWARRAY:

	case IFNULL:
		return &control.IfNull{}
	case IFNONNULL:
		return &control.IfNonNull{}
	case GOTO_W:
		return &control.GoToW{}
	//case JSR_W:
	//
	//case BREAKPOINT:
	//
	//case IMPDEP1:
	//
	//case IMPDEP:

	default:
		panic(fmt.Errorf("Unsupported opcode: 0x%x!", opcode))
	}
}
