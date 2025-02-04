#include "textflag.h"

TEXT ·AddInt64Scalar(SB), NOSPLIT, $32
	MOVQ v0+0(FP), AX
	MOVQ v1+8(FP), BX
	MOVQ vo+16(FP), CX
	MOVQ l+24(FP), DX

	XORQ SI, SI
loop:
	MOVQ (AX)(SI*8), DI
	ADDQ (BX)(SI*8), DI
	MOVQ DI, (CX)
	INCQ SI
	CMPQ SI, DX
	JB loop
	RET


TEXT ·AddInt64SSE2(SB), NOSPLIT, $32
	MOVQ v0+0(FP), AX
	MOVQ v1+8(FP), BX
	MOVQ vo+16(FP), CX
	MOVQ l+24(FP), DX

	MOVQ DX, DI
	ANDQ $0xFFFFFFFFFFFFFFFB, DX
	ANDQ $0x4, DI

	TESTQ DX, DX
	JZ scalar_part
	XORQ SI, SI

vector_loop:
	VMOVDQU (AX)(SI*8), X0
	VMOVDQU 16(AX)(SI*8), X1
	PADDQ (BX)(SI*8), X0
	PADDQ 16(BX)(SI*8), X1
	VMOVDQU X0, (CX)(SI*8)
	VMOVDQU X1, 16(CX)(SI*8)

	ADDQ $4, SI
	CMPQ SI, DX
	JB vector_loop

	// Can be here since we know for sure
	// slices len isn't 0
	TESTQ DI, DI
	JZ end
scalar_part:
	LEAQ (AX)(DX*8), AX
	LEAQ (BX)(DX*8), BX
	LEAQ (CX)(DX*8), CX

	MOVQ (AX), SI
	ADDQ (BX), SI
	MOVQ SI, (CX)
	CMPQ DI, $1
	JE end

	MOVQ 8(AX), SI
	ADDQ 8(BX), SI
	MOVQ SI, 8(CX)
	CMPQ DI, $2
	JE end

	MOVQ 16(AX), SI
	ADDQ 16(BX), SI
	MOVQ SI, 16(CX)

end:
	RET



TEXT ·AddInt64AVX2(SB), NOSPLIT, $32
	MOVQ v0+0(FP), AX
	MOVQ v1+8(FP), BX
	MOVQ vo+16(FP), CX
	MOVQ l+24(FP), DX

	MOVQ DX, DI
	ANDQ $0xFFFFFFFFFFFFFFF8, DX
	ANDQ $0x7, DI

	TESTQ DX, DX
	JZ scalar_part
	XORQ SI, SI

vector_loop:
	VMOVDQU (AX)(SI*8), Y0
	VMOVDQU 32(AX)(SI*8), Y1
	VPADDQ (BX)(SI*8), Y0, Y0
	VPADDQ 32(BX)(SI*8), Y1, Y1
	VMOVDQU Y0, (CX)(SI*8)
	VMOVDQU Y1, 32(CX)(SI*8)

	ADDQ $8, SI
	CMPQ SI, DX
	JB vector_loop

	// Can be here since we know for sure
	// slices len isn't 0
	TESTQ DI, DI
	JZ end
scalar_part:
	LEAQ (AX)(DX*8), AX
	LEAQ (BX)(DX*8), BX
	LEAQ (CX)(DX*8), CX

	MOVQ (AX), SI
	ADDQ (BX), SI
	MOVQ SI, (CX)
	CMPQ DI, $1
	JE end

	MOVQ 8(AX), SI
	ADDQ 8(BX), SI
	MOVQ SI, 8(CX)
	CMPQ DI, $2
	JE end

	MOVQ 16(AX), SI
	ADDQ 16(BX), SI
	MOVQ SI, 16(CX)
	CMPQ DI, $3
	JE end

	MOVQ 24(AX), SI
	ADDQ 24(BX), SI
	MOVQ SI, 24(CX)
	CMPQ DI, $4
	JE end

	MOVQ 32(AX), SI
	ADDQ 32(BX), SI
	MOVQ SI, 32(CX)
	CMPQ DI, $5
	JE end

	MOVQ 40(AX), SI
	ADDQ 40(BX), SI
	MOVQ SI, 40(CX)
	CMPQ DI, $6
	JE end

	MOVQ 48(AX), SI
	ADDQ 48(BX), SI
	MOVQ SI, 48(CX)

end:
	VZEROUPPER
	RET
