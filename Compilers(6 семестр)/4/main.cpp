#include <iostream>
#include <string>
#include <unicode/uchar.h>
using namespace std;
struct Position {
private:
    string text;
    int line, pos, index;

public:
    Position(string text) {
        this->text = text;
        line = pos = 1;
        index = 0;
    }
    int Line() { return line; }
    int Pos() { return pos; }
    int Index() { return index; }
    bool operator<(const Position& a1)
    {
        return this->index < a1.index ;
    }
    bool operator==(const Position& a1)
    {
        return this->index == a1.index ;
    }

    string ToString() {
        return "(" + to_string(line) + ", " + to_string(pos) + ")";
    }
    int Cp() const {
        return (index == text.length()) ? -1 : int(text[index]);
    }
    bool IsWhiteSpace(){
        return index != text.size() && text[index]== ' ';
    }
    bool IsDecimalDigit(){
        return index != text.size() && text[index]>='0'&& text[index]<='9';
    }
    bool IsLetter(){
        return index != text.size() && ((text[index]<='Z' && text[index]>='A') ||(text[index]<='z' && text[index]>='a'));
    }
    bool IsLetterOrDigit(){
        return IsLetter()||IsDecimalDigit();
    }
    bool IsNweLine(){
        if (index == text.size())
            return true;
        if ((text[index]=='\r')&& (index+1 < text.size()))
            return text[index] == '\n';
    }
    Position& operator++()
    {
        if (index<text.size()){
            if (IsNweLine()){
                if (text[index]=='\r')
                    index++;
                line++;
                pos=1;
            } else {

            }

        }
    }

};

int main() {
    std::cout << "Hello, World!" << std::endl;
    return 0;
}
