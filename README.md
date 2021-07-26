# Daily Coding Problem: Problem #746 [Easy]

Oddly, this is also _Daily Coding Problem: Problem #794 [Easy]_,
and _Daily Coding Problem: Problem #43_,
and it's one of the problems in the [Daily Coding Problem book](https://www.amazon.com/Daily-Coding-Problem-exceptionally-interviews/dp/1793296634/ref=sr_1_3?dchild=1&keywords=daily+coding+problem&qid=1627321879&sr=8-3).

This problem was asked by Amazon.

Implement a stack that has the following methods:

* `push(val)`, which pushes an element onto the stack
* `pop()`, which pops off and returns the topmost element of the stack.
If there are no elements in the stack,
then it should throw an error or return null.
* `max()`, which returns the maximum value in the stack currently.
If there are no elements in the stack,
then it should throw an error or return null.

Each method should run in constant time.

## Analysis

The specification is an abstract stack.
It doesn't specify the implementation,
which could be a fixed-size array, a linked list,
or a variably-sized array, like in Python or a Go slice.
I chose to do a linked list implementation.

The problem statement wants a `max()` method for the abstract stack.
That's where some implementation complications arise.
The problem statement doesn't specify a type for the values
that reside on the stack.
I chose integers because I didn't think about it too hard,
but strings have a lexicographic ordering: "aaa" < "aaaa" < "z"
for example.

The "run in constant time" constraint points at a particular
implementation for the `max()` method.
Pushing and popping elements from the head of a linked list
is the dictionary definition of "constant time".
Why not add a second linked list to the max-stack-type
whose values are the maximum value of the values pushed on the
abstract stack so far?
This doubles the time it takes to push or pop an element,
but that's still constant no matter how large the stack gets.

That's what I did: specify a max-stack type that contains
two different, internal stacks.
One stack keeps track of the values passed in via `push(value)`,
and returned via `pop()`,
the second internal stack keeps track of the maximum value of
the values in the first stack.

[Code](max1.go)

This is the Daily Coding Problem book's solution,
albeit they use Python's variable-length arrays instead of a linked list
for the two stacks.

## Interview Analysis
