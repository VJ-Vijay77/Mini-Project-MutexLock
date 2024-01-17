### Mini Project: Concurrent Task Scheduler with Mutex

#### Problem Statement:

Create a program that manages a list of tasks, each represented by a function. Multiple Goroutines will concurrently execute these tasks, and a mutex will be used to safely update a shared result map.

#### Requirements:

1. Define a list of tasks, each represented by a function. The tasks can perform various operations.
2. Create a shared result map where each task updates a specific key with its result.
3. Use a mutex to synchronize access to the shared result map.
4. Implement a Goroutine for each task to execute concurrently.
5. Wait for all Goroutines to finish and print the final results.

#### Tips:

* Utilize the `sync.Mutex` to protect the shared result map.
* Experiment with different types of tasks and their impact on concurrent execution.

This mini project allows you to practice working with Goroutines, mutexes, and concurrent data access in a more realistic scenario. You can experiment by adding more tasks or modifying existing tasks to see how the concurrent execution behaves.
