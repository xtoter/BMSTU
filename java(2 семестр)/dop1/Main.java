public class Main {
    public static void main(String[] args) {
        Particle temp[] = {new Particle(2, 4,3,0),new Particle(4, 1,0,0)};
        Universe x = new Universe(temp);
        for (int i = 0; i < Universe.getN(x); i++) {
            System.out.printf("m = %f vx = %f vy = %f vz = %f v = %f, ",Particle.getweight(Universe.getpart(x,i)),Particle.getspeedx(Universe.getpart(x,i)),Particle.getspeedy(Universe.getpart(x,i)),Particle.getspeedz(Universe.getpart(x,i)),Particle.getspeed(Universe.getpart(x,i)));
        }
        System.out.printf("\nKinetic energy = %f",Universe.getenergy(x));
    }
}
