# data pipeline - elevate the bottleneck

## Use-case
* This use-case is about task decomposition, fanning-out-fanning-in and sharing data between concurrent jobs.
* Imagine a task where:
    * we need to do a bunch of things that can be done independently of each other (eg. query the database for some information)
    * then we need to put it altogether synchronously (in one thread)
* So, the use-case is:
    * 100 independent tasks that take 1s with a standard deviation of 300ms.
    * The result of each task gets assembled in one thread; and assembling each piece takes 10ms