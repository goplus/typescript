/*
 * Copyright (c) 2026 The XGo Authors (xgo.dev). All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ast

import (
	"github.com/microsoft/typescript-go/internal/ast"
)

// -----------------------------------------------------------------------------

// NodeBase
type NodeBase = ast.NodeBase

// AST Node
// Interface values stored in AST nodes are never typed nil values. Construction code must ensure that
// interface valued properties either store a true nil or a reference to a non-nil struct.
type Node = ast.Node

type MutableNode = ast.MutableNode

// NodeList
type NodeList = ast.NodeList

// NodeFactory
type NodeFactory = ast.NodeFactory

// DeclarationBase
type DeclarationBase = ast.DeclarationBase

// Aliases for Node unions

type (
	Statement                      = Node // Node with StatementBase
	Declaration                    = Node // Node with DeclarationBase
	Expression                     = Node // Node with ExpressionBase
	TypeNode                       = Node // Node with TypeNodeBase
	TypeElement                    = Node // Node with TypeElementBase
	ClassElement                   = Node // Node with ClassElementBase
	NamedMember                    = Node // Node with NamedMemberBase
	ObjectLiteralElement           = Node // Node with ObjectLiteralElementBase
	BlockOrExpression              = Node // Block | Expression
	AccessExpression               = Node // PropertyAccessExpression | ElementAccessExpression
	DeclarationName                = Node // Identifier | PrivateIdentifier | StringLiteral | NumericLiteral | BigIntLiteral | NoSubstitutionTemplateLiteral | ComputedPropertyName | BindingPattern | ElementAccessExpression
	ModuleName                     = Node // Identifier | StringLiteral
	ModuleExportName               = Node // Identifier | StringLiteral
	PropertyName                   = Node // Identifier | StringLiteral | NoSubstitutionTemplateLiteral | NumericLiteral | ComputedPropertyName | PrivateIdentifier | BigIntLiteral
	ModuleBody                     = Node // ModuleBlock | ModuleDeclaration
	ForInitializer                 = Node // Expression | MissingDeclaration | VariableDeclarationList
	ModuleReference                = Node // Identifier | QualifiedName | ExternalModuleReference
	NamedImportBindings            = Node // NamespaceImport | NamedImports
	NamedExportBindings            = Node // NamespaceExport | NamedExports
	MemberName                     = Node // Identifier | PrivateIdentifier
	EntityName                     = Node // Identifier | QualifiedName
	BindingName                    = Node // Identifier | BindingPattern
	ModifierLike                   = Node // Modifier | Decorator
	JsxChild                       = Node // JsxText | JsxExpression | JsxElement | JsxSelfClosingElement | JsxFragment
	JsxAttributeLike               = Node // JsxAttribute | JsxSpreadAttribute
	JsxAttributeName               = Node // Identifier | JsxNamespacedName
	JsxAttributeValue              = Node // StringLiteral | JsxExpression | JsxElement | JsxSelfClosingElement | JsxFragment
	JsxTagNameExpression           = Node // IdentifierReference | KeywordExpression | JsxTagNamePropertyAccess | JsxNamespacedName
	ClassLikeDeclaration           = Node // ClassDeclaration | ClassExpression
	AccessorDeclaration            = Node // GetAccessorDeclaration | SetAccessorDeclaration
	LiteralLikeNode                = Node // StringLiteral | NumericLiteral | BigIntLiteral | RegularExpressionLiteral | TemplateLiteralLikeNode | JsxText
	LiteralExpression              = Node // StringLiteral | NumericLiteral | BigIntLiteral | RegularExpressionLiteral | NoSubstitutionTemplateLiteral
	UnionOrIntersectionTypeNode    = Node // UnionTypeNode | IntersectionTypeNode
	TemplateLiteralLikeNode        = Node // TemplateHead | TemplateMiddle | TemplateTail
	TemplateMiddleOrTail           = Node // TemplateMiddle | TemplateTail
	TemplateLiteral                = Node // TemplateExpression | NoSubstitutionTemplateLiteral
	TypePredicateParameterName     = Node // Identifier | ThisTypeNode
	ImportAttributeName            = Node // Identifier | StringLiteral
	LeftHandSideExpression         = Node // subset of Expression
	JSDocComment                   = Node // JSDocText | JSDocLink | JSDocLinkCode | JSDocLinkPlain;
	JSDocTag                       = Node // Node with JSDocTagBase
	SignatureDeclaration           = Node // CallSignatureDeclaration | ConstructSignatureDeclaration | MethodSignature | IndexSignatureDeclaration | FunctionTypeNode | ConstructorTypeNode | FunctionDeclaration | MethodDeclaration | ConstructorDeclaration | AccessorDeclaration | FunctionExpression | ArrowFunction;
	StringLiteralLike              = Node // StringLiteral | NoSubstitutionTemplateLiteral
	AnyValidImportOrReExport       = Node // (ImportDeclaration | ExportDeclaration | JSDocImportTag) & { moduleSpecifier: StringLiteral } | ImportEqualsDeclaration & { moduleReference: ExternalModuleReference & { expression: StringLiteral }} | RequireOrImportCall | ValidImportTypeNode
	ValidImportTypeNode            = Node // ImportTypeNode & { argument: LiteralTypeNode & { literal: StringLiteral } }
	NumericOrStringLikeLiteral     = Node // StringLiteralLike | NumericLiteral
	TypeOnlyImportDeclaration      = Node // ImportClause | ImportEqualsDeclaration | ImportSpecifier | NamespaceImport with isTypeOnly: true
	ObjectLiteralLike              = Node // ObjectLiteralExpression | ObjectBindingPattern
	ObjectTypeDeclaration          = Node // ClassLikeDeclaration | InterfaceDeclaration | TypeLiteralNode
	JsxOpeningLikeElement          = Node // JsxOpeningElement | JsxSelfClosingElement
	NamedImportsOrExports          = Node // NamedImports | NamedExports
	BreakOrContinueStatement       = Node // BreakStatement | ContinueStatement
	CallLikeExpression             = Node // CallExpression | NewExpression | TaggedTemplateExpression | Decorator | JsxCallLike | InstanceofExpression
	FunctionLikeDeclaration        = Node // FunctionDeclaration | MethodDeclaration | GetAccessorDeclaration | SetAccessorDeclaration | ConstructorDeclaration | FunctionExpression | ArrowFunction
	VariableOrParameterDeclaration = Node // VariableDeclaration | ParameterDeclaration
	VariableOrPropertyDeclaration  = Node // VariableDeclaration | PropertyDeclaration
	CallOrNewExpression            = Node // CallExpression | NewExpression
	ImportClauseOrBindingPattern   = Node // ImportClause | BindingPattern
	AnyImportSyntax                = Node // ImportDeclaration | ImportEqualsDeclaration
	AnyImportOrRequireStatement    = Node // AnyImportSyntax | RequireVariableStatement
)

// Aliases for node singletons

type (
	IdentifierNode                  = Node
	PrivateIdentifierNode           = Node
	TokenNode                       = Node
	StringLiteralNode               = Node
	TemplateHeadNode                = Node
	TemplateMiddleNode              = Node
	TemplateTailNode                = Node
	TemplateSpanNode                = Node
	TemplateLiteralTypeSpanNode     = Node
	BlockNode                       = Node
	CatchClauseNode                 = Node
	CaseBlockNode                   = Node
	CaseOrDefaultClauseNode         = Node
	CaseClauseNode                  = Node
	VariableDeclarationNode         = Node
	VariableDeclarationListNode     = Node
	BindingElementNode              = Node
	TypeParameterDeclarationNode    = Node
	ParameterDeclarationNode        = Node
	HeritageClauseNode              = Node
	ExpressionWithTypeArgumentsNode = Node
	EnumDeclarationNode             = Node
	EnumMemberNode                  = Node
	ModuleDeclarationNode           = Node
	FunctionDeclarationNode         = Node
	ImportClauseNode                = Node
	ImportAttributesNode            = Node
	ImportAttributeNode             = Node
	ImportSpecifierNode             = Node
	ExportSpecifierNode             = Node
	JsxAttributesNode               = Node
	JsxOpeningElementNode           = Node
	JsxClosingElementNode           = Node
	JsxOpeningFragmentNode          = Node
	JsxClosingFragmentNode          = Node
	SourceFileNode                  = Node
	PropertyAccessExpressionNode    = Node
	TypeLiteral                     = Node
	ObjectLiteralExpressionNode     = Node
	ConstructorDeclarationNode      = Node
	NamedExportsNode                = Node
	UnionType                       = Node
	LiteralType                     = Node
	JSDocNode                       = Node
	BindingPatternNode              = Node
	TypePredicateNodeNode           = Node
)

type (
	StatementList                   = NodeList // NodeList[*Statement]
	CaseClausesList                 = NodeList // NodeList[*CaseOrDefaultClause]
	VariableDeclarationNodeList     = NodeList // NodeList[*VariableDeclaration]
	BindingElementList              = NodeList // NodeList[*BindingElement]
	TypeParameterList               = NodeList // NodeList[*TypeParameterDeclaration]
	ParameterList                   = NodeList // NodeList[*ParameterDeclaration]
	HeritageClauseList              = NodeList // NodeList[*HeritageClause]
	ClassElementList                = NodeList // NodeList[*ClassElement]
	TypeElementList                 = NodeList // NodeList[*TypeElement]
	ExpressionWithTypeArgumentsList = NodeList // NodeList[*ExpressionWithTypeArguments]
	EnumMemberList                  = NodeList // NodeList[*EnumMember]
	ImportSpecifierList             = NodeList // NodeList[*ImportSpecifier]
	ExportSpecifierList             = NodeList // NodeList[*ExportSpecifier]
	TypeArgumentList                = NodeList // NodeList[*TypeNode]
	ArgumentList                    = NodeList // NodeList[*Expression]
	TemplateSpanList                = NodeList // NodeList[*TemplateSpan]
	ElementList                     = NodeList // NodeList[*Expression]
	PropertyDefinitionList          = NodeList // NodeList[*ObjectLiteralElement]
	TypeList                        = NodeList // NodeList[*TypeNode]
	ImportAttributeList             = NodeList // NodeList[*ImportAttributeNode]
	TemplateLiteralTypeSpanList     = NodeList // NodeList[*TemplateLiteralTypeSpan]
	JsxChildList                    = NodeList // NodeList[*JsxChild]
	JsxAttributeList                = NodeList // NodeList[*JsxAttributeLike]
)

// -----------------------------------------------------------------------------

type ExternalModuleIndicatorOptions = ast.ExternalModuleIndicatorOptions

type SourceFileParseOptions = ast.SourceFileParseOptions

// SourceFile
type SourceFile = ast.SourceFile

// -----------------------------------------------------------------------------
