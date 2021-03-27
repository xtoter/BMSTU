public class Particle {
    private double m;
    private double vx;
    private double vy;
    private double vz;
    public Particle(double m,double vx,double vy,double vz){
        this.m=m;
        this.vx=vx;
        this.vy=vy;
        this.vz=vz;
    }
    public static double getweight(Particle x){
        return x.m;
    }
    public static double getspeedx(Particle x){
        return x.vx;
    }
    public static double getspeedy(Particle x){
        return x.vy;
    }
    public static double getspeedz(Particle x){
        return x.vz;
    }
    public static double getspeed(Particle x){
        return Math.sqrt(Math.pow(x.vx,2)+Math.pow(x.vy,2)+Math.pow(x.vz,2));
    }
}
