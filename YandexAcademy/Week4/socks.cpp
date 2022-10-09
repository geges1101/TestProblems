#include <iostream>
#include <vector>
#include <unordered_map>

using namespace std;

int main(){
        int l,n,m;
        cin >> l >> n >> m;

        unordered_map<int,int> map;
        for(int i = 0; i != n; i++){
            int begin,end;
            cin >> begin >> end;

            while(begin <= end){
                map[begin]++;
                begin++;
            }
        }

        while(m){
            int point;
            cin >> point;
            cout << map[point] << endl;
            m--;
        }
    }
