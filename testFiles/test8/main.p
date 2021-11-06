def input param i as int.

def var i2 as int.

def temp-table ttTest no-undo
    field field1 as char.

do i2 = 1 to 50000:
    create ttTest.
    ttTest.field1 = string(i2).
end.

i = i + 1.

if i = 5
then return.

run test1.p(i, input table ttTest).
