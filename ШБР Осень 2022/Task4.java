import java.util.*;

public class Main {
    public static class Pair{
        Pair(int i, int j){
            this.i = i;
            this.j = j;
        }
        public int i;
        public int j;
    }
    public static void traverse(int i, int j, char c){
        if(i < 0 || i >= n ||
                j < 0 || j >= m) return;
        if(grid[i][j] == '#') return;

        grid[i][j] = c;
        Pair p = new Pair(i,j);
        visited.add(p);

        if(!visited.contains(new Pair(i - 1, j))) traverse(i - 1, j, 'R');
        if(!visited.contains(new Pair(i, j + 1))) traverse(i, j + 1, 'D');
        if(!visited.contains(new Pair(i + 1, j))) traverse(i + 1, j, 'L');
        if(!visited.contains(new Pair(i, j - 1))) traverse(i, j - 1, 'U');
    }

    public static int n;
    public static int m;

    public static HashSet<Pair> visited;

    public static char[][] grid;

    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        n = sc.nextInt();
        m = sc.nextInt();
        grid = new char[n][m];
        for (int i = 0; i != n; i++) {
            grid[i] = sc.next().toCharArray();
        }
        for (int i = 0; i != n; i++) {
            for (int j = 0; j != m; j++) {
                if (grid[i][j] == 'S') traverse(i, j, 'S');
            }
        }
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < m; j++) {
                System.out.print(grid[i][j] + " ");
            }
            System.out.println();
        }
    }
    }