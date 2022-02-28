# early bird gets the worm

* 2 services provide the same service (say we have to get the stock price)
* make a call to both simultaneously; whichever returns first - the result is used; the other request should be asked to stop
* Test setup:
    * 2 simultaneous calls - one 100ms; the other 1s
    * Expected result: result from first job is used; the other job is cancelled and finishes within 100ms