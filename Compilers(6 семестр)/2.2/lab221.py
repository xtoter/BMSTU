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
    
class Element(abc.ABC):
    pass

@dataclass
class Comment(Element):
    value : str
    
class Value(abc.ABC):
    pass

@dataclass
class Variable(Value):
    varname : str
    
@dataclass
class IntValue(Value):
    value : int

@dataclass
class PriorityVal(Value):
    vals : list[Value]
@dataclass
class Cons() :
    vals : list[Value]
@dataclass
class ValCortage(Value):
    vals : list[Cons]
    
@dataclass
class ValFunc(Value):
    funcname: str
    value: Value
    
@dataclass
class ValList(Value):
    vals : list[Cons]
    
class ExprElement(abc.ABC):
    pass
    
@dataclass
class ExprOp():
    op: str
    elem: ExprElement

@dataclass
class Expr():
    beg: ExprElement
    elems: list[ExprOp]
    
@dataclass
class PriorityExpr(Value):
    vals : list[Expr]
    
@dataclass
class ExprVal(ExprElement):
    value: Cons
    
@dataclass
class ExprBr(Value):
    value: Expr
    
@dataclass
class Pattern():
    sample : Cons
    result : Expr

class UserType(abc.ABC):
    pass

@dataclass
class DefaultType(UserType):
    typename : str
    
@dataclass
class TypeCortage(UserType):
    types : list[UserType]
    
@dataclass
class ScalaList(UserType):
    type: UserType
@dataclass
class Define(Element):
    funcname : str
    intype : UserType
    outtype: UserType
    patterns: list[Pattern]

@dataclass
class Program:
    defs : list[Element]

INT = pe.Terminal('INT', '[0-9]+', int, priority=7)
FUNCNAME = pe.Terminal('FUNCNAME', '[A-Z][A-Za-z0-9]*', str)
VARNAME = pe.Terminal('VARNAME', '[a-z][A-Za-z0-9]*', str)
STRING = pe.Terminal('STRING', '[@][^\n]*', str.upper)
INTEGER = pe.Terminal('INTEGER', 'int', str.upper)

def make_keyword(image):
    return pe.Terminal(image, image, lambda name: None,
                       re_flags=re.IGNORECASE, priority=10)
DOUBLECOLON, IS, END = \
    map(make_keyword, ':: is end'.split())


NElement, NComment, NDefine = \
    map(pe.NonTerminal, 'Element Comment Define'.split(' '))

NTypes, NType, NPatterns, NPattern, NVal, NExpr = \
    map(pe.NonTerminal, 'Types Type Patterns Pattern Val Expr'.split(' '))

NVals, NCons, NExprOp, NExprOps, NConsVal = \
    map(pe.NonTerminal, 'Vals Cons ExprOp ExprOps NConsVal'.split(' '))
    
NProgram, NElements, NOp, NExprElement = \
    map(pe.NonTerminal, 'Program Elements Op ExprElement'.split(' '))
    
NVarnames, NLCons, NLVal, NLVals, NLConsVal= \
    map(pe.NonTerminal, 'Varnames NLCons NLVal NLVals NLConsVal'.split(' '))

NProgram |= NElements, Program

NElements |= NElement, lambda x: [x]
NElements |= NElements, NElement, lambda xs, x:  xs + [x]

NElement |= NComment
NElement |= NDefine

NComment |= STRING, Comment

NDefine |= FUNCNAME, NType, DOUBLECOLON, NType, IS, NPatterns, END, Define

NType |= VARNAME, DefaultType
NType |= INTEGER, DefaultType
NType |= '(', NTypes, ')', TypeCortage
NType |= '*', NType, ScalaList

NTypes |= NType, lambda x: [x]
NTypes |= NTypes, ',', NType, lambda xs, x:xs+ [x]

NPatterns |= NPattern, lambda x: [x]
NPatterns |= NPatterns, ';', NPattern, lambda xs, x: xs + [x] 

NPattern |= NLConsVal, '=', NExpr, Pattern

NLConsVal |= NLVal, NLCons, lambda x, xs: [x] + [PriorityVal(xs)]
NLCons |= lambda: []
NLCons |= ':', NLVal, NLCons, lambda x, xs: [x] + [PriorityVal(xs)]


NLVal |= VARNAME, Variable
NLVal |= INT
NLVal |= '[', NLConsVal, ']'
NLVal |= '(', NLVals, ')', ValCortage
NLVal |= '{', NLVals, '}', ValList

NLVals |= lambda: []
NLVals |= NLConsVal, lambda x: [x]
NLVals |= NLVals, ',', NLConsVal, lambda xs, x: xs + [x] 

NVal |= FUNCNAME, NVal, ValFunc
NVal |= VARNAME, Variable
NVal |= INT
NVal |= '[', NExpr, ']', PriorityExpr
NVal |= '(', NVals, ')', ValCortage
NVal |= '{', NVals, '}', ValList

NVals |= lambda: []
NVals |= NConsVal, lambda x: [x]
NVals |= NVals, ',', NConsVal, lambda xs, x:xs+ [x] 

NConsVal |= NVal, NCons, lambda x, xs: [x] + [PriorityVal(xs)]
NCons |= lambda: []
NCons |= ':', NVal, NCons, lambda x, xs: [x] + [PriorityVal(xs)] 

NExpr |= NExprElement, NExprOps, Expr

NExprOps |= lambda: []
NExprOps |= NExprOps, NExprOp, lambda xs, x: xs + [x]

NExprOp |= NOp, NExprElement, ExprOp

NExprElement |= NConsVal

NOp |= '+', lambda: '+'
NOp |= '-', lambda: '-'
NOp |= '*', lambda: '*'
NOp |= '/', lambda: '/'

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
