public class Test {
    public static void main(String[] args) {
        Fraction tempa[] = {new Fraction(1,1),new Fraction(0,1),new Fraction(1,1)};
        Fraction tempb[] = {new Fraction(0, 1),new Fraction(1,1),new Fraction(0,1)};
        myVec a = new myVec(tempa);
        myVec b = new myVec(tempb);
        for (int i = 0; i < a.n; i++)
            System.out.printf("%d/%d ",a.vec[i].numerator,a.vec[i].denominator);
        System.out.printf("- вектор a\n");
        for (int i = 0; i < b.n; i++)
            System.out.printf("%d/%d ",b.vec[i].numerator,b.vec[i].denominator);
        System.out.printf("- вектор b\n");
        myVec z = myVec.sum( a , b );
        for (int i = 0; i < z.n; i++)
            System.out.printf("%d/%d ",z.vec[i].numerator,z.vec[i].denominator);
        System.out.printf("- вектор a+b\n");
        if (myVec.orthogonality(a,b) == 0)
            System.out.printf("Не ортогональны");
        else
            System.out.printf("Ортогональны");
    }
}
