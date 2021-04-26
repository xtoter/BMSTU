//
// Created by ter on 25.04.2021.
//
using namespace std;
#ifndef LAB8_STRUCTSTACK_H
#define LAB8_STRUCTSTACK_H
#include <iostream>
#include <cassert>

#include <iomanip>
template <class T> class stack {
    int size;
    T *vals;
    int top;
public:
    stack(int);
    void push(T i);
    T pop();
};
template <class T> stack<T>::stack(int s)
{
    size=s;
    vals= new T[size];
    top = 0;
    cout << "Initialized\n";
}
template <class T> void stack<T>::push(T i)
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
template <class T> T stack<T>::pop()
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
