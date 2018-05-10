#include "go_asm.h"
#include "funcdata.h"
#include "textflag.h"

DATA    asmGetStr_s0<>+0(SB)/8, $"This str"
DATA    asmGetStr_s0<>+8(SB)/8, $"ing is f"
DATA    asmGetStr_s0<>+16(SB)/8, $"rom asse"
DATA    asmGetStr_s0<>+24(SB)/8, $"mbly!\n"
GLOBL   asmGetStr_s0<>(SB), RODATA, $32

// func asmGetStr() string
TEXT	·asmGetStr(SB), NOSPLIT, $0-16
LEAQ	asmGetStr_s0<>(SB), AX
MOVQ	AX, ret+0(FP)
MOVQ	$30, ret+8(FP)
RET
