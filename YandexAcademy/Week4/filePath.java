import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
import java.util.Scanner;

public class filePath {
    public static void main(String[] arg){
        Scanner sc = new Scanner(System.in);
        String target = sc.next();
        int n = sc.nextInt();
        String[] all = new String[n];

        ArrayList<ArrayList<String>> paths = new ArrayList<>();
        for(int i = 0; i != n; i++){
            String curr = sc.next();
            all[i] = curr;
            if(curr.contains(".")){
                int pos = 0;
                while(curr.charAt(pos) == ' ')
                    pos++;

                curr = curr.substring(pos, curr.length());
                if(curr.equals(target)){
                    ArrayList<String> path = new ArrayList<>();
                    path.add(target);
                    int j = i - 1;
                    while(j >= 0){
                        curr = all[j];
                        int begin = 0;
                        while(curr.charAt(begin) == ' ')
                            begin++;
                        if(begin < pos){
                            path.add(curr.substring(begin, curr.length()));
                        }
                        pos = begin;
                        paths.add(path);
                        j--;
                    }
                }
            }
        }

        int min = Integer.MAX_VALUE;
        int idx = 0;
        for(int i = 0; i != paths.size(); i++){
            if(paths.get(i).size() < min){
                min = paths.get(i).size();
                idx = i;
            }
        }

        ArrayList<String> result_path = paths.get(idx);
        Collections.reverse(result_path);
        for(String s : result_path){
            System.out.print('/' + s);
        }
    }
}
