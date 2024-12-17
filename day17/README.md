# Regenerate code generator

./minilua dynasm/dynasm.lua -o codegen.c -D X64 codegen.dasc

# Caveats

Currently, this code assumes Win64 calling convention!