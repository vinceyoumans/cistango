# GoRoutine with Pointers

I thought an example of using goRoutines with pointers would also be an intersting example.
Passing parameters by reference.  I also thought an example of a unit test would be interesting.
It also demonstrates a JSON file, which is what you were asking for in the second challenge.

This code example shows reading a BASKET with n Items. In this case 3 Items.

The App reads the basket, then runs Go Routine on Each Item in Parallel. I am using pointers to the original Basket. Passing the values to be modified by reference.

When the WG.Done() is execute, the code Prints the modified Basket.

I also demonstrate MAXPROCS. In this case, I set MAXPROCS to 100, which is over kill. I could have set it to 4 and it should work fine. 1 for main process. 3 for each of the Items.

If I have time, I may also Docerfy this.



