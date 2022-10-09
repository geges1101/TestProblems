import java.util.*;

public class classRooms {
    public static void main(String[] arg){
        Scanner sc = new Scanner(System.in);
        int n = sc.nextInt();
        int[] groups = new int[n];
        for(int i = 0; i != n; i++) groups[i] = sc.nextInt();

        int m = sc.nextInt();
        int[] classes = new int[m];
        for(int i = 0; i != m; i++) classes[i] = sc.nextInt();

        Arrays.sort(groups);
        Arrays.sort(classes);

        int idx,count;
        idx = count = 0;
        for(int i = 0; i != n; i++){
            while(idx < m && classes[idx] < groups[i])
                idx++;
            if(idx != m){
                count++;
                idx++;
            }
        }

        System.out.println(count);
    }
}
