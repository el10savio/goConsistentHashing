package hashring

import (
	"hash/adler32"
	"sort"
)

// HashRing stores all the available
// nodes that we can access
type HashRing struct {
	NodeList Nodes
}

// Nodes describes the composite
// type of the slice of Nodes
type Nodes []Node

// Node stores the ID, usually the
// IP address of the node and
// its hash value
type Node struct {
	ID   string
	Hash int
}

// The utility functions are
// implemented so that
// Nodes can be sorted
func (nodes Nodes) Len() int           { return len(nodes) }
func (nodes Nodes) Swap(i, j int)      { nodes[i], nodes[j] = nodes[j], nodes[i] }
func (nodes Nodes) Less(i, j int) bool { return nodes[i].Hash < nodes[j].Hash }

// HashValue returns the integer
// hash using adler32
func HashValue(value string) int {
	return int(adler32.Checksum([]byte(value)))
}

// InitializeRing returns
// an empty HashRing
func InitializeRing() HashRing {
	return HashRing{}
}

// InitializeNode returns a node with the
// hash generated. In the case the ID is
// not present it returns -1 as the
// node's ID
func InitializeNode(ID string) Node {
	var newNode Node

	if ID == "" {
		newNode.ID = "-1"
		newNode.Hash = HashValue("-1")
		return newNode
	}

	newNode.ID = ID
	newNode.Hash = HashValue(ID)

	return newNode
}

// Sort sorts the nodes in the HashRing
// in ascending order based on
// the node's hash value
func (ring *HashRing) Sort() {
	sort.Sort(ring.NodeList)
}

// IndexOfNode iterates through the HashRing
// and returns the index of the node that
// is searched. In the case the node is
// not found it returns -1
func (ring *HashRing) IndexOfNode(nodeToBeIndexed Node) int {
	nodeIndex := -1

	for index, node := range ring.NodeList {
		if node.Hash == nodeToBeIndexed.Hash {
			nodeIndex = index
			break
		}
	}

	return nodeIndex
}

// InsertNode first appends the node to the
// HashRing.NodeList and the sorts it
func (ring *HashRing) InsertNode(node Node) {
	if node.ID != "-1" {
		ring.NodeList = append(ring.NodeList, node)
		ring.Sort()
	}
}

// RemoveNode searches for the node to be removed
// and then if found removes it from
// HashRing.NodeList
func (ring *HashRing) RemoveNode(node Node) {
	nodeIndex := ring.IndexOfNode(node)

	if nodeIndex != -1 {
		ring.NodeList = append(ring.NodeList[:nodeIndex], ring.NodeList[nodeIndex+1:]...)
		ring.Sort()
	}
}

// AddValue returns the node in the HashRing that the
// value can be added into based on its has. In the
// case where the hash is greater than the last
// node, it adds the value to the first node
// present, thus looping through
// all nodes like a ring
func (ring *HashRing) AddValue(value string) Node {
	entryHash := HashValue(value)

	for index := 1; index < len(ring.NodeList); index++ {
		if ring.NodeList[index-1].Hash <= entryHash && entryHash < ring.NodeList[index].Hash {
			return ring.NodeList[index]
		}
	}

	return ring.NodeList[0]
}
