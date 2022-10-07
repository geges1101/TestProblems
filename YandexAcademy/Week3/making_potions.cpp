#include <iostream>
#include <vector>
#include <stack>
#include <unordered_set>

using namespace std;

int main() {
    int n,m;
    cin >> n >> m;
    vector<int> ing(n);
    unordered_set<string> dna;

    for(int i = 0; i != n; i++) cin >> ing[i];

    if(n == 1){
        cout << ing[0];
        return 0;
    }

    std::sort(ing.begin(), ing.end());
    std::reverse(ing.begin(), ing.end());

    int l,r,one,max;
    l = one = max = 0;
    r = 1;

    while(m > 0){
        if(ing[l] + ing[r] > ing[one] && l < n - 1 && l != r){
            max += ing[l] + ing[r];
            m -= 2;
            if(r < n - 1) r++;
            else{
                l++;
                r = l + 1;
            }
        }
        else if(one < n - 1){
            max += ing[one];
            one++;
            m--;
        }
    }

    cout << max;
}
