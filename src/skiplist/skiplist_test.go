package skiplist

import (
	"testing"
	"math/rand"
	"time"
)

var MAX_NODE_NUM = 10
type Rank struct{
	Key int32
	Val int32
}

func TestInsert(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	sl := CreateSkipList()
	r := new(Rank)
	for i:=0;i<MAX_NODE_NUM;i++{
		r = &Rank{
			int32(i),
			int32(rand.Intn(100000)),
		}
		sl.Insert(r.Key,r.Val,r)
	}
	sl.Print()
}