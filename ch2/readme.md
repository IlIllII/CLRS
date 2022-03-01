# Chapter 2

This folder contains tests for all the algorithms mentioned in chapter 2. They are written in Go.

I have stubbed out functions for all the algorithms in `algorithms.go`. This file is where you will write your implementations of the algorithms from the chapter. I have written tests for them in `algorithms_test.go`.

To run the tests, navigate to this directory in your terminal and run the following:

```bash
go test
```

This will run all of the tests in `algorithms_test.go`. Many of the tests will fail until you have provided implementions of the algorithms. If you want to run a single test instead of all the tests, you can with the following command:

```bash
go test -run XXX
```

Where XXX is the name of the function you want to test.

You can look at the tests in `algorithms_test.go` to learn more about what the input and output look like for each function. Feel free to add new test cases if you run into edge cases that are not covered by the test suite.