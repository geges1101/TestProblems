import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        String[] arr = sc.nextLine().split(", ");
        int len = arr.length;

        int[] nums = new int[len];
        for(int i = 0; i != len; i++){
            nums[i] = Integer.parseInt(arr[i]);
        }

        int l = 0, r = 1, prev = 0, max = 0;
        int start = 0, end = 0;
        while(r < len){
            if(nums[prev] >= nums[r]){
                if(prev - l >= max){
                    max = prev - l;
                    start = l;
                    end = prev;
                }
                l = r;
            }
            prev++;
            r++;
        }
        System.out.print(start + ", " + end);
    }
}