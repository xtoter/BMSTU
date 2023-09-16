 #include<stdio.h>
 
 double power(double x, long n) {
    if(n == 0) return 1;
    if(n < 0) return power ( 1.0 / x, -n);
    return x * power(x, n - 1);
 }
 
 void main() {
     double x;
     long n;
     while (scanf ("%lf %ld", &x, &n) == 2) {
        printf("%lf\n", power (x, n));
     }
 }
