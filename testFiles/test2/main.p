def temp-table ttTest
    field field1 as char
    field field2 as char
    field field3 as char
    field field4 as char
    field field5 as char.

def var i as int.

do i = 1 to 5:
    run test1.p (input-output table ttTest).
end.

run test2.p(input table ttTest).
