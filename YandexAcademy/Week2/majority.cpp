#include <iostream>
#include <unordered_map>

using namespace std;

int main() {
    int n;
    cin >> n;

    unordered_map<int, int> map;

    for(int i = 0; i != n; i++){
        int curr;
        cin >> curr;
        map[curr]++;
    }

    for(auto i : map)
        if(i.second > n / 2) cout << i.first;
}
