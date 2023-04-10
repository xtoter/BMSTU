#include <iostream>
#include <string>
#include <map>
#include <vector>
#include <assert.h>
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
    bool IsNewLine(){
        if (index == text.size())
            return true;
        if ((text[index]=='\r')&& (index+1 < text.size()))
            return text[index] == '\n';
        return text[index] == '\n';
    }
    Position& operator++()
    {
        if (index<text.size()){
            if (IsNewLine()){
                if (text[index]=='\r')
                    index++;
                line++;
                pos=1;
            } else {
                pos++;
            }
            index++;

        }
        return *this;
    }

};
struct Fragment {
public:
  Position *Starting, *Following;
public:

    Fragment(Position Startingin,Position Followingin) {
       Starting=&Startingin;
       Following = &Followingin;
    }
    string ToString() {
        return Starting->ToString()+"-"+Following->ToString();
    }
};
class Message {
public:
  bool IsError;
  string Text;
  Message(bool err,string t){
      IsError = err;
      Text = t;
  }
};
enum DomainTag {
    IDENT = 0,       // Identifier
    NUMBER = 1,      // Number
    CHAR = 2,        // Character
    LPAREN = 3,      // Special '('
    RPAREN = 4,      // Special ')'
    PLUS = 5,        // Special '+'
    MINUS = 6,       // Special '-'
    MULTIPLY = 7,    // Special '*'
    DIVIDE = 8,      // Special '/'
    END_OF_PROGRAM = 9
};

class Token {
public:
    DomainTag Tag;
    Fragment *Coords;
    Token(){

    }
    Token(DomainTag tag,Position start,Position follow){
        Tag=tag;
        Coords =  new Fragment(start,follow);
    }
};
class IdentToken : public Token
{
public:
   int Code;
   IdentToken(int code, Position start,Position follow) {

       Coords =  new Fragment(start,follow);
       Code = code;
   };
};
class NumberToken : public Token
{
public:
   long Value;
   NumberToken(long val, Position start,Position follow) {

       Coords =  new Fragment(start,follow);
       Value = val;
   };
};
class CharToken : public Token
{
public:
   int CodePoint;
   CharToken(int val, Position start,Position follow) {

       Coords =  new Fragment(start,follow);
       CodePoint = val;
   };
};
class SpecToken : public Token
{
public:
   SpecToken(DomainTag tag, Position start,Position follow) {

       assert(tag ==  3||
      tag ==4 ||
      tag == 5||
      tag == 6||
      tag == 7||
      tag == 8||
      tag == 9 );
   };
};

class Compiler {
private:
    std::map<Position, Message> messages;
    std::map<std::string, int> nameCodes;
    std::vector<std::string> names;
public:
    Compiler() {
        messages = std::map<Position, Message>();
        nameCodes = std::map<std::string, int>();
        names = std::list<std::string>();
    }
    int AddName(std::string name) {
        if (nameCodes.count(name) > 0) {
            return nameCodes[name];
        }
        else {
            int code = names.size();
            names.push_back(name);
            nameCodes[name] = code;
            return code;
        }
    }
    std::string GetName(int code) {
        return names[code];
    }
    int AddName(std::string name) {
        if (nameCodes.count(name)) {
            return nameCodes[name];
        } else {
            int code = names.size();
            names.push_back(name);
            nameCodes[name] = code;
            return code;
        }
    }
    std::string GetName(int code) {
        return names[code];
    }
    void AddMessage(bool isError, Position c, std::string text) {
        messages[c] = Message(isError, text);
    }
    void OutputMessages() {
        for (const auto& p : messages) {
            std::cout << (p.second.isError ? "Error" : "Warning") << " " << p.first << " : " << p.second.text << std::endl;
        }
    }
    Scanner GetScanner(std::string program) {
        return Scanner(program, *this);
    }
    // ...
};
class Scanner {
private:
    std::string Program;
    Compiler *compiler;
    Position *cur;
    std::vector<Fragment> comments;
public:
    std::vector<Fragment> Comments()  {
        return comments;
    }
    Scanner(std::string program, Compiler *compiler) : Program(program), compiler(compiler) {
        cur = new Position(Program = program);
        comments = new  std::vector<Fragment>();
    }
    // ...
};
Token NextToken() {
    while (cur.Cp != -1) {
        while (cur.IsWhiteSpace())
            cur++;
        Position start = cur;
        switch (cur.Cp) {
            case '(':
                return SpecToken(DomainTag::LPAREN, start, ++cur);
            case ')':
                return SpecToken(DomainTag::RPAREN, start, ++cur);
            // ...
            case '*':
                return SpecToken(DomainTag::MULTIPLY, start, ++cur);
            // ...
            case '/':{
                if (++cur->Cp != '*')
                    return new SpecToken(DomainTag::DIVIDE, start, ++cur);
                do {
                    do cur++;
                    while (cur->Cp != '*' && cur->Cp != -1);
                    cur++;
                } while (cur->Cp != '/' && cur->Cp != -1);
                if (cur->Cp == -1)
                    compiler.AddMessage(true, *cur, "end of program found, '*/' expected");
                comments.Add(new Fragment(start, ++cur));
                break;
            }
            case '\\':
                if ((++cur).IsNewLine()) {
                    compiler.AddMessage(true, cur, "new line in constant");
                    return new CharToken(0, start, cur++);
                } else if (cur.Cp == '\'') {
                    compiler.AddMessage(true, cur, "empty character literal");
                    return new CharToken(0, start, cur++);
                } else {
                    int ch = cur.Cp;
                    if (ch == '\\') {
                        switch ((++cur).Cp) {
                            case 'n': ch = '\n'; break;
                            case '\'': ch = '\''; break;
                            case '\\': ch = '\\'; break;
                            default:
                                compiler.AddMessage(true, cur, "unrecognized Escape sequence");
                                break;
                        }
                    }
                    if ((++cur).Cp != '\'') {
                        compiler.AddMessage(true, cur, "too many characters in character literal");
                        while (cur.Cp != '\'' && !cur.IsNewLine()) {
                            cur++;
                        }
                        if (cur.Cp != '\'') {
                            compiler.AddMessage(true, cur, "newline in constant");
                        }
                    }
                    return new CharToken(ch, start, ++cur);
                }
            default:
                if (isalpha(cur.Cp))
                {
                    do
                    {
                        cur++;
                    } while (isalnum(cur.Cp));

                    string name = Program.substr(start.Index, cur.Index - start.Index);
                    return new IdentToken(compiler.AddName(name), start, cur);
                } else if (cur.IsDecimalDigit())
                {
                    long val = cur.Cp - '0';
                    try
                    {
                        while ((++cur).IsDecimalDigit())
                            val = __int64(val * 10 + cur.Cp - '0');
                    }
                    catch (const std::overflow_error&)
                    {
                        compiler.AddMessage(true, start, "integral constant is too large");
                        while ((++cur).IsDecimalDigit());
                    }
                    if (cur.IsLetter())
                        compiler.AddMessage(true, cur, "delimiter required");
                    return new NumberToken(val, start, cur);
                }
                else if (cur.Cp != -1)
                {
                    compiler.AddMessage(true, cur++, "unexpected character");
                    break;
                }
        }
    }
    return SpecToken(DomainTag::END_OF_PROGRAM, cur, cur);
}

int main() {
    std::cout << "Hello, World!" << std::endl;
    return 0;
}