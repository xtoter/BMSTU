%{
#include "parser.tab.h"
%}

%%

[ \t]+                  /* игнорирование пробелов и табуляций */
"var"                   { return VAR; }
"BEGIN_T"                 { return BEGIN_T; }
"end"                   { return END; }
"for"                   { return FOR; }
"to"                    { return TO; }
"do"                    { return DO; }
[0-9]+                  { return NUMBER; }
":="                    { return ASSIGN; }
[a-zA-Z][a-zA-Z0-9]*    { return IDENTIFIER; }
\n                      { return NEWLINE; }
.                       { /* игнорирование других символов */ }

%%

int yywrap() {
    return 1;
}
