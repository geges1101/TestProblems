import java.util.*;

public class Main {
    private static boolean isDigit(String s) throws NumberFormatException {
        try {
            Integer.parseInt(s);
            return true;
        } catch (NumberFormatException e) {
            return false;
        }
    }
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        while(sc.hasNext()){
            HashMap<String, Stack<String>> all = new HashMap<>();
            String line = sc.nextLine();
            if(Objects.equals(line, "{")){
                HashMap<String, Integer> map = new HashMap<>();
                while(line != "}"){
                    line = sc.nextLine();
                    int pivot = line.indexOf('=');
                    String var = line.substring(0, pivot);
                    String val = line.substring(pivot + 1, line.length());
                    int count = map.getOrDefault(line, 0);
                    map.put(var, count + 1);
                    if(isDigit(val)){
                        Stack<String> stack = all.get(var);
                        stack.push(val);
                        all.put(var, stack);
                        System.out.println(val);
                    }
                    else{
                        if(all.containsKey(val)){
                            String stack_val = all.get(val).peek();
                            Stack<String> stack_var = all.get(var);
                            stack_var.push(stack_val);
                            System.out.println(stack_val);
                        }
                        else {
                            Stack<String> stack_var = all.get(var);
                            stack_var.push("0");
                            System.out.println(0);
                        }
                    }
                }
                for (Map.Entry<String,Stack<String>> entry : all.entrySet())
                    if(map.containsKey(entry.getKey())){
                        int n = map.get(entry.getKey());
                        while(n > 0){
                            entry.getValue().pop();
                            n--;
                        }
                    }
            }
            else{
                line = sc.nextLine();
                int pivot = line.indexOf('=');
                String var = line.substring(0, pivot);
                String val = line.substring(pivot + 1, line.length());
                if(isDigit(val)){
                    Stack<String> stack = all.get(var);
                    stack.push(val);
                    all.put(var, stack);
                    System.out.println(val);
                }
                else{
                    if(all.containsKey(val)){
                        String stack_val = all.get(val).peek();
                        Stack<String> stack_var = all.get(var);
                        stack_var.push(stack_val);
                        System.out.println(stack_val);
                    }
                    else {
                        Stack<String> stack_var = all.get(var);
                        stack_var.push("0");
                        System.out.println(0);
                    }
                }
            }

        }

    }
}