# bencode

## Who

```
github.com/cristalhq/bencode v0.1.1
vs
github.com/IncSW/go-bencode v0.1.0
github.com/anacrolix/torrent v1.20.0
github.com/chihaya/chihaya v2.0.0-rc.2+incompatible
github.com/cuberat/go-bencode v1.0.0
github.com/ehmry/go-bencode v1.1.1
github.com/jackpal/bencode-go v1.0.0
github.com/lajide/bencode v0.0.0-20190629152703-fb09cf4e9a4a
github.com/lwch/bencode v0.0.0-20210121090207-4404582f634e
github.com/marksamman/bencode v0.0.0-20150821143521-dc84f26e086e
github.com/nabilanam/bencode v0.0.0-20190329141800-cde546b1530b
github.com/owenliang/dht v0.0.0-20180118074908-44baeeba7b13
github.com/stints/bencode v0.0.0-20160226205624-0ba65bd80165
github.com/tumdum/bencoding v0.0.0-20160911135503-fcfcb8bd55e9
github.com/zeebo/bencode v1.0.0
```

## Where

```
MacBook Pro (16-inch, 2019)
2,6 GHz 6-Core Intel Core i7
16 GB 2667 MHz DDR4
```

## How

```shell script
# make benchmarks as an executable
$ go test -c -o bencode-bench

# run them
$ time ./bencode-bench -test.v -test.benchmem -test.bench ^Benchmark -test.count 10 -test.run ^$ > bench.txt
```
