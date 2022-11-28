#include <iostream>
#include <vector>

using namespace std;

int main()
{
    vector<int> v;
    int n, curr;
    cin >>n;
    for (int i = 0; i < n; ++i){
        cin >> curr;
        v.push_back(curr) ;
    }
    cout << v[v.size()/2];
    return 0;
}