package chapter5_instructions

import (
	"GoVM/chapter5-instructions/base"
	"fmt"
	"GoVM/chapter5-instructions/constants"
	"GoVM/chapter5-instructions/loads"
	"GoVM/chapter5-instructions/stores"
	"GoVM/chapter5-instructions/stack"
	"GoVM/chapter5-instructions/math"
	"GoVM/chapter5-instructions/conversions"
	"GoVM/chapter5-instructions/comparisons"
	"GoVM/chapter5-instructions/control"
	"GoVM/chapter5-instructions/extended"
	"GoVM/chapter5-instructions/references"
)

// NoOperandsInstruction singletons
var (
	nop         = &constants.NOP{}
	aconst_null = &constants.ACONST_NULL{}
	iconst_m1   = &constants.ICONST_M1{}
	iconst_0    = &constants.ICONST_0{}
	iconst_1    = &constants.ICONST_1{}
	iconst_2    = &constants.ICONST_2{}
	iconst_3    = &constants.ICONST_3{}
	iconst_4    = &constants.ICONST_4{}
	iconst_5    = &constants.ICONST_5{}
	lconst_0    = &constants.LCONST_0{}
	lconst_1    = &constants.LCONST_1{}
	fconst_0    = &constants.FCONST_0{}
	fconst_1    = &constants.FCONST_1{}
	fconst_2    = &constants.FCONST_2{}
	dconst_0    = &constants.DCONST_0{}
	dconst_1    = &constants.DCONST_1{}
	iload_0     = &loads.ILOAD_0{}
	iload_1     = &loads.ILOAD_1{}
	iload_2     = &loads.ILOAD_2{}
	iload_3     = &loads.ILOAD_3{}
	lload_0     = &loads.LLOAD_0{}
	lload_1     = &loads.LLOAD_1{}
	lload_2     = &loads.LLOAD_2{}
	lload_3     = &loads.LLOAD_3{}
	fload_0     = &loads.FLOAD_0{}
	fload_1     = &loads.FLOAD_1{}
	fload_2     = &loads.FLOAD_2{}
	fload_3     = &loads.FLOAD_3{}
	dload_0     = &loads.DLOAD_0{}
	dload_1     = &loads.DLOAD_1{}
	dload_2     = &loads.DLOAD_2{}
	dload_3     = &loads.DLOAD_3{}
	aload_0     = &loads.ALOAD_0{}
	aload_1     = &loads.ALOAD_1{}
	aload_2     = &loads.ALOAD_2{}
	aload_3     = &loads.ALOAD_3{}
	// iaload      = &IALOAD{}
	// laload      = &LALOAD{}
	// faload      = &FALOAD{}
	// daload      = &DALOAD{}
	// aaload      = &AALOAD{}
	// baload      = &BALOAD{}
	// caload      = &CALOAD{}
	// saload      = &SALOAD{}
	istore_0 = &stores.ISTORE_0{}
	istore_1 = &stores.ISTORE_1{}
	istore_2 = &stores.ISTORE_2{}
	istore_3 = &stores.ISTORE_3{}
	lstore_0 = &stores.LSTORE_0{}
	lstore_1 = &stores.LSTORE_1{}
	lstore_2 = &stores.LSTORE_2{}
	lstore_3 = &stores.LSTORE_3{}
	fstore_0 = &stores.FSTORE_0{}
	fstore_1 = &stores.FSTORE_1{}
	fstore_2 = &stores.FSTORE_2{}
	fstore_3 = &stores.FSTORE_3{}
	dstore_0 = &stores.DSTORE_0{}
	dstore_1 = &stores.DSTORE_1{}
	dstore_2 = &stores.DSTORE_2{}
	dstore_3 = &stores.DSTORE_3{}
	astore_0 = &stores.ASTORE_0{}
	astore_1 = &stores.ASTORE_1{}
	astore_2 = &stores.ASTORE_2{}
	astore_3 = &stores.ASTORE_3{}
	// iastore  = &IASTORE{}
	// lastore  = &LASTORE{}
	// fastore  = &FASTORE{}
	// dastore  = &DASTORE{}
	// aastore  = &AASTORE{}
	// bastore  = &BASTORE{}
	// castore  = &CASTORE{}
	// sastore  = &SASTORE{}
	pop     = &stack.POP{}
	pop2    = &stack.POP2{}
	dup     = &stack.DUP{}
	dup_x1  = &stack.DUP_X1{}
	dup_x2  = &stack.DUP_X2{}
	dup2    = &stack.DUP2{}
	dup2_x1 = &stack.DUP2_X1{}
	dup2_x2 = &stack.DUP2_X2{}
	swap    = &stack.SWAP{}
	iadd    = &math.IADD{}
	ladd    = &math.LADD{}
	fadd    = &math.FADD{}
	dadd    = &math.DADD{}
	isub    = &math.ISUB{}
	lsub    = &math.LSUB{}
	fsub    = &math.FSUB{}
	dsub    = &math.DSUB{}
	imul    = &math.IMUL{}
	lmul    = &math.LMUL{}
	fmul    = &math.FMUL{}
	dmul    = &math.DMUL{}
	idiv    = &math.IDIV{}
	ldiv    = &math.LDIV{}
	fdiv    = &math.FDIV{}
	ddiv    = &math.DDIV{}
	irem    = &math.IREM{}
	lrem    = &math.LREM{}
	frem    = &math.FREM{}
	drem    = &math.DREM{}
	ineg    = &math.INEG{}
	lneg    = &math.LNEG{}
	fneg    = &math.FNEG{}
	dneg    = &math.DNEG{}
	ishl    = &math.ISHL{}
	lshl    = &math.LSHL{}
	ishr    = &math.ISHR{}
	lshr    = &math.LSHR{}
	iushr   = &math.IUSHR{}
	lushr   = &math.LUSHR{}
	iand    = &math.IAND{}
	land    = &math.LAND{}
	ior     = &math.IOR{}
	lor     = &math.LOR{}
	ixor    = &math.IXOR{}
	lxor    = &math.LXOR{}
	i2l     = &conversions.I2L{}
	i2f     = &conversions.I2F{}
	i2d     = &conversions.I2D{}
	l2i     = &conversions.L2I{}
	l2f     = &conversions.L2F{}
	l2d     = &conversions.L2D{}
	f2i     = &conversions.F2I{}
	f2l     = &conversions.F2L{}
	f2d     = &conversions.F2D{}
	d2i     = &conversions.D2I{}
	d2l     = &conversions.D2L{}
	d2f     = &conversions.D2F{}
	i2b     = &conversions.I2B{}
	i2c     = &conversions.I2C{}
	i2s     = &conversions.I2S{}
	lcmp    = &comparisons.LCMP{}
	fcmpl   = &comparisons.FCMPL{}
	fcmpg   = &comparisons.FCMPG{}
	dcmpl   = &comparisons.DCMPL{}
	dcmpg   = &comparisons.DCMPG{}
	ireturn = &control.IRETURN{}
	lreturn = &control.LRETURN{}
	freturn = &control.FRETURN{}
	dreturn = &control.DRETURN{}
	areturn = &control.ARETURN{}
	_return = &control.RETURN{}
	// arraylength   = &ARRAY_LENGTH{}
	// athrow        = &ATHROW{}
	// monitorenter  = &MONITOR_ENTER{}
	// monitorexit   = &MONITOR_EXIT{}
	// invoke_native = &INVOKE_NATIVE{}
)

func NewInstruction(opcode byte) base.Instruction {
	switch opcode {
	case 0x00:
		return nop
	case 0x01:
		return aconst_null
	case 0x02:
		return iconst_m1
	case 0x03:
		return iconst_0
	case 0x04:
		return iconst_1
	case 0x05:
		return iconst_2
	case 0x06:
		return iconst_3
	case 0x07:
		return iconst_4
	case 0x08:
		return iconst_5
	case 0x09:
		return lconst_0
	case 0x0a:
		return lconst_1
	case 0x0b:
		return fconst_0
	case 0x0c:
		return fconst_1
	case 0x0d:
		return fconst_2
	case 0x0e:
		return dconst_0
	case 0x0f:
		return dconst_1
	case 0x10:
		return &constants.BIPUSH{}
	case 0x11:
		return &constants.SIPUSH{}
	 case 0x12:
	 	return &constants.LDC{}
	 case 0x13:
	 	return &constants.LDC_W{}
	 case 0x14:
	 	return &constants.LDC2_W{}
	case 0x15:
		return &loads.ILOAD{}
	case 0x16:
		return &loads.LLOAD{}
	case 0x17:
		return &loads.FLOAD{}
	case 0x18:
		return &loads.DLOAD{}
	case 0x19:
		return &loads.ALOAD{}
	case 0x1a:
		return iload_0
	case 0x1b:
		return iload_1
	case 0x1c:
		return iload_2
	case 0x1d:
		return iload_3
	case 0x1e:
		return lload_0
	case 0x1f:
		return lload_1
	case 0x20:
		return lload_2
	case 0x21:
		return lload_3
	case 0x22:
		return fload_0
	case 0x23:
		return fload_1
	case 0x24:
		return fload_2
	case 0x25:
		return fload_3
	case 0x26:
		return dload_0
	case 0x27:
		return dload_1
	case 0x28:
		return dload_2
	case 0x29:
		return dload_3
	case 0x2a:
		return aload_0
	case 0x2b:
		return aload_1
	case 0x2c:
		return aload_2
	case 0x2d:
		return aload_3
	// case 0x2e:
	// 	return iaload
	// case 0x2f:
	// 	return laload
	// case 0x30:
	// 	return faload
	// case 0x31:
	// 	return daload
	// case 0x32:
	// 	return aaload
	// case 0x33:
	// 	return baload
	// case 0x34:
	// 	return caload
	// case 0x35:
	// 	return saload
	case 0x36:
		return &stores.ISTORE{}
	case 0x37:
		return &stores.LSTORE{}
	case 0x38:
		return &stores.FSTORE{}
	case 0x39:
		return &stores.DSTORE{}
	case 0x3a:
		return &stores.ASTORE{}
	case 0x3b:
		return istore_0
	case 0x3c:
		return istore_1
	case 0x3d:
		return istore_2
	case 0x3e:
		return istore_3
	case 0x3f:
		return lstore_0
	case 0x40:
		return lstore_1
	case 0x41:
		return lstore_2
	case 0x42:
		return lstore_3
	case 0x43:
		return fstore_0
	case 0x44:
		return fstore_1
	case 0x45:
		return fstore_2
	case 0x46:
		return fstore_3
	case 0x47:
		return dstore_0
	case 0x48:
		return dstore_1
	case 0x49:
		return dstore_2
	case 0x4a:
		return dstore_3
	case 0x4b:
		return astore_0
	case 0x4c:
		return astore_1
	case 0x4d:
		return astore_2
	case 0x4e:
		return astore_3
	// case 0x4f:
	// 	return iastore
	// case 0x50:
	// 	return lastore
	// case 0x51:
	// 	return fastore
	// case 0x52:
	// 	return dastore
	// case 0x53:
	// 	return aastore
	// case 0x54:
	// 	return bastore
	// case 0x55:
	// 	return castore
	// case 0x56:
	// 	return sastore
	case 0x57:
		return pop
	case 0x58:
		return pop2
	case 0x59:
		return dup
	case 0x5a:
		return dup_x1
	case 0x5b:
		return dup_x2
	case 0x5c:
		return dup2
	case 0x5d:
		return dup2_x1
	case 0x5e:
		return dup2_x2
	case 0x5f:
		return swap
	case 0x60:
		return iadd
	case 0x61:
		return ladd
	case 0x62:
		return fadd
	case 0x63:
		return dadd
	case 0x64:
		return isub
	case 0x65:
		return lsub
	case 0x66:
		return fsub
	case 0x67:
		return dsub
	case 0x68:
		return imul
	case 0x69:
		return lmul
	case 0x6a:
		return fmul
	case 0x6b:
		return dmul
	case 0x6c:
		return idiv
	case 0x6d:
		return ldiv
	case 0x6e:
		return fdiv
	case 0x6f:
		return ddiv
	case 0x70:
		return irem
	case 0x71:
		return lrem
	case 0x72:
		return frem
	case 0x73:
		return drem
	case 0x74:
		return ineg
	case 0x75:
		return lneg
	case 0x76:
		return fneg
	case 0x77:
		return dneg
	case 0x78:
		return ishl
	case 0x79:
		return lshl
	case 0x7a:
		return ishr
	case 0x7b:
		return lshr
	case 0x7c:
		return iushr
	case 0x7d:
		return lushr
	case 0x7e:
		return iand
	case 0x7f:
		return land
	case 0x80:
		return ior
	case 0x81:
		return lor
	case 0x82:
		return ixor
	case 0x83:
		return lxor
	case 0x84:
		return &math.IINC{}
	case 0x85:
		return i2l
	case 0x86:
		return i2f
	case 0x87:
		return i2d
	case 0x88:
		return l2i
	case 0x89:
		return l2f
	case 0x8a:
		return l2d
	case 0x8b:
		return f2i
	case 0x8c:
		return f2l
	case 0x8d:
		return f2d
	case 0x8e:
		return d2i
	case 0x8f:
		return d2l
	case 0x90:
		return d2f
	case 0x91:
		return i2b
	case 0x92:
		return i2c
	case 0x93:
		return i2s
	case 0x94:
		return lcmp
	case 0x95:
		return fcmpl
	case 0x96:
		return fcmpg
	case 0x97:
		return dcmpl
	case 0x98:
		return dcmpg
	case 0x99:
		return &comparisons.IFEQ{}
	case 0x9a:
		return &comparisons.IFNE{}
	case 0x9b:
		return &comparisons.IFLT{}
	case 0x9c:
		return &comparisons.IFGE{}
	case 0x9d:
		return &comparisons.IFGT{}
	case 0x9e:
		return &comparisons.IFLE{}
	case 0x9f:
		return &comparisons.IF_ICMPEQ{}
	case 0xa0:
		return &comparisons.IF_ICMPNE{}
	case 0xa1:
		return &comparisons.IF_ICMPLT{}
	case 0xa2:
		return &comparisons.IF_ICMPGE{}
	case 0xa3:
		return &comparisons.IF_ICMPGT{}
	case 0xa4:
		return &comparisons.IF_ICMPLE{}
	case 0xa5:
		return &comparisons.IF_ACMPEQ{}
	case 0xa6:
		return &comparisons.IF_ACMPNE{}
	case 0xa7:
		return &control.GOTO{}
	// case 0xa8:
	// 	return &JSR{}
	// case 0xa9:
	// 	return &RET{}
	case 0xaa:
		return &control.TABLE_SWITCH{}
	case 0xab:
		return &control.LOOKUP_SWITCH{}
	case 0xac:
	 	return ireturn
	case 0xad:
	 	return lreturn
	case 0xae:
	 	return freturn
	case 0xaf:
	 	return dreturn
	case 0xb0:
	 	return areturn
	case 0xb1:
	 	return _return
	case 0xb2:
		return &references.GET_STATIC{}
	 case 0xb3:
	 	return &references.PUT_STATIC{}
	 case 0xb4:
	 	return &references.GET_FIELD{}
	 case 0xb5:
	 	return &references.PUT_FIELD{}
	case 0xb6:
		return &references.INVOKE_VIRTUAL{}
	 case 0xb7:
	 	return &references.INVOKE_SPECIAL{}
	 case 0xb8:
		return &references.INVOKE_STATIC{}
	 case 0xb9:
	 	return &references.INVOKE_INTERFACE{}
	// case 0xba:
	// 	return &INVOKE_DYNAMIC{}
	 case 0xbb:
	 	return &references.NEW{}
	// case 0xbc:
	// 	return &NEW_ARRAY{}
	// case 0xbd:
	// 	return &ANEW_ARRAY{}
	// case 0xbe:
	// 	return arraylength
	// case 0xbf:
	// 	return athrow
	 case 0xc0:
	 	return &references.CHECK_CAST{}
	 case 0xc1:
	 	return &references.INSTANCE_OF{}
	// case 0xc2:
	// 	return monitorenter
	// case 0xc3:
	// 	return monitorexit
	case 0xc4:
		return &extended.WIDE{}
	// case 0xc5:
	// 	return &MULTI_ANEW_ARRAY{}
	case 0xc6:
		return &extended.IFNULL{}
	case 0xc7:
		return &extended.IFNONNULL{}
	case 0xc8:
		return &extended.GOTO_W{}
	// case 0xc9:
	// 	return &JSR_W{}
	// case 0xca: breakpoint
	// case 0xfe: impdep1
	// case 0xff: impdep2
	default:
		panic(fmt.Errorf("Unsupported opcode: 0x%x!", opcode))
	}
}
