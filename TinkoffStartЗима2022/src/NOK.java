import java.util.Scanner;

public class NOK {
    public static int gcd(int a, int b)
    {
        if (b==0)
            return a;
        else
            return gcd(b,a%b);
    }
    public static int lcm(int a, int b)
    {
        return a*b/gcd(a,b);
    }

    public static void main(String[] args){
        Scanner sc = new Scanner(System.in);
        int n = sc.nextInt();

        if(n % 2 == 0)System.out.printf("%s %s",n/2, n/2);
        else{
            int a = n /2;
            int b = a + 1;
            int l = a, r = b;
            int min_nok = Integer.MAX_VALUE;
            while(a > 0){
                if(b % a == 0){
                    int curr = lcm(a,b);
                    if(curr <=  min_nok){
                        l = a;
                        r = b;
                        min_nok = curr;
                    }
                }
                a--;
                b++;
            }
            System.out.printf("%s %s %s",l, r,"\n");
        }
    }
}