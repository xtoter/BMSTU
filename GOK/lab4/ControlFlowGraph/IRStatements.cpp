#include "IRStatements.hpp"
#include "../Parser/Expressions.hpp"
#include "BasicBlock.hpp"

std::string AssignStatement::Dump() {
    return var->stringValue() + " = " + rhs->stringValue();
}
std::string AssignStatement::DumpDot() {
    return "\"" + var->stringValue() + "\"" +" -> " + "\"" +rhs->stringValue() +"\"";
}
std::string BranchStatement::DumpDot() {
    if (isConditional) {
        std::string out;
        out+="\"" + condition->stringValue() + "\"" + " -> " + "\"" + firstBranchBB->stringValue() +"\"\n";
        out+="\"" + condition->stringValue() + "\"" + " -> " + "\"" + secondBranchBB->stringValue() +"\"\n";

        return out;
    } else {
        return "//branch to: " + firstBranchBB->stringValue();
    }
}


std::string BranchStatement::Dump() {
    if (isConditional) {
        return "branch on: " + condition->stringValue() + " to: " + firstBranchBB->stringValue() + " or: " + secondBranchBB->stringValue();
    } else {
        return firstBranchBB->stringValue();
    }
}
std::string PhiNodeStatement::Dump() {
    std::string argEnumeration;
    for (auto arg : bbToVarMap) {
        argEnumeration += arg.second->stringValue() + " " + arg.first->stringValue() + "; ";
    }
    return var->stringValue() + " = [" + argEnumeration + "]";
}

std::string PhiNodeStatement::DumpDot() {

    std::string out;
    std::string argprev;
    argprev = var->stringValue();
    for (auto arg : bbToVarMap) {
        
        std::string out1;
        out1 = arg.second->stringValue() + " " + arg.first->stringValue();
        out += "\"" + argprev + "\"" + " -> " + out1 + "; ";
        argprev = out1;
        //argEnumeration += arg.second->stringValue() + " " + arg.first->stringValue() + "; ";
    }
    return out;
}
