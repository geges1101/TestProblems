import java.util.Arrays;
import java.util.Scanner;

public class Annagram {
        public static void main(String[] arg){
            Scanner sc = new Scanner(System.in);
            String first = sc.next();
            String second = sc.next();

            char[] left = first.toCharArray();
            char[] right = second.toCharArray();

            Arrays.sort(left);
            Arrays.sort(right);

            if(Arrays.equals(left, right)) System.out.print("YES");
            else System.out.print("NO");
        }
}
