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

## How-to

```shell script
# make benchmarks as an executable
$ go test -c -o jwt-bench

# run them
$ time ./jwt-bench -test.v -test.bench ^Benchmark -test.count 10 -test.run ^$ > bench.txt

# split output into files
$ sed 's/_One_//' bench.txt > one.txt
$ sed 's/_Two_//' bench.txt > two.txt

# final results
$ benchstat one.txt two.txt > result.txt
```

## Result

```
name                  old time/op  new time/op  delta
ECDSA/sign-ES256-8    28.4µs ± 3%  28.3µs ± 1%     ~     (p=0.912 n=10+10)
ECDSA/sign-ES384-8    4.57ms ± 1%  4.59ms ± 1%   +0.45%  (p=0.035 n=10+10)
ECDSA/sign-ES512-8    7.99ms ± 1%  8.01ms ± 1%     ~     (p=0.280 n=10+10)
ECDSA/check-ES256-8   83.1µs ± 2%  78.2µs ± 1%   -5.91%  (p=0.000 n=9+8)
ECDSA/check-ES384-8   8.99ms ± 1%  8.95ms ± 1%     ~     (p=0.052 n=10+10)
ECDSA/check-ES512-8   15.7ms ± 1%  15.8ms ± 1%     ~     (p=0.353 n=10+10)
EdDSA/sign-EdDSA-8    52.7µs ± 1%  53.5µs ± 2%   +1.36%  (p=0.006 n=10+9)
EdDSA/check-EdDSA-8    141µs ± 1%   134µs ± 1%   -5.28%  (p=0.000 n=9+10)
HMAC/sign-HS256-8     2.19µs ± 2%  2.00µs ± 1%   -8.62%  (p=0.000 n=10+10)
HMAC/sign-HS384-8     2.51µs ± 0%  2.18µs ± 2%  -13.35%  (p=0.000 n=7+10)
HMAC/sign-HS512-8     2.56µs ± 1%  2.22µs ± 2%  -13.34%  (p=0.000 n=10+10)
HMAC/check-HS256-8    4.42µs ± 1%  1.08µs ± 3%  -75.67%  (p=0.000 n=10+10)
HMAC/check-HS384-8    4.87µs ± 1%  1.18µs ± 2%  -75.79%  (p=0.000 n=9+9)
HMAC/check-HS512-8    4.94µs ± 0%  1.20µs ± 2%  -75.66%  (p=0.000 n=9+9)
RSA/sign-1024-bit-8    340µs ± 5%   334µs ± 2%     ~     (p=0.315 n=10+8)
RSA/sign-2048-bit-8   1.53ms ± 1%  1.53ms ± 2%     ~     (p=0.739 n=10+10)
RSA/sign-4096-bit-8   8.49ms ± 1%  8.54ms ± 2%     ~     (p=0.143 n=10+10)
RSA/check-1024-bit-8  29.4µs ± 1%  26.2µs ± 3%  -10.75%  (p=0.000 n=10+10)
RSA/check-2048-bit-8  66.1µs ± 1%  68.2µs ±18%     ~     (p=0.481 n=10+10)
RSA/check-4096-bit-8   174µs ± 1%   190µs ± 5%   +9.42%  (p=0.000 n=8+10)

name                  old B/token  new B/token  delta
ECDSA/sign-ES256-8       166 ± 0%     173 ± 0%   +4.22%  (p=0.000 n=10+10)
ECDSA/sign-ES384-8       208 ± 0%     216 ± 0%   +3.85%  (p=0.000 n=10+10)
ECDSA/sign-ES512-8       256 ± 0%     262 ± 0%   +2.34%  (p=0.000 n=10+10)
EdDSA/sign-EdDSA-8       166 ± 0%     172 ± 0%   +3.61%  (p=0.000 n=10+10)
HMAC/sign-HS256-8        123 ± 0%     129 ± 0%   +4.88%  (p=0.000 n=10+10)
HMAC/sign-HS384-8        144 ± 0%     150 ± 0%   +4.17%  (p=0.000 n=10+10)
HMAC/sign-HS512-8        166 ± 0%     172 ± 0%   +3.61%  (p=0.000 n=10+10)
RSA/sign-1024-bit-8      251 ± 0%     257 ± 0%   +2.39%  (p=0.000 n=10+10)
RSA/sign-2048-bit-8      422 ± 0%     428 ± 0%   +1.42%  (p=0.000 n=10+10)
RSA/sign-4096-bit-8      763 ± 0%     769 ± 0%   +0.79%  (p=0.000 n=10+10)

```