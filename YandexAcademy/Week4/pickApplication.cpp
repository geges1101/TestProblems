#include <iostream>
#include <vector>

using namespace std;

bool cmp(pair<int, int> x, pair<int, int> y)
{
    if (x.second == y.second)
        return x.first < y.first;
    else
        return x.second < y.second;
}

int main() {
    int n;
    cin >> n;
    if(n <= 1){
        cout << n;
        return 0;
    }

    vector<pair<int,int>> classes;
    for(int i = 0; i != n; i++){
       pair<int,int> next;
       cin >> next.first;
       cin >> next.second;
       classes.push_back(next);
    }

    std::sort(classes.begin(), classes.end(), cmp);

    pair<int,int> curr = classes[0];
    int max = 1;

    for(int i = 1; i != n; i++){
        if(classes[i].first >= curr.second){
            curr = classes[i];
            max++;
        }
    }

    cout << max;
}
