def var i as int.
def var a as int.

a = 10.

do i = 1 to a:
    if i modulo 2 = 0
    then run test1.p(i).
    else run test2.p(i).
end.
