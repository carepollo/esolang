# ESOlang

This is a programming language made for fun as well to practice more the GO programming language. I call it Esoteric programming language but it actually **is not**, is pretty much like functional javascript with support for first class functions. Currectly there is no variables, only constants.

### Features
The language currently supports
- _string_ data type
- _int_ data type
- _functions_ data type (including first class)
- _array_ data structure and method to pop, push and get length
- _hashmap_ data structure
- conditionals
- closures

### Installation
Prerequisites:
- Having GO version 1.19 already installed
- Having git already installed.

Run the following commands
- `git clone https://github.com/carepollo/esolang.git`
- `cd esolang`
- `go build main.go`
- This step depends on your OS. On **Linux** or **MacOS** should be: `./esolang`
- On **Windows** maybe is going through folder and open the .exe

### How to use

##### Hello world
```js
puts("Hello World!");
```

##### Variables
```js
let message = "Hello World!";
let number = 1;
```

##### Arrays
```js
let array = [1, 2, 3];
array[0] = 2;
push(array, 1); // add value
rest(array); // elements starting by index 1
len(array); // amount of elements
```
##### HashMaps
```js
let hashmap = {"a": 1, "b": 2};
hashmap["a"] = 2;
```
#### Functions
```js
let myfunc = fn(a, b){ a + b }
let afunc = fn(){ return myfunc }
afunc()()
```
Supports implicit and explicit return
#### Conditionals
```js
if (true) {
    puts(1);
}
if (false) {
    puts("nothing");
} else {
    puts("something");
}
```


### Sidenote

In case you wonder why is called 'esolang' to begin with, is because it all started from an idea that was discard afterwards, it was merely shitpost so it not worth mentioning.
