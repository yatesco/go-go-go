It's a bizarre combination of a high-level lang and lower level C :-)

I'm sorely feeling the absence of Rust's shortcut error handling, Option monad, Error monad, traits etc.

Encapsulation feels hard to achieve, particularly a clean separation between "plain old structs" and their JSON representation. Making everything public feels whacky, but seems to be idiomatic.

Having unenforceable constructors is also pretty whack. (Unenforced because members must be `P`ublic if they need to be JSONified)

I don't know how to do:
 - Dependency Injection
 - System construction
 - macros to separate out enforcement of an API with the implementation being tested for example
 - web stuff - the tutorial talks about Gin
