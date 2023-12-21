# Atlas1.0: First Documentation Draft

## Introduction keys:
- JIT Interpreters.
- Tree-Walking Interpreters(walks the AST and interpreters it).

* Atlas1.0 list of features:
 - C-like Syntax.
 - Variable Bindings
 - Integeres and Booleans
 - Arithmetic Expressions
 - Built-in Functions
 - First-class and higher-order functions
 - Closures
 - A String data structure
 - An Array data structure
 - A Hash data structure

* Binding Variables and Data Structure exemples:

```javascript
    let num = 1;
    let str = "Atlas";
    let expr = 25 * (1999/12);
    let arr = [69, 420];
    
    //hash with key value
    let hsh = {"name": "Atlas", "Version": 1.0}

    //accessing elements
    arr[0] // => 69
    hsh["name"] // => Atlas

    //binding functions(Explicit and Implicit returns) to variable
    let addVal = add(a, b) { return a+b; };
    let addVal = add(a, b) { a+b; };

    //higher order functions(funcs that take a func as parameter and applies it on elements):
    let twice = fn(f, x) {
        return (f(f(x)));
    }

    let addTwo = fn(a) {
        return a + 2;
    }

    twice(add(2)); // => 6

```
