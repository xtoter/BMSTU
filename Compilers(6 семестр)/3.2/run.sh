#!/bin/bash
lex lexer.l
bison -d parser.y
gcc -o calc *.c
rm -f lex.yy.c parser.tab.?
./calc ./input.pas 80
rm -f  ./calc
