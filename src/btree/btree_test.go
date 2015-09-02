package btree

import (
	"testing"
)

func TestSearch(t *testing.T) {
	node, index := btree.Search(nil, &key{"C"})
	if node == nil || index < 0 {
		t.Error("没有找到关键字C")
	}
	t.Logf("C是第%d个关键字", index)

	node, index = btree.Search(nil, &key{"H"})
	if node != nil || index >= 0 {
		t.Errorf("不应该在%d位置上找到关键字H", index)
	}
}
