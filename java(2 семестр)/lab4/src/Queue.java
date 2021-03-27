public class Queue implements Comparable {
    private int sum;
    private int[] thisqueue;

    public Queue(int[] x) {
        this.thisqueue = x;
        this.sum=0;
        for (int i = 0; i < x.length; i++) {
            this.sum=this.sum+x[i];
        }
    }
    public int compareTo(Object obj) {
        Queue tmp = (Queue)obj;
        if(this.sum < tmp.sum) {
            return -1;
        }
        else if(this.sum > tmp.sum) {
            return 1;
        }
        return 0;
    }
    public static void print (Queue x){
        for (int i = 0; i < x.thisqueue.length; i++) {
            System.out.print(x.thisqueue[i] + " ");
        }
    }
}