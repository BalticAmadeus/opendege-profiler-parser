block-level on error undo, throw.

class Test1 inherits Main:
    method public override void test():
        pause 1.

        super:test().
    end method.
end.