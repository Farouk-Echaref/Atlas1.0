
![368685243_1983228972027832_5089135154179161756_n](https://github.com/Farouk-Echaref/Atlas1.0/assets/66710845/51156fca-05b6-4c16-ba05-dfe9f57255f2)

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
* Functions in Atlas are just values,like integers or strings. That feature is called “first class functions”.
* Interpreter will consists of these major parts:
 - Lexer.
 - Parser.
 - Abstract Syntax Tree (AST).
 - Internal Object System.
 - Evaluator.

## Lexing(Tokenizer):
### 1.1 Lexical Analysis:

- Source Code ==> Tokens ==> Abstract Syntax Tree

```javascript
// code like this:
let x = 69 + 420;

//Atlas doesn't evaluate whitespaces as an entity itself
//code can be like:

let         a =          13   +        4       ;
```

```go
//will be tokenized like this
[
LET,
IDENTIFIER("x"),
EQUAL_SIGN,
INTEGER(5),
PLUS_SIGN,
INTEGER(5),
SEMICOLON
]

```

* Lexer can also attach line number, column number and filename to token for future error code message:

```bash
error: expected semicolon token. line 42, column 23, program.ats
```

### 1.2 Defining our tokens:

* Subset of the Atlas code we are going to scan to define our first set of tokens:

```javascript
let five = 5;
let ten = 10;
let add = fn(x, y) {
    x + y;
};
let result = add(five, ten);
```

* types of tokens that we have:
- numbers(**TYPE**): 5, 10 (we gonna treat them as such, numbers)
- variable names(**IDENTIFIERS**): five, ten, add, result (functions are just variables)
- special expressions(part of the language => **KEYWORDS**): let, fn
- special characters: ( ) , { } ;
