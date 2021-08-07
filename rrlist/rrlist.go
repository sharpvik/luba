package rrlist

import (
	"errors"
	"sync"
)

/* Types. */

type RoundRobinList struct {
	sync.RWMutex
	length int
	head   *RoundRobinNode
}

type RoundRobinNode struct {
	next *RoundRobinNode
	addr string
}

/* Constructors. */

func New() *RoundRobinList {
	return &RoundRobinList{
		length: 0,
		head:   NewDummyRoundRobinNode(),
	}
}

func NewDummyRoundRobinNode() (node *RoundRobinNode) {
	return &RoundRobinNode{
		next: node,
		addr: "",
	}
}

func NewRoundRobinNode(
	next *RoundRobinNode, addr string) (node *RoundRobinNode) {
	return &RoundRobinNode{
		next: node,
		addr: addr,
	}
}

/* Methods. */

func (rrl *RoundRobinList) Add(addr string) {
	rrl.Lock()
	defer rrl.Unlock()
	if rrl.length == 0 {
		rrl.head.addr = addr
	} else {
		rrl.head.next = NewRoundRobinNode(rrl.head.next, addr)
	}
	rrl.length++
}

func (rrl *RoundRobinList) Head() (addr string, err error) {
	rrl.RLock()
	defer rrl.RUnlock()
	if rrl.length > 0 {
		addr = rrl.head.addr
	} else {
		err = errors.New("RoundRobinList is not ready")
	}
	return
}

func (rrl *RoundRobinList) Next() {
	rrl.Lock()
	defer rrl.Unlock()
	rrl.head = rrl.head.next
}

// TODO: perform availability check before revealing.
func (rrl *RoundRobinList) Reveal() (addr string, err error) {
	if addr, err = rrl.Head(); err == nil {
		rrl.Next()
	}
	return
}
