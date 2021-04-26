//
// Created by ter on 25.04.2021.
//

#ifndef LAB7_DECLARATION_H
#define LAB7_DECLARATION_H
#include <iostream>

class Fractions {
public:
    int numerator, denominator;

    Fractions(int n, int d) {
        if(d == 0) {
            std::cout << "net" << std::endl; //check if denom = 0;
        }
        numerator = n;
        denominator = d;

        int thisGcd = gcd(numerator, denominator);

        numerator = numerator/thisGcd;
        denominator = denominator/thisGcd;
    }

public:
    Fractions(int n) {
        numerator = n;
        denominator = 1;
    }

/*public:
    String toString() {
        int thisGcd = gcd(numerator, denominator);

        return (numerator/thisGcd + "/" + denominator/thisGcd);
    }*/
private:
    static int gcd (int n, int d) {
        if (d==0)
            return n;
        return gcd(d,n%d);
    }

public:
    double evaluate() {  // like 'eval'
        double n = numerator;
        double d = denominator;
        return (n / d);
    }

public:
    Fractions add(Fractions fract2) {
        Fractions *res;
        res = new Fractions((numerator * fract2.denominator) + (fract2.numerator * denominator),
                            (denominator * fract2.denominator));
        return *res;
    }

public:
    Fractions multiplication (Fractions fract2) {
        Fractions *res;
        res = new Fractions((numerator * fract2.numerator),
                            (fract2.denominator * denominator));
        return *res;
    }
};

class Matrix {
public:
    Fractions **matrix;
    int n, m;

public:
    Matrix(int n, int m);

    class Row {
    private:
        Matrix *matrix;
        int i;
    public:
        Row(Matrix *matrix, int i);
        Fractions& operator[] (int j);
    };

    Matrix::Row operator[] (int i){
        return Row(this, i);
    }
    Fractions static multiply(Matrix matrix, int m, int k);
    int static get_num_of_lines(Matrix matrix);
    int static get_num_of_columns(Matrix matrix);
};
#endif //LAB7_DECLARATION_H
