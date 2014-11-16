Brainfuck operations
====================

Brainfuck basic operations, such ass accumulation, substraction, multiplication, cyclic actions.

#Example: cyclic actions

```brainfuck
++++++++++++++++++++++++++++++++++++++++++++++ print char one
>
++++++++++++++++++++++++++++++++++++++++++ print char two
>
+++++++++++++++++++++ set loop num
> leave 0 for exit flag
>
+++ set cycle num
<< back to cycle
[
  <<.>> loop action
  >> cycle check
  [
    - decrement cycle num
    >+>+<< sets zero flags bit to 1
    [
      >[-]>[-]<< sets temp bit to 0 in case it's still not 0
      < go to exit this is why we need 2 zero
    ]
    >> check if zero flag is on
    [
      <<<<<.>>>>> cyclic action
      <<+++>> restore cycle number
      [-]<[-] zero the flags
    ]
    <
    <
  ]
  <-
]
```

Or simply:

```brainfuck
++++++++++++++++++++++++++++++++++++++++++++++>
++++++++++++++++++++++++++++++++++++++++++>
+++++++++++++++++++++>>+++<<
[
  <<.>>>>
  [
    ->+>+<<[>[-]>[-]<<<]>>
    [<<<<<.>>>>><<+++>>[-]<[-]]
    <<
  ]
  <-
]
```

Will print out:
```
...*...*...*...*...*...*...*
```