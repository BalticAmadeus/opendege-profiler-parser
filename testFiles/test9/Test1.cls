block-level on error undo, throw.

class Test1 implements iClass inherits Main:
    method public override void test(i as int):
        pause 1.
        
        if i = 1
        then return.

        run test1.p.
        test().
    end method.
end.