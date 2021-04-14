import java.util.*;
public class Matrixs {
    private int n;
    private Matrix[] currentQueues;

    public Matrixs(Matrix[] x) {
        this.currentQueues=x;
        this.n=x.length;

    }
    public static void sort (Matrixs x){
        Arrays.sort(x.currentQueues);
    }
    public static void print (Matrixs x){
        for (int i = 0; i < x.n; i++) {
            Matrix.print(x.currentQueues[i]);
            System.out.println();
        }
        System.out.println();
    }
}