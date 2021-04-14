public class Matrix implements Comparable {
    private int point;
    private int n;
    private int m;
    private int[][] thisMatrix;

    public Matrix(int[][] x,int n,int m) {
        this.n=n;
        this.m=m;
        this.thisMatrix = x;
        this.point=0;
        int condition=0;
        int[] points = new int[m];
        for (int i = 0; i < this.m; i++) {
            points[i]=-2147483648;
        }
        for (int i = 0; i < this.m; i++) {
            for (int j = 0; j < this.n; j++) {
                if (this.thisMatrix[j][i]>points[i])
                {
                    points[i]=this.thisMatrix[j][i];
                }
            }
        }
        for (int i = 0; i < this.n; i++) {
            int a;int b;int c;
            c=2147483647;a=-1;b=-1;
            for (int j = 0; j < this.m; j++) {
                if (this.thisMatrix[i][j]<c)
                {
                    c=this.thisMatrix[i][j];
                    a=i;
                    b=j;
                }
            }
            if (points[b]==c)
            {
                this.point = c;
                condition=1;
            }
        }
        if (condition == 0){
            int a=-2147483648;
            for (int i = 0; i < this.m; i++) {
                if (points[i]>a){
                    a=points[i];
                }
            }
            this.point =a;
        }
    }
    public int compareTo(Object obj) {
        Matrix tmp = (Matrix)obj;
        if(this.point < tmp.point) {
            return -1;
        }
        else if(this.point > tmp.point) {
            return 1;
        }
        return 0;
    }
    public static void print (Matrix x){
        for (int i = 0; i < x.n; i++) {
            for (int j = 0; j < x.m; j++) {
                System.out.print(x.thisMatrix[i][j] + " ");
            }
            System.out.println(" ");
        }
    }
}