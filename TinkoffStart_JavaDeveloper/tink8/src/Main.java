import java.util.*;

public class Main {

    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int n = sc.nextInt();
        int m = sc.nextInt();
        String[] domains = new String[n];
        for(int i = 0; i != n; i++){
            domains[i] = sc.next();
        }
        String[][] requests = new String[m][2];
        for(int i = 0; i != m; i++){
            requests[i][0] = sc.next();
            requests[i][1] = sc.next();
        }
        for(int i = 0; i != m; i++){
            int count = 0;
            for(int j = 0; j != n; j++){
                int length = domains[j].length();
                int length1 = requests[i][0].length();
                int length2 = requests[i][1].length();
                if(domains[j].substring(0, length1).equals(requests[i][0])
                && domains[j].substring(length - length2, length).equals(requests[i][1])) count++;
            }
            System.out.println(count);
        }
    }
}

