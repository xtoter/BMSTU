public class Test {
    public static void main(String[] args) {
        int [][] tempa1 = {{9,9,9,9}, {9,9,9,9}, {9,9,9,9}};
        int [][] tempa2 = {{1,1,1}, {1,1,1}};
        int [][] tempa3 = {{8,7,0,6}, {6,8,5,10}, {8,8,5,10}};
        Matrix a1 =new Matrix(tempa1,3,4);
        Matrix a2 =new Matrix(tempa2,2,3);
        Matrix a3 =new Matrix(tempa3,3,4);
        Matrix [] tempb ={a1,a2,a3};
        Matrixs b = new Matrixs(tempb);
        Matrixs.print(b);
        Matrixs.sort(b);
        Matrixs.print(b);
    }
}