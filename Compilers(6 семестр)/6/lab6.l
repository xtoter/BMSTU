%option noyywrap bison-bridge bison-locations
%{
#include <stdio.h>
#include <stdlib.h>

#define TAG_IDENT 1
#define TAG_DIRECTIVE 2
#define TAG_ERROR     3

char *tag_names[] = {
    "END_OF_PROGRAM", "IDENT", "DIRECTIVE", "ERROR"
};

typedef struct Position Position;

struct Position {
    int line, pos, index;
};

void print_pos(struct Position *p) {
    printf("(%d, %d)", p->line, p->pos);
}

struct Fragment {
    struct Position starting, following;
};

typedef struct Fragment YYLTYPE;

void print_frag(struct Fragment *f) {
    print_pos(&(f->starting));
    printf(" - ");
    print_pos(&(f->following));
}
union Token {
    char *tokenName;
};

typedef union Token YYSTYPE;


int continued;
struct Position cur;

# define YY_USER_ACTION                 \
{                                       \
    int i;                              \
    if (!continued)                     \
        yylloc->starting = cur;         \
    continued = 0;                      \
    for (i = 0; i < yyleng; i++) {      \
        if ('\n' == yytext[i]) {        \
            cur.line++;                 \
            cur.pos = 1;                \
        }                               \
        else                            \
            cur.pos++;                  \
        cur.index++;                    \
    }                                   \
    Position temp = cur;                \
    temp.pos--;                         \
    yylloc->following = temp;           \
}                                       \

void init_scanner(char *program) {
    continued = 0;
    cur.line = 1;
    cur.pos = 1;
    cur.index = 0;
    yy_scan_string(program);
}

void err(char *msg) {
    printf("Error ");
    print_pos(&cur);
    printf(": %s\n", msg);
}

%}

IDENT [A-Z][a-zA-Z0-9\-]*
DIRECTIVE [$€¥£][A-Z]+      


%%
[\n\t ]+


{IDENT}   {
                         yylval->tokenName = yytext;
                         return TAG_IDENT;
                      }
                      
{DIRECTIVE}               {
                         yylval->tokenName = yytext;
                         return TAG_DIRECTIVE;

                      }

.           {
                err("Unexpected character");
                return TAG_ERROR;
            }

<<EOF>>      {
               return 0;
            }          


%%


int main(){
    int tag;
    YYSTYPE value;
    YYLTYPE coords;

   	FILE *inputfile;
	long size_str;
	char *str;
	union Token token;
	inputfile = fopen("test.txt","r");
	if (inputfile == NULL) {
		fputs("File not found", stderr);
		exit(1);
	}
	fseek(inputfile, 0,SEEK_END);
	size_str = ftell(inputfile);
    rewind(inputfile);

    str=(char*)malloc(sizeof(char)*(size_str+1));
    if (str == NULL) {
		fputs("Memory error",stderr);
		exit(2);
	}    
    size_t n = fread(str,sizeof(char),size_str,inputfile);
    if (n != size_str) {
		fputs ("Reading error",stderr);
		exit (3);
	}
    str[size_str] = '\0';
    fclose (inputfile);
    init_scanner(str);
    do {
        tag = yylex(&value, &coords);
        if (0 != tag && TAG_ERROR != tag) {
		  printf("%s ", tag_names[tag]);
            print_frag(&coords);
            printf(": %s\n", value.tokenName);
        }
    } while (0 != tag);
	free(str);
    return 0;
}
