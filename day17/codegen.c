/*
** This file has been pre-processed with DynASM.
** https://luajit.org/dynasm.html
** DynASM version 1.5.0, DynASM x64 version 1.5.0
** DO NOT EDIT! The original file is in "codegen.dasc".
*/

#line 1 "codegen.dasc"
#include <stdio.h>

#include "dynasm/dasm_proto.h"
#include "dynasm/dasm_x86.h"
#include "codegen.h"

//|.arch x64
#if DASM_VERSION != 10500
#error "Version mismatch between DynASM and included encoding engine"
#endif
#line 8 "codegen.dasc"
//|.section code
#define DASM_SECTION_CODE	0
#define DASM_MAXSECTION		1
#line 9 "codegen.dasc"
//|.globals GLOB_
enum {
  GLOB__MAX
};
#line 10 "codegen.dasc"
//|.actionlist actionlist
static const unsigned char actionlist[92] = {
  177,235,255,76,137,193,255,76,137,201,255,76,137,209,255,77,137,203,77,137,
  194,73,137,200,73,137,209,255,249,255,73,211,232,255,73,129,252,241,239,255,
  72,131,225,7,73,137,201,255,77,33,192,15,133,245,255,77,49,209,255,128,225,
  7,65,136,11,73,252,255,195,255,77,137,193,73,211,252,233,255,77,137,194,73,
  211,252,234,255,249,76,137,216,195,255
};

#line 11 "codegen.dasc"

void encodeCombo(dasm_State** Dst, char operand) {
  if (operand >= '0' && operand <= '3') {
    //| mov cl, operand-'0'
    dasm_put(Dst, 0, operand-'0');
#line 15 "codegen.dasc"
  } else if (operand == '4') {
    //| mov rcx, r8
    dasm_put(Dst, 3);
#line 17 "codegen.dasc"
  } else if (operand == '5') {
    //| mov rcx, r9
    dasm_put(Dst, 7);
#line 19 "codegen.dasc"
  } else if (operand == '6') {
    //| mov rcx, r10
    dasm_put(Dst, 11);
#line 21 "codegen.dasc"
  } else {
    // throw error
    exit(2);
  }
}

// https://learn.microsoft.com/en-us/cpp/build/x64-calling-convention?view=msvc-170#calling-convention-defaults
//   integer args: RCX, RDX, R8, and R9
void* codegen(const unsigned char* opcodes, int opcodesLength, callbackFunctions* callbacks, size_t* codeSize) {
  dasm_State* state;
  dasm_State** Dst = &state;
  dasm_init(&state, DASM_MAXSECTION);
  dasm_setup(&state, actionlist);
  dasm_growpc(&state, opcodesLength);
  //| mov r11, r9 // <- output list
  //| mov r10, r8 // <- C register
  //| mov r8, rcx // <- A register
  //| mov r9, rdx // <- B register
  dasm_put(Dst, 15);
#line 39 "codegen.dasc"
  for (int i = 0; i < opcodesLength; i += 2) {
    //|=>i/2:
    dasm_put(Dst, 28, i/2);
#line 41 "codegen.dasc"
    char opcode = opcodes[i];
    char operand = opcodes[i + 1];
    switch (opcode) {
    case '0': // adv
      encodeCombo(Dst, operand);
      //| shr r8, cl
      dasm_put(Dst, 30);
#line 47 "codegen.dasc"
      break;
    case '1': // bxl
      //| xor r9, (operand-'0')
      dasm_put(Dst, 34, (operand-'0'));
#line 50 "codegen.dasc"
      break;
    case '2': // bst
      encodeCombo(Dst, operand);
      //| and rcx, 7
      //| mov r9, rcx
      dasm_put(Dst, 40);
#line 55 "codegen.dasc"
      break;
    case '3': // jnz (12 bytes)
      //| and r8, r8
      //| jnz =>operand-'0'
      dasm_put(Dst, 48, operand-'0');
#line 59 "codegen.dasc"
      break;
    case '4': // bxc
      //| xor r9, r10
      dasm_put(Dst, 55);
#line 62 "codegen.dasc"
      break;
    case '5': // out (12 bytes)
      encodeCombo(Dst, operand);
      //| and cl, 7
      //| mov [r11], cl
      //| inc r11
      dasm_put(Dst, 59);
#line 68 "codegen.dasc"
      break;
    case '6': // bdv
      encodeCombo(Dst, operand);
      //| mov r9, r8
      //| shr r9, cl
      dasm_put(Dst, 70);
#line 73 "codegen.dasc"
      break;
    case '7': // cdv
      encodeCombo(Dst, operand);
      //| mov r10, r8
      //| shr r10, cl
      dasm_put(Dst, 78);
#line 78 "codegen.dasc"
      break;
    }
  }
  //|=>opcodesLength/2:
  //| mov rax, r11
  //| ret
  dasm_put(Dst, 86, opcodesLength/2);
#line 84 "codegen.dasc"
  int status = dasm_link(&state, codeSize);
  void* code = callbacks->alloc(*codeSize);
  status = dasm_encode(&state, code);
  dasm_free(&state);
  return code;
}
