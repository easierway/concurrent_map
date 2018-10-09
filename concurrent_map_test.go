package concurrent_map

import (
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestBasicOPs(t *testing.T) {
	for i := 0; i < 10; i++ {
		testV := rand.Intn(1000)
		m := CreateConcurrentMap(99)
		v, ok := m.Get(StrKey("Hello"))
		if v != nil || ok != false {
			t.Error("init/get failed")
		}
		m.Set(StrKey("Hello"), testV)
		v, ok = m.Get(StrKey("Hello"))
		if v.(int) != testV || ok != true {
			t.Error("set/get failed.")
		}
		m.Del(StrKey("Hello"))
		v, ok = m.Get(StrKey("Hello"))
		if v != nil || ok != false {
			t.Error("del failed")
		}
	}
}

func TestInCurrenctEnv(t *testing.T) {
	m := CreateConcurrentMap(99)
	go func() {

		for i := 0; i < 100; i++ {
			if v, ok := m.Get(StrKey(strconv.Itoa(i))); ok {
				if v != i*i {
					t.Error("Fail")
				}
			} else {
				if v != nil {
					t.Error("Fail")
				}
			}

		}

	}()
	go func() {
		for i := 0; i < 100; i++ {
			m.Set(StrKey(strconv.Itoa(i)), i*i)
		}

	}()
	time.Sleep(time.Second * 1)
}
