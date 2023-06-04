%{
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "lexer.h"
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
void printTab(int length,long* env) {
    for (int i = 0; i < length; i++) {
        printf("    ");
        env[2]+=4;
    }
}
void checkNewLine(long* env) {
    if (env[2] > env[1]) {
        printf("\n");
        env[2]=0;
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
    compound_statement {printf("\n");env[2]=0;}
    | functioncall
    | for_loop
    | appropriation
    ;
appropriation:
    IDENTIFIER ASSIGN 
    {printTab(env[0],env); printf("%s", $1);env[2]+=strlen($1);checkNewLine(env);printf(" := ");env[2]+=4;checkNewLine(env);}
    expr SEMICOLON
    {printf(";\n");env[2]=0;}
    ;
expr:
    value
    | value OP {printf("%c", $OP);env[2]+=getLongLen($OP);checkNewLine(env);} expr 
    ;
value:
    IDENTIFIER  {printf("%s", $IDENTIFIER);env[2]+=strlen($IDENTIFIER);checkNewLine(env);}
    |NUMBER  {printf("%ld", $NUMBER);env[2]+=getLongLen($NUMBER);checkNewLine(env);}
    | LBRACKET  {printf("\(");env[2]+=1;checkNewLine(env);}
                expr RBRACKET 
                {printf("\)");env[2]+=1;checkNewLine(env);}
    ;
functioncall:
    IDENTIFIER LBRACKET 
    {printTab(env[0],env); printf("%s",$1);env[2]+=strlen($1);checkNewLine(env);printf("\(");env[2]+=1;checkNewLine(env);}
    variables 
    RBRACKET SEMICOLON
    {printf("\);\n");env[2]=0;}
    ;
statement_begin:
    variable_declaration
    | compound_statement_begin
    | function
    ;

variable_declaration:
    {printf("var\n");env[0]++;env[2]=0;}
    VAR variable_list {env[0]--;}
    ;

variable_list:
    variable_declaration_line SEMICOLON  {printf(";\n");env[2]=0;}
    | variable_list  variable_declaration_line SEMICOLON  {printf(";\n");env[2]=0;}
    ;

variable_declaration_line:
    {printTab(env[0],env);} variables COLON IDENTIFIER {printf(": ");env[2]+=2;checkNewLine(env);printf("%s", $IDENTIFIER);env[2]+=strlen($IDENTIFIER);checkNewLine(env);}
    ;

variables:
    IDENTIFIER {printf("%s", $IDENTIFIER);env[2]+=strlen($IDENTIFIER);checkNewLine(env);}
    | variables COMMA IDENTIFIER {printf(", ");env[2]+=2;checkNewLine(env);printf("%s", $IDENTIFIER);env[2]+=strlen($IDENTIFIER);checkNewLine(env);}
    ;
function:
    FUNCTION IDENTIFIER LBRACKET 
    {printf("function");env[2]+=9;checkNewLine(env);printf("%s\(", $2);env[2]+=1+strlen($2);checkNewLine(env);}
    variable_declaration_line 
    RBRACKET COLON IDENTIFIER SEMICOLON 
    {printf("): ");env[2]+=3;checkNewLine(env);printf("%s;\n", $8);}
    compound_statement SEMICOLON
    {printf(";\n");env[2]=0;}
    ;

compound_statement:
    {printTab(env[0],env); printf("begin\n"); env[0]++;env[2]=0;}
    BEGIN_T statement_list END {env[0]--; printTab(env[0],env);  printf("end");env[2]+=3;}
    ;
compound_statement_begin:
    {printf("begin\n");env[0]++;env[2]=0;}
    BEGIN_T statement_list END POINT {env[0]--;printf("end.\n");env[2]=0;}
    ;

for_loop:
    
    FOR IDENTIFIER ASSIGN NUMBER TO NUMBER DO 
    {printTab(env[0],env);printf("for ");env[2]+=4;checkNewLine(env);
                    printf("%s", $2);env[2]+=strlen($2);checkNewLine(env);
                    printf(" := ");env[2]+=4;checkNewLine(env);
                    printf("%ld", $4);env[2]+=getLongLen($4);checkNewLine(env);
                    printf(" to ");env[2]+=4;checkNewLine(env);
                    printf("%ld",$6);env[2]+=getLongLen($6);checkNewLine(env);
                    printf(" do\n" );env[2]=0;}
    statement
    ;

%%

int main(int argc, char *argv[]) {
   
    FILE *input = 0;
    long env[26] = { 0 };
    env[0] = 0;
    yyscan_t scanner;
    struct Extra extra;

    if (argc > 1) {
        printf("Read file %s\n", argv[1]);
        input = fopen(argv[1], "r");
        env[1] = atoi(argv[2]);
        printf("Len %ld\n", env[1]);
        
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
