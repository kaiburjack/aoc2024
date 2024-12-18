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
    for (start.numInsns = 0; *programLine; ) {
        if (*programLine != ',') {
            *out++ = *programLine;
            start.numInsns++;
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

void* part1Func(Input input, size_t* codeSize) {
    callbackFunctions callbacks = {my_alloc};
    void* code = codegen(input.insns, input.numInsns, &callbacks, codeSize);
    DWORD dwOld;
    VirtualProtect(code, *codeSize, PAGE_EXECUTE, &dwOld);
    FlushInstructionCache(GetCurrentProcess(), code, *codeSize);
    return code;
}

void* part2Func(Input input, size_t* codeSize) {
    callbackFunctions callbacks = {my_alloc};
    void* code = codegen2(input.insns, input.numInsns, &callbacks, codeSize);
    DWORD dwOld;
    VirtualProtect(code, *codeSize, PAGE_EXECUTE, &dwOld);
    FlushInstructionCache(GetCurrentProcess(), code, *codeSize);
    return code;
}

typedef struct ThreadParameter_ {
    unsigned char* expectedOutput;
    size_t expectedOutputSize;
    int (*func)(void*, void*, int64_t);
    int64_t a;
    int64_t increment;
} ThreadParameter;

DWORD WINAPI ThreadFunc(void* d) {
    ThreadParameter* data = d;
    int64_t a = data->a;
    int64_t counter = 0;
    while (!data->func(data->expectedOutput, data->expectedOutput+data->expectedOutputSize, a)) {
        if (counter > 1000000000LL) {
            printf("a: %lld\n", a);
            counter = 0;
        }
        counter++;
        a += data->increment;
    }
    fprintf(stderr, "Found a: %lld\n", a);
    exit(0);
}

int main(void) {
    Input start = parseInput("real.txt");
    size_t codeLen;
    {
        // part1
        void* code = part1Func(start, &codeLen);
        unsigned char* output = calloc(1024, 1);
        const unsigned char* (*func)(int64_t, int64_t, int64_t, void*) = code;
        const unsigned char* end = func(start.a, start.b, start.c, output);
        for (int i = 0; i < end-output; i++) {
            printf("%d", output[i]);
            if (i < end-output-1) {
                printf(",");
            }
        }
        printf("\n");
        free(output);
    }
    // part 2
    {
        void* code = part2Func(start, &codeLen);
        unsigned char* expectedOutput = calloc(1024, 1);
        memcpy(expectedOutput, start.insns, start.numInsns);
        int (*func)(void*, void*, int64_t) = code;
        const int N = 16;
        int64_t a = 55527112632703LL;
        HANDLE threads[N];
        // create 32 Win32 threads
        ThreadParameter data[N];
        for (int i = 0; i < N; i++) {
            data[i] = (ThreadParameter){
                .expectedOutput = expectedOutput,
                .expectedOutputSize = start.numInsns,
                .func = func,
                .a = a + i,
                .increment = N
            };
            threads[i] = CreateThread(NULL, 0, &ThreadFunc, &data[i], 0, NULL);
        }
        // wait for all threads to finish
        WaitForMultipleObjects(N, threads, TRUE, INFINITE);
        free(expectedOutput);
    }
}
