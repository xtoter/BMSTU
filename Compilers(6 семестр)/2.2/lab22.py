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
class ValCortage(Value):
    vals : list[Value]
    
@dataclass
class ValList(Value):
    vals : list[Value]
    
@dataclass
class Cons(Value) :
    vals : list[Value]
    
class Expr(abc.ABC):
    pass

@dataclass
class ExprFunction(Expr):
    funcname : str
    value: Value
    
@dataclass
class Operation(Expr):
    optype: str
    first : Expr
    second : Expr
@dataclass
class Pattern():
    sample : Value
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
VARNAME = pe.Terminal('VARNAME', '[A-Za-z][A-Za-z0-9]*', str.upper)
STRING = pe.Terminal('STRING', '[A-Za-za-яА-Я][A-Za-z0-9a-яА-Я]*', str.upper)
INTEGER = pe.Terminal('STRING', 'int', str.upper)

def make_keyword(image):
    return pe.Terminal(image, image, lambda name: None,
                       re_flags=re.IGNORECASE, priority=10)
DOUBLECOLON, IS, END = \
    map(make_keyword, ':: is end'.split())


NProgram, NElement, NComment, NDefine = \
    map(pe.NonTerminal, 'Program Element Comment Define'.split())

NTypes, NType, NPatterns, NPattern, NVal, NExpr = \
    map(pe.NonTerminal, 'Types Type Patterns Pattern Val Expr'.split())

NVals, NCons, NFunction = \
    map(pe.NonTerminal, 'Vals Cons Function'.split())
    
NElements = \
    map(pe.NonTerminal, 'Elements'.split())

#BLBLBLBL


NProgram |= NElements, Program

NElements |= NElement, lambda x: [x]
NElements |= NElement, NProgram, lambda x, xs: [x] + xs

NElement |= NComment
NElement |= NDefine

NComment |= '@', STRING, Comment

NDefine |= VARNAME, NType, DOUBLECOLON, NType, IS, NPatterns, END, Define

NType |= VARNAME, DefaultType
NType |= INTEGER, DefaultType
NType |= '(', NTypes, ')', TypeCortage
NType |= '*', NType, ScalaList

NTypes |= NType, lambda x: [x]
NTypes |= NTypes, ',', NType, lambda xs, x: [x] + xs

NPatterns |= NPattern, lambda x: [x]
NPatterns |= NPatterns, ';', NPattern, lambda xs, x: [x] + xs

NPattern |= NVal, '=', NExpr, Pattern

NVal |= VARNAME
NVal |= INT
NVal |= '(', NVals, ')', ValCortage
NVal |= '{', NVals, '}', ValList
NVal |= NCons, Cons

NVals |= NVal, lambda x: [x]
NVals |= NVals, ',', NVal, lambda xs, x: [x] + xs

NCons |= NVal, ':', NVal, lambda xf, xs: [xf] + [xs]
NCons |= NCons, ':', NVal, lambda xs, x: [x] + xs

NExpr |= NFunction
NExpr |= NExpr, '+', NExpr, Operation
NExpr |= NExpr, '-', NExpr, Operation
NExpr |= NExpr, '*', NExpr, Operation
NExpr |= NExpr, '/', NExpr, Operation

NFunction |= VARNAME, NVal, ExprFunction

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
