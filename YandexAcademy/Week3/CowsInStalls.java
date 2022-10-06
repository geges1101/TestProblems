import java.util.Scanner;

public class Main {
    public static boolean validate(int idx, int cows, int[] stalls){
        int count = 1;
        int last = stalls[0];
        for(int i = 0; i != stalls.length; i++){
            if(stalls[i] - last >= idx){
                count++;
                last = stalls[i];
            }
        }
        return count >= cows;
    }

    public static void main(String[] args){
        Scanner sc = new Scanner(System.in);
        int n = sc.nextInt();
        int cows = sc.nextInt();
        int[] stalls = new int[n];

        for(int i = 0; i != n; i++){
            stalls[i] = sc.nextInt();
        }

        int l = 0;
        int r = stalls[n - 1];

        while (l < r){
            int mid = (l + r + 1) / 2;
            if(validate(mid, cows, stalls)){
                l = mid;
            }
            else r = mid - 1;
        }

        System.out.println(l);
    }
}
