def var i as int.
def var testClass as class iClass no-undo.

testClass = new Test1().

do i = 1 to 2:
    testClass:test(1).
end.
