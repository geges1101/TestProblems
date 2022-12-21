import java.util.Scanner;

public class Colors {
    public static void main(String[] args){
        Scanner sc = new Scanner(System.in);
        int letters = Integer.parseInt(sc.nextLine());
        String words = sc.nextLine();
        String line = sc.nextLine();

        int pos = 0, count = 0, curr = 0;
        int l = 0, r = 1;

        while(r < letters - 1 && pos < letters - 1){
            if(words.charAt(r) == ' '){
                l = r + 1;
                r = l + 1;
                if(curr > 0) count++;
                curr = 0;
            }
            else{
                if(line.charAt(pos) == line.charAt(pos + 1))curr++;
                l++;
                r++;
                pos++;
            }
        }

        System.out.print(count);
    }
}

//7
//Tinkoff
//BYBYBYB

//27
//Algorithms and Data Structures
//BBBBBBBBBBBYBYYYYBBBBBBBBBB