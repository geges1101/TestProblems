import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        byte code = sc.nextByte();
        byte interactor = sc.nextByte();
        byte checker = sc.nextByte();

        byte result;
        switch (interactor){
            case(0):
                if(code != 0)result = 3;
                else result = checker;
                break;
            case(1):
                result = checker;
                break;
            case(4):
                if(code != 0) result = 3;
                else result = 4;
                break;
            case(6):
                result = 0;
                break;
            case(7):
                result = 1;
                break;
            default:
                result = interactor;
        }

        System.out.print(result);
    }
}