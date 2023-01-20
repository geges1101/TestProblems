import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class Main {
    static List<List<String>> process(String s){
        List<List<String>> res = new ArrayList<>();
        if(s.length() == 1) res.add(List.of(s));
        else {
            for(int i = 0; i < s.length() - 1; i++){
                String l = s.substring(0, i + 1);
                String r = s.substring(i + 1);
                for(List<String> combos : process(r)){
                    List<String> curr = new ArrayList<>();
                    curr.add(l);
                    curr.addAll(combos);
                    res.add(curr);
                }
            }
            res.add(List.of(s));
        }
        return res;
    }

    static void dfs(List<String> line, List<String> res, int idx,
                    int target, long sum, String curr){
        if(idx == line.size()){
            if(target == sum)
                if(curr.charAt(1) == '+')
                    res.add(curr.substring(2));
                else res.add(curr.substring(1));
        } else if (idx > line.size())return;
        else {
            long next = Integer.parseUnsignedInt(line.get(idx));
            dfs(line, res, idx+1, target, sum + next, curr + "+" + line.get(idx));
            dfs(line, res, idx+1, target, sum - next, curr + "-" + line.get(idx));
        }
    }


    public static void main(String[] args){
        String s = "9876543210";
        int target = 200;
        List<List<String>> combos = process(s);
        List<String> res = new ArrayList<>();
        for(List<String> c : combos){
            dfs(c, res, 0, target, 0, "0");
        }
        for (String curr : res){
            System.out.println(curr);
        }
    }
}
