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

## Analysis

Anybody who's ever written a program in [PostScript](https://en.wikipedia.org/wiki/PostScript)
or use an older Hewlett-Packard calculator
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
