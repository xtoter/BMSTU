import java.math.BigInteger;
import java.util.ArrayList;
import java.util.stream.Stream;

public class CreateStream {

    public static Stream CreateStream() {
        String b = "1234567890";
        ArrayList<BigInteger> vals = new ArrayList<>();
        int i =0;
        vals.add(new BigInteger(b));
        b = Newnumber.New(i++, b);
        vals.add(new BigInteger(b));
        b = Newnumber.New(i++, b);
        vals.add(new BigInteger(b));
        b = Newnumber.New(i++, b);
        vals.add(new BigInteger(b));
        /*for (int i = 0; i < 100; i++) {
            vals.add(new BigInteger(b));
            b = Newnumber.New(i, b);

        }*/
        return vals.stream();
    }

}
