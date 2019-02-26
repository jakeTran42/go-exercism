package tree

import (
	"errors"
	"sort"
)

// Record has two parts, an integer ID and an integer Parent
// that records its ID and the ID of its parent node.
type Record struct {
	ID     int
	Parent int
}

// Node has two parts, an integer ID and the Node's children, a
// slice of addresses for other Nodes.
type Node struct {
	ID       int
	Children []*Node
}

// Build takes in a slice of records and builds a tree. It returns
// the root node of the tree.
func Build(records []Record) (*Node, error) {
	recordMap := make(map[int][]int, 0)
	idTracker := make(map[int]bool, 0)
	highestID := 0

	if len(records) == 0 {
		return nil, nil
	}

	for _, record := range records {
		// error checking
		if record.ID == 0 && record.Parent != 0 {
			return nil, errors.New("Root record given a parent")
		} else if record.ID > 0 && record.ID <= record.Parent {
			return nil, errors.New("A record can't have a parent with a higher ID")
		} else if _, ok := idTracker[record.ID]; ok {
			return nil, errors.New("Duplicate node")
		}
		// keep track of IDs; false because we haven't built a node from it yet
		idTracker[record.ID] = false
		if record.ID > highestID {
			highestID = record.ID
		}
		// don't give the root node a parent
		if record.ID == 0 {
			continue
		}
		// put it on the map
		if value, ok := recordMap[record.Parent]; ok {
			recordMap[record.Parent] = append(value, record.ID)
		} else {
			recordMap[record.Parent] = []int{record.ID}
		}

	}

	// check if we have a continuous sequence of records
	for i := 0; i <= highestID; i++ {
		if _, ok := idTracker[i]; !ok {
			return nil, errors.New("Non-contiguous IDs")
		}
	}

	root := Node{ID: 0}
	idTracker[0] = true
	// recursively populate
	populate(&root, recordMap, idTracker)

	// check if all ids were hit
	for _, value := range idTracker {
		if value == false {
			return nil, errors.New("Discontinuous tree")
		}
	}

	return &root, nil
}

// populate is a private function that fills out a node and all its children recursively.
func populate(parent *Node, structure map[int][]int, tracker map[int]bool) {
	if value, ok := structure[parent.ID]; ok {
		sort.Ints(value)
		children := make([]*Node, len(value))
		for index, childID := range value {
			child := Node{ID: childID}
			if _, okay := structure[childID]; okay {
				populate(&child, structure, tracker)
			}
			children[index] = &child
			tracker[childID] = true
		}
		parent.Children = children
	}
}
