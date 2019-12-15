package hashring

import (
	"reflect"
	"testing"
)

func TestInitializeRing(t *testing.T) {
	expectedHashRing := HashRing{}

	actualHashRing := InitializeRing()

	if !reflect.DeepEqual(expectedHashRing, actualHashRing) {
		t.Fatalf("Expected:\n%v Got:\n%v", expectedHashRing, actualHashRing)
	}
}

func TestInitializeNode(t *testing.T) {
	ID := "121"

	expectedNode := Node{
		ID:   "121",
		Hash: 19595413,
	}

	actualNode := InitializeNode(ID)

	if !reflect.DeepEqual(expectedNode, actualNode) {
		t.Fatalf("Expected:\n%v Got:\n%v", expectedNode, actualNode)
	}
}

func TestInitializeNode_NoID(t *testing.T) {
	var ID string

	expectedNode := Node{
		ID:   "-1",
		Hash: 9240671,
	}

	actualNode := InitializeNode(ID)

	if !reflect.DeepEqual(expectedNode, actualNode) {
		t.Fatalf("Expected:\n%v Got:\n%v", expectedNode, actualNode)
	}
}

func TestIndexOfNode(t *testing.T) {
	ring := InitializeRing()

	node1 := InitializeNode("1")
	node2 := InitializeNode("10")
	node3 := InitializeNode("121")

	ring.InsertNode(node1)
	ring.InsertNode(node2)
	ring.InsertNode(node3)

	expectedIndex := 2

	actualIndex := ring.IndexOfNode(node3)

	if !reflect.DeepEqual(expectedIndex, actualIndex) {
		t.Fatalf("Expected:\n%v Got:\n%v", expectedIndex, actualIndex)
	}
}

func TestIndexOfNode_NoNode(t *testing.T) {
	ring := InitializeRing()

	node1 := InitializeNode("1")
	node2 := InitializeNode("10")
	node3 := InitializeNode("121")

	nodeX := InitializeNode("300")

	ring.InsertNode(node1)
	ring.InsertNode(node2)
	ring.InsertNode(node3)

	expectedIndex := -1

	actualIndex := ring.IndexOfNode(nodeX)

	if !reflect.DeepEqual(expectedIndex, actualIndex) {
		t.Fatalf("Expected:\n%v Got:\n%v", expectedIndex, actualIndex)
	}
}

func TestIndexOfNode_EmptyNodes(t *testing.T) {
	ring := InitializeRing()

	nodeX := InitializeNode("300")

	expectedIndex := -1

	actualIndex := ring.IndexOfNode(nodeX)

	if !reflect.DeepEqual(expectedIndex, actualIndex) {
		t.Fatalf("Expected:\n%v Got:\n%v", expectedIndex, actualIndex)
	}
}

func TestInsertNode(t *testing.T) {
	ring := InitializeRing()

	node1 := InitializeNode("1")

	expectedHashRing := HashRing{
		NodeList: Nodes{
			Node{
				ID:   "1",
				Hash: 3276850,
			},
		},
	}

	ring.InsertNode(node1)
	actualHashRing := ring

	if !reflect.DeepEqual(expectedHashRing, actualHashRing) {
		t.Fatalf("Expected:\n%v Got:\n%v", expectedHashRing, actualHashRing)
	}
}

func TestRemoveNode(t *testing.T) {
	ring := InitializeRing()

	expectedHashRing := HashRing{
		NodeList: Nodes{
			Node{
				ID:   "2",
				Hash: 3342387,
			},
		},
	}

	node1 := InitializeNode("1")
	node2 := InitializeNode("2")

	ring.InsertNode(node1)
	ring.InsertNode(node2)
	ring.RemoveNode(node1)

	actualHashRing := ring

	if !reflect.DeepEqual(expectedHashRing, actualHashRing) {
		t.Fatalf("Expected:\n%v Got:\n%v", expectedHashRing, actualHashRing)
	}
}

func TestRemoveNode_NoNode(t *testing.T) {
	ring := InitializeRing()

	expectedHashRing := HashRing{
		NodeList: Nodes{
			Node{
				ID:   "1",
				Hash: 3276850,
			},
		},
	}

	node1 := InitializeNode("1")
	node2 := InitializeNode("2")

	ring.InsertNode(node1)
	ring.RemoveNode(node2)

	actualHashRing := ring

	if !reflect.DeepEqual(expectedHashRing, actualHashRing) {
		t.Fatalf("Expected:\n%v Got:\n%v", expectedHashRing, actualHashRing)
	}
}

func TestAddValue(t *testing.T) {
	ring := InitializeRing()
	value := "5"

	expectedNode := Node{
		ID:   "10",
		Hash: 9699426,
	}

	node1 := InitializeNode("1")
	node2 := InitializeNode("10")
	node3 := InitializeNode("30")

	ring.InsertNode(node1)
	ring.InsertNode(node2)
	ring.InsertNode(node3)

	actualNode := ring.AddValue(value)

	if !reflect.DeepEqual(expectedNode, actualNode) {
		t.Fatalf("Expected:\n%v Got:\n%v", expectedNode, actualNode)
	}
}
