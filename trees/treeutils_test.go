package trees_test

import (
	"fmt"
	"testing"

	"github.com/jarkata/gotools/trees"
)

func TestTree(t *testing.T) {

	modal := []trees.TreeBO{
		{Key: "root0", ParentKey: "root0", Label: "一级0", Ext: "testext"},
		{Key: "root1", ParentKey: "root1", Label: "一级1", Ext: "testext"},
		{Key: "root11", ParentKey: "root1", Label: "一级11", Ext: "testext"},
		{Key: "root12", ParentKey: "root1", Label: "一级12", Ext: "testext"},
		{Key: "root2", ParentKey: "root2", Label: "一级2", Ext: "testext2"},
		{Key: "root21", ParentKey: "root2", Label: "二级21", Ext: "testext21"},
		{Key: "root22", ParentKey: "root2", Label: "二级22", Ext: "testext22"},
	}

	fmt.Printf("trees.ToTree(modal): %v\n", trees.ToTree(modal))
}
