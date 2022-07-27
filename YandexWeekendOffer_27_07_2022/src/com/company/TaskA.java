package com.company;

import java.util.Scanner;

public class TaskA {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int n = sc.nextInt();
        int q = sc.nextInt();

        StringBuilder sb = new StringBuilder();
        for(int i = 0; i != n; i++){
            sb.append('0');
        }

        StringBuilder res = new StringBuilder();
        for(int i = 0; i != n; i++){
            res.append('1');
        }

        int min = Integer.MAX_VALUE;
        String[] first = new String[q];
        String[] second = new String[q];

        for(int i = 0; i != q ; i++){
            first[i] = sc.next();
            second[i] = sc.next();
        }

        int[] ccc = new int[q];

        for(int i = 0; i != q ; i++){
            StringBuilder curr = sb;
            while(curr != res){
                int val = Math.max(hammingDist( first[i],curr.toString()), hammingDist(second[i], curr.toString()));
                if(val < min){
                    min = val;
                }
                for(int j = curr.length() - 1; j >= 0; j--){
                    if(curr.charAt(j) == '0'){
                        curr.replace(j,j + 1, "1");
                        break;
                    }
                }
            }
            ccc[i] = min;
            min = Integer.MAX_VALUE;
        }

        for(int i = 0; i != q; i++){
            System.out.println(ccc[i]);
        }
    }
    static int hammingDist(String str1, String str2)
    {
        int i = 0, count = 0;
        while (i < str1.length()) {
            if (str1.charAt(i) != str2.charAt(i))
                count++;
            i++;
        }
        return count;
    }
}
