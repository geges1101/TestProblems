import java.util.*;

public class Main {
    public static class Pair{
        Pair(int first, int second){
            this.first = first;
            this.second = second;
        }
        public int first;
        public int second;
    }

    public static boolean check(Pair[] elevators, int point){
        for(int i = 0; i != elevators.length; i++){
            if(elevators[i] != null && elevators[i].first == point) return true;
        }
        return false;
    }
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int n = sc.nextInt();
        Pair[] elevators = new Pair[7];
        int max = 0;
        for(int i = 0; i != n; i++){
            Pair elevator = new Pair(sc.nextInt(), sc.nextInt());
            elevators[i] = elevator;
            if(elevator.second > max) max = elevator.second;
        }
        int curr = 0;
        int count = 0;
        while(curr < max){
            int min = Integer.MAX_VALUE;
            int idx = 0;
            for(int i = 0; i != n; i++){
                if(elevators[i] != null
                        && elevators[i].second < min
                        && check(elevators, elevators[i].second)
                        && elevators[i].first == curr){
                    min = elevators[i].second;
                    idx = i;
                }
            }
            curr = min;
            elevators[idx] = null;
            count++;
        }
        System.out.println(count);
    }
}

