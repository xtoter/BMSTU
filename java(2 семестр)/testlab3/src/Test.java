public class Test {
    public static void main(String[] args) {
        int tempa[] = {1,2,3,6,8};
        int tempb[] = {1,2,3,5,77};
        int tempc[] = {0};
        Queue a[] ={new Queue(tempa),new Queue(tempb),new Queue(tempc)};
        Queues b = new Queues(a);
        Queues.print(b);
        Queues.sort(b);
        Queues.print(b);
    }
}