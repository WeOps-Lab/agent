package controller

import (
	"sync"

	"github.com/grafana/river/ast"
	"github.com/grafana/river/vm"
)

type DeclareNode struct {
	label         string
	nodeID        string
	componentName string
	declare       *Declare
	mut           sync.RWMutex
}

var _ BlockNode = (*DeclareNode)(nil)

// NewDeclareNode creates a new declare node with a content which will be loaded by declare component nodes.
func NewDeclareNode(declare *Declare) *DeclareNode {
	return &DeclareNode{
		label:         declare.block.Label,
		nodeID:        BlockComponentID(declare.block).String(),
		componentName: declare.block.GetBlockName(),
		declare:       declare,
	}
}

func (cn *DeclareNode) Declare() *Declare {
	cn.mut.Lock()
	defer cn.mut.Unlock()
	return cn.declare
}

// Evaluate does nothing for this node.
func (cn *DeclareNode) Evaluate(scope *vm.Scope) error {
	return nil
}

func (cn *DeclareNode) Label() string { return cn.label }

// Block implements BlockNode and returns the current block of the managed config node.
func (cn *DeclareNode) Block() *ast.BlockStmt {
	cn.mut.RLock()
	defer cn.mut.RUnlock()
	return cn.declare.block
}

// NodeID implements dag.Node and returns the unique ID for the config node.
func (cn *DeclareNode) NodeID() string { return cn.nodeID }
