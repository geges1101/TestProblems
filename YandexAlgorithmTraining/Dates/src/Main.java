import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        byte a = sc.nextByte();
        byte b = sc.nextByte();
        if(a == b) System.out.print(1);
        else if(a <= 12 && b <= 12) System.out.print(0);
        else System.out.print(1);
    }
}