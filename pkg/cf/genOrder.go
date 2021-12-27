package cf

import (
	"errors"
)

var ErrRefIsCyclic = errors.New("a ref can not be cyclic")
var ErrSelfReference = errors.New("a ref can not reference to itself")
var ErrRefInvalid = errors.New("invalid ref format")
var ErrRefToNotExist = errors.New("a ref can not point to not exist database and collection")

// dbIdx is index in config.databases
// graphIdx is index in graph
type genOrder struct {
	config *SSConfig

	nymMapDB

	dbIdxToGraphIdx map[int]int
	graph           []*orderRootNode

	order []int
	error
}

type orderRootNode struct {
	nym   string
	dbIdx int

	nodes []*orderNode
}

type orderNode struct {
	graphIdx int
	ref      string
}

// nym
type nymMapDB map[string]int

func (config *SSConfig) NewGenOrder() (*genOrder, error) {
	order := &genOrder{config: config}
	order.initNymMapDB()
	return order, order.error
}

func (order *genOrder) initNymMapDB() {
	order.nymMapDB = make(nymMapDB)
	order.dbIdxToGraphIdx = make(map[int]int)

	for dbIdx, db := range order.config.Databases {
		nym := db.CreateNym()
		if _, ok := order.nymMapDB[nym]; ok {
			continue
		}
		order.nymMapDB[nym] = dbIdx

		rootNode := &orderRootNode{nym: nym, dbIdx: dbIdx, nodes: make([]*orderNode, 0)}
		order.dbIdxToGraphIdx[dbIdx] = len(order.graph)
		order.graph = append(order.graph, rootNode)
	}

	order.setUpGraph()
	order.findOrder()
}

func (order *genOrder) setUpGraph() {
	for rootNodeGraphIdx, rootNode := range order.graph {
		if order.error != nil {
			break
		}

		db, _ := order.getDBByDBIndex(rootNode.dbIdx)
		refs := db.GetRef()

		for _, ref := range refs {
			nym, _, ok := SplitRef(ref)
			if !ok {
				order.error = ErrRefInvalid
				break
			}

			if db.isEqualToNym(nym) {
				order.error = ErrSelfReference
				break
			}

			precedingToGenDBIdx, ok := order.nymMapDB[nym]
			if !ok {
				order.error = ErrRefToNotExist
				break
			}

			precedingToGenGraphIdx := order.dbIdxToGraphIdx[precedingToGenDBIdx]
			node := &orderNode{graphIdx: rootNodeGraphIdx, ref: ref}
			order.graph[precedingToGenGraphIdx].nodes = append(order.graph[precedingToGenGraphIdx].nodes, node)
		}
	}
}

// Kahn's Algorithm
func (order *genOrder) findOrder() {
	if order.error != nil {
		return
	}

	degrees := make([]int, len(order.graph))

	for _, rootNode := range order.graph {
		for _, node := range rootNode.nodes {
			idx := node.graphIdx
			degrees[idx] += 1
		}
	}

	q := newQueue()
	for idx, degree := range degrees {
		if degree == 0 {
			q.push(idx)
		}
	}

	index := 0
	sortedOrder := make([]int, len(order.graph))

	for !q.isEmpty() {
		currentGraphIdx := q.pop()
		sortedOrder[index] = currentGraphIdx
		index++
		for _, successorNode := range order.graph[currentGraphIdx].nodes {
			successorIdx := successorNode.graphIdx
			degrees[successorIdx] -= 1
			if degrees[successorIdx] == 0 {
				q.push(successorIdx)
			}
		}
	}

	if index != len(order.graph) {
		order.error = ErrRefIsCyclic
		return
	}
	order.order = sortedOrder
}

func (order *genOrder) getDBByDBIndex(dbIdx int) (*Database, bool) {
	if 0 > dbIdx || dbIdx >= len(order.config.Databases) {
		return nil, false
	}

	return &order.config.Databases[dbIdx], true
}

func (order *genOrder) IterateDB(callback func(*Database)) {
	if order.error != nil {
		return
	}
	for _, orderGraphIdx := range order.order {
		dbIdx := order.graph[orderGraphIdx].dbIdx
		if db, ok := order.getDBByDBIndex(dbIdx); ok {
			callback(db)
		}
	}
}

type queue []int

func (q *queue) push(x int) {
	*q = append(*q, x)
}

func (q *queue) pop() int {
	if q.isEmpty() {
		return 0
	}

	var item int
	pointer := *q
	l := len(pointer)
	item, *q = pointer[0], pointer[1:l]
	return item
}

func (q *queue) size() int {
	return len(*q)
}

func (q *queue) isEmpty() bool {
	return q.size() == 0
}

func newQueue() *queue {
	return &queue{}
}
