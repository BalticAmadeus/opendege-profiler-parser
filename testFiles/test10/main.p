def input param a as int.
def input param i as int.

def var j as int.

pause 1.

if a = 5
then do:
    do j = 1 to 5:
        if i = 1
        then leave.

        run main.p(5, 1).
    end.
    
    return.
end.

run main.p(a + 1, 0).
