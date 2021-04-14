public class Newnumber {
    public static String New(int x,String S) {
        int n=x%9;
        char[] c = S.toCharArray();
        char temp = c[n];
        c[n] = c[n+1];
        c[n+1] = temp;

        return new String(c);

    }
}
