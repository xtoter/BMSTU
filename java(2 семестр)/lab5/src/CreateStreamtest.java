import java.math.BigInteger;
import java.util.ArrayList;
import java.util.stream.Stream;

public class CreateStreamtest {

    private ArrayList<BigInteger> data;
    private String b;
    private int i ;
    public CreateStreamtest (){
        this.data= new ArrayList<>();
        this.data.add(new BigInteger("22"));
        this.data.add(new BigInteger("23"));
        this.b = "1234567890";
        this.i=0;
    }
    public String NewInt(){
        this.b = Newnumber.New(this.i++, this.b);
        return this.b;
    }
    public Stream Cre() {
        return this.data.stream()
                .map(x -> x=new BigInteger(NewInt()));
    }

}
