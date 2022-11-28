import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        byte n = sc.nextByte();
        byte enter = sc.nextByte();
        byte exit = sc.nextByte();
        byte max = (byte) Math.max(enter,exit);
        byte min = (byte) Math.min(enter,exit);
        System.out.print(Math.min(max - min, n - max + min) - 1);
    }
}