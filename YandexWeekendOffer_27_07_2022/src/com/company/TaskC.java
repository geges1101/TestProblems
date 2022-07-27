package com.company;

import java.util.*;

class Main {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int n = sc.nextInt();
        String[] lines = new String[n];

        for(int i = 0; i != n; i++){
            lines[i] = sc.next();
        }

//        String[] post = new String[n];
//        HashMap<Integer, String> map = new HashMap<>();
        for(int j = 0; j != n; j++){
            int length = lines[j].length();
            length -= 2;

            StringBuilder sb = new StringBuilder();
            sb.append(lines[j].charAt(0));
            sb.append(length);
            sb.append(lines[j].charAt(lines.length - 1));

            lines[j] = sb.toString();
            System.out.println(lines[j]);
        }
    }
}
