package bencode

import (
	"bytes"
	"io"
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
	bencode4 "github.com/nabilanam/bencode/decoder"
	bencode9 "github.com/owenliang/dht"
	bencode10 "github.com/tumdum/bencoding"
	bencode3 "github.com/zeebo/bencode"
)

var unmarshalBenchData = []byte("d4:infod6:lengthi170917888e12:piece lengthi262144e4:name30:debian-8.8.0-arm64-netinst.isoe8:announce38:udp://tracker.publicbt.com:80/announce13:announce-listll38:udp://tracker.publicbt.com:80/announceel44:udp://tracker.openbittorrent.com:80/announceee7:comment33:Debian CD from cdimage.debian.orge")

func Benchmark_cristalhq_Unmarshal(b *testing.B) {
	var res interface{}
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		err := bencode1.NewDecodeBytes(unmarshalBenchData).Decode(&res)
		if err != nil {
			b.Fatal(err)
		}
		if res == nil {
			b.Fatal("is nil")
		}
	}
}

func Benchmark_cristalhq_UnmarshalReader(b *testing.B) {
	var res interface{}
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		r := bytes.NewReader(unmarshalBenchData)
		err := bencode1.NewDecoder(r).Decode(&res)
		if err != nil {
			b.Fatal(err)
		}
		if res == nil {
			b.Fatal("is nil")
		}
	}
}

func Benchmark_IncSW_Unmarshal(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		_, err := bencode2.Unmarshal(unmarshalBenchData)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_ZeeboBencode_Unmarshal(b *testing.B) {
	res := map[string]interface{}{}
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		err := bencode3.DecodeBytes(unmarshalBenchData, &res)
		if err != nil {
			b.Fatal(err)
		}
		if res == nil {
			b.Fatal("is nil")
		}
	}
}

func Benchmark_NabilanamBencode_Unmarshal(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		res := bencode4.New(unmarshalBenchData).Decode()
		if res == nil {
			b.Fatal("is nil")
		}
	}
}

func Benchmark_MarksammanBencode_Unmarshal(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		r := bytes.NewReader(unmarshalBenchData)
		_, err := bencode5.Decode(r)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_JackpalBencode_Unmarshal(b *testing.B) {
	b.Skip()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		r := bytes.NewReader(unmarshalBenchData)
		_, err := bencode6.Decode(r)
		if err != nil && err != io.EOF {
			b.Fatal(err)
		}
	}
}

func Benchmark_ChihayaBencode_Unmarshal(b *testing.B) {
	b.Skip()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		bencode7.Unmarshal(unmarshalBenchData)
	}
}

func Benchmark_AnacrolixTorrent_Unmarshal(b *testing.B) {
	b.Skip()
	res := map[string]interface{}{}
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		err := bencode8.Unmarshal(unmarshalBenchData, &res)
		if err != nil {
			b.Fatal(err)
		}
		if res == nil {
			b.Fatal("is nil")
		}
	}
}

func Benchmark_OwenliangDht_Unmarshal(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		res, err := bencode9.Decode(unmarshalBenchData)
		if err != nil {
			b.Fatal(err)
		}
		if res == nil {
			b.Fatal("is nil")
		}
	}
}

func Benchmark_TumdumBencoding_Unmarshal(b *testing.B) {
	b.Skip()
	res := map[string]interface{}{}
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		err := bencode10.Unmarshal(unmarshalBenchData, &res)
		if err != nil {
			b.Fatal(err)
		}
		if res == nil {
			b.Fatal("is nil")
		}
	}
}

func Benchmark_EhmryGoBencode_Unmarshal(b *testing.B) {
	b.Skip()
	res := map[string]interface{}{}
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		err := bencode11.Unmarshal(unmarshalBenchData, &res)
		if err != nil {
			b.Fatal(err)
		}
		if res == nil {
			b.Fatal("is nil")
		}
	}
}

func Benchmark_StintsBencode_Unmarshal(b *testing.B) {
	b.Skip()
}

func Benchmark_LwchBencode_Unmarshal(b *testing.B) {
	b.Skip()
	res := map[string]interface{}{}
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		err := bencode13.Decode(unmarshalBenchData, &res)
		if err != nil {
			b.Fatal(err)
		}
		if res == nil {
			b.Fatal("is nil")
		}
	}
}

func Benchmark_ClearcodecnBencode_Unmarshal(b *testing.B) {
	b.Skip()
}

func Benchmark_LajideBencode_Unmarshal(b *testing.B) {
	b.Skip()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		buf := bytes.NewBuffer(unmarshalBenchData)
		res, err := bencode15.NewDecoder(buf).Decode()
		if err != nil && err != io.EOF {
			b.Fatal(err)
		}
		if res == nil {
			// b.Fatal("is nil")
		}
	}
}

func Benchmark_CuberatGoBencode_Unmarshal(b *testing.B) {
	b.Skip()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		buf := bytes.NewBuffer(unmarshalBenchData)
		res, err := bencode16.NewDecoder(buf).Decode()
		if err != nil && err != io.EOF {
			b.Fatal(err)
		}
		if res == nil {
			// b.Fatal("is nil")
		}
	}
}
