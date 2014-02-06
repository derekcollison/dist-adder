Example for Gopher Academy Article on Distributed Systems
==========

## The Plumbing and Semantics: Communication Patterns in Distributed Systems

A simple example to show some distributed systems patterns for optimizing speed of responses by asking multiple responders to answer and only selecting the first, and fastest response.

Our example will have a simple adder  which can add multiple integers together. We could add more adders to present different “shards”, but I will leave that as an exercise for the user. We will replicate the adder classes and identify them with IDs that they generate on startup. Requests will be sent to all the replicas, and we will then answer the requests with both the answer and the Id of the responders so we can see who "wins".

# Install and Build
``` bash
git clone https://github.com/derekcollison/dist-adder
go get -d
go build
```

# Examples

``` bash
> ./dist-adder

Spinning up 10 responders.

Adder [5aa6ab3f] is ready
Adder [292ceae2] is ready
Adder [815ea650] is ready
Adder [6f9b9644] is ready
Adder [bff14500] is ready
Adder [bc73559e] is ready
Adder [043bd0aa] is ready
Adder [c6b15c38] is ready
Adder [2a84ec07] is ready
Adder [32fa2147] is ready

Sending 10 requests.

Request: {X:3 Y:48}	    Response: {Ans:51 Id:043bd0aa}
Request: {X:85 Y:89}	Response: {Ans:174 Id:043bd0aa}
Request: {X:55 Y:48}	Response: {Ans:103 Id:bff14500}
Request: {X:30 Y:65}	Response: {Ans:95 Id:6f9b9644}
Request: {X:5 Y:74}	    Response: {Ans:79 Id:043bd0aa}
Request: {X:99 Y:85}    Response: {Ans:184 Id:292ceae2}
Request: {X:96 Y:87}	Response: {Ans:183 Id:292ceae2}
Request: {X:47 Y:73}	Response: {Ans:120 Id:292ceae2}
Request: {X:68 Y:21}	Response: {Ans:89 Id:bff14500}
Request: {X:69 Y:50}	Response: {Ans:119 Id:5aa6ab3f}

> ./dist-adder -numRequests=10 -numResponders=30

Spinning up 30 responders.

Adder [f14c932c] is ready
Adder [a297343f] is ready
Adder [8fc56bbf] is ready
Adder [d30eae6e] is ready
Adder [d4225ee8] is ready
Adder [d4bffeff] is ready
Adder [c1ab2c68] is ready
Adder [21c9913d] is ready
Adder [65751e77] is ready
Adder [441900e0] is ready
Adder [785dc6c2] is ready
Adder [b22b1865] is ready
Adder [167905ce] is ready
Adder [69c94bc2] is ready
Adder [94261692] is ready
Adder [28d398bf] is ready
Adder [4eedcf8f] is ready
Adder [9de124c4] is ready
Adder [1dfec38e] is ready
Adder [08ffcadc] is ready
Adder [0ecf7aa1] is ready
Adder [f861dc2d] is ready
Adder [4a508dab] is ready
Adder [01bb4d08] is ready
Adder [ce6ea17e] is ready
Adder [84c8ec07] is ready
Adder [f6161c4d] is ready
Adder [4b7c5183] is ready
Adder [e042b2fc] is ready
Adder [01ea5556] is ready

Sending 10 requests.

Request: {X:90 Y:86}	Response: {Ans:176 Id:9de124c4}
Request: {X:70 Y:43}	Response: {Ans:113 Id:441900e0}
Request: {X:33 Y:70}	Response: {Ans:103 Id:65751e77}
Request: {X:0 Y:26}	    Response: {Ans:26 Id:441900e0}
Request: {X:10 Y:34}	Response: {Ans:44 Id:785dc6c2}
Request: {X:53 Y:68}	Response: {Ans:121 Id:65751e77}
Request: {X:93 Y:47}	Response: {Ans:140 Id:65751e77}
Request: {X:67 Y:3}	    Response: {Ans:70 Id:f861dc2d}
Request: {X:44 Y:85}	Response: {Ans:129 Id:ce6ea17e}
Request: {X:66 Y:39}	Response: {Ans:105 Id:28d398bf}

Finished

```

