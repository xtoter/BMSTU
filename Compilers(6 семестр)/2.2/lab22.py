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
class EPS:
    pass

@dataclass
class EMPTY(abc.ABC):
    pass


INT = pe.Terminal('INT', '[0-9]+', int, priority=7)
FLOAT = pe.Terminal('FLOAT', '[0-9]+(\\.[0-9]*)?(e[-+]?[0-9]+)?', float)

VARNAME = pe.Terminal('VARNAME', '[A-Za-z][A-Za-z0-9]*', str.upper)
STRING = pe.Terminal('STRING', '[A-Za-za-яА-Я][A-Za-z0-9a-яА-Я]*', str.upper)
ENDBR = pe.Terminal('ENDBR', '\)', str.upper,priority=7)

def make_keyword(image):
    return pe.Terminal(image, image, lambda name: None,
                       re_flags=re.IGNORECASE, priority=10)
KW_INTEGER, KW_FLOAT = \
    map(make_keyword, 'integer float'.split())


NProgram, NFunctionDef, NComment, NFunctionSignature, NType = \
    map(pe.NonTerminal, 'Program FunctionDef Comment FunctionSignature Type'.split())

NAddr, NAddr_, NAltexpr, NTypeList, NTypeName, NStatements, NStatement = \
    map(pe.NonTerminal, 'Addr Addr_ Altexpr TypeList TypeName Statements Statement'.split())

NSample, NVal, NPattermatch, NVals, NVals_, NExpr, NCallExpr = \
    map(pe.NonTerminal, 'Sample Val Pattermatch Vals Vals_ Expr CallExpr'.split())

NCalculate, NNum, NTuple, NValList, NOp , NDIGIT, NADDR,NCommentStr = \
    map(pe.NonTerminal, 'Calculate Num Tuple ValList Op DIGIT ADDR NCommentStr'.split())

#BLBLBLBL

NProgram |= EPS
NProgram |= NFunctionDef, NProgram

NFunctionDef |= NComment, NFunctionSignature, 'is', NStatements, 'end'

NComment |= '@', STRING, NCommentStr,'\n'
NComment |= EMPTY
NCommentStr |= EMPTY
NCommentStr |= STRING,NCommentStr

NFunctionSignature |= VARNAME, NType, '::', NType

NType |= NAddr, '(', NTypeList, ')'
NType |=  NAddr, NTypeName

NAddr |=  EPS
NAddr |= '*', NAddr

NTypeList |= NType, ', ', NType
NTypeList |= NType, ', ', NTypeList


NTypeName |= 'int'
NTypeName |= 'float'
NTypeName |= 'string'

NStatements |= NStatement
NStatements |= NStatement, ';', NStatements

NStatement |= NSample, '=', NAltexpr

NSample |= NPattermatch

NPattermatch |= NVal
NPattermatch |= NVal, ':', NPattermatch

NVal |= VARNAME
NVal |= '(', NVals, ')'
NVal |= '{', NVals, '}'
NVal |= '[', NPattermatch, ']'

NVals |= EPS
NVals |= NPattermatch
NVals |= NVals_
NVals_ |= NPattermatch, ',', NPattermatch
NVals_ |= NPattermatch, ',', NVals_

NAltexpr |= NExpr
NAltexpr |= NExpr, NOp,NAltexpr 
NExpr |= NPattermatch
NExpr |= NCallExpr
NExpr |= NNum
NExpr |= NTuple

NCallExpr |= VARNAME, NPattermatch



NNum |= NDIGIT

NTuple |= '(', NValList, ENDBR

NValList |= NPattermatch
NValList |= NPattermatch, ',', NValList

NOp |= '*'
NOp |= '/'
NOp |= '+'
NOp |= '-'

NDIGIT |= INT


p = pe.Parser(NProgram)
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
