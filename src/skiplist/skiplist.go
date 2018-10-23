package skiplist

import (
	"math/rand"
	"fmt"
	"time"
	"strconv"
)

//跳跃表，空间换时间
var (
	MAX_LEVEL = int32(32)
	HEAD_KEY = int32(-1)
	TAIL_KEY = int32(-2)
	)

type SkipList struct {
	header 	*Node	//
	tail *Node
}

type Node struct {
	Index int32      //序列
	Key  int32		 //主键
	Val	 int32		 //排序的标准
	Info interface{} //实际信息
	Forword	 []*Node //存储节点在每一层(该节点随机最高层直到底层)的指向的节点指针
}

func CreateSkipList() *SkipList{
	rand.Seed(time.Now().UnixNano())
	sl := new(SkipList)
	sl.header = CreateNode(HEAD_KEY,0, 1, nil)
	sl.tail = CreateNode(TAIL_KEY,0, 0, nil)
	sl.header.Forword[0] = sl.tail
	return sl
}

func CreateNode(key int32, val int32, level int32, info interface{}) *Node{
	return &Node{Key:key, Val:val, Info:info, Forword:make([]*Node, level)}
}

func (sl *SkipList)Search(key int32, val int32) *Node{
	pNode := sl.header
	maxLevel := int32(len(sl.header.Forword))
	for i := maxLevel-1; i >= 0; i--{
		for{
			if pNode.Val == val && pNode.Key == key{
				return pNode
			}else if TAIL_KEY != pNode.Forword[i].Key && pNode.Forword[i].Val<=val{
				pNode = pNode.Forword[i]
				continue
			}
			//如果到队伍或者该拍到当前指向node前面则跳到下一层的当前node位置继续查找
			break
		}
	}
	return nil
}

func (sl *SkipList)Insert(key int32, val int32, info interface{}){
	level := random_level()
	maxLevel := int32(len(sl.header.Forword))
	if level > maxLevel {
		for i:=level;i>maxLevel;i--{
			sl.header.Forword = append(sl.header.Forword, sl.tail)

		}
	}
	maxLevel = int32(len(sl.header.Forword))
	//fmt.Println("==================Insert===========Level=",maxLevel)
	//1.找到插入位置
	pNode := sl.header
	updateNodeList := make([]*Node, maxLevel)
	for i := maxLevel-1;i >= 0;i--{
		for{
			if TAIL_KEY != pNode.Forword[i].Key && pNode.Forword[i].Val<val{
				//pNode.Print("Next")
				pNode = pNode.Forword[i]
				continue
			}
			break
		}
		//pNode.Print(fmt.Sprintf("Update Update Node %v", i))
		//保存每层最后一个排到新增节点前的节点，方便调整指针
		updateNodeList[i] = pNode
	}

	//2.创建新节点
	newNode := CreateNode(key, val, level, info)
	//3.调整指针
	for i:=level-1;i>= 0;i--{
		updateNode := updateNodeList[i]
		//updateNode.Print("update")
		//newNode.Print("new")
		newNode.Forword[i] = updateNode.Forword[i]
		updateNode.Forword[i] = newNode
	}
	//4.调整Index
	node := newNode
	for {
		if nil == node.Forword[0].Info {
			break
		}
		node.Index = node.Forword[0].Index+1
		node = node.Forword[0]
	}
}

func (sl *SkipList)Delete(key int32, val int32) bool{
	fmt.Printf("Del key %v val %v\n", key, val)
	pNode := sl.header
	maxLevel := int32(len(sl.header.Forword))
	//1.找到删除位置
	updateNodeList := make([]*Node, maxLevel)
	for i := maxLevel - 1; i >= 0; i-- {
		for {
			if TAIL_KEY != pNode.Forword[i].Key && pNode.Forword[i].Val < val {
				pNode = pNode.Forword[i]
				continue
			}
			break
		}

		//保存每层最后一个排到删除节点前的节点，方便调整指针
		updateNodeList[i] = pNode
	}
	pDelNode := new(Node)
	if key == updateNodeList[0].Forword[0].Key && val == updateNodeList[0].Forword[0].Val{
		pDelNode = updateNodeList[0].Forword[0]
	}

	if nil == pDelNode {
		fmt.Println("can not find delete node!")
		return false
	}
	//3.调整指针
	level := len(pDelNode.Forword)
	for i:=level-1;i>= 0;i--{
		updateNode := updateNodeList[i]
		updateNode.Forword[i] = pDelNode.Forword[i]
	}
	//4.调整key
	node := pNode
	for {
		if nil == node.Forword[0].Info {
			break
		}
		node.Forword[0].Index = node.Index
		node = node.Forword[0]
	}

	return true
}
func (node *Node) Print(mark string){
	fmt.Printf("[Print Node(%s)]\n",mark)
	if nil == node{
		fmt.Println("error nil node!")
		return
	}
	level := len(node.Forword)
	fmt.Printf("Index[%v] Key[%v] Val[%v] Level[%v]\n",node.Index, node.Key, node.Val, level)
	for i:=level-1;i>= 0;i--{
		if node.Forword[i] == nil{
			fmt.Printf("-----ForwordNode[%d] {nil}\n", i)
		}else{
			fmt.Printf("-----ForwordNode[%d] {Index[%v] Key[%v] Val[%v]}\n", i, node.Forword[i].Index, node.Forword[i].Key, node.Forword[i].Val)
		}
	}

}
func (sl *SkipList) Print(){
	fmt.Println("[**Print List Begin**]")
	node := sl.header
	for {
		if node.Key == TAIL_KEY{
			break
		}
		node.Print(strconv.Itoa(int(node.Key)))
		node = node.Forword[0]
	}
	fmt.Println("[**Print List End**]")
}
func random_level() int32{
	level := int32(1)
	for{
		r := rand.Intn(2)
		if  r > 0{
			level ++
		}else{
			break
		}
		if level >= MAX_LEVEL{
			break
		}
	}

	return level
}