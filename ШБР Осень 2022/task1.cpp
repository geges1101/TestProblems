#include <iostream>
#include <unordered_map>

using namespace std;

int main() {
    string a,b, res;
    cin >> a >> b;
    res = b;
    std::fill(res.begin(), res.end(), '!');
    unordered_map<char,int> letters;

    for(char & i : a){
        letters[i]++;
    }

    for(int i = 0; i != b.length(); i++) {
        if (a[i] == b[i]) {
            res[i] = 'P';
            letters[a[i]]--;
        }
    }

    for(int i = 0; i != b.length(); i++){
        if(res[i] != 'P' && letters[b[i]] > 0){
            letters[b[i]]--;
            res[i] = 'S';
        }
        else if(res[i] != 'P')res[i] = 'I';
    }

    cout << res;
}
