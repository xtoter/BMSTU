#include <iostream>
#include "structStack.h"
using namespace std;
int main()
{
    stack<bool> a(5);
    cout << a.pop() << " ";
    a.push (0);
    a.push(1);
    a.push(1);
    a.push (0);
    a.push(1);
    a.push(1);

    a.push(1);
    cout << a.pop() << " ";
    cout << a.pop() <<  " ";
    cout << a.pop() << " ";
    cout << a.pop() <<  " ";
    cout << a.pop() << " ";
    cout << a.pop() <<  " ";
    cout << a.pop() << " ";
    return 0;
}


