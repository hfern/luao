package parse

import (
//"github.com/hfern/luao/lexer"
)

type Rule func() nil

type RuleId int

const (
	NChunk RuleId = iota
	NBlock
	NStat
	NLastStat
	NFuncName
	NVarList
	NVar
	NNameList
	NExpList
	NExp
	NPrefixExp
	NFunctionCall
	NArgs
	NFunction
	NFuncBody
	NParlist
	NTableConstructor
	NFieldList
	NField
	NFieldSep
	NBinOp
	NUnOp
)
