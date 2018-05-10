#include "go_asm.h"
#include "funcdata.h"
#include "textflag.h"

DATA    asmGetStr_s0<>+0(SB)/8, $"This str"
DATA    asmGetStr_s0<>+8(SB)/8, $"ing is f"
DATA    asmGetStr_s0<>+16(SB)/8, $"rom asse"
DATA    asmGetStr_s0<>+24(SB)/8, $"mbly!\n"
GLOBL   asmGetStr_s0<>(SB), RODATA, $32

// func asmGetStr() string
TEXT	·asmGetStr(SB), NOSPLIT, $-4-8
MOVW	$asmGetStr_s0<>(SB), R0
MOVW	R0, ret+0(FP)
MOVW	$30, R0
MOVW	R0, ret+4(FP)
RET
