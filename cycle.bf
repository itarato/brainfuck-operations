runs a cyclic action
this case in every 3rd iteration there is a special action
needs 5 cells:
loop number
exit place (0)
cycle number
zero flag 1
zero flag 2

+++++++++++++++++++++++++++++++++
>
++++++++++++++++++++++++++++++++++++++++++++++ print char one
>
++++++++++++++++++++++++++++++++++++++++++ print char two

>
+++++++++++++++++++++ set loop num

> leave 0 for exit flag

>
+++ set cycle num

>>>>
+++++ set second cycle num
<<<<

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
    < back to cycle num
    < back to exit flag
  ]

  >>>>> cycle check
  [
    - decrement cycle num

    >+>+<< sets zero flags bit to 1
    [
      >[-]>[-]<< sets temp bit to 0 in case it's still not 0
      < go to exit this is why we need 2 zero
    ]

    >> check if zero flag is on
    [
      <<<<<<<<<<<.>>>>>>>>>>> cyclic action

      <<+++++>> restore cycle number
      [-]<[-] zero the flags
    ]
    < back to cycle num
    < back to exit flag
  ]
  <<<<

  <-
]
