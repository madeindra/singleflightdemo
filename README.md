# Singleflight Workshop

This repo is a demo project from Tokopedia Tech Workshop on How to Handle Million of Requests Using Singleflight Mechanism.

## Singleflight Mechanism

## What is singleflight mechanism?

Singleflight mechanism is a mechanism that limit simultaneous calls to a function so it will only be called once.

This can only be done for simultaneous calls to an identical function that have identical result.

## When should I use singleflight mechanism?

If there is a possibility that the function called more than once at a time.

## Let's take an example!

You have a function for checking today's weather, this function work by calling third-party API.

If there are 5 simultaneous call to this function, you will need to call third-party API for 5 times, even though the API have a same result.

What if you can call the API just 1 time and give the result to the former 5 simultaneous call?

This is where singleflight mechanism comes into play.

## Dependencies

- [Golang Singleflight module](https://pkg.go.dev/golang.org/x/sync/singleflight).

## How to Run

There are 2 ways to run the project.

### Run without Build

```
go run main.go
```

### Run with Build

```
go build -o main
./main
```

P.S. This project will run on localhost:8080 by default. 

## Endpoints

There are two endpoints on this project.

1. Normal
```
http://localhost:8080/normal
```

2. Singleflight
```
http://localhost:8080/singleflight
```

## Benchmark

To benchmark this project, we will use **Apache Benchmark**, [click here for tutorial on how install it on Linux](https://www.tutorialspoint.com/apache_bench/index.htm).

P.S. Apache Benchmark is already installed on MacOS by default.

1. Run the project

2. Run Apache Benchmark on terminal

```
ab -n 100 -c 100  -t 10 http://localhost:8080/singleflight
```