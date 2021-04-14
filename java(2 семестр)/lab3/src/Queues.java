import java.util.*;
public class Queues {
    private int n;
    private Queue[] currentQueues;

    public Queues(Queue[] x) {
        this.currentQueues=x;
        this.n=x.length;

    }
    public static void sort (Queues x){
        Arrays.sort(x.currentQueues);
    }
    public static void print (Queues x){
        for (int i = 0; i < x.n; i++) {
            Queue.print(x.currentQueues[i]);
            System.out.println();
        }
        System.out.println();
    }
}