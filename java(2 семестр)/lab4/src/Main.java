public class Main {

    public static void main(String[] args) {
        Line[] templine = {new Line(1, 2, 2, 3), new Line(2, 3, 4, 5)};
        Brokenline brokenline = new Brokenline(templine);
        Brokenline.Print(brokenline);
        for (Vector s : brokenline) Vector.Print(s);
    }

}
