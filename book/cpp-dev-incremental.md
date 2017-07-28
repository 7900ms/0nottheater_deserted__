
迭代式开发，不用于展示 仅自己看得到

(正式作业 而非展示型作业)

```
incremental
partial
lib
```

```c++
// === header.cpp ===
// g++ header.cpp -o header
// ./header
// g++ -std=c++11 header.cpp -o header
// ./header

#include <iostream>
using namespace std;

int main(){






    // <% yield %>
    // cout << "a" << endl;








    return 0;
}


```

```c++
// === var-var.cpp ===
#include <iostream>
using namespace std;

// 5.3
// string getGreeting53(string name, int age){
//     string s = "";
//     s = "hi, " + name + ", I know you are " + to_string(age) + " years old. ";
//     return s;
// }

// 5.4
string getGreeting(string name, int age);

// 5.5
string getGreetingStr(string name, string age);

int main(){

    // 1

    // int i = 1;
    // int j = 2;

    // cout << i << j << endl;



    // 2

    // int i = 1;
    // int j = 2;

    // cout << "hi" << endl;
    // cout << "we got " << i << " and " << j << ". " << endl;



    // 3

    // int i = 1;
    // int j = 2;

    // cout << "hi," << endl;
    // cout << "how are you?" << endl;
    // cout << "we got " << i << " and " << j << ". " << endl;



    // 4

    // int i = 1;
    // int j = 2;

    // cout << "hi, ";
    // cout << "how are you?" << endl;
    // cout << "we got " << i << " and " << j << ". " << endl;



    // 5.1

    // string name = "Andrew";
    // int age = 12;

    // cout << "hi, " << name << endl;
    // cout << "I know you are " << age << " years old." << endl;



    // 5.2

    // string name = "Andrew";
    // int age = 12;
    // string greeting = "hi, " + name + ", I know you are " + to_string(age) + " years old. ";
    // greeting += "GDay!";
    // cout << greeting << endl;
    // 注意 实际上只要是为了打印（而非计算），直接用 string age
    // 只要是为了打印 全string (name, age, address, score)
    // string getGreetingStr(string name, string age);



    // 5.3

    // string greeting = "";
    // greeting = getGreeting53("Sandro", 14);
    // cout << greeting << endl;



    // 5.4
    // getGreeting(declared)

    // string greeting1, greeting2, greeting3;
    // greeting1 = getGreeting("Sandro", 14);
    // greeting2 = getGreeting("Andrew", 14);
    // greeting3 = getGreeting("Mathew", 18);
    // cout << greeting1 << endl;
    // cout << greeting2 << endl;
    // cout << greeting3 << endl;

    // cout << endl;
    // cout << endl;

    // cout << greeting1 << greeting2 << greeting3 << endl;



    // 5.4.1
    // getGreeting, asking

    // string name = "";
    // int age = 0;
    // cout << "Good morning, ";
    // cout << "please enter your name: " << endl;
    // cin >> name;
    // cout << "please enter your age: " << endl;
    // cin >> age;

    // cout << "processing ..." << endl;
    // cout << getGreeting(name, age) << endl;



    // 5.4.2
    // getGreeting, asking, getline

    string name = "";
    int age = 0;
    cout << "Good morning, ";
    cout << "please enter your name: " << endl;
    getline(cin, name);
    cout << "please enter your age: " << endl;
    cin >> age; // 不用  getline(cin, age) 因为 getline(cin, xxx) 只能赋值给string类型的变量
    cout << "processing ..." << endl;
    cout << getGreeting(name, age) << endl;



    // 5.5
    // getGreetingStr, asking, getline

    // string name = "";
    // string age = "";
    // cout << "Good morning, ";
    // cout << "please enter your name: " << endl;
    // getline(cin, name);
    // cout << "please enter your age: " << endl;
    // getline(cin, age);
    // cout << "processing ..." << endl;
    // cout << getGreetingStr(name, age) << endl;






    return 0;
}

// 5.4
string getGreeting(string name, int age){
    string s = "";
    s = "hi, " + name + ", I know you are " + to_string(age) + " years old. ";
    return s;
}

// 5.5
string getGreetingStr(string name, string age){
    string s = "";
    s = "hi, " + name + ", I know you are " + age + " years old. ";
    return s;
}


```




-
