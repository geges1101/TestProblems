import java.util.Arrays;
import java.util.Map;
import java.util.Scanner;

public class timeDifference {

    public static void main(String[] arg) {
        Scanner sc = new Scanner(System.in);
        int n = sc.nextInt();
        int[] trains = new int[n];

        for(int i = 0; i != n; i++){
            String s = sc.next();
            String[] curr = s.split(":");
            trains[i] = Integer.parseInt(curr[0]) * 60 + Integer.parseInt(curr[1]);
        }
        Arrays.sort(trains);

        int min = 24 * 60 + trains[0] - trains[n - 1];
        for(int i = 0; i != n; i++){
            min = Math.min(min, trains[i] = trains[i] - 1);
        }
        System.out.println(min);
    }

}
