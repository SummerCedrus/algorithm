package main

import (
	"skiplist"
	"math/rand"
	"bufio"
	"os"
	"fmt"
	"strings"
	"strconv"
	"time"
)
var MAX_NODE_NUM = 10
var TEST_NODE_NUM = 1000
type Rank struct{
	Key int32
	Val int32
}

//func main() {
//	rand.Seed(time.Now().UnixNano())
//	sl := skiplist.CreateSkipList()
//	r := new(Rank)
//	sKey := int32(3)
//	sVal := int32(0)
//	for i:=0;i<MAX_NODE_NUM;i++{
//		r = &Rank{
//			int32(i),
//			int32(rand.Intn(100000)),
//		}
//		if r.Key == sKey{
//			sVal = r.Val
//		}
//		sl.Insert(r.Key,r.Val,r)
//	}
//	sl.Print()
//
//	sl.Delete(sKey,sVal)
//
//	sl.Print()
//}
func initList(sl *skiplist.SkipList){
	r := new(Rank)
	for i := 0; i < MAX_NODE_NUM; i++ {
		r = &Rank{
			int32(i),
			int32(rand.Intn(100000)),
		}
		sl.Insert(r.Key, r.Val, r)
	}
}
func main()  {
	sl := skiplist.CreateSkipList()
	initList(sl)
	reader := bufio.NewReader(os.Stdin)
	for{
		fmt.Println("input cmd:")
		line, _, err := reader.ReadLine()
		if nil != err{
			fmt.Printf("read input error [%v]", err.Error())
			continue
		}
		vStr := strings.Split(string(line),"#")

		switch vStr[0] {
		case "insert","delete","search":
			if len(vStr) < 3{
				fmt.Printf("param not enough, need 3 but %v\n",len(vStr))
				continue
			}
			key, _ := strconv.Atoi(vStr[1])
			val, _ := strconv.Atoi(vStr[2])
			switch vStr[0] {
			case "insert":
				r := &Rank{
					int32(key),
					int32(val),
				}
				sl.Insert(int32(key), int32(val), r)
				sl.Print()
			case "delete":
				res := sl.Delete(int32(key), int32(val))
				if res{
					fmt.Println("delete success")
				}
				sl.Print()
			case "search":
				if len(vStr) < 3{
					fmt.Printf("param not enough, need 3 but %v",len(vStr))
					continue
				}
				node := sl.Search(int32(key),int32(val))
				node.Print("find node:")
			}
		case "insert_test":
			r := new(Rank)
			begin := time.Now().UnixNano()
			for i := 0; i < TEST_NODE_NUM; i++ {
				r = &Rank{
					int32(rand.Intn(100000)),
					int32(rand.Intn(100000)),
				}
				sl.Insert(r.Key, r.Val, r)
			}
			end := time.Now().UnixNano()
			fmt.Printf("insert begin[%d] end[%d] total_use[%d] pre_use[%d/op]\n",begin, end, end-begin, (end-begin)/int64(TEST_NODE_NUM))
		}
	}
}