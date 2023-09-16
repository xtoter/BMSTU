#include <iostream>
#include <string>
#include <fstream>
#include <vector>
#include <map>
#include <algorithm>
#include <cctype>
#include <cstdio>
#include <cstdlib>
#include <memory>
#include "llvm/ADT/APInt.h"
#include "llvm/ADT/STLExtras.h"
#include "llvm/IR/BasicBlock.h"
#include "llvm/IR/Constants.h"
#include "llvm/IR/DerivedTypes.h"
#include "llvm/IR/Function.h"
#include "llvm/IR/IRBuilder.h"
#include "llvm/IR/LLVMContext.h"
#include "llvm/IR/Module.h"
#include "llvm/IR/Type.h"
#include "llvm/IR/Verifier.h"


static llvm::LLVMContext context;
static llvm::IRBuilder<> builder(context);
static std::unique_ptr<llvm::Module> TheModule = std::make_unique<llvm::Module>("lab3", context);
llvm::FunctionType *FT = llvm::FunctionType::get(llvm::Type::getInt64Ty(context), false);
llvm::Function *func = llvm::Function::Create(FT, llvm::Function::ExternalLinkage, "main", TheModule.get());
bool debug;
using namespace std;

enum Tag {
    MUL,  
    ADD,
    SUB,
    IF,
    INT, 
    WHILE,
    Ret,
    END_OF_PROGRAM,
    NTERM,
    TERM,
    MAIN,
    LEFT,
    RIGHT,
    LPAREN,
    RPAREN,
    Equal,
    More,
    End,
};

struct tree {
    string type;
    vector<tree*> children;
};

struct token
{
   Tag type;
   string data;
};


map<string, llvm::Value*>  keywords = {};
string buffer;
int cur = 0;

bool IsWhiteSpace(char s) {
    return s == ' ' || s == '\n';
}

void add(int c) {
    if (c + 1 > buffer.size()) {
        cur = -1;
        return;
    }
    cur++;
}

bool IsNumber(char s) {
    if (s >= '0' && s <= '9') 
        return true;
    return false;
}

bool IsLetter(char s) {
    if ((s >= 'a' && s <= 'z') || (s >= 'A' && s <= 'Z')) {
        return true;
    }
    return false;
}
token TokenType(){
        int start = cur;
        string name;
        switch (buffer[cur])
        {
        case '*':
            add(cur);
            return token{Tag::MUL, string("*")};
        case '+':
            add(cur);
            return token{Tag::ADD, string("+")};
        case '-':
            add(cur);
            return token{Tag::SUB, string("-")};
        case '(':
            add(cur);
            return token{Tag::LEFT, string("(")};
        case ')':
            add(cur);
            return token{Tag::RIGHT, string(")")};
        case '{':
            add(cur);
            return token{Tag::LPAREN, string("{")};
        case '}':
            add(cur);
            return token{Tag::RPAREN, string("}")};
        case ';':
            add(cur);
            return token{Tag::End, string(";")};
        case '=':
            add(cur);
            return token{Tag::Equal, string("=")};
        default:
            int back;
            if (IsLetter(buffer[cur])) {
                do
                {   back = cur;
                    add(cur);
                    if (debug) {cout << buffer[cur] << " " << cur << IsLetter(buffer[cur]) << endl;}
                } while (cur != -1 && IsLetter(buffer[cur]));
                int k = cur;
                if (k == -1) {
                    k = buffer.size();
                }
                string name = buffer.substr(start,k - start);
                if (debug) {cout << name << " NAME" << endl;}
                if (name == "if") {
                    return  token{Tag::IF, name};
                }
                if (name == "while") {
                    return  token{Tag::WHILE, name};
                }
                if (name == "return") {
                    if (debug) {cout << "RETURN\n" << endl;}
                    return  token{Tag::Ret, name};
                }
                if (name == "int") {
                    return  token{Tag::INT, name};
                }
                if (name == "main") {
                    return  token{Tag::MAIN, name};
                }
                return token{Tag::NTERM, name};
            } else if (cur != -1)  {
                if (IsNumber(buffer[cur])) {
                    do
                    {   back = cur;
                        add(cur);
                    } while (cur != -1 && IsNumber(buffer[cur]));
                    int k = cur;
                    if (k == -1) {
                        k = buffer.size();
                    }
                    string name = buffer.substr(start,k - start);
                    return token{Tag::TERM, name};
                }
            }
        }
}
token NextToken() {
    if (debug) {cout << "NextToken " << buffer[cur] << endl;}
    int k = 0;
    while (cur != -1)
    {   
        if (debug) {cout << buffer[cur] << " " << cur << endl;}

        while (IsWhiteSpace(buffer[cur]))
        {
            add(cur);
            if (debug) {cout << buffer[cur] << " " << cur << endl;}
        }
        if (debug) {cout << buffer[cur] << " " << cur << endl;}
        return TokenType();
    }
    return token{Tag::END_OF_PROGRAM, ""};
}

llvm::Value* Expr() {
    token t = NextToken();
    llvm::Value *value_1_p;
    llvm::Value *value_2_p;
    switch (t.type)
    {
        case Tag::TERM:
            value_1_p = llvm::ConstantInt::get(
                                        context,
                                        llvm::APInt(64, stoi(t.data)));
            break;
        case Tag::NTERM:
            value_1_p = keywords[t.data];
            break;
        default:
            break;
    }
    token z = NextToken(); // + - *
    if (z.type == Tag::End) {
        return value_1_p;
    }
    t = NextToken(); // Operand

    switch (t.type)
    {
        case Tag::TERM:
        
             value_2_p = llvm::ConstantInt::get(
                                        context,
                                        llvm::APInt(64, stoi(t.data)));
            break;
        case Tag::NTERM:
            value_2_p = keywords[t.data];
            break;
        default:
            break;
    }
    llvm::Value* val  = builder.CreateFAdd(
                value_1_p,
                value_2_p,
                "res"); ;
    switch (z.type)
    {
        case Tag::ADD:
            val = builder.CreateFAdd(
                value_1_p,
                value_2_p,
                "res"); 
            break;
        case Tag::SUB:
            val = builder.CreateFSub(
            value_1_p,
            value_2_p,
            "res");
            break;
        case Tag::MUL:
            val = builder.CreateFMul(
            value_1_p,
            value_2_p,
            "res");
            break;
        default:
            break;
    }
    NextToken(); // ;
    return val;
}



token Expr(tree* root) {
    token tt = NextToken();
    if (debug) {cout << "Expr " << tt.data << " " << tt.type << endl;}
    token t;
    token out;
    token nt;
    llvm::Value *value_ptr_p;
    llvm::Value *value_to_be_stored;
    llvm::Value *Cmp;
    llvm::AllocaInst* alloca;
    llvm::BasicBlock *block_in_if;
    llvm::BasicBlock *block_after_if;
    llvm::BasicBlock *block_before_while;
    if (debug) {cout << "Expr start" << endl;}
    switch (tt.type)
    {
    case Tag::NTERM:
      
        if (NextToken().type == Tag::Equal) {
            llvm::Value* v = Expr();
            if (debug) {cout << "NTERM = " << tt.data << endl;}
            builder.CreateStore(v, keywords[tt.data]);
            
        }
        break;
    case Tag::IF:
        nt = NextToken(); // NTerm
        if (debug) {cout << "IF " << nt.data << endl;}
        Cmp = builder.CreateICmpNE(keywords[nt.data], llvm::ConstantInt::get(context, llvm::APInt(64, 0)));
        block_in_if = llvm::BasicBlock::Create(context, "in_if", func);
        block_after_if = llvm::BasicBlock::Create(context, "after_if", func);
        builder.CreateCondBr(Cmp, block_in_if, block_after_if);
        builder.SetInsertPoint(block_in_if);
        NextToken(); // Expr
        Expr(nullptr); 
        builder.CreateBr(block_after_if);
        NextToken(); // }
        builder.SetInsertPoint(block_after_if);
        
        break;
    case Tag::INT:
        nt = NextToken(); /* NTERM */
        alloca = builder.CreateAlloca(
                    llvm::Type::getInt64Ty(context), // type
                    nullptr,
                    nt.data ); // name
        value_ptr_p = builder.CreateLoad(alloca);
        keywords[nt.data] = value_ptr_p;
        value_to_be_stored = llvm::ConstantInt::get(context, llvm::APInt(64, 0));
        builder.CreateStore(value_to_be_stored, value_ptr_p);
        NextToken(); // ;
        break;
    case Tag::WHILE:
        nt = NextToken(); // NTerm
        if (debug) {cout << "WHILE " << nt.data << endl;}
        Cmp = builder.CreateICmpNE(keywords[nt.data], llvm::ConstantInt::get(context, llvm::APInt(64, 0)));
        block_before_while = llvm::BasicBlock::Create(context, "before_while", func);
        builder.CreateBr(block_before_while);
        block_in_if = llvm::BasicBlock::Create(context, "in_while", func);
        block_after_if = llvm::BasicBlock::Create(context, "after_while", func);
        builder.CreateCondBr(Cmp, block_in_if, block_after_if);
        builder.SetInsertPoint(block_in_if);
        NextToken(); // Expr
        Expr(nullptr); 
        builder.CreateBr(block_before_while);
        NextToken(); // }
        builder.SetInsertPoint(block_after_if);
        break;
    case Tag::Ret:
        if (debug) {cout << "-------\n";}
        out = tt;
        t = NextToken();
        if (debug) {cout << "-------\n";}
        llvm::Value *RetVal;
        if (t.type == Tag::TERM) {
            RetVal = builder.getInt64(stoi(t.data));
        } else if (t.type == Tag::NTERM) {

            RetVal = keywords[t.data];
        }
        builder.CreateRet(RetVal);
        if (debug) {cout << t.data << " ------" << endl;}
        t = NextToken(); // ;
        if (debug) {cout << t.data << " ------" << endl;}
        break;
    default:
        break;
    }
    return out;
}

string Exprs(tree* root) {
    tree* t = new tree();
    t->type = "Exprs";
    if (debug) {cout << "Exprs start" << endl;}
    token tt = Expr(root);
    while (true) {
        if (tt.type == Tag::Ret) { 
            if (debug) {cout << "Exprs End" << endl;}
            return tt.data;
        }
        tt = Expr(root);
        

    }
}

auto* Program() {
    if (NextToken().type == Tag::INT) {
        if (NextToken().type == Tag::MAIN) {
            if (NextToken().type == Tag::LEFT) {
                NextToken();
                if (NextToken().type == Tag::LPAREN) {
                    tree* t = new tree();
                    
                    llvm::BasicBlock *BB = llvm::BasicBlock::Create(context, "main", func);
                    builder.SetInsertPoint(BB);
                    if (debug) {cout << "Program start" << endl;}
                    Exprs(t);
                   
                    verifyFunction(*func);
                    return func;
                }
            }
        }    
    } else {
        cout << "Error INT " << buffer.substr(0,cur) << endl;
    }
}




int main(int argc, char *argv[]) {
    if (argc == 2 && argv[1][0]=='d'){
        debug = true;
    } else {
        debug = false;
    }
    ifstream file("./../test.txt");
    if (file.is_open()) {
        buffer = string((istreambuf_iterator<char>(file)), (istreambuf_iterator<char>()));
        file.close();
        if (debug) {cout << buffer << " " << buffer[0] << endl;}
        
    } else {
        cout << "Put test.txt in .. \n";
        return 1;
    }
    llvm::Function *Fn = Program();
    Fn->print(llvm::errs());

    return 0;
}
