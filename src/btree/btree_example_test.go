package btree

import (
	"bytes"
	"fmt"
	"util"
)

type key struct {
	k string
}

func (k *key) CompareTo(other Key) int {
	return bytes.Compare([]byte(k.k), []byte(other.String()))
}

func (k *key) String() string {
	return k.k
}

var btree *BTree

func init() {
	btree = NewBTree(3)
	btree.Insert(&key{"A"})
	btree.Insert(&key{"C"})
	btree.Insert(&key{"G"})
	btree.Insert(&key{"J"})
	btree.Insert(&key{"K"})
	btree.Insert(&key{"D"})
	btree.Insert(&key{"E"})
	btree.Insert(&key{"M"})
	btree.Insert(&key{"N"})
	btree.Insert(&key{"O"})
	btree.Insert(&key{"P"})
	btree.Insert(&key{"R"})
	btree.Insert(&key{"S"})
	btree.Insert(&key{"X"})
	btree.Insert(&key{"Y"})
	btree.Insert(&key{"Z"})
	btree.Insert(&key{"T"})
	btree.Insert(&key{"U"})
	btree.Insert(&key{"V"})
}

// 测试插入，选择的例子见《算法导论》P285

func ExampleInsert() {
	str := util.TreeString(btree.root, "\n")
	fmt.Println(str)
	// Output:
	// G, M, P, X
	// |-- A, C, D, E
	// |-- J, K
	// |-- N, O
	// |-- R, S, T, U, V
	// `-- Y, Z
}

func ExampleInsertB() {
	btree.Insert(&key{"B"})
	fmt.Println(util.TreeString(btree.root, "\n"))
	// Output:
	// G, M, P, X
	// |-- A, B, C, D, E
	// |-- J, K
	// |-- N, O
	// |-- R, S, T, U, V
	// `-- Y, Z
}

func ExampleInsertQ() {
	btree.Insert(&key{"Q"})
	fmt.Println(util.TreeString(btree.root, "\n"))
	// Output:
	// G, M, P, T, X
	// |-- A, B, C, D, E
	// |-- J, K
	// |-- N, O
	// |-- Q, R, S
	// |-- U, V
	// `-- Y, Z
}

func ExampleInsertL() {
	btree.Insert(&key{"L"})
	fmt.Println(util.TreeString(btree.root, "\n"))
	// Output:
	// P
	// |-- G, M
	// |   |-- A, B, C, D, E
	// |   |-- J, K, L
	// |   `-- N, O
	// `-- T, X
	//     |-- Q, R, S
	//     |-- U, V
	//     `-- Y, Z
}

func ExampleInsertF() {
	btree.Insert(&key{"F"})
	fmt.Println(util.TreeString(btree.root, "\n"))
	// Output:
	// P
	// |-- C, G, M
	// |   |-- A, B
	// |   |-- D, E, F
	// |   |-- J, K, L
	// |   `-- N, O
	// `-- T, X
	//     |-- Q, R, S
	//     |-- U, V
	//     `-- Y, Z
}
