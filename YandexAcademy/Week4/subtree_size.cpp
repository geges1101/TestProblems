#include <iostream>
#include <vector>
#include <unordered_map>

using namespace std;

int dfs(int idx, unordered_map<int, vector<int>>& nb, vector<int>& sub){
        sub[idx] = 1;
        for(int i : nb[idx]){
            if(sub[i] == -1)
                sub[idx] += dfs(i, nb, sub);
        }
        return  sub[idx];
    }

int main(){
        int n;
        cin >> n;
        unordered_map<int, vector<int>> neighbors(n + 1);

        for(int i = 1; i != n; i++){
            int l,r;
            cin >> l >> r;
            neighbors[l].push_back(r);
            neighbors[r].push_back(l);
        }

    vector<int> subtreeSize(n + 1);
    std::fill(subtreeSize.begin(), subtreeSize.end(),-1);
        dfs(1, neighbors, subtreeSize);
        subtreeSize.erase(subtreeSize.begin());
        for(auto i : subtreeSize)
            cout << i << " ";
    }

