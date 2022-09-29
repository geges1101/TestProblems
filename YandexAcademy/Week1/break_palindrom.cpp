#include <iostream>

using namespace std;

int main() {
    string s;
    cin >> s;
    int n = s.length();

    string min = s;
    bool f = false;
    std::fill(min.begin(), min.end(),'z');

    for(int i = 0; i != n; i++){
        char c = 'a';
        while(c <= 'z'){
            string temp = s;
            temp[i] = c;
            if(temp[i] != temp[n - i - 1] && temp < min){
                min = temp;
                f = true;
            }

            c++;
        }
    }

    if(!f) min = "";
    cout << min;
}
