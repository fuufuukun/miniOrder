package util

import "fmt"

const (
	MAXUINT32              = 4294967295
	DEFAULT_UUID_CNT_CACHE = 512
)

type UUIDGenerator struct {
	Prefix       string
	idGen        uint32
	internalChan chan uint32
}

func NewUUIDGenerator(prefix string) *UUIDGenerator {
	gen := &UUIDGenerator{
		Prefix:       prefix,
		idGen:        0,
		internalChan: make(chan uint32, DEFAULT_UUID_CNT_CACHE),
	}

	gen.startGen()
	return gen
}

func (u *UUIDGenerator) startGen() {
	go func() {
		for {
			if u.idGen == MAXUINT32 {
				u.idGen = 1
			} else {
				u.idGen += 1
			}
			u.internalChan <- u.idGen
		}
	}()
}

func (u *UUIDGenerator) Get() string {
	idgen := <-u.internalChan
	return fmt.Sprintf("%s%d", u.Prefix, idgen)
}

func (u *UUIDGenerator) GetUint32() uint32 {
	return <-u.internalChan
}
