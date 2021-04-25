//
// Created by ter on 25.04.2021.
//
using namespace std;
#ifndef LAB8_STRUCTSTACK_H
#define LAB8_STRUCTSTACK_H
#include <iostream>
#include <cassert>

#include <iomanip>
template <class SType> class stack {
    int size;
    SType *vals;
    int top;
public:
    stack(int);
    void push(SType i);
    SType pop();
};
template <class SType> stack<SType>::stack(int s)
{
    size=s;
    vals= new SType[size];
    top = 0;
    cout << "Initialized\n";
}
template <class SType> void stack<SType>::push(SType i)
{
    if (top == size) {
        cout << "error - stack is full. \n";
        return;
    } else {
        cout << "add " << i << "\n";
        vals[top] = i;
        top++;
    }
}
template <class SType> SType stack<SType>::pop()
{
    if(top == 0) {
        cout << "error - stack empty \n";
        return 0;
    } else {
        top--;
        return vals[top];
    }
}
#endif //LAB8_STRUCTSTACK_H
