def var i as int.

def temp-table ttTest no-undo
    field field1 as char.

do i = 1 to 50000:
    create ttTest.
    ttTest.field1 = string(i).
end.

do i = 1 to 5:
    run test1.p(i, table ttTest).
    pause 1.
end.
