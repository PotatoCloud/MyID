package MyID

import (
	"crypto/rand"
	"math/big"
	"os"
	"time"
)

const (
	SaltBit          uint  = 7
	SaltShift        uint  = 4
	IncrShift        uint  = SaltBit + SaltShift
	DefaultIncrValue int64 = 100
	MaxUint8         int64 = 1<<8 - 1
)

var IdService *Metadata

type Metadata struct {
	Incr int64
}

func New(incr int64) *Metadata {
	return &Metadata{Incr: incr}
}

func (m *Metadata) Generate() int64 {
	m.Incr++
	a, b := wrapRandI64(MaxUint8), wrapRandI64(MaxUint8)
	return (m.Incr << IncrShift) | (a << SaltShift) | b
}

func wrapRandI64(max int64) int64 {
	w, _ := rand.Int(rand.Reader, big.NewInt(max-1))
	return w.Int64()
}

func asyncWrite(interval time.Duration) {
	var (
		data []byte
		err  error
	)
	for {
		time.Sleep(interval)
		data = IntToBytes(IdService.Incr)
		if err = SmartWrite("id_data", data); err != nil {
			panic(err)
		}
	}
}

func init() {
	var incr int64 = DefaultIncrValue
	if IsExist("id_data") {
		data, err := os.ReadFile("id_data")
		if err != nil {
			panic(err)
		}
		incr = BytesToInt(data)
	}
	IdService = New(incr)
	go asyncWrite(time.Minute * 10)
}
