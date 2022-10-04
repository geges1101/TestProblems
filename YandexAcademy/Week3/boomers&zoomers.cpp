#include <iostream>
#include <vector>

using namespace std;

int main() {
    int n;
    cin >> n;

    vector<long> nums(n);
    for(int i = 0; i != n; i++) cin >> nums[i];

    std::sort(nums.begin(), nums.end());

    int l,r;
    long count;
    count = 0;
    l = r =0;

    for(int i = 0; i != n; i++){
        while(l < n && nums[l] <= nums[i]/2 + 7){
            l++;
        }
        while(r < n && nums[r] <= nums[i]){
            r++;
        }
        if(r > l + 1) count += r - l - 1;
    }

    cout << count;

}
