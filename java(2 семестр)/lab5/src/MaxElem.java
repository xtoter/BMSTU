import java.math.*;
import java.util.Comparator;
import java.util.List;
import java.util.Optional;
import java.util.Set;
import java.util.stream.Collectors;
import java.util.stream.Stream;
/*class Comparatortest implements Comparator<BigInteger> {

    public int compare(BigInteger a, BigInteger b){

        return a.compareTo(b);
    }
}*/
public class MaxElem {


    public static Optional MaxElem (Stream x){
        Optional temp = x.sorted().skip(4-1).findFirst();
        return temp;
    }
}
