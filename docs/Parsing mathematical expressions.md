# Parsing mathematical expressions

Possible inputs

> f(x) = x^2 + 2x -4

and

> x^2 + 2x -4


## EBNF grammar

<equation> ::= <function_definition> <equals_sign> <term> | <term>
<function_definition> ::= <identifier> <opening_parenthesis> <variable_list> <closing_parenthesis>
<variable_list> ::= <identifier> | <identifier> "," <variable_list>+

<term> ::= <number> | <identifier> | <term> <operator> <term> | <term> <term> | <opening_parenthesis> <term> <closing_parenthesis>
<operator> ::= "+" | "-" | "*" | "/" | "^"

<identifier> ::= <alpha>+
<number> ::= <sign>? <non_zero_digit> <digit>* | <sign>? <non_zero_digit> <digit>* "." <digit>+ | <sign>? <non_zero_digit> <digit>* "," <digit>+

<alpha> ::= "a" | "b" | "c" | "d" | "e" | "f" | "g" | "h" | "i" | "j" | "k" | "l" | "m" | "o" | "p" | "q" | "r" | "s" | "t" | "u" | "v" | "w" | "x" | "y" | "z" | "A" | "B" | "C" | "D" | "E" | "F" | "G" | "H" | "I" | "J" | "K" | "L" | "M" | "O" | "P" | "Q" | "R" | "S" | "T" | "U" | "V" | "W" | "X" | "Y" | "Z"
<sign> ::= "+" | "-"
<digit> ::= "0" | <non_zero_digit>
<non_zero_digit> ::= "1" | "2" | "3" | "4" | "5" | "6" | "7" | "8" | "9"

<equals_sign> ::= "="
<opening_parenthesis> ::= "("
<closing_parenthesis> ::= ")"
