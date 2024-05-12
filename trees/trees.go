package trees

type TreeBO struct {
	Key       string
	Label     string
	ParentKey string
	Children  []TreeBO
	Ext       interface{}
}

func ToTree(trees []TreeBO) []TreeBO {
	treeMap := make(map[string][]TreeBO)
	var rootItems []TreeBO
	for index := range trees {
		tree := trees[index]
		if tree.Key == tree.ParentKey || tree.ParentKey == "root" {
			rootItems = append(rootItems, tree)
			continue
		}
		tmpBo := treeMap[tree.ParentKey]
		if len(tmpBo) == 0 {
			treeMap[tree.ParentKey] = append(treeMap[tree.ParentKey], tree)
		} else {
			treeMap[tree.ParentKey] = append(tmpBo, tree)
		}
	}
	return makeTreeDO(rootItems, treeMap)
}

func makeTreeDO(parentItem []TreeBO, treeMap map[string][]TreeBO) []TreeBO {
	var treeItems []TreeBO
	for _, item := range parentItem {
		pchildren := treeMap[item.Key]
		if len(pchildren) == 0 {
			treeItems = append(treeItems, item)
			continue
		}
		pchildren = makeTreeDO(pchildren, treeMap)
		item.Children = pchildren
		treeItems = append(treeItems, item)
	}
	return treeItems
}
