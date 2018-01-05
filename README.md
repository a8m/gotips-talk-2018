## "Did you know that..." talk.

0. [Introduction - "feel the audience"](https://github.com/a8m/gotips-talk-2018/blob/master/intro/1.md).

1. Make the nil receiver useful - [Example](https://github.com/a8m/gotips-talk-2018/tree/master/code/1).

2. `copy` - The most "unpopular" built-in function, after `complex`/`imag`/`real` - [Examples](https://github.com/a8m/gotips-talk-2018/tree/master/code/2).

3. [Anonymous structs](https://github.com/a8m/gotips-talk-2018/blob/master/code/3/3.md), [anonymous interfaces](https://github.com/a8m/gotips-talk-2018/blob/master/code/3/interface.go), and [interface validation](https://github.com/a8m/gotips-talk-2018/blob/master/code/3/impl.go).

4. Typed function that implements an interface (show only if more than 5 people in the
   audience don't know that) - [Example](https://github.com/a8m/gotips-talk-2018/blob/master/code/4/4.go).

5. Control parallelism. limit the number of goroutines - [Example](https://github.com/a8m/gotips-talk-2018/tree/master/code/5).

6. Relative imports - [Example](https://github.com/a8m/gotips-talk-2018/tree/master/code/6).

7. Shadowing. run vet -shadow - [Example](https://github.com/a8m/gotips-talk-2018/blob/master/code/7/7.go).

8. close channel to notify many. but be careful from the gotcha
   (closed channel always true in select) - [Example](https://github.com/a8m/gotips-talk-2018/blob/master/code/8/8.go).

9. Panic-Recover flow control example - [Example](https://github.com/a8m/gotips-talk-2018/blob/master/code/9/9.go).

10. Embed interface. sort.Reverse example - [Example](https://github.com/a8m/gotips-talk-2018/blob/master/code/10/10.go).

11. Few words about 1.10 release.  
    - strings.Builder - why, and what it solves.
    - File.SetDeadline
    - json.IgnoreUnknownFields
    - And more..

Suggestions for next meetups:
- Some tips to organize your table driven tests.
- Go schedular more cooperative than preemptive.
- Explain the `reflect` package.
- Compiler optimization.
