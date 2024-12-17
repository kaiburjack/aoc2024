#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <windows.h>

#include "codegen.h"

typedef struct Input_ {
    int64_t a, b, c;
    unsigned char insns[1024];
    int numInsns;
} Input;

Input parseInput(const char* fileName) {
    FILE* f = fopen(fileName, "r");
    if (f == NULL) {
        perror("fopen");
        exit(1);
    }
    fseek(f, 0, SEEK_END);
    long fsize = ftell(f);
    fseek(f, 0, SEEK_SET);
    char* buffer = calloc(fsize + 1, 1);
    fread(buffer, 1, fsize, f);
    fclose(f);
    Input start = {0};
    sscanf(buffer,
        "Register A: %lld\n"\
               "Register B: %lld\n"\
               "Register C: %lld\n\n",
               &start.a, &start.b, &start.c);
    const char *programLine = strstr(buffer, "Program: ");
    programLine += strlen("Program: "); // Skip "Program: "
    char *out = start.insns; // Output pointer to the destination buffer
    for (start.numInsns = 0; *programLine; start.numInsns++) {
        if (*programLine != ',') {
            *out++ = *programLine;
        }
        programLine++;
    }
    free(buffer);
    return start;
}

void* my_alloc(size_t size) {
    void* executableCode = VirtualAlloc(0, size, MEM_RESERVE | MEM_COMMIT, PAGE_READWRITE);
    return executableCode;
}

void* genCallableFunction(Input input, size_t* codeSize) {
    callbackFunctions callbacks = {my_alloc};
    void* code = codegen(input.insns, input.numInsns, &callbacks, codeSize);
    DWORD dwOld;
    VirtualProtect(code, *codeSize, PAGE_EXECUTE, &dwOld);
    FlushInstructionCache(GetCurrentProcess(), code, *codeSize);
    return code;
}

int main(void) {
    Input start = parseInput("real.txt");
    size_t codeLen;
    void* code = genCallableFunction(start, &codeLen);
    unsigned char* output = calloc(1024, 1);
    const unsigned char* (*func)(int64_t, int64_t, int64_t, void*) = code;
    const unsigned char* end = func(start.a, start.b, start.c, output);
    for (int i = 0; i < end-output; i++) {
        printf("%d", output[i]);
        if (i < end-output-1) {
            printf(",");
        }
    }
    free(output);
}
