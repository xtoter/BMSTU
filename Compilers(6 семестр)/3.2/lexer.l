%option reentrant noyywrap bison-bridge bison-locations
%option extra-type="struct Extra *"

/* Подавление предупреждений для -Wall */
%option noinput nounput

%{

#include <stdio.h>
#include <stdlib.h>
#include "parser.tab.h"  /* файл генерируется Bison’ом */
#include <stdbool.h>
#include <ctype.h>

#ifndef LEXER_H
#define LEXER_H




#ifndef YY_TYPEDEF_YY_SCANNER_T
#define YY_TYPEDEF_YY_SCANNER_T
typedef void *yyscan_t;
#endif /* YY_TYPEDEF_YY_SCANNER_T */

struct Extra {
    bool continued;
    int cur_line;
    int cur_column;
};

void init_scanner(FILE *input, yyscan_t *scanner, struct Extra *extra);
void destroy_scanner(yyscan_t);

#endif /* LEXER_H */
#define YY_USER_ACTION \
  { \
    int i; \
    struct Extra *extra = yyextra; \
    if (! extra->continued ) { \
      yylloc->first_line = extra->cur_line; \
      yylloc->first_column = extra->cur_column; \
    } \
    extra->continued = false; \
    for (i = 0; i < yyleng; ++i) { \
      if (yytext[i] == '\n') { \
        extra->cur_line += 1; \
        extra->cur_column = 1; \
      } else { \
        extra->cur_column += 1; \
      } \
    } \
    yylloc->last_line = extra->cur_line; \
    yylloc->last_column = extra->cur_column; \
  }

void yyerror(YYLTYPE *loc, yyscan_t scanner, long env[26], const char *message) {
    printf("Error (%d,%d): %s\n", loc->first_line, loc->first_column, message);
}
char* filterAlphanumeric(char* input) {
    int length = strlen(input);
    char* result = (char*) malloc((length + 1) * sizeof(char));  // Выделяем память для результирующей строки
    int resultIndex = 0;

    for (int i = 0; i < length; i++) {
        if (!(isalpha(input[i]) || (isalnum(input[i]) && resultIndex > 0))) {
            goto stop;  // Прекращаем фильтрацию, если встречается символ, не принадлежащий `[a-zA-Z][a-zA-Z0-9]*`
        }

        result[resultIndex] = input[i];
        resultIndex++;
    }
    stop:
    result[resultIndex] = '\0';  // Добавляем завершающий нулевой символ
    return result;
}
%}

%%

[ \t]+                  /* игнорирование пробелов и табуляций */
"var"                   { return VAR; }
"function"              { return FUNCTION; }
"begin"                 { return BEGIN_T; }
"end"                   { return END; }
"for"                   { return FOR; }
"to"                    { return TO; }
"do"                    { return DO; }
[0-9]+                  {  yylval->number = atoi(yytext);
                          return NUMBER; }
":="                    { return ASSIGN; }
[a-zA-Z][a-zA-Z0-9]*    { yylval->indentifier = filterAlphanumeric(yytext);
                          return IDENTIFIER; }
\.                      { return POINT; }
\n                      { }
\;                      { return SEMICOLON; }
\:                      { return COLON; }
\+|\-|\*|\/             { yylval->variable = yytext[0];
                          return OP; }
\(                      { return LBRACKET; }
\)                      { return RBRACKET; }
\,                      { return COMMA; }
.                       { /* игнорирование других символов */ }

%%

void init_scanner(FILE *input, yyscan_t *scanner, struct Extra *extra) {
    extra->continued = false;
    extra->cur_line = 1;
    extra->cur_column = 1;

    yylex_init(scanner);
    yylex_init_extra(extra, scanner);
    yyset_in(input, *scanner);
}

void destroy_scanner(yyscan_t scanner) {
    yylex_destroy(scanner);
}
