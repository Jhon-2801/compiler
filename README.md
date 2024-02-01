
# Teenytinycompiler

Simple compiler implemented in Golang for learning basic compilation concepts. It supports a limited set of features and is designed to be an educational project.

## It supports:

 - Numerical variables
 - Basic arithmetic
 - If statements
 - While loops
 - Print text and numbers
 - Input numbers
 - Labels and goto
 - Comments


## Usage/Examples

``` 
PRINT "How many fibonacci numbers do you want?"
INPUT nums
PRINT ""

LET a = 0
LET b = 1
WHILE nums > 0 REPEAT
    PRINT a
    LET c = a + b
    LET a = b
    LET b = c
    LET nums = nums - 1
ENDWHILE
```

