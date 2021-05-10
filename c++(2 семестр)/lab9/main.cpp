#include <assert.h>
#include <iostream>
template <class T>
class Array
{
private:
    int m_length;
    T *m_data;

public:
    Array()
    {
        m_length = 0;
        m_data = nullptr;
    }
    Array(int a)
    {
        m_data = new T[a];
        m_length = a;
    }
    Array(int a,int b)
    {
        m_data = new T[b-a+1];
        m_length = b-a+1;
        for (int count = 0; count <= b-a; ++count)
        {
            m_data[count] = count+a;
        }
    }




    T& operator[](int index)
    {
        assert(index >= 0 && index < m_length);
        return m_data[index];
    }
    int getLength();
    bool operator ()(int b)
    {
        bool c = 0;
        for (int count = 0; count < m_length; ++count){
            if (b == m_data[count]) {
                c = 1;
            }
        }
        return c;
    }
};

template <typename T>
Array<T> operator +(Array<T> &a, Array<T> &b)
{
    int lena  = a.getLength();
    int lenb  = b.getLength();


    if (a[0]<b[0]) {
        if (a[lena-1]<b[0]){
            Array<T> cur(lena+lenb);
            for (int count = 0; count < lena; ++count){
                cur[count]=a[count];
            }
            for (int count = 0; count < lenb; ++count){
                cur[count+lena]=b[count];
            }
            return(cur);
        } else {
            Array<T> cur(b[lenb-1]-a[0]+1);
            for (int count = 0; count < lena; ++count) {
                cur[count] = a[count];
            }
            int t = 0;
            for (int count = 0; b[t] != a[lena-1]; ++count){
                t=count;
            }
            for (int count = t; count<lenb ; ++count){
                cur[count-t+lena-1]=b[count];
            }
            return(cur);
        }
    } else {
        if (b[lenb-1]<a[0]) {
            Array<T> cur(lena+lenb);
            for (int count = 0; count < lenb; ++count){
                cur[count]=b[count];
            }
            for (int count = 0; count < lena; ++count){
                cur[count+lenb]=a[count];
            }
            return(cur);
        } else{
            Array<T> cur(a[lena-1]-b[0]+1);
            for (int count = 0; count < lenb; ++count) {
                cur[count] = b[count];
            }
            int t = 0;
            for (int count = 0; a[t] != b[lenb-1]; ++count){
                t=count;
            }
            for (int count = t; count<lena ; ++count){
                cur[count-t+lenb-1]=a[count];
            }
            return(cur);
        }
    }
    return(0);
}

template <typename T>
Array<T> operator *(Array<T> &a, Array<T> &b)
{
    int lena  = a.getLength();
    int lenb  = b.getLength();
    if (a[0]<b[0]) {
        if (a[lena-1]<b[0]){
            Array<T> cur(0);
            return(cur);
        } else {
            int t2 = 0;
            for (int count = 0; a[t2] != b[0]; ++count){
                t2=count;
            }
            Array<T> cur(lena-t2);
            for (int count = 0; count <lena-t2 ; ++count){
                cur[count]=a[count+t2];
            }
            return(cur);
        }
    } else {
        if (b[lenb-1]<a[0]) {
            Array<T> cur(0);
            return(cur);
        } else{
            int t2 = 0;
            for (int count = 0; b[t2] != a[0]; ++count){
                t2=count;
            }
            Array<T> cur(lenb-t2);
            for (int count = 0; count <lenb-t2 ; ++count){
                cur[count]=b[count+t2];
            }
            return(cur);
        }
    }
    return(0);
}

template <typename T>
int Array<T>::getLength() { return m_length; }

int main()
{
    Array<int> test2(13,20);
    Array<int> test1(10,13);

    for (int count = 0; count < test1.getLength(); ++count)
        std::cout << test1[count] << '\n';
    Array<int> temp = test1*test2;
    std::cout << "done" << '\n';
    for (int count = 0; count < temp.getLength(); ++count)
        std::cout << temp[count] << '\n';
    std::cout << "done" << '\n';
    std::cout << test1.operator()(13) << '\n';
    return(0);
}