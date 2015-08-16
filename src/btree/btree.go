package btree

import (
	"bytes"
)

// 保存在B树中的关键字，同时也可以保存其他数据
type Key interface {
	CompareTo(other Key) int
	String() string
}

type BTree struct {
	root *BTreeNode
	t    int // 最小度数，除根节点外的内部节点至少有t个孩子，至多2t个孩子
}

type BTreeNode struct {
	keys     []Key        // 关键字
	children []*BTreeNode // 孩子节点
	isLeaf   bool         // 是否是叶子节点
}

// NewBTree 创建一颗空的B树，只有根节点
func NewBTree(minDegree int) *BTree {
	t := &BTree{
		t: minDegree,
	}

	x := t.allocateNode()
	x.isLeaf = true

	t.diskWrite(x)
	t.root = x

	return t
}

// allocateNode 创建一个空节点
func (t *BTree) allocateNode() *BTreeNode {
	return &BTreeNode{
		keys:     make([]Key, 0),        // 关键字最多为2t-1个
		children: make([]*BTreeNode, 0), // 最多2t个孩子
	}
}

// diskWrite 将节点n写入到磁盘
func (t *BTree) diskWrite(n *BTreeNode) {
	// do nothing
}

// 从磁盘上读取n节点下的child节点
func (t *BTree) diskRead(n *BTreeNode) {
	// do nothing
}

/************** 插入关键字 ******************/

// Insert插入key
func (t *BTree) Insert(key Key) {
	r := t.root
	if len(r.keys) == 2*t.t-1 { // 根节点满了
		s := t.allocateNode() // 新的根节点
		t.root = s
		s.isLeaf = false
		s.children = append(s.children, r)
		t.splitChild(s, 0) // 分裂r
		t.insertNotFull(s, key)
	} else {
		t.insertNotFull(r, key)
	}
}

// splitChild分裂x的第i个子节点，x是非满的内部节点，x的children[i]已满，现在要分裂children[i]节点
func (t *BTree) splitChild(x *BTreeNode, i int) {
	z := t.allocateNode() // 分裂出来的节点
	y := x.children[i]    // z将是y的兄弟节点
	z.isLeaf = y.isLeaf
	d := t.t

	// y后半部分关键字分给z
	// y和z各有d-1个关键字，y为keys[0..d-1),z为keys[d,2d-1),keys[d-1]被提到父节点中
	for j := d; j < 2*d-1; j++ {
		z.keys = append(z.keys, y.keys[j])
	}
	upKey := y.keys[d-1] // 将要提升的关键字
	y.keys = y.keys[0 : d-1]

	// 如果y不是叶子，将y后半部分的孩子节点也分给z,分t个
	if !y.isLeaf {
		for j := d; j < 2*d; j++ {
			z.children = append(z.children, y.children[j])
		}
		y.children = y.children[0:d]
	}

	// 将z插入到x.children中
	// y是x.children[i],那么z现在是x.children[i+1]
	x.children = append(x.children, nil)
	for j := len(x.children) - 1; j > i+1; j-- {
		// x有n个关键字，必然有n+1个子结点
		x.children[j] = x.children[j-1]
	}
	x.children[i+1] = z

	// 将提升上来的关键字插入到x.keys中
	// 分裂前y中所有关键字都比x.keys[i]小，分裂后提升上来的关键字也比x.keys[i]小，所以插入到x.keys[i]之前
	x.keys = append(x.keys, nil)
	for j := len(x.keys) - 1; j >= i+1; j-- {
		x.keys[j] = x.keys[j-1]
	}
	x.keys[i] = upKey

	t.diskWrite(y)
	t.diskWrite(z)
	t.diskWrite(x)
}

// insertNotFull: 将k插入到x中，x不满。
func (t *BTree) insertNotFull(x *BTreeNode, k Key) {
	i := len(x.keys) - 1
	if x.isLeaf {
		x.keys = append(x.keys, nil)
		// 从后向前遍历，找到第一个小于或等于k的位置，将k插入到该位置后
		for ; i >= 0 && k.CompareTo(x.keys[i]) < 0; i-- {
			x.keys[i+1] = x.keys[i]
		}
		x.keys[i+1] = k
		t.diskWrite(x)
	} else {
		// 从后向前遍历，找到第一个小于或等于k的位置
		for ; i >= 0 && k.CompareTo(x.keys[i]) < 0; i-- {
		}
		i++
		t.diskRead(x.children[i])
		if len(x.children[i].keys) == 2*t.t-1 { // 满节点
			t.splitChild(x, i)
			// 分裂后x.keys[i]已经被替换成提升上来的那个关键字
			if k.CompareTo(x.keys[i]) > 0 {
				i++
			}
		}
		t.insertNotFull(x.children[i], k) // 尾递归，可优化
	}
}

/************** 实现util.TreeNodePrinter接口，打印树到字符串 ******************/

func (node *BTreeNode) String() string {
	s := new(bytes.Buffer)
	n := len(node.keys)
	for i := 0; i < n-1; i++ {
		s.WriteString(node.keys[i].String() + ", ")
	}
	s.WriteString(node.keys[n-1].String())
	return s.String()
}

func (n *BTreeNode) Children() []interface{} {
	leng := len(n.children)
	chil := make([]interface{}, leng)
	for i := 0; i < leng; i++ {
		chil[i] = n.children[i]
	}
	return chil
}
