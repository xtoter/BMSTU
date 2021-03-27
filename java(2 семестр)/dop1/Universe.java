public class Universe {
    private int n;
    private Particle universe[];
    public Universe(Particle[] x) {
        this.n = x.length;;
        this.universe = x;
    }
    public static Particle getpart (Universe x,int n) {
    return x.universe[n];
    }
    public static int getN (Universe x) {
        return x.n;
    }
    public static double getenergy(Universe x){
        double tempenergy = 0;
        for (int i = 0; i < x.n; i++){
        tempenergy = tempenergy + Particle.getweight(Universe.getpart(x,i))*Math.pow(Particle.getspeed(Universe.getpart(x,i)),2)/2;
        }
        return tempenergy;
    }
}
