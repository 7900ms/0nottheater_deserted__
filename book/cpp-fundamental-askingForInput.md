
1.1
cin
```
// g++ helloworld.cpp -o hello
// ./hello

#include <iostream>
using namespace std;

int main(){

    int i = 0;
    int j = 0;
    cout << "please enter two numbers: " << endl; // newline

    cout << "value for i: "; // continues
    cin >> i;

    cout << "value for j: ";
    cin >> j;

    cout << "now we got 2 variables:"; // continues
    cout << " i as " << i << ", and j as " << j << ". " << endl;

    return 0;
}
```
可能从 theBuffer 里读取 （输入 “5 44” 空格后面的44会直接复制给了j）（办法1：用getline ，办法2：cin并清除theBuffer）

1.2
getline
```
// g++ helloworld.cpp -o hello
// ./hello

#include <iostream>
using namespace std;

int main(){
	
	cout << "Please input your name: ";
	string name = "";
	getline(cin, name);
	cout << "hello, " << name << "! " << endl;

	cout << "Please input your address: " << endl;
	string address = "";
	getline(cin, address);
	cout << "So " << name << " lives in " << address << "!! " << endl;
    
    return 0;
}
```
