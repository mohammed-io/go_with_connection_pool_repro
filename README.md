# README

## Install Dependencies

`go mod download`

### Install `wrk`

`brew install wrk`

## Start Server

`./serve.sh`

## Run Benchmark

`./benchmark.sh`

## Benchmark Result for Pool of 20 connections

```
Running 10s test @ http://localhost:3040
  1 threads and 32 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.28s   431.15ms   2.00s    72.84%
    Req/Sec    44.87     53.35   190.00     86.67%
  200 requests in 10.09s, 23.63KB read
  Socket errors: connect 0, read 0, write 0, timeout 38
Requests/sec:     19.82
Transfer/sec:      2.34KB
```

