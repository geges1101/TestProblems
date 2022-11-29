import java.util.ArrayList;
import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        ArrayList<Integer> arr = new ArrayList<>();
        while(sc.hasNextInt()){
            int curr = sc.nextInt();
            if(curr == 0) break;
            arr.add(curr);
        }

        int max = Integer.MIN_VALUE;
        for(Integer i : arr){
            max = Math.max(max, i);
        }

        int count = 0;
        for(Integer i : arr){
            if(i == max) count++;
        }

        System.out.print(count);
    }
}