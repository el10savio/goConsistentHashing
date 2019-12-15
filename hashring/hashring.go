package hashring

import (
	"hash/adler32"
	"sort"
)

type HashRing struct {
	NodeList Nodes
}

type Nodes []Node

type Node struct {
	ID   string
	Hash int
}

func (nodes Nodes) Len() int           { return len(nodes) }
func (nodes Nodes) Swap(i, j int)      { nodes[i], nodes[j] = nodes[j], nodes[i] }
func (nodes Nodes) Less(i, j int) bool { return nodes[i].Hash < nodes[j].Hash }

func HashValue(value string) int {
	return int(adler32.Checksum([]byte(value)))
}

func InitializeRing() HashRing {
	return HashRing{}
}

func InitializeNode(ID string) Node {
	var newNode Node

	newNode.ID = ID
	newNode.Hash = HashValue(ID)

	return newNode
}

func (ring *HashRing) Sort() {
	sort.Sort(ring.NodeList)
}

func (ring *HashRing) indexOfNode(nodeToBeIndexed Node) int {
	nodeIndex := -1

	for index, node := range ring.NodeList {
		if node.Hash == nodeToBeIndexed.Hash {
			nodeIndex = index
			break
		}
	}

	return nodeIndex
}

func (ring *HashRing) InsertNode(node Node) {
	ring.NodeList = append(ring.NodeList, node)
	ring.Sort()
}

func (ring *HashRing) RemoveNode(node Node) {
	nodeIndex := ring.indexOfNode(node)

	if nodeIndex != -1 {
		ring.NodeList = append(ring.NodeList[:nodeIndex], ring.NodeList[nodeIndex+1:]...)
		ring.Sort()
	}
}

// Only tells which node to add to
func (ring *HashRing) AddValue(value string) string {
	entryHash := HashValue(value)

	for index := 1; index < len(ring.NodeList); index++ {
		if ring.NodeList[index-1].Hash <= entryHash && entryHash < ring.NodeList[index].Hash {
			return ring.NodeList[index].ID
		}
	}

	return ring.NodeList[0].ID
}