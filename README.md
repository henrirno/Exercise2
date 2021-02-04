Exercise 2: Bottlenecks
=======================

1: Some Theory
--------------

Edit in your answers to [the theory questions in this file](/theory.md). You may have to come back to the last two questions after you have done the programming part of this exercise.

2: Shared variable
------------------

In this part you will solve [the concurrent access problem from Exercise 1](https://github.com/TTK4145/Exercise1/#4-finally-some-code), such that the final result is always zero. You can choose to either use [the provided starter code](./2%20-%20shared%20variable), or to copy your solution to Exercise 1 into the suitable directories.

In your solution, make sure that the two threads intermingle. Running them one after the other would somewhat defeat the purpose. It may be useful to change the number of iterations in one of the threads, such that the expected final result is not zero (say, -1). This way it is easier to see that your solution actually works, and isn't just printing the initial value.


### C

 - POSIX has both mutexes ([`pthread_mutex_t`](http://pubs.opengroup.org/onlinepubs/7990989775/xsh/pthread.h.html)) and semaphores ([`sem_t`](http://pubs.opengroup.org/onlinepubs/7990989775/xsh/semaphore.h.html)). Which one should you use? Add a comment (anywhere in the C file) explaining why your choice is the correct one.
 - Acquire the lock, do your work in the critical section, and release the lock.


### Go

Using shared variable synchronization is possible, but not the idiomatic approach in Go. You should instead create a server that is responsible for its own data, [`select{}`](http://golang.org/ref/spec#Select_statements)s messages, and perform different actions on its data when it receives a corresponding message. 

The server should have three actions it can perform: Increment, decrement, and read (or "get"). Two other goroutines should send the increment and decrement requests to the server, and main should read out the final value after these two goroutines are done.

Before attempting to do the exercise, it is reccomended to have a look at the following chapters of the interactive go tutorial:
- [Goroutines](https://tour.golang.org/concurrency/1)
- [Channels](https://tour.golang.org/concurrency/2)
- [Select](https://tour.golang.org/concurrency/5)


Remember from Exercise 1 where we had no good way of waiting for a goroutine to finish? Try sending a "finished"/"worker done" message from the workers back to main on a separate channel. If you use different channels for the two threads, you will have to use `select { /*case...*/ }` so that it doesn't matter what order they arrive in, but you could also have multiple senders on the same channel.

3: Bounded buffer
-----------------

From the previous part, it may appear that message passing requires a lot more code to do the same work - so naturally, in this part the opposite will be the case. In the folder [3 - bounded buffer](./3%20-%20bounded%20buffer) you will find the starting point for a *bounded buffer* problem.

The bounded buffer should work as follows:
 - The `push` functionality should put one data item into the buffer - unless it is full, in which case it should block (think "pause", "wait") until room becomes available.
 - The `pop` functionality should return one data item, and block until one becomes available if necessary.

### C

The actual buffer part is already provided (as a ring buffer, see `ringbuf.c` if you are interested, but you do not have to edit - or even look at - this file), and your task is to use semaphores and mutexes to complete the synchronization required to make this work with multiple threads. If you run it as-is, it should crash when the consumer tries to read from an empty buffer.

(If you need a C compiler online, [try this](https://repl.it/@klasbo/ScientificArcticInstance#main.c))

The expected behavior (dependent on timing from the sleeps, so may not be completely consistent):
```
[producer]: pushing 0
[producer]: pushing 1
[producer]: pushing 2
[producer]: pushing 3
[producer]: pushing 4
[producer]: pushing 5
[consumer]: 0
[consumer]: 1
[producer]: pushing 6
[consumer]: 2
[consumer]: 3
[producer]: pushing 7
[consumer]: 4
[consumer]: 5
[producer]: pushing 8
[consumer]: 6
[consumer]: 7
[producer]: pushing 9
   -- program terminates here(-ish) --
```


### Go

Read [the documentation for `make`](https://golang.org/pkg/builtin/#make) carefully. Hint: making a bounded buffer is one line of code. 

Your task: Make a bounded buffer that can hold 5 elements, and use it in the producer and consumer.

The program will deadlock at the end (main is waiting forever - as it should, and the consumer is waiting for a channel no one is sending on). Since this is a toy example, don't worry about it. Really, it's fine.


4: Choosing a language
----------------------

In Exercise 3 & 4 (Network exercises) and the project, you will be using a language of your own choice. You are of course free to change your mind at any time, but to help avoid this situation (and all its associated costs) it is worth doing some research already now.

Here are a few things you should consider:
 - Think about how want to move data around (reading buttons, network, setting motor & lights, state machines, etc). Do you think in a shared-variable way or a message-passing way? Will you be using concurrency at all?
 - How will you split into modules? Functions, objects, threads? Think about what modules you need, and how they need to interact. This is an iterative design process that will take you many tries to get "right".
 - The networking part is often difficult. Can you find anything useful in the standard libraries, or other libraries?
 - While working on new sections on the project you'll want to avoid introducing bugs to the parts that already work properly. Does the language have a framework for making and running tests, or can you create one? Testing multithreaded code is especially difficult.
 - Code analysis/debugging/IDE support?

Create a new file with some reflections on language choice (somewhere in the 100-300 words range (-ish)), and push it to GitHub. We don't expect a full design, just some preliminary thoughts and ideas, all subject to change as you learn more.



NaN: Multithreading in other languages
--------------------------------------

This is an optional exercise. You are not recommended to do this for "completion" or "achievement points". You should only do it if you're interested in learning more about how different languages can protect against data races, or you're considering to use one of these languages in your project. The languages chosen is not a complete list, but rather some examples that might be useful in the context of this course.

### Erlang
Erlang disallows mutability of variables completely, a new state will instead be reached by calling into a different function (or the same function with different arguments). This means it will be impossible to solve the task from Part2 with lock based synchronization. Instead the "go-channel" approach needs to be taken, and a server needs to be created. 

These servers are so common in Erlang that they have been made an [OTP design pattern](http://erlang.org/doc/design_principles/gen_server_concepts.html). To not obfuscate the Erlang code, this approach has not been taken in the starter code. Instead a program very similar to the Go solution has been made. Complete the program and verify that the answer is 0.


### Rust
Rust uses its static type system to make sure that no data races are possible. This is possible by using the [marker traits](https://doc.rust-lang.org/std/marker/) [`Send` and `sync`](https://doc.rust-lang.org/beta/nomicon/send-and-sync.html). A data type is `Send` if you're allowed to send the data to another thread. If a data type is not marked as `Send` it's statically (i.e. at compile-time) guaranteed that it will never be sent between threads. A data type is `Sync` if it's safe to share between threads (safe to access concurrently).

The primitive integer types in rust are not "thread safe" and thus not `Sync`, but there is no reason they can't be sent between threads, so `Send` is implemented. Since Rust doesn't take a stance in which concurrency model to use (as long as you avoid undefined behavior) both the "channel" and "lock" solutions are possible. 

A [`Mutex`](https://doc.rust-lang.org/std/sync/struct.Mutex.html) takes something that is `Send` and makes it `Sync`, while [`mpsc`](https://doc.rust-lang.org/std/sync/mpsc/index.html) allows you to create "channels" for data types that implement `Send`.

The lock based approach has been taken in the starter code, you are of course free to re-write it into the `mpsc` approach if you feel like it.





