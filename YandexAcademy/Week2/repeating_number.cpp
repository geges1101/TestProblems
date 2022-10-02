#include <iostream>
#include <unordered_set>
#include <vector>

using namespace std;
int main() {
    int n;
    int k;

    cin >> n;
    cin >> k;

    vector<int> seq(n);
    unordered_set<int> visited;

    for(int i = 0; i != n; i++){
        cin >> seq[i];
    }

    string message = "NO";
    for(int i = 0; i != n; i++){
        if(!visited.contains(seq[i])){
            visited.insert(seq[i]);
            int last = i;
            for(int j = i + 1; j != n; j++){
                if(seq[j] == seq[last]){
                    if(j - last <= k){
                        message = "YES";
                        last = j;
                    }
                    else{
                        cout << "NO";
                        return 0;
                    }
                }
            }
        }
    }

    cout << message;
}
