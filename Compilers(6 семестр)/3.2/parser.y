%{
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "lexer.h"
int indent;
int maxlen;
int curlen;
%}

%define api.pure
%locations
%lex-param {yyscan_t scanner}  /* параметр для yylex() */
/* параметры для yyparse() */
%parse-param {yyscan_t scanner}
%parse-param {long env[26]}

%union {
    char* indentifier;
    char variable;
    long number;
}

%token VAR BEGIN_T END FOR TO DO ASSIGN SEMICOLON COLON COMMA POINT LBRACKET RBRACKET FUNCTION 
%token <indentifier> IDENTIFIER
%token <number> NUMBER
%token <variable> OP

%{
int yylex(YYSTYPE *yylval_param, YYLTYPE *yylloc_param, yyscan_t scanner);
void yyerror(YYLTYPE *loc, yyscan_t scanner, long env[26], const char *message);
void printTab(int length) {
    for (int i = 0; i < length; i++) {
        printf("    ");
        curlen+=4;
    }
}
void checkNewLine() {
    if (curlen > maxlen) {
        printf("\n");
        curlen=0;
    }
}
int getLongLen(long number) {
    char buffer[20];  // Буфер для хранения строки
    int length = snprintf(buffer, sizeof(buffer), "%ld", number);
    return length;
}
%}

%%

program:
    statement_list_begin 
    ;

statement_list_begin:
    statement_begin
    | statement_list_begin  statement_begin
    ;

statement_list:
    statement
    | statement_list  statement
    ;

statement:
    compound_statement {printf("\n");curlen=0;}
    | functioncall
    | for_loop
    | appropriation
    ;
appropriation:
    IDENTIFIER ASSIGN 
    {printTab(indent); printf("%s", $1);curlen+=strlen($1);checkNewLine();printf(" := ");curlen+=4;checkNewLine();}
    expr SEMICOLON
    {printf(";\n");curlen=0;}
    ;
expr:
    value
    | value OP {printf("%c", $OP);curlen+=getLongLen($OP);checkNewLine();} value 
    ;
value:
    IDENTIFIER  {printf("%s", $IDENTIFIER);curlen+=strlen($IDENTIFIER);checkNewLine();}
    |NUMBER  {printf("%ld", $NUMBER);curlen+=getLongLen($NUMBER);checkNewLine();}
    | LBRACKET  {printf("\(");curlen+=1;checkNewLine();}
                expr RBRACKET 
                {printf("\)");curlen+=1;checkNewLine();}
    ;
functioncall:
    IDENTIFIER LBRACKET 
    {printTab(indent); printf("%s",$1);curlen+=strlen($1);checkNewLine();printf("\(");curlen+=1;checkNewLine();}
    variables 
    RBRACKET SEMICOLON
    {printf("\);\n");curlen=0;}
    ;
statement_begin:
    variable_declaration
    | compound_statement_begin
    | function
    ;

variable_declaration:
    {printf("var\n");indent++;curlen=0;}
    VAR variable_list {indent--;}
    ;

variable_list:
    variable_declaration_line SEMICOLON  {printf(";\n");curlen=0;}
    | variable_list  variable_declaration_line SEMICOLON  {printf(";\n");curlen=0;}
    ;

variable_declaration_line:
    {printTab(indent);} variables COLON IDENTIFIER {printf(": ");curlen+=2;checkNewLine();printf("%s", $IDENTIFIER);curlen+=strlen($IDENTIFIER);checkNewLine();}
    ;

variables:
    IDENTIFIER {printf("%s", $IDENTIFIER);curlen+=strlen($IDENTIFIER);checkNewLine();}
    | variables COMMA IDENTIFIER {printf(", ");curlen+=2;checkNewLine();printf("%s", $IDENTIFIER);curlen+=strlen($IDENTIFIER);checkNewLine();}
    ;
function:
    FUNCTION IDENTIFIER LBRACKET 
    {printf("function");curlen+=9;checkNewLine();printf("%s\(", $2);curlen+=1+strlen($2);checkNewLine();}
    variable_declaration_line 
    RBRACKET COLON IDENTIFIER SEMICOLON 
    {printf("): ");curlen+=3;checkNewLine();printf("%s;\n", $8);}
    compound_statement SEMICOLON
    {printf(";\n");curlen=0;}
    ;

compound_statement:
    {printTab(indent); printf("begin\n"); indent++;curlen=0;}
    BEGIN_T statement_list END {indent--; printTab(indent);  printf("end");curlen+=3;}
    ;
compound_statement_begin:
    {printf("begin\n");indent++;curlen=0;}
    BEGIN_T statement_list END POINT {indent--;printf("end.\n");curlen=0;}
    ;

for_loop:
    
    FOR IDENTIFIER ASSIGN NUMBER TO NUMBER DO 
    {printTab(indent);printf("for ");curlen+=4;checkNewLine();
                    printf("%s", $2);curlen+=strlen($2);checkNewLine();
                    printf(" := ");curlen+=4;checkNewLine();
                    printf("%ld", $4);curlen+=getLongLen($4);checkNewLine();
                    printf(" to ");curlen+=4;checkNewLine();
                    printf("%ld",$6);curlen+=getLongLen($6);checkNewLine();
                    printf(" do\n" );curlen=0;}
    statement
    ;

%%

int main(int argc, char *argv[]) {
    indent = 0;
    FILE *input = 0;
    long env[26] = { 0 };
    yyscan_t scanner;
    struct Extra extra;

    if (argc > 1) {
        printf("Read file %s\n", argv[1]);
        input = fopen(argv[1], "r");
        maxlen = atoi(argv[2]);
        printf("Len %d\n", maxlen);
        
    } else {
        printf("No file in command line, use stdin\n");
        input = stdin;
    }

    init_scanner(input, &scanner, &extra);
    yyparse(scanner, env);
    destroy_scanner(scanner);

    if (input != stdin) {
        fclose(input);
    }

    return 0;
}
