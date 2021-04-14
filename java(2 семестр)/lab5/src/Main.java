import java.util.*;
import java.util.stream.Stream;

public class Main {

    public static void main(String[] args) {
        Stream stream=CreateStream.CreateStream();
        Optional result = MaxElem.MaxElem(stream);
        System.out.println(result.get());

    }
}