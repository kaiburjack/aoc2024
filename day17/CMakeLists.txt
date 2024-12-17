cmake_minimum_required(VERSION 3.30)
project(dynasmtest C)

set(CMAKE_C_STANDARD 11)

add_executable(dynasmtest main.c codegen.c)
