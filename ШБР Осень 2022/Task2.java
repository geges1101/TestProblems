import java.util.*;

public class Main {
    public static class Candidate{
        static class StockComparator implements Comparator<Candidate> {
            public int compare(Candidate c1, Candidate c2)
            {
                if(Objects.equals(c1.score, c2.score)){
                    return c1.fine.compareTo(c2.fine);
                }
                else if (c1.score < c2.score)
                    return 1;
                else
                    return -1;
            }
        }
        Candidate(String name, String group, Integer score, Integer fine){
            this.name = name;
            this.group = group;
            this.score = score;
            this.fine = fine;
        }
        String name;
        String group;
        Integer score;
        Integer fine;
    }

    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int n = sc.nextInt();
        String[][] groups = new String[n][2];

        for(int i = 0; i != n; i++){
            String[] group = sc.next().split(",");
            groups[i] = group;
        }

        n = sc.nextInt();
        String[][] candidates = new String[n][4];

        for(int i = 0; i != n; i++){
            String[] candidate = sc.next().split(",");
            candidates[i] = candidate;
        }

        HashMap<String, ArrayList<Candidate>> db = new HashMap<>();
        for(int i = 0 ; i != n; i++){
            ArrayList<Candidate> updated = db.getOrDefault(candidates[i][1], new ArrayList<>());
            Candidate c = new Candidate(candidates[i][0], candidates[i][1],
                    Integer.parseInt(candidates[i][2]), Integer.parseInt(candidates[i][3]));
            updated.add(c);
            updated.sort(new Candidate.StockComparator());
            db.put(candidates[i][1], updated);
        }

        ArrayList<String> res = new ArrayList<>();
        for(int i = 0; i != groups.length; i++){
            int count = Integer.parseInt(groups[i][1]);
            for(Candidate c : db.get(groups[i][0])){
                if(count > 0){
                    res.add(c.name);
                    count--;
                }
                else break;
            }
        }

        Collections.sort(res);
        for(String s : res) System.out.println(s);
    }
}