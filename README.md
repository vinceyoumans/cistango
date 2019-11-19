# CISCO Code Challenge


## assignment one

### Write a go code using goroutine and channels to send numbers sequentially from 1 to 10 and print.
I wanted to demonstrate a couple of ways to do this.

#### ass01 
is an unbuffered channel, which, I think is what is requested in assignment. 

#### ass02
Version 2, Is using GoRoutines, but Not using Channels, Instead, I am using waitgroups.  

#### ass03
Version 3, is completely not what your asking for but thought I would do it anyway.  Demonstrating multiple channels and a SELECT statement.  The results are completely random.

If I have time, I will also do a mutex example...
NOTE:   I have been doing Lambdas for last couple of months.  They tend to be very short lived by design and avoid anything that blocks.  I dont think we used goRoutines or channels as we had a couple strategies that are more predictable.  1...We had a memcache DB, 2...passed values as parameters.. and 3... the DynamoDB was super tuned to get values back.  I don't think I saw a single goRoutine or channel at capitalOne, which is why I had to think about it.

#### ass04
Version 4
is also a goRoutine example, but here I use Pointers to pass parameters by reference.
I also include a main_test.go

``` run test ```
part of the test fails intentionally


## assignment 

### Write a go code to create a struct type with 3 fields date, time and zone and print the values in a json format.

#### ass20
In this assignment I was not sure if objective was to show Structures or Understanding of Time package. If Doing Structs, then I thought the concept of a slice of structs was also important.  Thus, in Ass20 I went off range a bit to demonstrate everything.

1. struct
2. Slice
3. time package

###  TODO
1. I will be adding some Docker and Docker-Compose examples to the above.
2. Add Unit testing

