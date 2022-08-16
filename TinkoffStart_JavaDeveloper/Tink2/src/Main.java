import java.util.*;

public class Main {

    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int n = sc.nextInt();
        Map<String , Integer> result = new HashMap<>();
        String s;
        for (int i = 0 ; i < n; i++){
            s = sc.nextLine();
            s = s.chars()        // IntStream
                    .sorted()
                    .collect(StringBuilder::new,
                            StringBuilder::appendCodePoint,
                            StringBuilder::append)
                    .toString();
            int count = result.getOrDefault(s, 0);
            result.put(s, count++);
        }
        int max = 0;
        for (Map.Entry<String,Integer> entry : result.entrySet()){
            if(max < entry.getValue()) max = entry.getValue();
        }
        System.out.println(max);
    }
}