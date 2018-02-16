Language Learning Challenge #1
==============================

Objectives
----------
- Set up a development environment for the new language
- Learn basic language syntax and data structures
- Create a working executable

Instructions
------------
Complete as many of the challenges below as you can.  If you don't have time to do all of them, that's okay.  Our main objective in this challenge is to get acquainted with the new language. Be prepared to share your code for at least one solution and to talk about what you found interesting in the new language.  Use our Slack channel for discussion as you go along.  

### Task 1: Convert Roman Numerals to Arabic Numerals
Create an application that will convert a number from Roman numerals to Arabic numerals.
```
> Enter Roman Numerals: xxiv
> Arabic Numerals Equivalent: 24
```

### Task 2: Convert Arabic Numerals to Roman Numerals
Create an application that will convert a number from Arabic numerals into Roman numerals.

```
> Enter Number: 1971
> Roman Numeral: MCMLXXI
```

### Task 3: Create a Roman Numerals calculator
Create an application that will perform standard calculations (add, subtract, multiply, divide) using Roman numerals.

```
Enter Calculation: XXI + XXV
Answer: XLVI
```

Notes
-----
* While it is possible to express numbers above 3,999 in Roman Numerals by using bars over the letters, let's restrict our applications to numbers smaller than 3,999.
*	There is no zero in Roman numerals, so your calculator will need to find a way to express a zero when it is the solution to an equation (such as V - V = ??).   Be creative.


Resources
---------

### Roman Numerals Quick Guide

| Character | N   | I   | V   | X   | L   | C   | D   | M    |  
| --------- | --- | --- | --- | --- | --- | --- | --- | ---- |  
| Value     | 0   | 1   | 5   | 10  | 50  | 100 | 500 | 1000 |

- - -

"When you're learning a new language, there's a natural tendency to write code as you would have written it in a language you already know.  Be aware of this bias … and try to avoid it."

_From The Go Programming Language by Alan Donovan and Brian Kernighan_

- - -

In general, the number zero did not have its own Roman numeral, but the concept of zero as a number was known by medieval computists (responsible for calculating the date of Easter). They included zero (via the Latin word `nullus` meaning none) as one of nineteen epacts, or the age of the moon on March 22. The first three epacts were nullae, xi, and xxii (written in minuscule or lower case). The first known computist to use zero was Dionysius Exiguus in 525. Only one instance of a Roman numeral for zero is known. About 725, Bede or one of his colleagues used the letter `N`, the initial of `nullae`, in a table of epacts, all written in Roman numerals.

### Installation

```
$ go get github.com/kgust/go_challenge/roman_calculator
$ cd $GOPATH/github.com/kgust/go_challenge/roman_calculator
$ go test
$ go install ./...

$ roman  # To see usage and examples
```
