package tree

import (
	"fmt"
	"sort"
)

type Record struct {
	ID     int
	Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

func Build(records []Record) (*Node, error) {
	// sort the Records slice first
	sort.Slice(records, func(i, j int) bool { return records[i].ID < records[j].ID })
	fmt.Println("slice", records)
	var n *Node
	for i, r := range records {
		// root element
		if records[i] == 0 {
			n := &Node {
				ID: i,
				Children: nil 
			}
		}
		n := &Node{
			ID:       r.ID,
			Children: []*Node{},
		}
		for _, c := range records {
			if r.Parent == c.Parent {
				cn := &Node{
					ID:       c.ID,
					Children: []*Node{},
				}
				n.Children = append(n.Children, cn)
			}
		}
	}
	return n, nil

}
