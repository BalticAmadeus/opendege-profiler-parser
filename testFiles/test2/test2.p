def temp-table ttTest
    field field1 as char
    field field2 as char
    field field3 as char
    field field4 as char
    field field5 as char.

def input param table for ttTest.

for each ttTest:
    ttTest.field1 = subst("&1,&1", ttTest.field1).
end.
