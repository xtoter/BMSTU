#include <iostream>
#include "declaration.h"

std::ostream &operator<<(std::ostream &os, Matrix &matr) {
    for (int i = 0; i < matr.m; i++) {
        for (int j = 0; j < matr.n; j++) {
            os << matr[i][j].numerator << "/" << matr[i][j].denominator << " ";
        }
        os << std::endl;
    }
    return os;
}

int main() {
    int n, m, a, b;
    std::cin >> n >> m;
    Matrix matrix(n, m);
    for (int i = 0; i < m; i++) {
        for (int j = 0; j < n; j++) {
            matrix[i][j] = *new Fractions(rand(), rand());
        }
    }
    std::cout << matrix;
    Fractions i = matrix[0][0];
    std::cout << i.numerator << "/" << i.denominator << std::endl;
    std::cout << Matrix::get_num_of_lines(matrix) << " ";
    std::cout << Matrix::get_num_of_columns(matrix) << std::endl;
    return 0;
}
