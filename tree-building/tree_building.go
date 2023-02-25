package tree

import (
	"fmt"
	"sort"
)

// Record represents record type
type Record struct {
	ID     int
	Parent int
}

// Node represents node type
type Node struct {
	ID       int
	Children []*Node
}

// Build creates a tree of records
func Build(records []Record) (*Node, error) {
	var n *Node
	// address empty record list
	if len(records) == 0 {
		return nil, nil
	}
	root, err := findRootRecord(records)
	//no root record
	if err != nil {
		return nil, fmt.Errorf("No root record found")
	}
	// wrong parent
	if wrongRootParent(root) {
		return nil, fmt.Errorf("Wrong root parent")
	}
	// duplicate nodes
	if duplicateNodes(records) {
		return nil, fmt.Errorf("Duplicate nodes")
	}
	//non-continuous nodes
	if nonContinuous(records) {
		return nil, fmt.Errorf("Non-continous nodes")
	}
	// nodes cycle directly
	if cycleDirectly(records) {
		return nil, fmt.Errorf("Cycle directly nodes")

	}
	//higherIDParent
	if higherIDParent(records) {
		return nil, fmt.Errorf("Higher ID parent than the node's ID")
	}

	n = &Node{
		ID:       0,
		Children: findAllChildren(0, records),
	}
	return n, nil
}

func findAllChildren(p int, records []Record) []*Node {
	var ch []*Node
	for _, r := range records {
		// exclude root element
		if r.Parent == p && r.ID != 0 {
			n := &Node{
				ID:       r.ID,
				Children: findAllChildren(r.ID, records),
			}
			ch = append(ch, n)
		}
	}
	sort.Slice(ch, func(i, j int) bool {
		return ch[i].ID < ch[j].ID
	})
	return ch
}

func wrongRootParent(root Record) bool {
	return root.Parent != 0

}

func findRootRecord(records []Record) (Record, error) {
	for _, r := range records {
		if r.ID == 0 {
			root := r
			return root, nil
		}
	}
	return Record{}, fmt.Errorf("No root element found in records")
}

func duplicateNodes(records []Record) bool {
	seen := make(map[int]bool)
	for _, r := range records {
		if seen[r.ID] == true {
			return true
		}
		seen[r.ID] = true
	}
	return false
}

func nonContinuous(records []Record) bool {
	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})
	// max element ID must match last index of records
	return (records[len(records)-1].ID == len(records))

}

func cycleDirectly(records []Record) bool {
	for _, r := range records {
		if r.ID == r.Parent && r.ID != 0 {
			return true
		}
	}
	return false
}

func higherIDParent(records []Record) bool {
	for _, r := range records {
		if r.ID < r.Parent {
			return true
		}
	}
	return false
}
