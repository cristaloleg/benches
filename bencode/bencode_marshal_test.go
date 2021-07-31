package bencode

import (
	"bytes"
	"testing"

	bencode2 "github.com/IncSW/go-bencode"
	bencode8 "github.com/anacrolix/torrent/bencode"
	bencode7 "github.com/chihaya/chihaya/frontend/http/bencode"
	bencode1 "github.com/cristalhq/bencode"
	bencode16 "github.com/cuberat/go-bencode"
	bencode11 "github.com/ehmry/go-bencode"
	bencode6 "github.com/jackpal/bencode-go"
	bencode15 "github.com/lajide/bencode"
	bencode13 "github.com/lwch/bencode"
	bencode5 "github.com/marksamman/bencode"
	bencode4 "github.com/nabilanam/bencode/encoder"
	bencode9 "github.com/owenliang/dht"
	bencode12 "github.com/stints/bencode"
	bencode10 "github.com/tumdum/bencoding"
	bencode3 "github.com/zeebo/bencode"
)

var marshalBenchData = map[string]interface{}{
	"announce": ("udp://tracker.publicbt.com:80/announce"),
	"announce-list": []interface{}{
		[]interface{}{("udp://tracker.publicbt.com:80/announce")},
		[]interface{}{[]byte("udp://tracker.openbittorrent.com:80/announce")},
		[]interface{}{
			"udp://tracker.openbittorrent.com:80/announce",
			"udp://tracker.openbittorrent.com:80/announce",
		},
	},
	"comment": []byte("Debian CD from cdimage.debian.org"),
	"info": map[string]interface{}{
		"name":         []byte("debian-8.8.0-arm64-netinst.iso"),
		"length":       170917888,
		"piece length": 262144,
	},
}

func Benchmark_cristalhq_Marshal(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		_, err := bencode1.Marshal(marshalBenchData)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_cristalhq_MarshalTo(b *testing.B) {
	dst := make([]byte, 0, 1<<12)
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		_, err := bencode1.MarshalTo(dst, marshalBenchData)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_IncSW_Marshal(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		_, err := bencode2.Marshal(marshalBenchData)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_IncSW_MarshalTo(b *testing.B) {
	dst := make([]byte, 0, 1<<12)
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		_, err := bencode2.MarshalTo(dst, marshalBenchData)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_Zeebo_Marshal(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		_, err := bencode3.EncodeBytes(marshalBenchData)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_Marksamman_Marshal(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		err := bencode5.Encode(marshalBenchData)
		if err == nil {
			b.Fatal("got nil")
		}
	}
}

func Benchmark_Anacrolix_Marshal(b *testing.B) {
	w := bytes.NewBuffer(make([]byte, 0, 1<<12))
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		err := bencode8.NewEncoder(w).Encode(marshalBenchData)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_Nabilanam_Marshal(b *testing.B) {
	b.Skip()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		res := bencode4.New(marshalBenchData).Encode()
		if res != "" && b.N%3 == 66666 {
			b.Fatal("empty")
		}
	}
}

func Benchmark_Jackpal_Marshal(b *testing.B) {
	b.Skip()
	w := bytes.NewBuffer(make([]byte, 0, 1<<12))
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		err := bencode6.Marshal(w, marshalBenchData)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_Chihaya_Marshal(b *testing.B) {
	b.Skip()
	w := bytes.NewBuffer(make([]byte, 0, 1<<12))
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		err := bencode7.NewEncoder(w).Encode(marshalBenchData)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_Owenliang_Marshal(b *testing.B) {
	b.Skip()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		_, err := bencode9.Encode(marshalBenchData)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_Tumdum_Marshal(b *testing.B) {
	b.Skip()
	w := bytes.NewBuffer(make([]byte, 0, 1<<12))
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		err := bencode10.NewEncoder(w).Encode(marshalBenchData)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_Ehmry_Marshal(b *testing.B) {
	b.Skip()
	w := bytes.NewBuffer(make([]byte, 0, 1<<12))
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		err := bencode11.NewEncoder(w).Encode(marshalBenchData)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_Stints_Marshal(b *testing.B) {
	b.Skip()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		d := bencode12.NewEncoder().Encode(marshalBenchData)
		if d == nil {
			b.Fatal("is nil")
		}
	}
}

func Benchmark_Lwch_Marshal(b *testing.B) {
	b.Skip()
	w := bytes.NewBuffer(make([]byte, 0, 1<<12))
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		err := bencode13.NewEncoder(w).Encode(marshalBenchData)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_Lajide_Marshal(b *testing.B) {
	b.Skip()
	buf := bytes.NewBuffer(make([]byte, 0, 1<<12))
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		err := bencode15.NewEncoder(buf).Encode(marshalBenchData)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_Cuberat_Marshal(b *testing.B) {
	b.Skip()
	buf := bytes.NewBuffer(make([]byte, 0, 1<<12))
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		err := bencode16.NewEncoder(buf).Encode(marshalBenchData)
		if err != nil {
			b.Fatal(err)
		}
	}
}
