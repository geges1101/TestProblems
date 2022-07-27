package com.company;

import java.security.Key;
import java.util.*;

public class TaskB {

    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        String s = sc.nextLine();
        HashMap<Integer, Integer> map = new HashMap<>();
        StringBuilder sb = new StringBuilder();
        for(int i = 0; i != s.length(); i++){
            if(s.charAt(i) != '[' && s.charAt(i) != ']'
                    && s.charAt(i) != ','  && s.charAt(i) != ' '){
                sb.append(s.charAt(i));
            }
            else{
                if(sb.length() > 0){
                    if(map.containsKey(Integer.parseInt(sb.toString()))){
                        int count = map.get(Integer.parseInt(sb.toString()));
                        map.put(Integer.parseInt(sb.toString()), count + 1);
                    }
                    else{
                        map.put(Integer.parseInt(sb.toString()), 1);
                    }
                    sb.setLength(0);
                }
            }
        }
        int maxValueInMap=(Collections.max(map.values()));
        ArrayList<Integer> entries = new ArrayList<>();
        for (Map.Entry<Integer, Integer> entry : map.entrySet()) {
            if (entry.getValue()==maxValueInMap) {
                entries.add(entry.getKey());
            }
        }
        Collections.sort(entries);
        for (Integer elem: entries
             ) {
            System.out.println(elem);
        }
    }
}
