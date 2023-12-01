# Go

Go is a statically typed, compiled high-level programming language designed at Google by Robert Griesemer, Rob Pike, and Ken Thompson. It is syntactically similar to C, but also has memory safety, garbage collection, structural typing, and CSP-style concurrency. It is often referred to as Golang because of its former domain name, golang.org, but its proper name is Go.

There are two major implementations:
- Google's self-hosting "gc" compiler toolchain, targeting multiple operating systems and WebAssembly.
- gofrontend, a frontend to other compilers, with the libgo library. With GCC the combination is gccgo; with LLVM the combination is gollvm.

A third-party source-to-source compiler, GopherJS, compiles Go to JavaScript for front-end web development.

# History

Go was designed at Google in 2007 to improve programming productivity in an era of multicore, networked machines and large codebases. The designers wanted to address criticism of other languages in use at Google, but keep their useful characteristics:

- Static typing and run-time efficiency (like C)
- Readability and usability (like Python)
- High-performance networking and multiprocessing

Its designers were primarily motivated by their shared dislike of C++.

Go was publicly announced in November 2009, and version 1.0 was released in March 2012. Go is widely used in production at Google and in many other organizations and open-source projects.

# Branding and styling

The Gopher mascot was introduced in 2009 for the open source launch of the language. The design, by Renée French, borrowed from a c. 2000 WFMU promotion.

In November 2016, the Go and Go Mono fonts were released by type designers Charles Bigelow and Kris Holmes specifically for use by the Go project. Go is a humanist sans-serif resembling Lucida Grande, and Go Mono is monospaced. Both fonts adhere to the WGL4 character set and were designed to be legible with a large x-height and distinct letterforms. Both Go and Go Mono adhere to the DIN 1450 standard by having a slashed zero, lowercase l with a tail, and an uppercase I with serifs.

In April 2018, the original logo was replaced with a stylized GO slanting right with trailing streamlines. (The Gopher mascot remained the same.)

# Generics

The lack of support for generic programming in initial versions of Go drew considerable criticism. The designers expressed an openness to generic programming and noted that built-in functions were in fact type-generic, but are treated as special cases; Pike called this a weakness that might be changed at some point. The Google team built at least one compiler for an experimental Go dialect with generics, but did not release it.

In August 2018, the Go principal contributors published draft designs for generic programming and error handling and asked users to submit feedback. However, the error handling proposal was eventually abandoned.

In June 2020, a new draft design document was published that would add the necessary syntax to Go for declaring generic functions and types. A code translation tool, go2go, was provided to allow users to try the new syntax, along with a generics-enabled version of the online Go Playground.

Generics were finally added to Go in version 1.18.

# Versioning

Go 1 guarantees compatibility for the language specification and major parts of the standard library. All versions up to the current Go 1.21 release have maintained this promise.

Each major Go release is supported until there are two newer major releases.

**A big thanks to you for read this type!**
