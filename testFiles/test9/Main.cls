block-level on error undo, throw.

class Main implements iClass:
    method public void test():
        test1(1).
    end method.

    method public void test(i as int):
    end method.

    method public void test1(i as int):
        if i = 3
        then return.

        test1(getPar(i)).

        pause i.
    end method.

    method public int getPar(i as int):
        return i + 1.
    end method.
end.