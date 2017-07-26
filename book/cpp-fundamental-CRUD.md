cpp-fundamental-CRUD.md

#### RHEL-Desktop

1.1
Newbie test
```c++
// g++ helloworld.cpp
// ./a.out

// g++ helloworld.cpp -o hello
// ./hello

# include <iostream>
using namespace std;
int main(){

	// print hello world
	cout << "hello world" << endl;

	return 0;
}

```

#### RHEL-Desktop2

1.1
to directly use "std::cout"
```c++
// g++ helloworld.cpp -o hello
// ./hello

#include <iostream>

int main(){
	if(true){
		std::cout << "this is true" << std::endl;
	}
	
	std::cout << "hello world" << std::endl;
	
	return 0;
}
```

2.1
to cut off "std::"
```c++
// g++ helloworld.cpp -o hello
// ./hello

#include <iostream>
using namespace std;

int main(){
	if(true){
		cout << "this is true" << endl;
	}
	
	cout << "hello world" << endl;
	
	return 0;
}
```

3.1
to ask for input
```c++
// g++ helloworld.cpp -o hello
// ./hello

#include <iostream>
using namespace std;

int main(){
	if(true){
		cout << "this is true" << endl;
	}

	cout << "hello world" << endl;
	cout << endl;
	cout << endl;



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

4.1
to use function
```c++
// g++ helloworld.cpp -o hello
// ./hello

#include <iostream>
using namespace std;

int printAandB(int a, int b){
	cout << "A as " << a << " and B as " << b << ". "<< endl;
	return 0;
}

int main(){
	if(true){
		cout << "this is true" << endl;
	}

	cout << "hello world" << endl;
	cout << endl;
	cout << endl;



	int i = 0;
	int j = 0;
	cout << "please enter two numbers: " << endl; // newline

	cout << "value for i: "; // continues
	cin >> i;

	cout << "value for j: ";
	cin >> j;

	cout << "now we got 2 variables:"; // continues
	cout << " i as " << i << ", and j as " << j << ". " << endl;

	cout << endl;
	cout << endl;



	int m = 7;
	int n = 25;
	printAandB(m, n);


	return 0;
}
```
4.2
```
// g++ helloworld.cpp -o hello
// ./hello

// g++ -std=c++11 helloworld.cpp -o hello
// ./hello

#include <iostream>
using namespace std;

int printAandB(int a, int b){
	cout << "A as " << a << " and B as " << b << ". "<< endl;
	return 0;
}

string getStr(int a, int b){
	string result = "";
	result = "A as " + to_string(a) + " and B as " + to_string(b) + "! ";
	return result;
}

int main(){
	if(true){
		cout << "this is true" << endl;
	}

	cout << "hello world" << endl;
	cout << endl;
	cout << endl;



	int i = 0;
	int j = 0;
	cout << "please enter two numbers: " << endl; // newline

	cout << "value for i: "; // continues
	cin >> i;

	cout << "value for j: ";
	cin >> j;

	cout << "now we got 2 variables:"; // continues
	cout << " i as " << i << ", and j as " << j << ". " << endl;

	cout << endl;
	cout << endl;



	int m = 7;
	int n = 25;
	cout << "(this is hardcoded)" << endl;
	printAandB(m, n);
	cout << endl;
	cout << endl;

	// string s = "";
	// s = getStr(m, n);
	// cout << s << endl;

	string s = "";
	s = getStr(i, j);
	cout << s << endl;

	return 0;
}
```
