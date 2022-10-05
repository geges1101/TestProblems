#include <iostream>
#include <vector>

using namespace std;

int main() {
    int n;
    cin >> n;
    vector<int> nums(n);
    vector<int> positive(n + 1);

    for(int i = 0; i != n; i++){
        cin >> nums[i];
        if(nums[i] > 0) positive[i + 1] = positive[i] + 1;
        else positive[i + 1] = positive[i];
    }

    int q;
    cin >> q;

    for(int i = 0; i != q; i++){
        int begin, end;
        cin >> begin >> end;
        cout << positive[end] - positive[begin - 1] << endl;
    }
}
