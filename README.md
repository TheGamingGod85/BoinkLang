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
- **REPL Mode** – Allows interaction with BoinkLang by printing generated tokens in real time.  
  - Greets the user upon startup (because it’s polite).  
  - Reads commands from standard input and processes them.  

### Planned Features:  
- **Parser** – To analyze and construct the Abstract Syntax Tree (AST) from tokens (recursive descent parser).  
- **Evaluator** – To execute parsed expressions and statements (because what’s the point of writing code if it doesn’t do anything?).  
- **Variable Storage** – Implement a symbol table for storing and retrieving variables, because remembering values is too mainstream.  
- **Function Calls** – Support defining and calling functions within BoinkLang, making it *almost* look like a real language.  

## Installation & Usage  
### Prerequisites:  
- Go (latest version recommended, but hey, live dangerously).  

### Clone the Repository:  
```sh
git clone https://github.com/TheGamingGod85/BoinkLang.git
cd BoinkLang
```

### Build & Run:  
#### Currently Starts BoinkLang in REPL Mode  
```sh
go run main.go
```

Upon starting, BoinkLang will greet the user by name and allow typing in commands, which it will probably misunderstand spectacularly.

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
