cmake_minimum_required(VERSION 3.15)
project(lab4)

set(CMAKE_CXX_STANDARD 14)



include_directories(ControlFlowGraph)
include_directories(IRGeneration)
include_directories(Lexer)
include_directories(Parser)




add_executable(lab4 main.cpp)


add_library(lexer
    Lexer/Lexer.cpp
    Lexer/Lexer.hpp
    Lexer/Token.cpp
    Lexer/Token.hpp
)
add_library(parser
    Parser/Expressions.cpp
    Parser/Expressions.hpp
    Parser/Parser.cpp
    Parser/Parser.hpp
)
add_library(irgeneration
    IRGeneration/AbstractVisitor.hpp
    IRGeneration/IRGenerator.hpp
    IRGeneration/IRGenerator.cpp
)

add_library(cfg
    ControlFlowGraph/BasicBlock.hpp
    ControlFlowGraph/BasicBlock.cpp
    ControlFlowGraph/ControlFlowGraph.hpp
    ControlFlowGraph/ControlFlowGraph.cpp
    ControlFlowGraph/IRStatements.hpp
    ControlFlowGraph/IRStatements.cpp
)
target_link_libraries(lab4  parser lexer irgeneration cfg )
