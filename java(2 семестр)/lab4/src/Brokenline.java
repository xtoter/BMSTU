import java.util.Iterator;
public class Brokenline implements Iterable <Vector>{
    private Line[] brokenline;
    private int n;
    public Brokenline(Line[] x) {
        this.brokenline=x;
        this.n= x.length;
    }
    public Iterator <Vector> iterator ( ) {return new BrokenlineIterator( ) ; }
    public static void Print (Brokenline x) {
        for (int i = 0; i < x.n; i++) {
            Line.Print(x.brokenline[i]);
        }
    }
    private class BrokenlineIterator implements Iterator <Vector> {
        int pos=0;
        public BrokenlineIterator ( ) { pos = 0;}
        public boolean hasNext ( ) {return pos<n;}
        public Vector next () {
            return new Vector(brokenline[pos++]);
        }
    }
}
