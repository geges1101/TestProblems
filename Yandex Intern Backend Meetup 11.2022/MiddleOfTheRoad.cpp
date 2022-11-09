#include <iostream>

using namespace std;

int main() {
    int first, second, third;
    cin >> first >> second >> third;

    int a,b = 0;

    if((first > third && first > second) || (first < third && first < second)){
        if(abs(first - second) < abs(first - third))
            a = abs(first - second);
        else a = abs(first - third);
    }
    else a = 0;
    if((second > third && second > first) || (second < third && second < first)) b++;
    if((third > first && third > second) || (third < first && third < second)) b++;


    cout << a << ' ' << b << endl;

    b = 0;

    if((second > third && second > first) || (second < third && second < first)){
        if(abs(first - second) < abs(second - third))
            a = abs(first - second);
        else a = abs(second - third);
    }
    else a = 0;
    if((first > third && first > second) || (first < third && first < second)) b++;
    if((third > first && third > second) || (third < first && third < second)) b++;

    cout << a << ' ' << b << endl;

    b = 0;

    if((third > first && third > second) || (third < first && third < second)){
        if(abs(first - third) < abs(second - third))
            a = abs(first - third);
        else a = abs(second - third);
    }
    else a = 0;
    if((first > third && first > second) || (first < third && first < second)) b++;
    if((second > third && second > first) || (second < third && second < first)) b++;

    cout << a << ' ' << b << endl;
}
