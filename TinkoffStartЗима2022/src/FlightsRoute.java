import java.util.ArrayList;
import java.util.HashMap;
import java.util.Scanner;

public class FlightsRoute {
    public static int n;
    public static boolean reached;
    public static HashMap<Integer, ArrayList<Flight>> map;
    static class Flight{
        int begin;
        int end;
        byte odd;
        Flight(int b, int e, byte o){
            this.begin = b;
            this.end = e;
            this.odd = o;
        }
    }

    static int travel(int curr, int count, int ignore){
        if(curr == n){
            reached = true;
            return count;
        }
        ArrayList<Flight> flights = map.get(curr);
        ArrayList<Flight> airport = new ArrayList<>();
        for(Flight f : flights){
            if(f.odd != ignore) airport.add(f);
        }

        if(airport.size() == 0) return -1;
        int min = Integer.MAX_VALUE;
        for(Flight f : airport){
            int min_now = Math.min(travel(f.end, count + 1, 0),
                    travel(f.end, count + 1, 1));
            min = Math.min(min, min_now);
        }
        return -1;
    }

    public static void main(String args[]) {
        Scanner sc = new Scanner(System.in);
        n = sc.nextInt();
        int m = sc.nextInt();

        for(int i = 0; i != m; i++){
            int begin = sc.nextInt();
            int end = sc.nextInt();
            Flight curr = new Flight(begin,end, sc.nextByte());
            ArrayList<Flight> b = map.getOrDefault(begin, new ArrayList<>());
            b.add(curr);
            map.put(begin, b);
        }

        int lhs = travel(1, 0, 0);
        int rhs = travel(1, 0, 1);
        System.out.print(Math.min(lhs,rhs));
    }
}