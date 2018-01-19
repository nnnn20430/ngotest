#include "go_asm.h"
#include "funcdata.h"
#include "textflag.h"

DATA    asm_GetStr_s0<>+0(SB)/8, $"This str"
DATA    asm_GetStr_s0<>+8(SB)/8, $"ing is f"
DATA    asm_GetStr_s0<>+16(SB)/8, $"rom asse"
DATA    asm_GetStr_s0<>+24(SB)/8, $"mbly!\n"
GLOBL   asm_GetStr_s0<>(SB), RODATA, $32

// func asm_GetStr() string
TEXT	·asm_GetStr(SB), NOSPLIT, $-8-16
MOVD	$asm_GetStr_s0<>(SB), R0
MOVD	R0, ret+0(FP)
MOVD	$30, R0
MOVD	R0, ret+8(FP)
RET
