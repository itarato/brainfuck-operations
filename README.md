Brainfuck operations
====================

Brainfuck basic operations, such ass accumulation, substraction, multiplication, cyclic actions.

#Example: cyclic actions

(See the implementation file for comments.)

```brainfuck
++++++++++++++++++++++++++++++++++++++++++++++>
++++++++++++++++++++++++++++++++++++++++++>
+++++++++++++++++++++>>
+++<<
[
  <<.>>
  >>
  [
    -
    >+>+<<
    [>[-]>[-]<<<]
    >>
    [
      <<<<<.>>>>>
      <<+++>>
      [-]<[-]
    ]
    <<
  ]
  <-
]
```
