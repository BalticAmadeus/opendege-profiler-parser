def temp-table ttTest
    field field1 as char
    field field2 as char
    field field3 as char
    field field4 as char
    field field5 as char.

def input-output param table for ttTest.

def var i as int.

do i = 1 to 30000:
    create ttTest.
    ttTest.field1 = string(i).
    ttTest.field2 = string(i).
    ttTest.field3 = string(i).
    ttTest.field4 = string(i).
    ttTest.field5 = string(i).
end.
