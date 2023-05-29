%{
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
%}

%token VAR BEGIN_T END FOR TO DO NUMBER IDENTIFIER NEWLINE ASSIGN

%%

program:
    statement_list
    ;

statement_list:
    statement
    | statement_list NEWLINE statement
    ;

statement:
    variable_declaration
    | compound_statement
    | for_loop
    ;

variable_declaration:
    VAR variable_list
    ;

variable_list:
    variable_declaration_line
    | variable_list ';' variable_declaration_line
    ;

variable_declaration_line:
    IDENTIFIER ':' IDENTIFIER
    ;

compound_statement:
    BEGIN_T statement_list END
    ;

for_loop:
    FOR IDENTIFIER ASSIGN NUMBER TO NUMBER DO statement
    ;

%%

int main() {
    yyparse();
    return 0;
}

void yyerror(const char *s) {
    fprintf(stderr, "Ошибка синтаксического анализа: %s\n", s);
    exit(1);
}
