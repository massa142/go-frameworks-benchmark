# go-frameworks-benchmark

Benchmark for golang web frameworks.

## Date
Jan 11, 2017

## Results

### Bench

```
BenchmarkBoneParam-4       	 1000000	      1363 ns/op	    1088 B/op	       9 allocs/op
BenchmarkChiParam-4        	 2000000	       962 ns/op	     752 B/op	       8 allocs/op
BenchmarkEchoParam-4       	 2000000	       845 ns/op	     496 B/op	       7 allocs/op
BenchmarkGinParam-4        	 2000000	       619 ns/op	     416 B/op	       3 allocs/op
BenchmarkGocraftParam-4    	 1000000	      1457 ns/op	    1064 B/op	      12 allocs/op
BenchmarkGojiParam-4       	 1000000	      1112 ns/op	     768 B/op	       6 allocs/op
BenchmarkGorillaParam-4    	  500000	      2596 ns/op	    1432 B/op	      15 allocs/op
BenchmarkHttpParam-4       	 1000000	      1546 ns/op	     848 B/op	       7 allocs/op
BenchmarkKamiParam-4       	 1000000	      1459 ns/op	    1072 B/op	       9 allocs/op
BenchmarkMartiniParam-4    	  200000	      6962 ns/op	    1728 B/op	      22 allocs/op
BenchmarkSiestaParam-4     	  300000	      4141 ns/op	    2041 B/op	      24 allocs/op
```

### Memory

```
bone: 3584 Bytes
chi: 1120 Bytes
echo: 2576 Bytes
gin: 712 Bytes
gocraft: 1616 Bytes
goji: 40216 Bytes
gorilla: 5464 Bytes
http: 624 Bytes
kami: 1936 Bytes
martini: 82032 Bytes
siesta: 680 Bytes
```

## env

go version: go1.7.4 darwin/amd64

MacBook Pro, 2.7 GHz Intel Core i5, 8 GB 1867 MHz DDR3

## How to run

```
go test -bench .
```
