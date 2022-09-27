#include <iostream>
#include <vector>

using namespace std;
int main() {
    int n;
    cin >> n;
    vector<int> days(n);

    for(int i = 0; i != n; i++){
        cin >> days[i];
    }

    int buy,sell, max = 0;
    int curr = days[0];

    for(int i = 0; i != n; i++){
        if(days[i] - curr <= 0) curr = days[i];
        if(days[i] - curr > max){
            buy = curr;
            sell = i;
        }
    }

    cout << buy << ' ' << sell;
}
