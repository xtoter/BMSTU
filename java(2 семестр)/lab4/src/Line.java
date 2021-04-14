public class Line {
    private int x1;
    private int y1;
    private int x2;
    private int y2;
    public Line(int x1,int y1,int x2,int y2) {
        this.x1=x1;
        this.y1=y1;
        this.y2=y2;
        this.x2=x2;
    }

    public int getX1() {
        return x1;
    }

    public int getX2() {
        return x2;
    }
    public int getY1() {
        return y1;
    }

    public int getY2() {
        return y2;
    }
    public static void Print (Line x){
            System.out.println(x.x1 + " " + x.y1 + " " + x.x2 + " " + x.y2 + " ");
    }
}
