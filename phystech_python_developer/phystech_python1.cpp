#include <iostream>
#include <string>
#include <vector>
#include <algorithm>

using namespace std;

int calculate(const string& s){
    vector<int> v;
    int l = 0, r = 1, len = s.length();
    while(r <= len){
        if(r == len){
            v.push_back(stoi(s.substr(l , r)));
            break;
        }
        if(s[r] == '+'){
            v.push_back(stoi(s.substr(l , r)));
            l = r + 1;
            r ++;
        }
        r++;
    }
    int res = 0;
    for(int i : v) res += i;
    return res;
}

int main() {
    int n,m;
    cin >> n >> m;

    string result = "";
    char c = '1';
    for(int i = 1; i <= n; i++){
        result.push_back(c);
        c++;
    }

    int pos = 1, last_pos = 1;
    int curr = calculate(result);
    while(curr != m){
        if(curr > m){
            result.insert(pos,"+");
            last_pos = pos;
            pos += 2;
        }
        else if(curr < m){
            result.erase(last_pos, 1);
            int len = result.length();
            for(int i = len; i >= 0; i--){
                if(result[i] == '+'){
                    last_pos = i;
                    break;
                }
            }
            while(calculate(result) > m && last_pos <= len - 1){
                swap(result[last_pos], result[last_pos + 1]);
                last_pos++;
            }
            pos = last_pos++;
        }
        curr = calculate(result);
    }
    cout << result;
}
