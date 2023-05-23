#include <iostream>
#include <string>
#include <map>
#include <vector>
#include <assert.h>
#include <fstream>

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

    int Line() const { return line; }
    int Pos() const { return pos; }
    int Index() const { return index; }

    bool operator<(const Position& a1) const {
        return this->index < a1.index;
    }

    bool operator==(const Position& a1) const {
        return this->index == a1.index;
    }

    string ToString() const {
        return "(" + to_string(line) + ", " + to_string(pos) + ")";
    }

    int Cp() const {
        return (index == text.length()) ? -1 : int(text[index]);
    }

    bool IsWhiteSpace() const {
        return index != text.size() && (text[index] == ' '||text[index] == '\t');
    }

    bool IsDecimalDigit() const {
        return index != text.size() && text[index] >= '0' && text[index] <= '9';
    }
    bool IsDecimalDigitOr_() const {
        return index != text.size() && ((text[index] >= '0' && text[index] <= '9')|| text[index] == '_');
    }

    bool IsLetter() const {
        return index != text.size() && ((text[index] <= 'Z' && text[index] >= 'A') || (text[index] <= 'z' && text[index] >= 'a'));
    }

    bool IsLetterOrDigit() const {
        return IsLetter() || IsDecimalDigit();
    }

    bool IsNewLine() const {
        if (index == text.size())
            return true;
        if ((text[index] == '\r') && (index + 1 < text.size()))
            return text[index + 1] == '\n';
        return text[index] == '\n';
    }

    Position& operator++() {
        if (index < text.size()) {
            if (IsNewLine()) {
                if (text[index] == '\r')
                    index++;
                line++;
                pos = 1;
            } else {
                pos++;
            }
            index++;
        }
        return *this;
    }
    Position operator++ (int)
    {
        Position copy {*this};
        ++(*this);
        return copy;
    }

};

struct Fragment {
public:
    Position Starting,  Following;
public:

    Fragment(Position Startingin, Position Followingin): Starting(""), Following("") {
        Starting = Startingin;
        Following = Followingin;
    }

    string ToString() const {
        return Starting.ToString() + "-" + Following.ToString();
    }
};

class Message {
public:
    bool IsError;
    string Text;

    Message(bool err, string t) {
        IsError = err;
        Text = t;
    }
};

enum DomainTag {
    STRING = 0,       // Identifier
    NUMBER = 1,      // Number
    END_OF_PROGRAM = 2
};

std::ostream& operator << (std::ostream& out, const DomainTag& t)
{
    switch(t) {
        case NUMBER: return (out << "NUMBER");
        case STRING: return (out << "STRING");
        case END_OF_PROGRAM: return (out << "END_OF_PROGRAM");
    }
    return (out);
}
class Token {
public:
    DomainTag Tag;
    Fragment* Coords;
    string Str = "";
    long Value = 0;
    Token() {}

    Token(DomainTag tag, Position start, Position follow) {
        Tag = tag;
        Coords = new Fragment(start, follow);
    }
};

class StringToken : public Token {
public:
    StringToken(string str, Position start, Position follow) {
        Coords = new Fragment(start, follow);
        Str = str;
        Tag = STRING;
    }
};

class NumberToken : public Token {
public:
    NumberToken(long val, Position start, Position follow) {
        Coords = new Fragment(start, follow);
        Value = val;
        Tag = NUMBER;
    }
};

class SpecToken : public Token {
public:
    SpecToken(DomainTag tag, Position start, Position follow) {
        Coords = new Fragment(start, follow);
        Tag = tag;
    }
};
class Compiler;

class Scanner {
private:
    std::string Program;
    Compiler* compiler;
    Position cur;
    std::vector<Fragment> comments;

public:
    std::vector<Fragment> Comments() {
        return comments;
    }

    Scanner(std::string program, Compiler* compiler) : Program(program), compiler(compiler), cur(program) {
        comments = std::vector<Fragment>();
    }

    Token NextToken();
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
        names = std::vector<std::string>();
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

    void AddMessage(bool isError, Position c, std::string text) {
        messages.insert({c , Message(isError, text)});
    }

    void OutputMessages() {
        for (const auto& p : messages) {
            std::cout << (p.second.IsError ? "Error" : "Warning") << " ";
            std::cout << p.first.ToString() << " : " << p.second.Text << std::endl;
        }
    }

    Scanner GetScanner(std::string program) {
        return Scanner(program, this);
    }

};
Token Scanner::NextToken() {
    while (cur.Cp() != -1) {
        while (cur.IsWhiteSpace()|| cur.Cp() == '\n'|| cur.Cp() == '\t')
            cur.operator++();
        Position start = cur;
        switch (cur.Cp()) {
            case '\"': {
                string newstr;
                while ((++cur).Cp() != '\"') {

                    newstr += cur.Cp();
                }
                return StringToken(newstr, start, cur++);
            }
            default:
                if (cur.IsDecimalDigit())
                {
                    long number = cur.Cp() - '0';
                    Position cursymb= cur;
                    try
                    {
                        while ((++cur).IsDecimalDigitOr_()) {
                            cursymb= cur;
                            int val = cur.Cp();
                            if (val != (int)'_') {
                                number =number*10 + (val - '0');
                            }

                        }
                    }
                    catch (const std::overflow_error&)
                    {
                        compiler->AddMessage(true, start, "integral constant is too large");
                        while ((++cur).IsDecimalDigit());
                    }
                    if (cur.IsLetter()) {
                        compiler->AddMessage(true, cur, "delimiter required");
                    }
                    return  NumberToken(number, start, cursymb);
                }
                else if (cur.Cp() != -1)
                {
                    compiler->AddMessage(true, cur++, "unexpected character");
                    break;
                }
        }
    }

    return SpecToken(DomainTag::END_OF_PROGRAM, cur, cur);

}

/*
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
 */
int main() {
    std::string program;
    std::string line;

    std::ifstream in("/home/xtoter/github/BMSTU/Compilers(6 семестр)/untitled3/1.txt"); // окрываем файл для чтения
    if (in.is_open())
    {
        while (getline(in, line)) {

            program += line + "\n";

        }
    }
    program.erase(program.size()-1,1);
    in.close();     // закрываем файл
    //cout<< program<<"\n";
    Compiler compiler;
    Scanner scanner(program, &compiler);
    Token token;
    do {
        token = scanner.NextToken();
        cout<< token.Coords->ToString()<<" ";
        DomainTag tag = token.Tag;
        cout << tag;
        switch (tag){
            case STRING:
                cout<<": "<<token.Str <<"\n";
                break;
            case NUMBER:
                cout<<": "<<token.Value <<"\n";
                break;
            default:
                cout<<"\n";
                break;
        }


    } while (token.Tag != END_OF_PROGRAM);
    compiler.OutputMessages();
    scanner.Comments();
    return 0;
}
