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
    
func = {} # key = f_name : (in_type, out_type)
checkVariablesMap = {}
boolVariable1 = False
class SemanticError(pe.Error):
	pass
class Element(abc.ABC):
    pass
class VarRedefinition(SemanticError):
    def __init__(self, pos, type):
        self.pos = pos
        self.type = type
 
    @property
    def message(self):
        return f'Переопределение переменной: {self.type}'
class VarNotExist(SemanticError):
    def __init__(self, pos, type):
        self.pos = pos
        self.type = type
 
    @property
    def message(self):
        return f'Неопределенная переменная: {self.type}'
 
@dataclass
class Comment(Element):
    value : str
    def check(self):
        
        print("comment")
        pass
    
class Value(abc.ABC):
    def getType(self):
            pass
    def check(self):
        pass

@dataclass
class Variable(Value):
    varname : str
    typeVarCoords: pe.Fragment
    @pe.ExAction
    def create(attrs, coords, res_coord):
        varname= attrs
        name = ""
        for cur in varname:
            name += cur
        typeVarCoords= coords
        return Variable(name,typeVarCoords)
    def check(self,example):
        
        print("variable",example)
        print(boolVariable1)
        if boolVariable1:
            if self.varname in checkVariablesMap:
                print("printerr")
                raise VarRedefinition(self.typeVarCoords,self.varname)
            else:
                print("printcreat")
                checkVariablesMap[self.varname]=True
        else:
            print("printreadnot")
            if self.varname not in checkVariablesMap:
               raise VarNotExist(self.typeVarCoords,self.varname)
            
        
        #print ("check intval")
    def getType(self):
        return 'INT'
    
@dataclass
class IntValue(Value):
    value : int
    def check(self):
        print ("check intval")
        pass
    def getType(self):
        return 'INT'

@dataclass
class PriorityVal(Value):
    vals : list[Value]
    def check(self,example):
        
        print("priorityval")
        for cur in self.vals:
            print(type(cur),"PriorityVal")
            cur.check(example)
        #print (self.vals,"check priorityval")
    def getType(self):
        if len(self.vals)==0:
            return ''
        curType = "prior"#self.vals[0].getType()
        #Обработка когда типы разные
        for cur in self.vals:
            curType+= cur.getType()+"* "
        #print(self)
        return curType
    
        
    
@dataclass
class Cons() :
    vals : list[Value]
    def check(self,example):
        print("Cons",example)
        for cur in self.vals:
            ## example is type cons
            self.vals[1].check(example)
            self.vals[0].check(example - 1) ## example - 1 ne ebu poka no potom kto-to sdelae
            if type(cur) != list:
                print(type(cur),"Cons")
                cur.check(example)
    def getType(self):
        str="Переменные("
        for cur in self.vals:
            str+=cur.getType()+" "
        str+=")"
        return str
@dataclass
class ValCortage(Value):
    vals : list[Cons]
    def check(self,example):
        print("valcortage",type(self.vals),example)
        index = 0
        for cur in self.vals:
            print(type(cur),"ValCortage")
            if type(cur)  is list: # nahuya => delete
                for temp in cur:
                    print(type(temp),"valcortage")
                    temp.check(example)
            else:
                print(type(cur),"valcortage")
                cur.check(example[index]) # dolgen otpravlatsa i type value
            index += 1
    def getType(self):
        str=""
        for cur in self.vals:
            if type(cur)  is list:
                for temp in cur:
                    str+=temp.getType()
            else:
                str+=cur.getType()
        return str
    
    
@dataclass
class ValFunc(Value):
    funcname: str
    value: Value
    def check(self,example):
        print("Valfunc",example)
        self.value.check(example)
    def getType(self):
        return "ValFunc"
    
@dataclass
class ValList(Value):
    vals : list[Cons]
    def check(self,example):
        print("ValList",example)
        for cur in self.vals:
            if type(cur)  is list:
                for temp in cur:
                    if type(temp) != int:
                        temp.check(example)
            else:
                cur.check(example)
    def getType(self):
        return "ValList"
    
class ExprElement(abc.ABC):
    def check(self,example):
        print("ExprElement",example)
    
@dataclass
class ExprOp():
    op: str
    elem: ExprElement
    def check(self,example):
        print("ExprElement",example)
        for cur in self.elem:
            cur.check(example)

@dataclass
class Expr():
    beg: ExprElement
    elems: list[ExprOp]
    def check(self,example):
        print("Expr",example)
        for cur in self.beg:
            if type(cur) != int:
                cur.check(example)
        for cur in self.elems:
            if type(cur) != int:
                cur.check(example)
    
    
@dataclass
class PriorityExpr(Value):
    vals : list[Expr]
    def check(self,example):
        print("PriorityExpr",example)
        if type(self.vals) == list: 
            for cur in self.vals:
                cur.check(example)
        else:
            self.vals.check(example)
    def getType(self):
        return "PriorityExpr"
            
    
@dataclass
class ExprVal(ExprElement):
    value: Cons
    def check(self,example):
        print("ExprVal",example)
        self.value.check(example)
    
@dataclass
class ExprBr(Value):
    value: Expr
    def getType(self):
        return "ExprBr"
    
    
@dataclass
class Pattern():
    sample : Cons
    result : Expr
    def check(self,funcname,example):
        print("Pattern",funcname,example)
        
        checkVariablesMap.clear()
        global boolVariable1
        boolVariable1 = True
        self.sample.check(func[funcname][0]) #(in_type, out_type)[0]
        print(self.sample)
        boolVariable1 = False
        self.result.check(func[funcname][1])
        #print("\n проверка типов \n")
        #print(self.sample.getType())

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
class FuncRedefinition(SemanticError):
    def __init__(self, pos, type):
        self.pos = pos
        self.type = type
 
    @property
    def message(self):
        return f'Переопределение функции: {self.type}'
 
@dataclass
class Define(Element):
    funcname : str
    intype : UserType
    outtype: UserType
    patterns: list[Pattern]
    variants: dict[str, tuple[UserType,UserType]]
    typeDefineCoords: pe.Fragment
    @pe.ExAction
    def create(attrs, coords, res_coord):
        funcname, intype, outtype,patterns = attrs
        
        temp = func.copy()
        if funcname not in func:
            func[funcname]=(intype,outtype)
        typeDefineCoords, _,_,_,_,_,_ = coords
        return Define(funcname,intype,outtype,patterns,temp,typeDefineCoords)
    def check(self):
        
        print("define")
        if self.funcname in self.variants:
            raise FuncRedefinition(self.typeDefineCoords, self.funcname)
        for cur in self.patterns:
            print(type(cur),"check")
            cur.check(self.funcname)
    
            

@dataclass
class Program:
    defs : list[Element]
    def check(self):
        for name in self.defs:
            
            print(type(name),"program")
            name.check()

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

NDefine |= FUNCNAME, NType, DOUBLECOLON, NType, IS, NPatterns, END, Define.create

NType |= VARNAME, DefaultType
NType |= INTEGER, DefaultType
NType |= '(', NTypes, ')', TypeCortage
NType |= '*', NType, ScalaList

NTypes |= NType, lambda x: [x]
NTypes |= NTypes, ',', NType, lambda xs, x:xs+ [x]

NPatterns |= NPattern, lambda x: [x]
NPatterns |= NPatterns, ';', NPattern, lambda xs, x: xs + [x] 

NPattern |= NLConsVal, '=', NExpr, Pattern

NLConsVal |= NLVal, NLCons, lambda x, xs: Cons([x] + [(xs)])
NLCons |= lambda: []
NLCons |= ':', NLVal, NLCons, lambda x, xs: [x] + [(xs)]


NLVal |= VARNAME, Variable.create
NLVal |= INT
NLVal |= '[', NLConsVal, ']', PriorityVal
NLVal |= '(', NLVals, ')', ValCortage
NLVal |= '{', NLVals, '}', ValList

NLVals |= lambda: []
NLVals |= NLConsVal, lambda x: [x]
NLVals |= NLVals, ',', NLConsVal, lambda xs, x: xs + [x] 

NVal |= FUNCNAME, NVal, ValFunc
NVal |= VARNAME, Variable.create
NVal |= INT
NVal |= '[', NExpr, ']', PriorityExpr
NVal |= '(', NVals, ')', ValCortage
NVal |= '{', NVals, '}', ValList

NVals |= lambda: []
NVals |= NConsVal, lambda x: [x]
NVals |= NVals, ',', NConsVal, lambda xs, x:xs+ [x] 

NConsVal |= NVal, NCons, lambda x, xs: [x] + [(xs)]
NCons |= lambda: []
NCons |= ':', NVal, NCons, lambda x, xs: [x] + [(xs)] 

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
            #pprint(tree)
            tree.check()
    except pe.Error as e:
        print(f'Ошибка {e.pos}: {e.message}')
    except Exception as e:
        print(e)
