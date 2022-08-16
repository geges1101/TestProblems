import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int x1 = sc.nextInt();
        int y1 = sc.nextInt();
        int x2 = sc.nextInt();
        int y2 = sc.nextInt();

        int new_x1 = sc.nextInt();
        int new_y1 = sc.nextInt();
        int new_x2 = sc.nextInt();
        int new_y2 = sc.nextInt();
        int max = Math.max(Math.max(Math.abs(new_y2 - y1), Math.abs(y2 - new_y1)), Math.max(Math.abs(new_x2 - x1), Math.abs(x2 - new_x1)));
        System.out.println(max * max);
    }
}
