#ifndef BasicBlock_hpp
#define BasicBlock_hpp

#include <stdio.h>
#include <vector>
#include <string>
#include <algorithm>


class AbstractStatement;
class BasicBlock;

class BasicBlock {
private:
    int _index;
    std::string _label;
public:
    std::string stringValue();
    std::vector<AbstractStatement *> statements;
    std::vector<BasicBlock *> succs;
    std::vector<BasicBlock *> preds;
    
    BasicBlock *dominator;
    std::vector<BasicBlock *> domimatingBlocks;

    void AddStatement(AbstractStatement *statement);
    static void AddLink(BasicBlock *pred, BasicBlock *succ);
    
    BasicBlock(int index, std::string label): _index(index), _label(label) {};
};

#endif
