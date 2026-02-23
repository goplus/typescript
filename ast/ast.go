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

// Kind represents the kind of a syntax node.
type Kind = ast.Kind

const (
	KindUnknown                       = ast.KindUnknown
	KindEndOfFile                     = ast.KindEndOfFile
	KindSingleLineCommentTrivia       = ast.KindSingleLineCommentTrivia
	KindMultiLineCommentTrivia        = ast.KindMultiLineCommentTrivia
	KindNewLineTrivia                 = ast.KindNewLineTrivia
	KindWhitespaceTrivia              = ast.KindWhitespaceTrivia
	KindConflictMarkerTrivia          = ast.KindConflictMarkerTrivia
	KindNonTextFileMarkerTrivia       = ast.KindNonTextFileMarkerTrivia
	KindNumericLiteral                = ast.KindNumericLiteral
	KindBigIntLiteral                 = ast.KindBigIntLiteral
	KindStringLiteral                 = ast.KindStringLiteral
	KindJsxText                       = ast.KindJsxText
	KindJsxTextAllWhiteSpaces         = ast.KindJsxTextAllWhiteSpaces
	KindRegularExpressionLiteral      = ast.KindRegularExpressionLiteral
	KindNoSubstitutionTemplateLiteral = ast.KindNoSubstitutionTemplateLiteral
	// Pseudo-literals
	KindTemplateHead   = ast.KindTemplateHead
	KindTemplateMiddle = ast.KindTemplateMiddle
	KindTemplateTail   = ast.KindTemplateTail
	// Punctuation
	KindOpenBraceToken                         = ast.KindOpenBraceToken
	KindCloseBraceToken                        = ast.KindCloseBraceToken
	KindOpenParenToken                         = ast.KindOpenParenToken
	KindCloseParenToken                        = ast.KindCloseParenToken
	KindOpenBracketToken                       = ast.KindOpenBracketToken
	KindCloseBracketToken                      = ast.KindCloseBracketToken
	KindDotToken                               = ast.KindDotToken
	KindDotDotDotToken                         = ast.KindDotDotDotToken
	KindSemicolonToken                         = ast.KindSemicolonToken
	KindCommaToken                             = ast.KindCommaToken
	KindQuestionDotToken                       = ast.KindQuestionDotToken
	KindLessThanToken                          = ast.KindLessThanToken
	KindLessThanSlashToken                     = ast.KindLessThanSlashToken
	KindGreaterThanToken                       = ast.KindGreaterThanToken
	KindLessThanEqualsToken                    = ast.KindLessThanEqualsToken
	KindGreaterThanEqualsToken                 = ast.KindGreaterThanEqualsToken
	KindEqualsEqualsToken                      = ast.KindEqualsEqualsToken
	KindExclamationEqualsToken                 = ast.KindExclamationEqualsToken
	KindEqualsEqualsEqualsToken                = ast.KindEqualsEqualsEqualsToken
	KindExclamationEqualsEqualsToken           = ast.KindExclamationEqualsEqualsToken
	KindEqualsGreaterThanToken                 = ast.KindEqualsGreaterThanToken
	KindPlusToken                              = ast.KindPlusToken
	KindMinusToken                             = ast.KindMinusToken
	KindAsteriskToken                          = ast.KindAsteriskToken
	KindAsteriskAsteriskToken                  = ast.KindAsteriskAsteriskToken
	KindSlashToken                             = ast.KindSlashToken
	KindPercentToken                           = ast.KindPercentToken
	KindPlusPlusToken                          = ast.KindPlusPlusToken
	KindMinusMinusToken                        = ast.KindMinusMinusToken
	KindLessThanLessThanToken                  = ast.KindLessThanLessThanToken
	KindGreaterThanGreaterThanToken            = ast.KindGreaterThanGreaterThanToken
	KindGreaterThanGreaterThanGreaterThanToken = ast.KindGreaterThanGreaterThanGreaterThanToken
	KindAmpersandToken                         = ast.KindAmpersandToken
	KindBarToken                               = ast.KindBarToken
	KindCaretToken                             = ast.KindCaretToken
	KindExclamationToken                       = ast.KindExclamationToken
	KindTildeToken                             = ast.KindTildeToken
	KindAmpersandAmpersandToken                = ast.KindAmpersandAmpersandToken
	KindBarBarToken                            = ast.KindBarBarToken
	KindQuestionToken                          = ast.KindQuestionToken
	KindColonToken                             = ast.KindColonToken
	KindAtToken                                = ast.KindAtToken
	KindQuestionQuestionToken                  = ast.KindQuestionQuestionToken
	/** Only the JSDoc scanner produces BacktickToken. The normal scanner produces NoSubstitutionTemplateLiteral and related kinds. */
	KindBacktickToken = ast.KindBacktickToken
	/** Only the JSDoc scanner produces HashToken. The normal scanner produces PrivateIdentifier. */
	KindHashToken = ast.KindHashToken
	// Assignments
	KindEqualsToken                                  = ast.KindEqualsToken
	KindPlusEqualsToken                              = ast.KindPlusEqualsToken
	KindMinusEqualsToken                             = ast.KindMinusEqualsToken
	KindAsteriskEqualsToken                          = ast.KindAsteriskEqualsToken
	KindAsteriskAsteriskEqualsToken                  = ast.KindAsteriskAsteriskEqualsToken
	KindSlashEqualsToken                             = ast.KindSlashEqualsToken
	KindPercentEqualsToken                           = ast.KindPercentEqualsToken
	KindLessThanLessThanEqualsToken                  = ast.KindLessThanLessThanEqualsToken
	KindGreaterThanGreaterThanEqualsToken            = ast.KindGreaterThanGreaterThanEqualsToken
	KindGreaterThanGreaterThanGreaterThanEqualsToken = ast.KindGreaterThanGreaterThanGreaterThanEqualsToken
	KindAmpersandEqualsToken                         = ast.KindAmpersandEqualsToken
	KindBarEqualsToken                               = ast.KindBarEqualsToken
	KindBarBarEqualsToken                            = ast.KindBarBarEqualsToken
	KindAmpersandAmpersandEqualsToken                = ast.KindAmpersandAmpersandEqualsToken
	KindQuestionQuestionEqualsToken                  = ast.KindQuestionQuestionEqualsToken
	KindCaretEqualsToken                             = ast.KindCaretEqualsToken
	// Identifiers and PrivateIdentifier
	KindIdentifier            = ast.KindIdentifier
	KindPrivateIdentifier     = ast.KindPrivateIdentifier
	KindJSDocCommentTextToken = ast.KindJSDocCommentTextToken
	// Reserved words
	KindBreakKeyword      = ast.KindBreakKeyword
	KindCaseKeyword       = ast.KindCaseKeyword
	KindCatchKeyword      = ast.KindCatchKeyword
	KindClassKeyword      = ast.KindClassKeyword
	KindConstKeyword      = ast.KindConstKeyword
	KindContinueKeyword   = ast.KindContinueKeyword
	KindDebuggerKeyword   = ast.KindDebuggerKeyword
	KindDefaultKeyword    = ast.KindDefaultKeyword
	KindDeleteKeyword     = ast.KindDeleteKeyword
	KindDoKeyword         = ast.KindDoKeyword
	KindElseKeyword       = ast.KindElseKeyword
	KindEnumKeyword       = ast.KindEnumKeyword
	KindExportKeyword     = ast.KindExportKeyword
	KindExtendsKeyword    = ast.KindExtendsKeyword
	KindFalseKeyword      = ast.KindFalseKeyword
	KindFinallyKeyword    = ast.KindFinallyKeyword
	KindForKeyword        = ast.KindForKeyword
	KindFunctionKeyword   = ast.KindFunctionKeyword
	KindIfKeyword         = ast.KindIfKeyword
	KindImportKeyword     = ast.KindImportKeyword
	KindInKeyword         = ast.KindInKeyword
	KindInstanceOfKeyword = ast.KindInstanceOfKeyword
	KindNewKeyword        = ast.KindNewKeyword
	KindNullKeyword       = ast.KindNullKeyword
	KindReturnKeyword     = ast.KindReturnKeyword
	KindSuperKeyword      = ast.KindSuperKeyword
	KindSwitchKeyword     = ast.KindSwitchKeyword
	KindThisKeyword       = ast.KindThisKeyword
	KindThrowKeyword      = ast.KindThrowKeyword
	KindTrueKeyword       = ast.KindTrueKeyword
	KindTryKeyword        = ast.KindTryKeyword
	KindTypeOfKeyword     = ast.KindTypeOfKeyword
	KindVarKeyword        = ast.KindVarKeyword
	KindVoidKeyword       = ast.KindVoidKeyword
	KindWhileKeyword      = ast.KindWhileKeyword
	KindWithKeyword       = ast.KindWithKeyword
	// Strict mode reserved words
	KindImplementsKeyword = ast.KindImplementsKeyword
	KindInterfaceKeyword  = ast.KindInterfaceKeyword
	KindLetKeyword        = ast.KindLetKeyword
	KindPackageKeyword    = ast.KindPackageKeyword
	KindPrivateKeyword    = ast.KindPrivateKeyword
	KindProtectedKeyword  = ast.KindProtectedKeyword
	KindPublicKeyword     = ast.KindPublicKeyword
	KindStaticKeyword     = ast.KindStaticKeyword
	KindYieldKeyword      = ast.KindYieldKeyword
	// Contextual keywords
	KindAbstractKeyword    = ast.KindAbstractKeyword
	KindAccessorKeyword    = ast.KindAccessorKeyword
	KindAsKeyword          = ast.KindAsKeyword
	KindAssertsKeyword     = ast.KindAssertsKeyword
	KindAssertKeyword      = ast.KindAssertKeyword
	KindAnyKeyword         = ast.KindAnyKeyword
	KindAsyncKeyword       = ast.KindAsyncKeyword
	KindAwaitKeyword       = ast.KindAwaitKeyword
	KindBooleanKeyword     = ast.KindBooleanKeyword
	KindConstructorKeyword = ast.KindConstructorKeyword
	KindDeclareKeyword     = ast.KindDeclareKeyword
	KindGetKeyword         = ast.KindGetKeyword
	KindImmediateKeyword   = ast.KindImmediateKeyword
	KindInferKeyword       = ast.KindInferKeyword
	KindIntrinsicKeyword   = ast.KindIntrinsicKeyword
	KindIsKeyword          = ast.KindIsKeyword
	KindKeyOfKeyword       = ast.KindKeyOfKeyword
	KindModuleKeyword      = ast.KindModuleKeyword
	KindNamespaceKeyword   = ast.KindNamespaceKeyword
	KindNeverKeyword       = ast.KindNeverKeyword
	KindOutKeyword         = ast.KindOutKeyword
	KindReadonlyKeyword    = ast.KindReadonlyKeyword
	KindRequireKeyword     = ast.KindRequireKeyword
	KindNumberKeyword      = ast.KindNumberKeyword
	KindObjectKeyword      = ast.KindObjectKeyword
	KindSatisfiesKeyword   = ast.KindSatisfiesKeyword
	KindSetKeyword         = ast.KindSetKeyword
	KindStringKeyword      = ast.KindStringKeyword
	KindSymbolKeyword      = ast.KindSymbolKeyword
	KindTypeKeyword        = ast.KindTypeKeyword
	KindUndefinedKeyword   = ast.KindUndefinedKeyword
	KindUniqueKeyword      = ast.KindUniqueKeyword
	KindUnknownKeyword     = ast.KindUnknownKeyword
	KindUsingKeyword       = ast.KindUsingKeyword
	KindFromKeyword        = ast.KindFromKeyword
	KindGlobalKeyword      = ast.KindGlobalKeyword
	KindBigIntKeyword      = ast.KindBigIntKeyword
	KindOverrideKeyword    = ast.KindOverrideKeyword
	KindOfKeyword          = ast.KindOfKeyword
	KindDeferKeyword       = ast.KindDeferKeyword // LastKeyword and LastToken and LastContextualKeyword
	// Parse tree nodes
	// Names
	KindQualifiedName        = ast.KindQualifiedName
	KindComputedPropertyName = ast.KindComputedPropertyName
	// Signature elements
	KindTypeParameter = ast.KindTypeParameter
	KindParameter     = ast.KindParameter
	KindDecorator     = ast.KindDecorator
	// TypeMember
	KindPropertySignature           = ast.KindPropertySignature
	KindPropertyDeclaration         = ast.KindPropertyDeclaration
	KindMethodSignature             = ast.KindMethodSignature
	KindMethodDeclaration           = ast.KindMethodDeclaration
	KindClassStaticBlockDeclaration = ast.KindClassStaticBlockDeclaration
	KindConstructor                 = ast.KindConstructor
	KindGetAccessor                 = ast.KindGetAccessor
	KindSetAccessor                 = ast.KindSetAccessor
	KindCallSignature               = ast.KindCallSignature
	KindConstructSignature          = ast.KindConstructSignature
	KindIndexSignature              = ast.KindIndexSignature
	// Type
	KindTypePredicate           = ast.KindTypePredicate
	KindTypeReference           = ast.KindTypeReference
	KindFunctionType            = ast.KindFunctionType
	KindConstructorType         = ast.KindConstructorType
	KindTypeQuery               = ast.KindTypeQuery
	KindTypeLiteral             = ast.KindTypeLiteral
	KindArrayType               = ast.KindArrayType
	KindTupleType               = ast.KindTupleType
	KindOptionalType            = ast.KindOptionalType
	KindRestType                = ast.KindRestType
	KindUnionType               = ast.KindUnionType
	KindIntersectionType        = ast.KindIntersectionType
	KindConditionalType         = ast.KindConditionalType
	KindInferType               = ast.KindInferType
	KindParenthesizedType       = ast.KindParenthesizedType
	KindThisType                = ast.KindThisType
	KindTypeOperator            = ast.KindTypeOperator
	KindIndexedAccessType       = ast.KindIndexedAccessType
	KindMappedType              = ast.KindMappedType
	KindLiteralType             = ast.KindLiteralType
	KindNamedTupleMember        = ast.KindNamedTupleMember
	KindTemplateLiteralType     = ast.KindTemplateLiteralType
	KindTemplateLiteralTypeSpan = ast.KindTemplateLiteralTypeSpan
	KindImportType              = ast.KindImportType
	// Binding patterns
	KindObjectBindingPattern = ast.KindObjectBindingPattern
	KindArrayBindingPattern  = ast.KindArrayBindingPattern
	KindBindingElement       = ast.KindBindingElement
	// Expression
	KindArrayLiteralExpression      = ast.KindArrayLiteralExpression
	KindObjectLiteralExpression     = ast.KindObjectLiteralExpression
	KindPropertyAccessExpression    = ast.KindPropertyAccessExpression
	KindElementAccessExpression     = ast.KindElementAccessExpression
	KindCallExpression              = ast.KindCallExpression
	KindNewExpression               = ast.KindNewExpression
	KindTaggedTemplateExpression    = ast.KindTaggedTemplateExpression
	KindTypeAssertionExpression     = ast.KindTypeAssertionExpression
	KindParenthesizedExpression     = ast.KindParenthesizedExpression
	KindFunctionExpression          = ast.KindFunctionExpression
	KindArrowFunction               = ast.KindArrowFunction
	KindDeleteExpression            = ast.KindDeleteExpression
	KindTypeOfExpression            = ast.KindTypeOfExpression
	KindVoidExpression              = ast.KindVoidExpression
	KindAwaitExpression             = ast.KindAwaitExpression
	KindPrefixUnaryExpression       = ast.KindPrefixUnaryExpression
	KindPostfixUnaryExpression      = ast.KindPostfixUnaryExpression
	KindBinaryExpression            = ast.KindBinaryExpression
	KindConditionalExpression       = ast.KindConditionalExpression
	KindTemplateExpression          = ast.KindTemplateExpression
	KindYieldExpression             = ast.KindYieldExpression
	KindSpreadElement               = ast.KindSpreadElement
	KindClassExpression             = ast.KindClassExpression
	KindOmittedExpression           = ast.KindOmittedExpression
	KindExpressionWithTypeArguments = ast.KindExpressionWithTypeArguments
	KindAsExpression                = ast.KindAsExpression
	KindNonNullExpression           = ast.KindNonNullExpression
	KindMetaProperty                = ast.KindMetaProperty
	KindSyntheticExpression         = ast.KindSyntheticExpression
	KindSatisfiesExpression         = ast.KindSatisfiesExpression
	// Misc
	KindTemplateSpan          = ast.KindTemplateSpan
	KindSemicolonClassElement = ast.KindSemicolonClassElement
	// Element
	KindBlock                      = ast.KindBlock
	KindEmptyStatement             = ast.KindEmptyStatement
	KindVariableStatement          = ast.KindVariableStatement
	KindExpressionStatement        = ast.KindExpressionStatement
	KindIfStatement                = ast.KindIfStatement
	KindDoStatement                = ast.KindDoStatement
	KindWhileStatement             = ast.KindWhileStatement
	KindForStatement               = ast.KindForStatement
	KindForInStatement             = ast.KindForInStatement
	KindForOfStatement             = ast.KindForOfStatement
	KindContinueStatement          = ast.KindContinueStatement
	KindBreakStatement             = ast.KindBreakStatement
	KindReturnStatement            = ast.KindReturnStatement
	KindWithStatement              = ast.KindWithStatement
	KindSwitchStatement            = ast.KindSwitchStatement
	KindLabeledStatement           = ast.KindLabeledStatement
	KindThrowStatement             = ast.KindThrowStatement
	KindTryStatement               = ast.KindTryStatement
	KindDebuggerStatement          = ast.KindDebuggerStatement
	KindVariableDeclaration        = ast.KindVariableDeclaration
	KindVariableDeclarationList    = ast.KindVariableDeclarationList
	KindFunctionDeclaration        = ast.KindFunctionDeclaration
	KindClassDeclaration           = ast.KindClassDeclaration
	KindInterfaceDeclaration       = ast.KindInterfaceDeclaration
	KindTypeAliasDeclaration       = ast.KindTypeAliasDeclaration
	KindEnumDeclaration            = ast.KindEnumDeclaration
	KindModuleDeclaration          = ast.KindModuleDeclaration
	KindModuleBlock                = ast.KindModuleBlock
	KindCaseBlock                  = ast.KindCaseBlock
	KindNamespaceExportDeclaration = ast.KindNamespaceExportDeclaration
	KindImportEqualsDeclaration    = ast.KindImportEqualsDeclaration
	KindImportDeclaration          = ast.KindImportDeclaration
	KindImportClause               = ast.KindImportClause
	KindNamespaceImport            = ast.KindNamespaceImport
	KindNamedImports               = ast.KindNamedImports
	KindImportSpecifier            = ast.KindImportSpecifier
	KindExportAssignment           = ast.KindExportAssignment
	KindExportDeclaration          = ast.KindExportDeclaration
	KindNamedExports               = ast.KindNamedExports
	KindNamespaceExport            = ast.KindNamespaceExport
	KindExportSpecifier            = ast.KindExportSpecifier
	KindMissingDeclaration         = ast.KindMissingDeclaration
	// Module references
	KindExternalModuleReference = ast.KindExternalModuleReference
	// JSX
	KindJsxElement            = ast.KindJsxElement
	KindJsxSelfClosingElement = ast.KindJsxSelfClosingElement
	KindJsxOpeningElement     = ast.KindJsxOpeningElement
	KindJsxClosingElement     = ast.KindJsxClosingElement
	KindJsxFragment           = ast.KindJsxFragment
	KindJsxOpeningFragment    = ast.KindJsxOpeningFragment
	KindJsxClosingFragment    = ast.KindJsxClosingFragment
	KindJsxAttribute          = ast.KindJsxAttribute
	KindJsxAttributes         = ast.KindJsxAttributes
	KindJsxSpreadAttribute    = ast.KindJsxSpreadAttribute
	KindJsxExpression         = ast.KindJsxExpression
	KindJsxNamespacedName     = ast.KindJsxNamespacedName
	// Clauses
	KindCaseClause     = ast.KindCaseClause
	KindDefaultClause  = ast.KindDefaultClause
	KindHeritageClause = ast.KindHeritageClause
	KindCatchClause    = ast.KindCatchClause
	// Import attributes
	KindImportAttributes = ast.KindImportAttributes
	KindImportAttribute  = ast.KindImportAttribute
	// Property assignments
	KindPropertyAssignment          = ast.KindPropertyAssignment
	KindShorthandPropertyAssignment = ast.KindShorthandPropertyAssignment
	KindSpreadAssignment            = ast.KindSpreadAssignment
	// Enum
	KindEnumMember = ast.KindEnumMember
	// Top-level nodes
	KindSourceFile = ast.KindSourceFile
	// JSDoc nodes
	KindJSDocTypeExpression  = ast.KindJSDocTypeExpression
	KindJSDocNameReference   = ast.KindJSDocNameReference
	KindJSDocMemberName      = ast.KindJSDocMemberName // C#p
	KindJSDocAllType         = ast.KindJSDocAllType    // The * type
	KindJSDocNullableType    = ast.KindJSDocNullableType
	KindJSDocNonNullableType = ast.KindJSDocNonNullableType
	KindJSDocOptionalType    = ast.KindJSDocOptionalType
	KindJSDocVariadicType    = ast.KindJSDocVariadicType
	KindJSDoc                = ast.KindJSDoc
	KindJSDocText            = ast.KindJSDocText
	KindJSDocTypeLiteral     = ast.KindJSDocTypeLiteral
	KindJSDocSignature       = ast.KindJSDocSignature
	KindJSDocLink            = ast.KindJSDocLink
	KindJSDocLinkCode        = ast.KindJSDocLinkCode
	KindJSDocLinkPlain       = ast.KindJSDocLinkPlain
	KindJSDocTag             = ast.KindJSDocTag
	KindJSDocAugmentsTag     = ast.KindJSDocAugmentsTag
	KindJSDocImplementsTag   = ast.KindJSDocImplementsTag
	KindJSDocDeprecatedTag   = ast.KindJSDocDeprecatedTag
	KindJSDocPublicTag       = ast.KindJSDocPublicTag
	KindJSDocPrivateTag      = ast.KindJSDocPrivateTag
	KindJSDocProtectedTag    = ast.KindJSDocProtectedTag
	KindJSDocReadonlyTag     = ast.KindJSDocReadonlyTag
	KindJSDocOverrideTag     = ast.KindJSDocOverrideTag
	KindJSDocCallbackTag     = ast.KindJSDocCallbackTag
	KindJSDocOverloadTag     = ast.KindJSDocOverloadTag
	KindJSDocParameterTag    = ast.KindJSDocParameterTag
	KindJSDocReturnTag       = ast.KindJSDocReturnTag
	KindJSDocThisTag         = ast.KindJSDocThisTag
	KindJSDocTypeTag         = ast.KindJSDocTypeTag
	KindJSDocTemplateTag     = ast.KindJSDocTemplateTag
	KindJSDocTypedefTag      = ast.KindJSDocTypedefTag
	KindJSDocSeeTag          = ast.KindJSDocSeeTag
	KindJSDocPropertyTag     = ast.KindJSDocPropertyTag
	KindJSDocSatisfiesTag    = ast.KindJSDocSatisfiesTag
	KindJSDocImportTag       = ast.KindJSDocImportTag
	// Synthesized list
	KindSyntaxList = ast.KindSyntaxList
	// Reparsed JS nodes
	KindJSTypeAliasDeclaration = ast.KindJSTypeAliasDeclaration
	KindJSExportAssignment     = ast.KindJSExportAssignment
	KindCommonJSExport         = ast.KindCommonJSExport
	KindJSImportDeclaration    = ast.KindJSImportDeclaration
	// Transformation nodes
	KindNotEmittedStatement          = ast.KindNotEmittedStatement
	KindPartiallyEmittedExpression   = ast.KindPartiallyEmittedExpression
	KindCommaListExpression          = ast.KindCommaListExpression
	KindSyntheticReferenceExpression = ast.KindSyntheticReferenceExpression
	KindNotEmittedTypeElement        = ast.KindNotEmittedTypeElement
	// Enum value count
	KindCount = ast.KindCount
	// Markers
	KindFirstAssignment         = ast.KindFirstAssignment
	KindLastAssignment          = ast.KindLastAssignment
	KindFirstCompoundAssignment = ast.KindFirstCompoundAssignment
	KindLastCompoundAssignment  = ast.KindLastCompoundAssignment
	KindFirstReservedWord       = ast.KindFirstReservedWord
	KindLastReservedWord        = ast.KindLastReservedWord
	KindFirstKeyword            = ast.KindFirstKeyword
	KindLastKeyword             = ast.KindLastKeyword
	KindFirstFutureReservedWord = ast.KindFirstFutureReservedWord
	KindLastFutureReservedWord  = ast.KindLastFutureReservedWord
	KindFirstTypeNode           = ast.KindFirstTypeNode
	KindLastTypeNode            = ast.KindLastTypeNode
	KindFirstPunctuation        = ast.KindFirstPunctuation
	KindLastPunctuation         = ast.KindLastPunctuation
	KindFirstToken              = ast.KindFirstToken
	KindLastToken               = ast.KindLastToken
	KindFirstLiteralToken       = ast.KindFirstLiteralToken
	KindLastLiteralToken        = ast.KindLastLiteralToken
	KindFirstTemplateToken      = ast.KindFirstTemplateToken
	KindLastTemplateToken       = ast.KindLastTemplateToken
	KindFirstBinaryOperator     = ast.KindFirstBinaryOperator
	KindLastBinaryOperator      = ast.KindLastBinaryOperator
	KindFirstStatement          = ast.KindFirstStatement
	KindLastStatement           = ast.KindLastStatement
	KindFirstNode               = ast.KindFirstNode
	KindFirstJSDocNode          = ast.KindFirstJSDocNode
	KindLastJSDocNode           = ast.KindLastJSDocNode
	KindFirstJSDocTagNode       = ast.KindFirstJSDocTagNode
	KindLastJSDocTagNode        = ast.KindLastJSDocTagNode
	KindFirstContextualKeyword  = ast.KindFirstContextualKeyword
	KindLastContextualKeyword   = ast.KindLastContextualKeyword
	KindComment                 = ast.KindComment
	KindFirstTriviaToken        = ast.KindFirstTriviaToken
	KindLastTriviaToken         = ast.KindLastTriviaToken
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

// Symbol
type Symbol = ast.Symbol

// SymbolTable
type SymbolTable = ast.SymbolTable

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
