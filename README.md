# What ?

This projet can transpile the custom (fr*nch) programming language used at the IUT Aix-Marseille, into C++.

This project works as a CLI (available on Linux/Windows/MacOS), and as an online website available at https://thomasrubini.github.io/algo-iut

# How to build
- Install the Go language toolchain
- Run `go build` in the root of this repository
- A new binary `algo-iut` (or `algo-iut.exe` on Windows) will have been created.

# FAQ
## Non-acceptd syntaxes
We aren't sure if these syntaxes are valid or not (ask Casali)

### Declaration without type
e.g. `declarer a <- rand(1, 2);`

### `pour` loop as a C `while`
e.g. `pour (j < taille(voyelle)-1)`

# Resources/references
https://craftinginterpreters.com  
https://github.com/flouksac/DOC-ALGO-PAPIER/blob/main/syntaxe/tout_en_un.md  
