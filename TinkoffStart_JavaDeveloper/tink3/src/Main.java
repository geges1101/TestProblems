import java.util.*;

public class Main {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int n = sc.nextInt();
        int sub;
        int val = 0;
        List<Integer> on = new ArrayList<>();
        List<Integer> off = new ArrayList<>();
        if (n == 1){
            System.out.println(sc.nextInt());
            return;
        } else if (n == 2){
            for (int i =0; i < n; i++) {
                sub = sc.nextInt();
                if (i % 2 == 0) val += sub;
                else val -= sub;
            }
            System.out.println(Math.abs(val));
            return;
        }
        for (int i =0; i < n; i++) {
            sub = sc.nextInt();
            if (i % 2 == 0) {
                val += sub;
                on.add(sub);
            }
            else {
                val -= sub;
                off.add(sub);
            }
        }
        Arrays.sort(on.toArray());
        Arrays.sort(off.toArray());
        int result = val + 2 * off.get(off.size() - 1) - 2 * on.get(0);
        System.out.println(Math.max(val, result));
    }
}

