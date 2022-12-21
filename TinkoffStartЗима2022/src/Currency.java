import java.util.Arrays;
import java.util.HashSet;
import java.util.Scanner;

public class Currency {
    public static int[] values;
    public static void helper(int[] wallet, HashSet<String> combos, int idx, int other){
        if(other < 0) other = 2;
        else if(other > 2) other = 0;

        if(wallet[idx] > 0){
            wallet[idx] -= values[idx];
            wallet[other] += values[other];
            if(!combos.contains(Arrays.toString(wallet))){
                combos.add(Arrays.toString(wallet));
                for(int i = 0; i != 3; i++){
                    helper(wallet.clone(),combos, i, i - 1);
                    helper(wallet.clone(),combos, i, i + 1);
                }
            }
        }
    }

    public static void main(String[] args){
        Scanner sc = new Scanner(System.in);
        values = new int[]{sc.nextInt(), sc.nextInt(), sc.nextInt()};
        int[] wallet = {sc.nextInt(), sc.nextInt(), sc.nextInt()};
        HashSet<String> combos = new HashSet<>();

        for(int i = 0; i != 3; i++){
            helper(wallet.clone(),combos, i, i - 1);
            helper(wallet.clone(), combos, i, i + 1);
        }

        System.out.print(combos.size());
    }
}

//1 1 1
//1 0 2

//1 2 3
//3 5 4
