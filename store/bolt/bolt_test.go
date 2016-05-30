package bolt

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math/rand"
	"testing"
	"time"

	"github.com/boltdb/bolt"
	"github.com/jubalh/dropf/store"
	"github.com/pborman/uuid"
)

var bs *Bolt

var f = store.File{
	Metadata: store.MetaData{
		Expires: time.Now().Add(2 * time.Hour),
		Public:  true,
	},
}

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func TestMain(m *testing.M) {

	f, err := ioutil.TempFile("", "dropf-bolt")

	if err != nil {
		panic("Can not create Tempfile: " + err.Error())
	}
	f.Close()

	// defer os.Remove(f.Name())

	db, err := bolt.Open(f.Name(), 0600, nil)

	if err != nil {
		panic("Can not create bolt database: " + err.Error())
	}
	defer db.Close()

	bs, err = NewBoltFileStore(db)

	if err != nil {
		panic(err)
	}

	m.Run()
	db.Sync()

}

func TestBoltStore(t *testing.T) {

	buf := genRandomBytes(1048576)

	f.Content = bytes.NewBuffer(buf)
	f.Metadata.Name = "test.txt"

	id, err := bs.Store(&f)

	if err != nil {
		t.Fatal("Could not store file", err)
	}

	fn, err := bs.Get(id)

	if err != nil {
		t.Fatal("Could not retrieve file", err)
	}

	res, err := ioutil.ReadAll(fn.Content)

	if err != nil {
		t.Fatal("Could not read file content", err)
	}

	if fn.Metadata.Expires != f.Metadata.Expires {
		t.Log("Expiry dates do not match")
		t.Fail()
	}

	if fn.Metadata.Public != f.Metadata.Public {
		t.Log("Public flags do not match")
		t.Fail()
	}

	if string(res) != string(buf) {
		t.Log("Content of original and retrieved not identical")
		t.FailNow()
	}
}

func benchStore(b *testing.B, size int) {
	buf := genRandomBytes(size)

	f.Content = bytes.NewBuffer(buf)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		f.Metadata.Name = fmt.Sprintf("%s.txt", uuid.NewRandom().String())
		bs.Store(&f)
	}

}

func Benchmark1024(b *testing.B)     { benchStore(b, 1024) }
func Benchmark2048(b *testing.B)     { benchStore(b, 2048) }
func Benchmark4096(b *testing.B)     { benchStore(b, 4096) }
func Benchmark8192(b *testing.B)     { benchStore(b, 8192) }
func Benchmark16384(b *testing.B)    { benchStore(b, 16384) }
func Benchmark32768(b *testing.B)    { benchStore(b, 32768) }
func Benchmark65536(b *testing.B)    { benchStore(b, 65536) }
func Benchmark131072(b *testing.B)   { benchStore(b, 131072) }
func Benchmark262144(b *testing.B)   { benchStore(b, 262144) }
func Benchmark524288(b *testing.B)   { benchStore(b, 524288) }
func Benchmark1048576(b *testing.B)  { benchStore(b, 1048576) }
func Benchmark2097152(b *testing.B)  { benchStore(b, 2097152) }
func Benchmark4194304(b *testing.B)  { benchStore(b, 4194304) }
func Benchmark8388608(b *testing.B)  { benchStore(b, 8388608) }
func Benchmark16777216(b *testing.B) { benchStore(b, 16777216) }
func Benchmark33554432(b *testing.B) { benchStore(b, 33554432) }
func Benchmark67108864(b *testing.B) { benchStore(b, 67108864) }

func genRandomBytes(size int) (b []byte) {
	b = make([]byte, size, size)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return
}
