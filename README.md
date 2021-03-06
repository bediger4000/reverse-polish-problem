# Daily Coding Problem: Problem #481 [Hard] 

This problem was asked by Jane Street.

Given an arithmetic expression in Reverse Polish Notation, write a program to
evaluate it.

The expression is given as a list of numbers and operands. For example:

	[5, 3, '+'] should return 5 + 3 = 8.

For example,

    [15, 7, 1, 1, '+', '-', '/', 3, '*', 2, 1, 1, '+', '+', '-']

should return 5, since it is equivalent to

    ((15 / (7 - (1 + 1))) * 3) - (2 + (1 + 1)) = 5.

You can assume the given expression is always valid.

## Building and Running

This is a command line Go program.
You need to have a Go compiler installed.

```sh
$ go build revpolish.go
$ ./revpolish 5 3 -
2
2
```

The program contains 2 different evaluation functions,
and thus should output the same number (the answer) twice.

The program uses `*` to denote the multiplication operator.
You'll need to escape it on the Linux command line,
otherwise your shell will expand it,
creating invalid input for the program.

```sh
$ ./revpolish 5 3 \*
15
15
```

I did a minimalistic program, too:

```sh
$ go build simple.go
$ ./simple 5 3 \*
Push 5 on stack
Push 3 on stack
stack 2 deep, [5 3]
Push 15 on stack
Answer: 15
Stack remaining: [15]
$
```

This minimalistic program is only 42 lines of code,
including printfs that might help debug or explain its operations.
This program would probably be more typical of what a job candidate
would produce in an interview.

## Analysis

Anybody who's ever written a program in [PostScript](https://en.wikipedia.org/wiki/PostScript)
or used an [older Hewlett-Packard calculator](https://www.hpmuseum.org/rpn.htm)
won't have any trouble with this problem,
supposing they can recognize the in-order inputs as analogous
to the PostScript program,
or keyboard entries on a calculator.

I believe the program needs a stack (first-in, first-out) on which to push input
numbers and intermediate results,
so there's a sub-algorithm for an interview candidate to grapple with,
even if that candidate already understands RPN.
If a program doesn't have an explicit stack type,
it's possible to use the inputs as an implicit stack,
modifying them in-place.
I wrote a second evaluation function that does this.

This seems like a decent interview question for a variety of skill levels.
Someone who's written PostScript, or used an HP RPN calculator,
can concentrate on the implementation details.
Someone who has to grapple with learning RPN can at least verbally walk through
what needs to happen to evaluate an expression.

The only thing hard about this is whether or not you've worked with anything RPN in the past.
This keeps up the Daily Coding Problem tradition of mis-rating most problems.
