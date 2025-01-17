package ir

import (
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	ir_pb "github.com/unpackdev/protos/dist/go/ir"
	"github.com/unpackdev/solgo/ast"
)

// Override represents an Override in the Abstract Syntax Tree.
type Override struct {
	Unit                    *ast.OverrideSpecifier `json:"ast"`
	Id                      int64                  `json:"id"`
	NodeType                ast_pb.NodeType        `json:"nodeType"`
	Name                    string                 `json:"name"`
	ReferencedDeclarationId int64                  `json:"referencedDeclarationId"`
	TypeDescription         *ast.TypeDescription   `json:"typeDescription"`
}

// GetAST returns the underlying AST node for the Override.
func (m *Override) GetAST() *ast.OverrideSpecifier {
	return m.Unit
}

// GetId returns the ID of the Override.
func (m *Override) GetId() int64 {
	return m.Id
}

// GetName returns the name of the Override.
func (m *Override) GetName() string {
	return m.Name
}

// GetNodeType returns the AST node type of the Override.
func (m *Override) GetNodeType() ast_pb.NodeType {
	return m.NodeType
}

// GetReferencedDeclarationId returns the ID of the referenced declaration for the Override.
func (m *Override) GetReferencedDeclarationId() int64 {
	return m.ReferencedDeclarationId
}

// GetTypeDescription returns the type description of the Override.
func (m *Override) GetTypeDescription() *ast.TypeDescription {
	return m.TypeDescription
}

// GetSrc returns the source node of the Override.
func (m *Override) GetSrc() ast.SrcNode {
	return m.Unit.GetSrc()
}

// ToProto converts the Override to its corresponding protobuf representation.
func (m *Override) ToProto() *ir_pb.Override {
	proto := &ir_pb.Override{
		Id:                      m.GetId(),
		NodeType:                m.GetNodeType(),
		Name:                    m.GetName(),
		ReferencedDeclarationId: m.GetReferencedDeclarationId(),
		TypeDescription:         m.GetTypeDescription().ToProto(),
	}
	return proto
}
