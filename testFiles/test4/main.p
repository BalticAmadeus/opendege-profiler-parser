def var i as int.

def var testClass as class Test1 no-undo.

testClass = new Test1().

do i = 1 to 5:
    testClass:test().
end.
