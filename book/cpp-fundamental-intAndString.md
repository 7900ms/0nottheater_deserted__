1.1
```c++
// g++ string.cpp -o string (NO, In need of C++11)
// ./string

// g++ -std=c++11 string.cpp -o string
// ./string

#include <iostream>
using namespace std;

int main(){

	int a = 99;
	cout << "a is " << a << endl;
	cout << endl;
	cout << endl;

	string s = "Player";
	cout << s << endl;
	cout << endl;
	cout << endl;

	cout << s << a << endl;
	cout << endl;
	cout << endl;

	string str1 = "Player";
	int a1 = 100;
	str1 += to_string(a1);
	cout << "str1: " << str1 << endl;
	cout << endl;
	cout << endl;

	string str2 = "";
	str2 = s + to_string(a);
	cout << "str2: " << str2 << endl;
	cout << endl;
	cout << endl;

	return 0;
}



// G cpp string int print
// https://stackoverflow.com/questions/64782/how-do-you-append-an-int-to-a-string-in-c

// 说明：<iostream> 包含了 <string> . 在 c++11 里可以直接用 to_string


```
