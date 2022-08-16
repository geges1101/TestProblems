#include <iostream>
#include <vector>
#include <string>

using namespace std;

        int main() {
        int n , q;
        cin >> n >> q;
        string val = "";
        vector<string> values;
        for (int i = 0 ; i < n ; i++){
        cin >> val;
        items.push_back(val);
        }
        string s = "";
        int res = 0;
        int count = 0;
        for (int i = 0 ; i < q; i++){
        cin >> count >> s;
        res = 0;
        for (int j = 0 ; j < n; j++){
        if (s == items[j].substr(0 , s.length()) && count + 1 > 0){
        count--;
        res = j;
        }
        }
        if (res == 0) {
        cout << -1;
        } else{
        cout << res + 1 << endl;
        }
        }

        return 0;
        }