package jwt_test

import (
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"testing"
	"time"

	jwt1 "github.com/pascaldekloe/jwt"

	jwt2 "github.com/cristalhq/jwt/v3"
)

var mybenchClaims = &struct {
	jwt2.StandardClaims
}{
	StandardClaims: jwt2.StandardClaims{
		Issuer:   "benchmark",
		IssuedAt: jwt2.NewNumericDate(time.Now()),
	},
}

var benchClaims = &jwt1.Claims{
	Registered: jwt1.Registered{
		Issuer: "benchmark",
		Issued: jwt1.NewNumericTime(time.Now()),
	},
}

func Benchmark_One_ECDSA(b *testing.B) {
	tests := []struct {
		key *ecdsa.PrivateKey
		alg string
	}{
		{testKeyEC256, jwt1.ES256},
		{testKeyEC384, jwt1.ES384},
		{testKeyEC521, jwt1.ES512},
	}

	for _, test := range tests {
		b.Run("sign-"+test.alg, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, err := benchClaims.ECDSASign(test.alg, test.key)
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}

	for _, test := range tests {
		token, err := benchClaims.ECDSASign(test.alg, test.key)
		if err != nil {
			b.Fatal(err)
		}

		b.Run("check-"+test.alg, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, err := jwt1.ECDSACheck(token, &test.key.PublicKey)
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

func Benchmark_One_EdDSA(b *testing.B) {
	b.Run("sign-"+jwt1.EdDSA, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := benchClaims.EdDSASign(testKeyEd25519Private)
			if err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("check-"+jwt1.EdDSA, func(b *testing.B) {
		token, err := benchClaims.EdDSASign(testKeyEd25519Private)
		if err != nil {
			b.Fatal(err)
		}

		for i := 0; i < b.N; i++ {
			_, err := jwt1.EdDSACheck(token, testKeyEd25519Public)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func Benchmark_One_HMAC(b *testing.B) {
	// 512-bit key
	secret := make([]byte, 64)
	algs := []string{jwt1.HS256, jwt1.HS384, jwt1.HS512}

	for _, alg := range algs {
		b.Run("sign-"+alg, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, err := benchClaims.HMACSign(alg, secret)
				if err != nil {
					b.Fatal(err)
				}
			}
		})

		b.Run("sign-"+alg+"-reuse", func(b *testing.B) {
			h, err := jwt1.NewHMAC(alg, secret)
			if err != nil {
				b.Fatal(err)
			}
			for i := 0; i < b.N; i++ {
				_, err := h.Sign(benchClaims)
				if err != nil {
					b.Fatal(err)
				}
			}
		})

		token, err := benchClaims.HMACSign(alg, secret)
		if err != nil {
			b.Fatal(err)
		}

		b.Run("check-"+alg, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, err := jwt1.HMACCheck(token, secret)
				if err != nil {
					b.Fatal(err)
				}
			}
		})

		b.Run("check-"+alg+"-reuse", func(b *testing.B) {
			h, err := jwt1.NewHMAC(alg, secret)
			if err != nil {
				b.Fatal(err)
			}
			for i := 0; i < b.N; i++ {
				_, err := h.Check(token)
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

func Benchmark_One_RSA(b *testing.B) {
	keys := []*rsa.PrivateKey{testKeyRSA1024, testKeyRSA2048, testKeyRSA4096}

	for _, key := range keys {
		b.Run(fmt.Sprintf("sign-%d-bit", key.Size()*8), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, err := benchClaims.RSASign(jwt1.RS384, key)
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}

	for _, key := range keys {
		token, err := benchClaims.RSASign(jwt1.RS384, key)
		if err != nil {
			b.Fatal(err)
		}

		b.Run(fmt.Sprintf("check-%d-bit", key.Size()*8), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, err := jwt1.RSACheck(token, &key.PublicKey)
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

func Benchmark_Two_ECDSA(b *testing.B) {
	tests := []struct {
		key *ecdsa.PrivateKey
		alg jwt2.Algorithm
	}{
		{testKeyEC256, jwt2.ES256},
		{testKeyEC384, jwt2.ES384},
		{testKeyEC521, jwt2.ES512},
	}

	for _, test := range tests {
		signer, err := jwt2.NewSignerES(test.alg, test.key)
		if err != nil {
			b.Fatal(err)
		}
		bui := jwt2.NewBuilder(signer)
		b.Run("sign-"+test.alg.String(), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, err := bui.BuildBytes(mybenchClaims)
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}

	for _, test := range tests {
		signer, err := jwt2.NewSignerES(test.alg, test.key)
		if err != nil {
			b.Fatal(err)
		}
		bui := jwt2.NewBuilder(signer)
		token, err := bui.BuildBytes(mybenchClaims)
		if err != nil {
			b.Fatal(err)
		}

		verifier, err := jwt2.NewVerifierES(test.alg, &test.key.PublicKey)
		if err != nil {
			b.Fatal(err)
		}
		b.Run("check-"+test.alg.String(), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				obj, err := jwt2.ParseAndVerify(token, verifier)
				if err != nil {
					b.Fatal(err)
				}
				err = json.Unmarshal(obj.RawClaims(), new(map[string]interface{}))
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

func Benchmark_Two_EdDSA(b *testing.B) {
	signer, err := jwt2.NewSignerEdDSA(testKeyEd25519Private)
	if err != nil {
		b.Fatal(err)
	}
	bui := jwt2.NewBuilder(signer)
	b.Run("sign-"+jwt2.EdDSA.String(), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := bui.BuildBytes(mybenchClaims)
			if err != nil {
				b.Fatal(err)
			}
		}
	})

	token, err := bui.BuildBytes(mybenchClaims)
	if err != nil {
		b.Fatal(err)
	}

	verifier, err := jwt2.NewVerifierEdDSA(testKeyEd25519Public)
	if err != nil {
		b.Fatal(err)
	}
	b.Run("check-"+jwt2.EdDSA.String(), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			obj, err := jwt2.ParseAndVerify(token, verifier)
			if err != nil {
				b.Fatal(err)
			}
			err = json.Unmarshal(obj.RawClaims(), new(map[string]interface{}))
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func Benchmark_Two_HMAC(b *testing.B) {
	// 512-bit key
	secret := make([]byte, 64)
	algs := []jwt2.Algorithm{jwt2.HS256, jwt2.HS384, jwt2.HS512}

	for _, alg := range algs {
		b.Run("sign-"+alg.String()+"-reuse", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				signer, err := jwt2.NewSignerHS(alg, secret)
				if err != nil {
					b.Fatal(err)
				}
				bui := jwt2.NewBuilder(signer)
				_, err = bui.BuildBytes(mybenchClaims)
				if err != nil {
					b.Fatal(err)
				}
			}
		})
		b.Run("sign-"+alg.String()+"-reuse", func(b *testing.B) {
			signer, err := jwt2.NewSignerHS(alg, secret)
			if err != nil {
				b.Fatal(err)
			}
			bui := jwt2.NewBuilder(signer)
			for i := 0; i < b.N; i++ {
				_, err := bui.BuildBytes(mybenchClaims)
				if err != nil {
					b.Fatal(err)
				}
			}
		})

		signer, err := jwt2.NewSignerHS(alg, secret)
		if err != nil {
			b.Fatal(err)
		}
		token, err := jwt2.NewBuilder(signer).BuildBytes(mybenchClaims)
		if err != nil {
			b.Fatal(err)
		}

		b.Run("check-"+alg.String(), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				verifier, err := jwt2.NewVerifierHS(alg, secret)
				if err != nil {
					b.Fatal(err)
				}
				obj, err := jwt2.ParseAndVerify(token, verifier)
				if err != nil {
					b.Fatal(err)
				}
				err = json.Unmarshal(obj.RawClaims(), new(map[string]interface{}))
				if err != nil {
					b.Fatal(err)
				}
			}
		})
		b.Run("check-"+alg.String()+"-reuse", func(b *testing.B) {
			verifier, err := jwt2.NewVerifierHS(alg, secret)
			if err != nil {
				b.Fatal(err)
			}
			for i := 0; i < b.N; i++ {
				obj, err := jwt2.ParseAndVerify(token, verifier)
				if err != nil {
					b.Fatal(err)
				}
				err = json.Unmarshal(obj.RawClaims(), new(map[string]interface{}))
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

func Benchmark_Two_RSA(b *testing.B) {
	keys := []*rsa.PrivateKey{testKeyRSA1024, testKeyRSA2048, testKeyRSA4096}

	for _, key := range keys {
		signer, _ := jwt2.NewSignerRS(jwt2.RS384, key)
		bui := jwt2.NewBuilder(signer)
		b.Run(fmt.Sprintf("sign-%d-bit", key.Size()*8), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, err := bui.BuildBytes(mybenchClaims)
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}

	for _, key := range keys {
		signer, err := jwt2.NewSignerRS(jwt2.RS384, key)
		if err != nil {
			b.Fatal(err)
		}
		token, err := jwt2.NewBuilder(signer).BuildBytes(mybenchClaims)
		if err != nil {
			b.Fatal(err)
		}

		verifier, err := jwt2.NewVerifierRS(jwt2.RS384, &key.PublicKey)
		if err != nil {
			b.Fatal(err)
		}
		b.Run(fmt.Sprintf("check-%d-bit", key.Size()*8), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				obj, err := jwt2.ParseAndVerify(token, verifier)
				if err != nil {
					b.Fatal(err)
				}
				err = json.Unmarshal(obj.RawClaims(), new(map[string]interface{}))
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

func mustParseECKey(s string) *ecdsa.PrivateKey {
	block, _ := pem.Decode([]byte(s))
	if block == nil {
		panic("invalid PEM")
	}

	key, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	return key
}

func mustParseRSAKey(s string) *rsa.PrivateKey {
	block, _ := pem.Decode([]byte(s))
	if block == nil {
		panic("invalid PEM")
	}

	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	return key
}

var testKeyEC256 = mustParseECKey(`-----BEGIN EC PRIVATE KEY-----
MHcCAQEEIBOm12aaXvqSzysOSGV2yL/xKY3kCtaOfAPY1KQN2sTJoAoGCCqGSM49
AwEHoUQDQgAEX0iTLAcGqlWeGIRtIk0G2PRgpf/6gLxOTyMAdriP4NLRkuu+9Idt
y3qmEizRC0N81j84E213/LuqLqnsrgfyiw==
-----END EC PRIVATE KEY-----`)

var testKeyEC384 = mustParseECKey(`-----BEGIN EC PRIVATE KEY-----
MIGkAgEBBDBluSyfK9BEPc9y944ZLahd4xHRVse64iCeEC5gBQ4UM1961bsEthUC
NKXyTGTBuW2gBwYFK4EEACKhZANiAAR3Il6V61OwAnb6oYm4hQ4TVVaGQ2QGzrSi
eYGoRewNhAaZ8wfemWX4fww7yNi6AmUzWV8Su5Qq3dtN3nLpKUEaJrTvfjtowrr/
ZtU1fZxzI/agEpG2+uLFW6JNdYzp67w=
-----END EC PRIVATE KEY-----`)

var testKeyEC521 = mustParseECKey(`-----BEGIN EC PRIVATE KEY-----
MIHcAgEBBEIBH31vhkSH+x+J8C/xf/PRj81u3MCqgiaGdW1S1jcjEuikczbbX689
9ETHGCPtHEWw/Il1RAFaKMvndmfDVd/YapmgBwYFK4EEACOhgYkDgYYABAGNpBDA
Lx6rKQXWdWQR581uw9dTuV8zjmkSpLZ3k0qLHVlOqt00AfEL4NO+E7fxh4SuAZPb
RDMu2lx4lWOM2EyFvgFIyu8xlA9lEg5GKq+A7+y5r99RLughiDd52vGnudMspHEy
x6IpwXzTZR/T8TkluL3jDWtVNFxGBf/aEErnpeLfRQ==
-----END EC PRIVATE KEY-----`)

var testKeyRSA1024 = mustParseRSAKey(`-----BEGIN RSA PRIVATE KEY-----
MIICXAIBAAKBgQDCzQ4MMppUkCXTi/BjPWO2gLnaVmPhyMdo7rnccfoBnH5lCTdY
x2aK2vNkVVLi4w8zITBXAXwKB7O5iQaaXImnUD2KPReRKbyGbvkGwQGpU1UsZjzZ
uPFfbDtdWr+d2CxQUdPjKu886Lad4BsJFWSJYt06K1byYCGAYyN5hosmOQIDAQAB
AoGAO5EIYqJ2nrUVXALGlxIGk5/5NNKF6FzE3UlifA4+LI/19l9DFVqj+IHLOzr8
BXT5COF1LqW9kDOauXk1E66ISJ/vAFYvS+hIugKDqUhpBTpgPa2nyJGOjUHScvIP
sVdo1unpYU40bvhhy7HD4kwQvohYq9w5KW732jpqPJK5TKECQQD3XpZGlXAJ+O/5
p97Xwt6Rz7peG1Aqx3TlzVUvOPCXT8rnycEub0j52sYZUwg3dtf763R385pJmBJs
TJc2oN9PAkEAyZjyDqGUM6IJy7O55Ylsy3dxply7NIym+BM4p8MiEwzHZb5dXgX3
pxuPlLX3DojlGWNcLB5+gw1ZSq9Y5dz/9wJBAOQoQtUBemBIUhbj5d795sl4Xn30
FUIPy9s1Qy+WBhqZxx148gxBKn8BcRvkgLyfieDasAb/Ebx1XfCzx/jj8nMCQBNr
WT3RkL4ciMcHjAuxXjqHSfpVim74cYkKCPYYFOsy2u5RFRtehcmiHQWdNaw/wZnd
eV6CnXswSP6pv219CWcCQBv3wKhme0RkuPuyG4MUFFeHxOcilasHx/nWiz8U90Tm
hP30X1iUlekEFj/2oneT6qWqtH4nVX18/WehPQoDoLg=
-----END RSA PRIVATE KEY-----`)

var testKeyRSA2048 = mustParseRSAKey(`-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEA8NBRypbuyT1o2p7Ze94kmZTrwu2TsqZ1u7BOcY97xn6cc7/e
c9aZI+S4Ure57XNvKQAZlULWjWhEfY8vhP1m2hDzVCV0DnNRCPmMJxx212b2iTmA
1IsMmRYFHOYgVVUdx5QzS1xIQMZgyLP++CBkYJXZZCC1MBqyW93BkBcNzt0+70ZT
mMpXOYKoq/pFcxVMllKY41JCcDqpKcJnSmWyS+DQX5X4CcNecXxCMoL7WGeMVrng
7NTFJmv4Iyh19/WRERqQUqlPPQoWd0Wrw/Ih+p38PlxvdxxcGIgG8gZC1eZ441MR
4KeHEnx7nQ08TtzdsTULYlx3kM173h1yI+HuBwIDAQABAoIBAQCyX5w2E9aL+ZDR
Xxh5R/KUUFrR6Giey+4pOE7ijwV/4gjBND3yT+LfU2u02aI+4GJWXFyW0wtZcwJI
fucT+x9UJ3oVuihdC83ad/34en0M0JeMzas/xD9wpX7kCRGqI4ILcxsLly9ty4Ol
Jq6V3Gh9ooGESTXsi9nRclEOCgWQU6F8BeDGbI19aqkFi67wZqvOYrlUXfznRwQQ
iaffaeh+wH5qp79dd+MoSPJLmhuhNH6Q/T70tVqTvlslufcro1/7YYuq9X2/IO+u
O1Nd1/nyT46xYQ16HqLdH2KPN2jsmbWFCMM4leTbyk18ldDnU3LG7BMwwoW7vemE
gU8KuX4BAoGBAP4zMVT+M1421fXUTyxK1ViQprdqT5zksFwK0cMdy1upE0EFqUrr
TtN5mao+7rGFp7R/0xuVSwYs+LX7jsrRPXn5JgB2JdPg9UakdKkULAfGVCLJouXm
/32C9YlFuqPjJWxr5Ndb1aqvPNfIvsfmys1O+GJ39x9R+iFezvuKN2BvAoGBAPKE
3E9fSWjXg9N+y2QazeU6wJjJYhIGtceuTwPPW1n3IfOzgB1QHXZhH7YM07OoI2jF
NFBM99ygjdfRbKCosEQoUQCF78avHYJJDhdPhjAWiaIZg7X4gfgWqEMJ0SWXyCAM
cxQ0XEC0AHocWNipWv8zVFEC62K3omMXS/9leefpAoGAB/eGxkkpRvyk/A1pZdP6
l8oAz6LPV/V66YeVR245n2fPKKyKv8RcNhiLjmBmjr3HocqXzTeCoHDsYpe9w/GG
4bnDTSRmzxsv1MT2uw3cy2mV3XlAV8BDpaVjGKhMzzIhTCKdi3pfWfggCgtKn21G
UeT1t/BWmG6zTjRwfEW6spUCgYAUsXF69E53O6xr523DZOYcoR696rELiLcKCr2D
PbY1vviOqspLtgJNj4v9JKsLsVUUI3+LOoYLtUdlGuGB8+LWbfo7aTJEabzC2Sjy
pD526/Vid3rdlA7C9Gv3DGdkJcdVtLo9Bxq4CqPfx3ttQUYacG7JWs5q5fBdNCev
6yCzwQKBgHZRiC82Bzd10OgIL4WadlNphmMnGgROgNhwBu2bd5loPc+26omBAVtC
mQ9Ug7u6QOshlvxmqrgRFlWkLAwozqvS6RC4yru8FRqYnmtW7QgxO1pOj9VEzHSw
iugbqlkWvaTnn5JZoHZ+60PZc8Z4UJvzi0/h9ksnWhp5l6u1KBmc
-----END RSA PRIVATE KEY-----`)

var testKeyRSA4096 = mustParseRSAKey(`-----BEGIN RSA PRIVATE KEY-----
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
-----END RSA PRIVATE KEY-----`)

// example from RFC 8037, appendix A.1
var testKeyEd25519Private = ed25519.PrivateKey([]byte{
	0x9d, 0x61, 0xb1, 0x9d, 0xef, 0xfd, 0x5a, 0x60,
	0xba, 0x84, 0x4a, 0xf4, 0x92, 0xec, 0x2c, 0xc4,
	0x44, 0x49, 0xc5, 0x69, 0x7b, 0x32, 0x69, 0x19,
	0x70, 0x3b, 0xac, 0x03, 0x1c, 0xae, 0x7f, 0x60,
	// public key suffix
	0xd7, 0x5a, 0x98, 0x01, 0x82, 0xb1, 0x0a, 0xb7,
	0xd5, 0x4b, 0xfe, 0xd3, 0xc9, 0x64, 0x07, 0x3a,
	0x0e, 0xe1, 0x72, 0xf3, 0xda, 0xa6, 0x23, 0x25,
	0xaf, 0x02, 0x1a, 0x68, 0xf7, 0x07, 0x51, 0x1a,
})

// example from RFC 8037, appendix A.1
var testKeyEd25519Public = ed25519.PublicKey([]byte{
	0xd7, 0x5a, 0x98, 0x01, 0x82, 0xb1, 0x0a, 0xb7,
	0xd5, 0x4b, 0xfe, 0xd3, 0xc9, 0x64, 0x07, 0x3a,
	0x0e, 0xe1, 0x72, 0xf3, 0xda, 0xa6, 0x23, 0x25,
	0xaf, 0x02, 0x1a, 0x68, 0xf7, 0x07, 0x51, 0x1a,
})
