# BoinkLang  
BoinkLang is a lightweight, interpreted programming language designed as a learning project to explore how interpreters function and to gain hands-on experience with Go. Or in simpler terms, I wanted to create a language where I could make the computer suffer through my poorly written code. By building BoinkLang, I aim to deepen my understanding of language processing, tokenization, and interpreter architecture while also reinforcing my Go skills—because debugging in Go wasn’t painful enough already.

## Why Choose Go?  

Go was chosen as the implementation language for BoinkLang due to its simplicity, performance, and ability to make developers question their life choices. Here’s why:  

- **Ease of Use** – Go’s clean syntax makes it a great choice for interpreters, ensuring that the real complexity comes from my own coding mistakes.  
- **Strong Standard Library** – Go provides robust built-in packages, saving me from reinventing the wheel (although I probably will, just for fun).  
- **Garbage Collection** – Automatic memory management means I can focus on making logic errors instead of memory leaks.  
- **Efficient Performance** – Fast execution speed ensures that my terrible programs fail instantly, instead of making me wait.  
- **Static Typing** – Helps catch errors at compile-time, but unfortunately, it doesn’t prevent bad life decisions.  

## Features  
### Implemented:  
- **Tokenizer (Lexer)** – Converts BoinkLang code into tokens faster than you can say “syntax error.”  
- **Parser** – Constructs an Abstract Syntax Tree (AST) from tokens using Top-Down Operator Precedence (Vaughan Pratt Parsing). Now my code at least pretends to be structured.  
- **Function Calls** – Yes, BoinkLang now supports functions, which means you can write spaghetti code with slightly more organization.  
- **Control Flow** – Implements if-else conditions and comparison operators, ensuring that logic errors are now officially my fault.  
- **REPL Mode (Revamped!)** – Upon startup, BoinkLang now politely asks if you want to:  
  - Run **R-Lex-PL** (just lexing, because why not?)  
  - Run **R-Parse-PL** (lexing + parsing, for when you feel fancy).  
  - I haven’t added a full evaluator yet, so execution is still a distant dream.  

### Planned Features:  
- **Evaluator** – To actually execute parsed expressions and statements (because what’s the point of writing code if it doesn’t do anything?).  
- **Variable Storage (Symbol Table)** – While the AST currently holds parsed expressions, a proper symbol table will be needed for managing scope and variable resolution.  
- **Loops** – Implementing while and for loops so BoinkLang can suffer through infinite iterations.  

## Installation & Usage  
### Prerequisites:  
- Go (latest version recommended, but hey, live dangerously).  

### Clone the Repository:  
```sh
git clone https://github.com/TheGamingGod85/BoinkLang.git
cd BoinkLang
```

### Build & Run:  
#### Starts BoinkLang with a choice of REPL mode  
```sh
go run main.go
```

Upon starting, BoinkLang will now let you choose between lexing or parsing before misunderstanding your commands spectacularly.

## Sample Code  
Here’s some BoinkLang code. If it breaks, that’s a feature, not a bug.  
```boink
let five = 5;
let ten = 10;
let add = fn(x, y) {
    x + y;
};
let result = add(five, ten);

!-/*5;  // Because why not?
5 < 10 > 5;  // Just to confuse people.

if (5 < 10) {
    return true;  // Of course, 5 is less than 10. Duh.
} else {
    return false;  // This should never happen. If it does, blame floating-point precision.
}

10 == 10;  // As expected.
10 != 9;   // Big surprise.
```

## References  
- **Writing an Interpreter in Go** by Thorsten Ball – A great resource for understanding interpreter development. Also known as *The Book That Made Me Do This.*  

<!--  
## Contributing  
If you’d like to contribute, feel free to fork this repository and submit a pull request. Feedback and improvements are always welcome! But no promises I’ll accept them—BoinkLang is already weird enough.  
-->

## License  
This project is licensed under the MIT License – meaning you are free to use and modify it, but I take zero responsibility if it makes your computer cry.

