package ast

import (
	"fmt"
	"strings"

	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

type PayableConversionExpression struct {
	*ASTBuilder

	Id                    int64              `json:"id"`
	NodeType              ast_pb.NodeType    `json:"node_type"`
	Src                   SrcNode            `json:"src"`
	Arguments             []Node[NodeType]   `json:"arguments"`
	ArgumentTypes         []*TypeDescription `json:"argument_types"`
	ReferencedDeclaration int64              `json:"referenced_declaration,omitempty"`
	TypeDescription       *TypeDescription   `json:"type_description"`
	Payable               bool               `json:"payable"`
}

func NewPayableConversionExpression(b *ASTBuilder) *PayableConversionExpression {
	return &PayableConversionExpression{
		ASTBuilder:    b,
		Id:            b.GetNextID(),
		NodeType:      ast_pb.NodeType_PAYABLE_CONVERSION,
		ArgumentTypes: []*TypeDescription{},
	}
}

// SetReferenceDescriptor sets the reference descriptions of the PayableConversionExpression node.
func (p *PayableConversionExpression) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	p.ReferencedDeclaration = refId
	p.TypeDescription = refDesc
	return false
}

func (p *PayableConversionExpression) GetId() int64 {
	return p.Id
}

func (p *PayableConversionExpression) GetType() ast_pb.NodeType {
	return p.NodeType
}

func (p *PayableConversionExpression) GetSrc() SrcNode {
	return p.Src
}

func (p *PayableConversionExpression) GetTypeDescription() *TypeDescription {
	return p.TypeDescription
}

func (p *PayableConversionExpression) GetArgumentTypes() []*TypeDescription {
	return p.ArgumentTypes
}

func (p *PayableConversionExpression) GetArguments() []Node[NodeType] {
	return p.Arguments
}

func (p *PayableConversionExpression) IsPayable() bool {
	return p.Payable
}

func (p *PayableConversionExpression) GetNodes() []Node[NodeType] {
	return nil
}

func (p *PayableConversionExpression) ToProto() NodeType {
	return ast_pb.PayableConversion{}
}

func (p *PayableConversionExpression) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	vDeclar *VariableDeclaration,
	exprNode Node[NodeType],
	ctx *parser.PayableConversionContext,
) Node[NodeType] {
	p.Src = SrcNode{
		Id:     p.GetNextID(),
		Line:   int64(ctx.GetStart().GetLine()),
		Column: int64(ctx.GetStart().GetColumn()),
		Start:  int64(ctx.GetStart().GetStart()),
		End:    int64(ctx.GetStop().GetStop()),
		Length: int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: func() int64 {
			if vDeclar != nil {
				return vDeclar.GetId()
			}

			if exprNode != nil {
				return exprNode.GetId()
			}

			return bodyNode.GetId()
		}(),
	}
	p.Payable = ctx.Payable() != nil

	expression := NewExpression(p.ASTBuilder)

	typeStrings := []string{}
	typeIdentifiers := []string{}

	if ctx.CallArgumentList() != nil {
		for _, expressionCtx := range ctx.CallArgumentList().AllExpression() {
			expr := expression.Parse(unit, contractNode, fnNode, bodyNode, nil, p, expressionCtx)
			p.Arguments = append(
				p.Arguments,
				expr,
			)

			typeStrings = append(typeStrings, expr.GetTypeDescription().TypeString)
			typeIdentifiers = append(typeIdentifiers, expr.GetTypeDescription().TypeIdentifier)

			p.ArgumentTypes = append(
				p.ArgumentTypes,
				expr.GetTypeDescription(),
			)
		}
	}

	p.TypeDescription = &TypeDescription{
		TypeString: func() string {
			return fmt.Sprintf(
				"function(%s) payable",
				strings.Join(typeStrings, ","),
			)
		}(),
		TypeIdentifier: func() string {
			return fmt.Sprintf(
				"t_function_payable$_%s$",
				strings.Join(typeIdentifiers, "$_"),
			)
		}(),
	}

	return p
}