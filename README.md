## "Did you know that..." talk.

0. Introduction - "feel the audience".

1. Make the nil receiver useful.

2. `copy` - The most "unpopular" built-in function, after `complex`/`imag`/`real`.

3. Anonymous structs, anonymous interfaces, and interface validation.

4. Typed function that implements an interface (show only if more than 5 people in the
   audience don't know that).

5. Control parallelism. limit the number of goroutines.

6. Relative imports.

7. Shadowing. run vet -shadow.

8. close channel to notify many. but be careful from the gotcha
   (closed channel always true in select).

9. Panic-Recover flow control example.

10. Embed interface. sort.Reverse example.

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
