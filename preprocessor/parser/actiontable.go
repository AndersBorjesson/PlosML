package parser

import "ploshml/preprocessor/token"

type (
	actionTable [numStates]actionRow
	actionRow   struct {
		canRecover bool
		actions    map[token.Type]action
	}
)

var actionTab = actionTable{
	actionRow{ // S0
		canRecover: false,
		actions: map[token.Type]action{
			token.T_0: shift(3), /* \include{ */
		},
	},
	actionRow{ // S1
		canRecover: false,
		actions: map[token.Type]action{
			token.EOF: accept(true), /* $ */
		},
	},
	actionRow{ // S2
		canRecover: false,
		actions: map[token.Type]action{
			token.T_1: shift(4), /* filename */
		},
	},
	actionRow{ // S3
		canRecover: false,
		actions: map[token.Type]action{
			token.T_1: reduce(2), /* filename, reduce: Start */
		},
	},
	actionRow{ // S4
		canRecover: false,
		actions: map[token.Type]action{
			token.T_2: shift(6), /* } */
		},
	},
	actionRow{ // S5
		canRecover: false,
		actions: map[token.Type]action{
			token.EOF: reduce(1), /* $, reduce: Expr */
		},
	},
	actionRow{ // S6
		canRecover: false,
		actions: map[token.Type]action{
			token.EOF: reduce(3), /* $, reduce: Stop */
		},
	},
}
