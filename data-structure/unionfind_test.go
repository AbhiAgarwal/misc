package datastructure

import "testing"

func TestUnionFind(t *testing.T) {
	unionfind := NewUnionFind()
	unionfind.UnionFind(4)
	if unionfind.NumSets != 4 {
		t.Error("Size initialized is incorrect")
	}
}

func TestFindSet(t *testing.T) {
	unionfind := NewUnionFind()
	unionfind.UnionFind(4)
	if unionfind.FindSet(1) != 1 {
		t.Error("Set cannot be found")
	}
}

func TestIsSameSet(t *testing.T) {
	unionfind := NewUnionFind()
	unionfind.UnionFind(4)
	if !unionfind.IsSameSet(1, 1) {
		t.Error("Same Set is not working")
	}
}

func TestIsUnionSet(t *testing.T) {
	unionfind := NewUnionFind()
	unionfind.UnionFind(4)
	// 1 and 2 are not the same set
	unionfind.IsSameSet(1, 2)
	unionfind.UnionSet(1, 2)
	if !unionfind.IsSameSet(1, 2) {
		t.Error("Union operation is not working")
	}
}

func TestIsNumDisjointSets(t *testing.T) {
	unionfind := NewUnionFind()
	unionfind.UnionFind(4)
	// 1 and 2 are not the same set
	unionfind.IsSameSet(1, 2)
	unionfind.UnionSet(1, 2)
	if unionfind.NumDisjointSets() != 3 {
		t.Error("Wrong number of disjoint sets")
	}
}

func TestIsSizeOfSet(t *testing.T) {
	unionfind := NewUnionFind()
	unionfind.UnionFind(4)
	// 1 and 2 are not the same set
	unionfind.IsSameSet(1, 2)
	unionfind.UnionSet(1, 2)
	if unionfind.SizeOfSet(1) != 2 {
		t.Error("Wrong number within set 1")
	}
}
