import java.util.Scanner;

public class BuildingStairs {
    public static void main(String[] arg){
        Scanner sc = new Scanner(System.in);
        int n = sc.nextInt();
        int curr = 0;
        int idx = 1;

        while (idx <= n){
            n -= idx;
            curr++;
            idx++;
        }
        System.out.println(idx - 1);
    }
}
