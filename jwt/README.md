# jwt

## Who

```
github.com/pascaldekloe/jwt v1.9.0
vs
github.com/cristalhq/jwt/v3 v3.0.1
```

## Where

```
MacBook Pro (15-inch, 2017)
2,8 GHz Quad-Core Intel Core i7
16 GB 2133 MHz LPDDR3
```

## How

```shell script
# make benchmarks as an executable
$ go test -c -o jwt-bench

# run them
$ time ./jwt-bench -test.v -test.benchmem -test.bench ^Benchmark -test.count 10 -test.run ^$ > bench.txt

# split output into files
$ sed 's/_One_//' bench.txt > one.txt
$ sed 's/_Two_//' bench.txt > two.txt

# final results
$ benchstat one.txt two.txt > result.txt
```

## Well

```
name                  old time/op    new time/op    delta
ECDSA/sign-ES256-8      27.8µs ± 3%    28.5µs ± 3%   +2.38%  (p=0.007 n=10+10)
ECDSA/sign-ES384-8      4.68ms ± 2%    4.60ms ± 1%   -1.84%  (p=0.000 n=10+10)
ECDSA/sign-ES512-8      9.10ms ±13%    8.06ms ± 1%  -11.46%  (p=0.000 n=10+10)
ECDSA/check-ES256-8     94.9µs ±13%    79.3µs ± 2%  -16.47%  (p=0.000 n=10+10)
ECDSA/check-ES384-8     9.28ms ± 1%    9.03ms ± 1%   -2.73%  (p=0.000 n=9+9)
ECDSA/check-ES512-8     16.4ms ± 2%    16.0ms ± 3%   -2.07%  (p=0.027 n=8+10)
EdDSA/sign-EdDSA-8      54.6µs ± 5%    53.9µs ± 2%     ~     (p=0.247 n=10+10)
EdDSA/check-EdDSA-8      138µs ± 3%     138µs ± 3%     ~     (p=0.853 n=10+10)
HMAC/sign-HS256-8       2.21µs ± 1%    2.02µs ± 2%   -8.65%  (p=0.000 n=10+10)
HMAC/sign-HS384-8       2.51µs ± 1%    2.19µs ± 2%  -12.95%  (p=0.000 n=10+10)
HMAC/sign-HS512-8       2.56µs ± 1%    2.23µs ± 2%  -13.00%  (p=0.000 n=10+10)
HMAC/check-HS256-8      4.44µs ± 1%    1.09µs ± 4%  -75.51%  (p=0.000 n=9+9)
HMAC/check-HS384-8      5.03µs ± 7%    1.20µs ± 4%  -76.22%  (p=0.000 n=10+9)
HMAC/check-HS512-8      5.04µs ± 2%    1.20µs ± 2%  -76.12%  (p=0.000 n=10+10)
RSA/sign-1024-bit-8      349µs ± 5%     335µs ± 2%   -3.91%  (p=0.003 n=10+10)
RSA/sign-2048-bit-8     1.54ms ± 4%    1.54ms ± 2%     ~     (p=0.739 n=10+10)
RSA/sign-4096-bit-8     8.81ms ± 5%    8.53ms ± 2%   -3.13%  (p=0.029 n=10+10)
RSA/check-1024-bit-8    29.8µs ± 5%    26.1µs ± 2%  -12.35%  (p=0.000 n=9+10)
RSA/check-2048-bit-8    66.1µs ± 2%    62.9µs ± 2%   -4.80%  (p=0.000 n=10+10)
RSA/check-4096-bit-8     176µs ± 3%     170µs ± 2%   -3.16%  (p=0.000 n=10+10)

name                  old B/token    new B/token    delta
ECDSA/sign-ES256-8         166 ± 0%       173 ± 0%   +4.22%  (p=0.000 n=10+10)
ECDSA/sign-ES384-8         208 ± 0%       216 ± 0%   +3.85%  (p=0.000 n=10+10)
ECDSA/sign-ES512-8         256 ± 0%       262 ± 0%   +2.34%  (p=0.000 n=10+10)
EdDSA/sign-EdDSA-8         166 ± 0%       172 ± 0%   +3.61%  (p=0.000 n=10+10)
HMAC/sign-HS256-8          123 ± 0%       129 ± 0%   +4.88%  (p=0.000 n=10+10)
HMAC/sign-HS384-8          144 ± 0%       150 ± 0%   +4.17%  (p=0.000 n=10+10)
HMAC/sign-HS512-8          166 ± 0%       172 ± 0%   +3.61%  (p=0.000 n=10+10)
RSA/sign-1024-bit-8        251 ± 0%       257 ± 0%   +2.39%  (p=0.000 n=10+10)
RSA/sign-2048-bit-8        422 ± 0%       428 ± 0%   +1.42%  (p=0.000 n=10+10)
RSA/sign-4096-bit-8        763 ± 0%       769 ± 0%   +0.79%  (p=0.000 n=10+10)

name                  old alloc/op   new alloc/op   delta
ECDSA/sign-ES256-8      3.04kB ± 0%    3.38kB ± 0%  +11.08%  (p=0.000 n=10+10)
ECDSA/sign-ES384-8      1.75MB ± 0%    1.75MB ± 0%     ~     (p=0.360 n=10+8)
ECDSA/sign-ES512-8      3.03MB ± 0%    3.03MB ± 0%     ~     (p=0.218 n=10+10)
ECDSA/check-ES256-8     2.54kB ± 0%    1.25kB ± 0%  -50.95%  (p=0.000 n=10+10)
ECDSA/check-ES384-8     3.49MB ± 0%    3.48MB ± 0%   -0.25%  (p=0.000 n=9+10)
ECDSA/check-ES512-8     6.12MB ± 0%    6.10MB ± 0%   -0.35%  (p=0.000 n=9+10)
EdDSA/sign-EdDSA-8        672B ± 0%      912B ± 0%  +35.71%  (p=0.000 n=10+10)
EdDSA/check-EdDSA-8     1.62kB ± 0%    0.29kB ± 0%  -82.18%  (p=0.000 n=10+10)
HMAC/sign-HS256-8         656B ± 0%      400B ± 0%  -39.02%  (p=0.000 n=10+10)
HMAC/sign-HS384-8         992B ± 0%      432B ± 0%  -56.45%  (p=0.000 n=10+10)
HMAC/sign-HS512-8       1.02kB ± 0%    0.46kB ± 0%  -54.69%  (p=0.000 n=10+10)
HMAC/check-HS256-8      1.78kB ± 0%    0.03kB ± 0%  -98.20%  (p=0.000 n=10+10)
HMAC/check-HS384-8      2.21kB ± 0%    0.05kB ± 0%  -97.83%  (p=0.000 n=10+10)
HMAC/check-HS512-8      2.22kB ± 0%    0.06kB ± 0%  -97.12%  (p=0.000 n=10+10)
RSA/sign-1024-bit-8     16.7kB ± 0%    16.9kB ± 0%   +1.56%  (p=0.000 n=10+10)
RSA/sign-2048-bit-8     31.4kB ± 0%    31.6kB ± 0%   +0.74%  (p=0.000 n=10+10)
RSA/sign-4096-bit-8     77.6kB ± 0%    77.9kB ± 0%   +0.46%  (p=0.000 n=10+8)
RSA/check-1024-bit-8    4.35kB ± 0%    2.90kB ± 0%  -33.26%  (p=0.000 n=10+10)
RSA/check-2048-bit-8    7.04kB ± 0%    5.47kB ± 0%  -22.39%  (p=0.000 n=10+10)
RSA/check-4096-bit-8    17.1kB ± 0%    15.3kB ± 0%  -10.74%  (p=0.000 n=9+9)

name                  old allocs/op  new allocs/op  delta
ECDSA/sign-ES256-8        35.0 ± 0%      42.0 ± 0%  +20.00%  (p=0.000 n=10+10)
ECDSA/sign-ES384-8       14.4k ± 0%     14.4k ± 0%     ~     (p=0.303 n=10+8)
ECDSA/sign-ES512-8       19.6k ± 0%     19.6k ± 0%     ~     (p=0.171 n=10+10)
ECDSA/check-ES256-8       46.0 ± 0%      21.0 ± 0%  -54.35%  (p=0.000 n=10+10)
ECDSA/check-ES384-8      28.9k ± 0%     28.8k ± 0%   -0.24%  (p=0.000 n=9+8)
ECDSA/check-ES512-8      39.6k ± 0%     39.5k ± 0%   -0.37%  (p=0.000 n=10+10)
EdDSA/sign-EdDSA-8        7.00 ± 0%     11.00 ± 0%  +57.14%  (p=0.000 n=10+10)
EdDSA/check-EdDSA-8       28.0 ± 0%       2.0 ± 0%  -92.86%  (p=0.000 n=10+10)
HMAC/sign-HS256-8         7.00 ± 0%      6.00 ± 0%  -14.29%  (p=0.000 n=10+10)
HMAC/sign-HS384-8         7.00 ± 0%      6.00 ± 0%  -14.29%  (p=0.000 n=10+10)
HMAC/sign-HS512-8         7.00 ± 0%      6.00 ± 0%  -14.29%  (p=0.000 n=10+10)
HMAC/check-HS256-8        31.0 ± 0%       1.0 ± 0%  -96.77%  (p=0.000 n=10+10)
HMAC/check-HS384-8        32.0 ± 0%       1.0 ± 0%  -96.88%  (p=0.000 n=10+10)
HMAC/check-HS512-8        32.0 ± 0%       1.0 ± 0%  -96.88%  (p=0.000 n=10+10)
RSA/sign-1024-bit-8        106 ± 0%       110 ± 0%   +3.77%  (p=0.000 n=10+10)
RSA/sign-2048-bit-8        112 ± 0%       116 ± 0%   +3.57%  (p=0.000 n=10+10)
RSA/sign-4096-bit-8        126 ± 0%       130 ± 0%   +3.17%  (p=0.000 n=10+8)
RSA/check-1024-bit-8      39.0 ± 0%      13.0 ± 0%  -66.67%  (p=0.000 n=10+10)
RSA/check-2048-bit-8      39.0 ± 0%      13.0 ± 0%  -66.67%  (p=0.000 n=10+10)
RSA/check-4096-bit-8      40.0 ± 0%      14.0 ± 0%  -65.00%  (p=0.000 n=10+10)
```