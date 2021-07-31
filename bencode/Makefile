do: build bench

build:
	go test -c -o bencode-bench

bench:
	time ./bencode-bench -test.v -test.benchmem -test.bench ^Benchmark_cristalhq_Unmarshal -test.count 10 -test.run ^$ > bench1.txt