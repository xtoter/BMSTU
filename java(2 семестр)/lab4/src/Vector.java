public class Vector {
    private int x;
    private int y;
    public Vector(Line z) {
        int x1= z.getX1();
        int x2= z.getX2();
        int y1= z.getY1();
        int y2= z.getY2();
        this.x=-1*(y2-y1);
        this.y=x2-x1;
    }
    public static void Print (Vector x){
        System.out.println(x.x + " " + x.y + " " );
    }
}
