//
// Created by ter on 25.04.2021.
//

#include <algorithm>
#include <iostream>
#include "declaration.h"

Fractions& Matrix::Row::operator[](int j) {
    return this->matrix->matrix[i][j];
}

Matrix::Matrix(int m, int n) {
    this->n = n;
    this->m = m;
    matrix = new Fractions*[m];
    for (int i = 0; i < n; i++) {
        matrix[i] = new Fractions(n);
    }
}

/*Matrix::~Matrix() {
    for (int i = 0; i < m; i++) {
        delete[] matrix[i];
    }
    delete[] matrix;
}*/

Matrix::Row::Row(Matrix *matrix, int i) {
    this->matrix = matrix;
    this->i = i;
}



Fractions Matrix::multiply(Matrix matrix, int m, int k) {
    for (int i=0; i<matrix.n; i++) {
        matrix[m][i] = matrix[m][i].multiplication(k);
    }

}

int Matrix::get_num_of_columns(Matrix matrix) {
    return matrix.n;
}

int Matrix::get_num_of_lines(Matrix matrix) {
    return matrix.m;
}