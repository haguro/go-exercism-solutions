package tree

import (
	"errors"
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

	for i, record := range records {
		if record.Parent > record.ID ||
			record.ID != i ||
			(record.ID != 0 && record.ID == record.Parent) {
			return nil, errors.New("invalid record set")
		}

		nodes[record.ID] = &Node{ID: record.ID}
		p := nodes[record.Parent]
		if record.ID != 0 {
			p.Children = append(p.Children, nodes[record.ID])
		}
	}

	return nodes[0], nil
}
