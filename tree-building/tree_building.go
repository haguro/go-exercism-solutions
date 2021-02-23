package tree

import (
	"fmt"
	"sort"
)

//Record represent a post entry in the database
type Record struct {
	ID     int
	Parent int
}

//Node represents a node in the tree
type Node struct {
	ID       int
	Children []*Node
}

//Build builds the posts tree, returning a pointer to the parent node
//and an error
func Build(records []Record) (*Node, error) {
	nodes := make(map[int]*Node, len(records))
	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	for i, r := range records {
		if r.ID != i || r.Parent > r.ID || r.ID > 0 && r.Parent == r.ID {
			return nil, fmt.Errorf("not in sequence or has bad parent: %v", r)
		}
		nodes[r.ID] = &Node{ID: r.ID}
		p := nodes[r.Parent]
		if r.ID != 0 {
			p.Children = append(p.Children, nodes[r.ID])
		}
	}

	return nodes[0], nil
}
