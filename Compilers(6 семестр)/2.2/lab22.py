import abc
import enum
import parser_edsl as pe
import sys
import re
import typing
from dataclasses import dataclass
from pprint import pprint

from dataclasses import dataclass
from typing import List, Union
from enum import Enum

@dataclass
class Program:
    functions: List['FunctionDef']

@dataclass
class FunctionDef:
    comment: Union['Comment', None]
    signature: 'FunctionSignature'
    statements: 'Statements'
@dataclass
class EMPTY(abc.ABC):
    pass
@dataclass
class Comment:
    string: str

@dataclass
class FunctionSignature:
    funcname: str
    input_type: 'Type'
    output_type: 'Type'

@dataclass
class Type:
    address: str
    typename: Union[str, None]
    type_list: Union[List['Type'], None]

@dataclass
class Statements:
    statements: List['Statement']

@dataclass
class Statement:
    sample: 'Sample'
    expression: 'Expr'

@dataclass
class Sample:
    value: Union[str, 'Vals']

@dataclass
class Vals:
    values: List[Union[str, 'Vals']]

@dataclass
class Expr:
    value: Union['Val', 'CallExpr', 'Calculate', 'Num', 'Tuple']

@dataclass
class CallExpr:
    varname: str
    value: 'Val'

@dataclass
class Calculate:
    left_expr: 'Expr'
    operator: str
    right_expr: 'Expr'

@dataclass
class Num:
    value: int

@dataclass
class Tuple:
    values: 'ValList'

@dataclass
class ValList:
    values: List['Val']

@dataclass
class Val:
    value: Union[str, 'ValList']

@dataclass
class VariableType(Enum):
    INT = 'INT'
    FLOAT = 'FLOAT'
    STRING = 'STRING'


INT = pe.Terminal('INT', '[0-9]+', int, priority=7)
FLOAT = pe.Terminal('FLOAT', '[0-9]+(\\.[0-9]*)?(e[-+]?[0-9]+)?', float)
STRING = pe.Terminal('STRING', '[A-Za-z][A-Za-z0-9 ]*', str.upper)
ENDBR = pe.Terminal('ENDBR', '\)', str.upper,priority=7)

def make_keyword(image):
    return pe.Terminal(image, image, lambda name: None,
                       re_flags=re.IGNORECASE, priority=10)
#KW_VAR, KW_BEGIN, KW_END, KW_INTEGER, KW_REAL, KW_BOOLEAN = \
#    map(make_keyword, 'var begin end integer real boolean'.split())

Program, FunctionDef, Comment, FunctionSignature, Type = \
    map(pe.NonTerminal, 'Program FunctionDef Comment FunctionSignature Type'.split())

Addr,Addr_, Altexpr,TypeList, TypeName, Statements, Statement = \
    map(pe.NonTerminal, 'Addr Addr_ Altexpr TypeList TypeName Statements Statement'.split())

Sample, Val,Pattermatch, Vals, Vals_, Expr, CallExpr = \
    map(pe.NonTerminal, 'Sample Val Pattermatch Vals Vals_ Expr CallExpr'.split())

Calculate, Num, Tuple, ValList, Op ,DIGIT, VARNAME, EPS = \
    map(pe.NonTerminal, 'Calculate Num Tuple ValList Op DIGIT VARNAME EPS'.split())

#BLBLBLBL

Program |= EPS
Program |= FunctionDef, Program

FunctionDef |= Comment, FunctionSignature, 'is', Statements, 'end'

Comment |= '@', STRING
Comment |= EMPTY

FunctionSignature |= VARNAME, Type, '::', Type

Type |= '(', TypeList, ')'
Type |=  TypeName

#Addr |= Star
#Addr |= EmptyStar
#Addr |=  EPS

TypeList |= Type, ',', TypeList
TypeList |= Type, ',', Type

TypeName |= 'INT'
TypeName |= 'FLOAT'
TypeName |= 'STRING'

Statements |= Statement
Statements |= Statement, ';', Statements

Statement |= Sample, '=', Altexpr

Sample |= Pattermatch

Pattermatch |= Val
Pattermatch |= Val, ':', Pattermatch

Val |= VARNAME
Val |= '(', Vals, ')'
Val |= '{', Vals, '}'
Val |= '[', Pattermatch, ']'

Vals |= EPS
Vals |= Pattermatch
Vals |= Vals_
Vals_ |= Pattermatch, ',', Pattermatch
Vals_ |= Pattermatch, ',', Vals_

Altexpr |= Expr
Altexpr |= Expr, Op,Altexpr 
Expr |= Pattermatch
Expr |= CallExpr
Expr |= Num
Expr |= Tuple

CallExpr |= VARNAME, Pattermatch



Num |= DIGIT

Tuple |= '(', ValList, ENDBR

ValList |= Pattermatch
ValList |= Pattermatch, ',', ValList

Op |= '*'
Op |= '/'
Op |= '+'
Op |= '-'

DIGIT |= INT
VARNAME |= STRING


p = pe.Parser(Program)
p.add_skipped_domain('\\s')        # пробельные символы
#p.print_table()
assert p.is_lalr_one()



for filename in sys.argv[1:]:
    try:
        with open(filename) as f:
            tree = p.parse(f.read())
            pprint(tree)
    except pe.Error as e:
        print(f'Ошибка {e.pos}: {e.message}')
    except Exception as e:
        print(e)
