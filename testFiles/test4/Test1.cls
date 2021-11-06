block-level on error undo, throw.

class Test1:
    def var i as int init 0.

    method public void test():
        test(i).
    end method.

    method private void test(i as int):
        if i = 5
        then return.

        i = i + 1.

        test(i).
    end method.
end.