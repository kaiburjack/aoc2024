#ifndef CODEGEN_H
#define CODEGEN_H
typedef struct callbackFunctions_ {
    void* (*alloc)(size_t size);
} callbackFunctions;
extern void* codegen(const unsigned char* opcodes, int opcodesLength, callbackFunctions* callbacks, size_t* codeSize);
#endif
