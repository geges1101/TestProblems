#include <iostream>
#include <unordered_map>
#include <unordered_set>
#include <vector>

using namespace std;

int main() {
    unordered_map<int,int> map;
    for(int i = 0; i != 3; i++){
        int n;
        cin >> n;
        unordered_set<int> visited;
        for(int j = 0; j != n; j++){
            int curr;
            cin >> curr;
            if(!visited.contains(curr))
                map[curr]++;
            visited.insert(curr);
        }
    }

    vector<int> result;
    for(auto i : map){
        if(i.second >= 2) result.push_back(i.first);
    }

    std::sort(result.begin(), result.end());
    for(auto i : result){
        cout << i << ' ';
    }
}