public class myVec  {
    int n;
    Fraction vec[];
    public myVec(Fraction[] vec) {
        this.n = vec.length;
        this.vec = vec;
    }
    public myVec(int n) {
        this.n = n;
        this.vec = new Fraction[n];;
    }
    public static myVec sum (myVec a, myVec b){
        myVec z = new myVec(a.n);
        if (a.n == b.n)
        {
            for (int i = 0; i < a.n; i++){
                z.vec[i]=new Fraction(0,0);
                z.vec[i].denominator = a.vec[i].denominator*b.vec[i].denominator;
                z.vec[i].numerator = (a.vec[i].numerator * b.vec[i].denominator) + (b.vec[i].numerator * a.vec[i].denominator);
            }
        }
        else {
            System.out.println("error");
            System.exit(0);
        }
        return z;
    }
    public static int orthogonality (myVec a, myVec b){
        int temp = 0;
        if (a.n == b.n) {
            for (int i = 0; i < a.n; i++) {
                temp = temp + a.vec[i].numerator * b.vec[i].numerator;
            }
        }
        else {
            System.out.println("error");
            System.exit(0);
        }
        if (temp == 0)
            return 1;
        else
            return 0;
    }
}
