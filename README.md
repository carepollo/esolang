# ESOlang

This is a programming language made for fun as well to practice more the GO programming language. I call it Esoteric programming language but it actually **is not**, is pretty much like functional javascript with support for first class functions. Currectly there is no variables, only constants.

### Features
The language currently supports
- _string_ data type
- _int_ data type
- _functions_ data type (including high order functions)
- _array_ data structure and method to pop, push and get length
- _hashmap_ data structure

### Installation
Prerequisites:
- Having GO version 1.19 already installed
- Having git already installed.

Run the following commands
- `git clone https://github.com/carepollo/sexlang.git`
- `cd sexlang`
- `go build`
- This step depends on your OS. On **Linux** or **MacOS** should be: `./sexlang`
- On **Windows** maybe is going through folder and open the .exe

### How to use

##### Hello world
```
print("Hello World!");
```

##### Variables
```
let message = "Hello World!";
```

##### Arrays
```
let array = [1, 2, 3];
array[0] = 2;
array.push(1);
array.pop();
len(array);
```
##### HashMaps
```
let hashmap = {"a": 1, "b": 2};
hashmap["a"] = 2;
```
#### Functions
```
let myfunc = fn(a, b){ a + b }
let afunc = fn(){ return myfunc }
afunc()()
```
supports implicit and explicit return
