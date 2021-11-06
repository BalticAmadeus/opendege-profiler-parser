block-level on error undo, throw.

class Main:
    method public void test1():
    end method.

    method public void test2():
        test1().
    end method.

    method public void test3():
        test2().
        test1().
    end method.

    method public void test4():
        test3().
        test2().
        test1().
    end method.

    method public void test5():
        test4().
        test3().
        test2().
        test1().
    end method.
end.