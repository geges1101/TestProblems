import java.util.Scanner;

public class wordChange {
    public static void main(String[] arg){
        Scanner sc = new Scanner(System.in);
        String[] small = sc.nextLine().split(" ");
        String[] big = sc.nextLine().split(" ");

        int n;

        for(int i = 0; i != big.length; i++){
            n = big[i].length();
            for(int j = 0; j != small.length; j++){
                if(big[i].startsWith(small[j])
                && small[j].length() < n)
                    big[i] = small[j];
            }
        }

        n = big.length;
        for (int i = 0; i != n; i++) {
            System.out.print(big[i] + ' ');
        }
    }
}
