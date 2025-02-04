#include "textflag.h"

TEXT ·CpuId(SB), NOPTR|NOSPLIT, $0
	MOVD flag+0(FP), AX
	XORL CX, CX
	CPUID
	MOVD AX, eax+8(FP)
	MOVD BX, ebx+12(FP)
	MOVD CX, ecx+16(FP)
	MOVD DX, edx+24(FP)
	RET
