block-level on error undo, throw.

class Test2:
    method public void test():
        def var i as int init 0.
        
        test(i).
    end method.

    method public void test(i as int):
        if i = 5
        then return.

        i = i + 1.
        
        test(i).
    end method.
end.