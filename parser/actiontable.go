
package parser

import "ploshml/token"

type(
    actionTable [numStates]actionRow
    actionRow struct {
        canRecover bool
        actions map[token.Type]action
    }
)

var actionTab = actionTable{ 
	actionRow{ // S0
        canRecover: false,
		actions: map[token.Type]action{ 
			token.T_0:shift(9),		/* #pragma */
			token.T_9:shift(10),		/* call */
			token.T_10:shift(11),		/* class */
			token.T_16:shift(12),		/* key_component */
			token.T_20:shift(13),		/* subroutine */
			token.T_22:shift(14),		/* zoomin */
        },

	},
	actionRow{ // S1
        canRecover: false,
		actions: map[token.Type]action{ 
			token.T_16:shift(15),		/* key_component */
        },

	},
	actionRow{ // S2
        canRecover: false,
		actions: map[token.Type]action{ 
			token.EOF:accept(true),		/* $ */
        },

	},
	actionRow{ // S3
        canRecover: false,
		actions: map[token.Type]action{ 
			token.T_1:shift(19),		/* , */
			token.T_2:shift(20),		/* -< */
			token.T_3:shift(21),		/* -> */
			token.T_4:shift(22),		/* :: */
			token.T_5:shift(23),		/* <- */
			token.T_6:shift(24),		/* >- */
			token.T_7:shift(25),		/* aggregates */
			token.T_11:shift(26),		/* exhibits */
			token.T_12:shift(27),		/* generalizes */
			token.T_13:shift(28),		/* handles */
			token.T_14:shift(29),		/* implements */
			token.T_17:shift(30),		/* owns */
			token.T_18:shift(31),		/* relates */
			token.T_19:shift(32),		/* requires */
			token.T_21:shift(33),		/* triggers */
        },

	},
	actionRow{ // S4
        canRecover: false,
		actions: map[token.Type]action{ 
			token.EOF:reduce(6),		/* $, reduce: EndExpr */
			token.T_2:shift(20),		/* -< */
			token.T_3:shift(21),		/* -> */
			token.T_5:shift(23),		/* <- */
			token.T_6:shift(24),		/* >- */
			token.T_7:shift(25),		/* aggregates */
			token.T_11:shift(26),		/* exhibits */
			token.T_12:shift(27),		/* generalizes */
			token.T_13:shift(28),		/* handles */
			token.T_14:shift(29),		/* implements */
			token.T_17:shift(30),		/* owns */
			token.T_18:shift(31),		/* relates */
			token.T_19:shift(32),		/* requires */
			token.T_21:shift(33),		/* triggers */
        },

	},
	actionRow{ // S5
        canRecover: false,
		actions: map[token.Type]action{ 
			token.T_16:shift(35),		/* key_component */
        },

	},
	actionRow{ // S6
        canRecover: false,
		actions: map[token.Type]action{ 
			token.EOF:reduce(7),		/* $, reduce: EndExpr */
        },

	},
	actionRow{ // S7
        canRecover: false,
		actions: map[token.Type]action{ 
			token.T_16:shift(36),		/* key_component */
        },

	},
	actionRow{ // S8
        canRecover: false,
		actions: map[token.Type]action{ 
			token.EOF:reduce(1),		/* $, reduce: EndExpr */
        },

	},
	actionRow{ // S9
        canRecover: false,
		actions: map[token.Type]action{ 
			token.T_16:reduce(42),		/* key_component, reduce: Pragma */
        },

	},
	actionRow{ // S10
        canRecover: false,
		actions: map[token.Type]action{ 
			token.T_16:reduce(43),		/* key_component, reduce: Call */
        },

	},
	actionRow{ // S11
        canRecover: false,
		actions: map[token.Type]action{ 
			token.T_16:reduce(36),		/* key_component, reduce: Structural */
        },

	},
	actionRow{ // S12
        canRecover: false,
		actions: map[token.Type]action{ 
			token.T_1:shift(19),		/* , */
			token.T_2:shift(20),		/* -< */
			token.T_3:shift(21),		/* -> */
			token.T_4:shift(22),		/* :: */
			token.T_5:shift(23),		/* <- */
			token.T_6:shift(24),		/* >- */
			token.T_7:shift(25),		/* aggregates */
			token.T_8:shift(42),		/* as */
			token.T_11:shift(26),		/* exhibits */
			token.T_12:shift(27),		/* generalizes */
			token.T_13:shift(28),		/* handles */
			token.T_14:shift(29),		/* implements */
			token.T_15:shift(43),		/* is */
			token.T_17:shift(30),		/* owns */
			token.T_18:shift(31),		/* relates */
			token.T_19:shift(32),		/* requires */
			token.T_21:shift(33),		/* triggers */
        },

	},
	actionRow{ // S13
        canRecover: false,
		actions: map[token.Type]action{ 
			token.T_16:reduce(35),		/* key_component, reduce: Structural */
        },

	},
	actionRow{ // S14
        canRecover: false,
		actions: map[token.Type]action{ 
			token.T_16:reduce(34),		/* key_component, reduce: Structural */
        },

	},
	actionRow{ // S15
        canRecover: false,
		actions: map[token.Type]action{ 
			token.T_8:shift(42),		/* as */
        },

	},
	actionRow{ // S16
        canRecover: false,
		actions: map[token.Type]action{ 
			token.T_16:shift(46),		/* key_component */
        },

	},
	actionRow{ // S17
        canRecover: false,
		actions: map[token.Type]action{ 
			token.T_16:shift(47),		/* key_component */
        },

	},
	actionRow{ // S18
        canRecover: false,
		actions: map[token.Type]action{ 
			token.T_16:shift(48),		/* key_component */
        },

	},
	actionRow{ // S19
        canRecover: false,
		actions: map[token.Type]action{ 
			token.T_16:reduce(20),		/* key_component, reduce: Separator */
        },

	},
	actionRow{ // S20
        canRecover: false,
		actions: map[token.Type]action{ 
			token.T_16:reduce(32),		/* key_component, reduce: Dualsidedkey */
        },

	},
	actionRow{ // S21
        canRecover: false,
		actions: map[token.Type]action{ 
			token.T_16:reduce(30),		/* key_component, reduce: Dualsidedkey */
        },

	},
	actionRow{ // S22
        canRecover: false,
		actions: map[token.Type]action{ 
			token.T_16:reduce(39),		/* key_component, reduce: Property */
        },

	},
	actionRow{ // S23
        canRecover: false,
		actions: map[token.Type]action{ 
			token.T_16:reduce(33),		/* key_component, reduce: Dualsidedkey */
        },

	},
	actionRow{ // S24
        canRecover: false,
		actions: map[token.Type]action{ 
			token.T_16:reduce(31),		/* key_component, reduce: Dualsidedkey */
        },

	},
	actionRow{ // S25
        canRecover: false,
		actions: map[token.Type]action{ 
			token.T_16:reduce(26),		/* key_component, reduce: Dualsidedkey */
        },

	},
	actionRow{ // S26
        canRecover: false,
		actions: map[token.Type]action{ 
			token.T_16:reduce(28),		/* key_component, reduce: Dualsidedkey */
        },

	},
	actionRow{ // S27
        canRecover: false,
		actions: map[token.Type]action{ 
			token.T_16:reduce(29),		/* key_component, reduce: Dualsidedkey */
        },

	},
	actionRow{ // S28
        canRecover: false,
		actions: map[token.Type]action{ 
			token.T_16:reduce(22),		/* key_component, reduce: Dualsidedkey */
        },

	},
	actionRow{ // S29
        canRecover: false,
		actions: map[token.Type]action{ 
			token.T_16:reduce(27),		/* key_component, reduce: Dualsidedkey */
        },

	},
	actionRow{ // S30
        canRecover: false,
		actions: map[token.Type]action{ 
			token.T_16:reduce(23),		/* key_component, reduce: Dualsidedkey */
        },

	},
	actionRow{ // S31
        canRecover: false,
		actions: map[token.Type]action{ 
			token.T_16:reduce(24),		/* key_component, reduce: Dualsidedkey */
        },

	},
	actionRow{ // S32
        canRecover: false,
		actions: map[token.Type]action{ 
			token.T_16:reduce(25),		/* key_component, reduce: Dualsidedkey */
        },

	},
	actionRow{ // S33
        canRecover: false,
		actions: map[token.Type]action{ 
			token.T_16:reduce(21),		/* key_component, reduce: Dualsidedkey */
        },

	},
	actionRow{ // S34
        canRecover: false,
		actions: map[token.Type]action{ 
			token.T_16:shift(50),		/* key_component */
        },

	},
	actionRow{ // S35
        canRecover: false,
		actions: map[token.Type]action{ 
			token.EOF:reduce(19),		/* $, reduce: PragmaExpr */
        },

	},
	actionRow{ // S36
        canRecover: false,
		actions: map[token.Type]action{ 
			token.EOF:reduce(8),		/* $, reduce: StructuralExpr */
        },

	},
	actionRow{ // S37
        canRecover: false,
		actions: map[token.Type]action{ 
			token.T_16:shift(51),		/* key_component */
        },

	},
	actionRow{ // S38
        canRecover: false,
		actions: map[token.Type]action{ 
			token.T_16:shift(53),		/* key_component */
        },

	},
	actionRow{ // S39
        canRecover: false,
		actions: map[token.Type]action{ 
			token.T_16:shift(55),		/* key_component */
        },

	},
	actionRow{ // S40
        canRecover: false,
		actions: map[token.Type]action{ 
			token.T_16:shift(56),		/* key_component */
        },

	},
	actionRow{ // S41
        canRecover: false,
		actions: map[token.Type]action{ 
			token.T_16:shift(57),		/* key_component */
        },

	},
	actionRow{ // S42
        canRecover: false,
		actions: map[token.Type]action{ 
			token.T_16:reduce(40),		/* key_component, reduce: Alias */
        },

	},
	actionRow{ // S43
        canRecover: false,
		actions: map[token.Type]action{ 
			token.T_16:reduce(41),		/* key_component, reduce: Attribute */
        },

	},
	actionRow{ // S44
        canRecover: false,
		actions: map[token.Type]action{ 
			token.T_16:shift(58),		/* key_component */
        },

	},
	actionRow{ // S45
        canRecover: false,
		actions: map[token.Type]action{ 
			token.EOF:reduce(11),		/* $, reduce: NormalExpr */
			token.T_1:shift(19),		/* , */
			token.T_2:reduce(11),		/* -<, reduce: NormalExpr */
			token.T_3:reduce(11),		/* ->, reduce: NormalExpr */
			token.T_4:shift(22),		/* :: */
			token.T_5:reduce(11),		/* <-, reduce: NormalExpr */
			token.T_6:reduce(11),		/* >-, reduce: NormalExpr */
			token.T_7:reduce(11),		/* aggregates, reduce: NormalExpr */
			token.T_11:reduce(11),		/* exhibits, reduce: NormalExpr */
			token.T_12:reduce(11),		/* generalizes, reduce: NormalExpr */
			token.T_13:reduce(11),		/* handles, reduce: NormalExpr */
			token.T_14:reduce(11),		/* implements, reduce: NormalExpr */
			token.T_17:reduce(11),		/* owns, reduce: NormalExpr */
			token.T_18:reduce(11),		/* relates, reduce: NormalExpr */
			token.T_19:reduce(11),		/* requires, reduce: NormalExpr */
			token.T_21:reduce(11),		/* triggers, reduce: NormalExpr */
        },

	},
	actionRow{ // S46
        canRecover: false,
		actions: map[token.Type]action{ 
			token.EOF:reduce(9),		/* $, reduce: NormalExpr */
			token.T_1:shift(19),		/* , */
			token.T_2:reduce(9),		/* -<, reduce: NormalExpr */
			token.T_3:reduce(9),		/* ->, reduce: NormalExpr */
			token.T_4:shift(22),		/* :: */
			token.T_5:reduce(9),		/* <-, reduce: NormalExpr */
			token.T_6:reduce(9),		/* >-, reduce: NormalExpr */
			token.T_7:reduce(9),		/* aggregates, reduce: NormalExpr */
			token.T_11:reduce(9),		/* exhibits, reduce: NormalExpr */
			token.T_12:reduce(9),		/* generalizes, reduce: NormalExpr */
			token.T_13:reduce(9),		/* handles, reduce: NormalExpr */
			token.T_14:reduce(9),		/* implements, reduce: NormalExpr */
			token.T_17:reduce(9),		/* owns, reduce: NormalExpr */
			token.T_18:reduce(9),		/* relates, reduce: NormalExpr */
			token.T_19:reduce(9),		/* requires, reduce: NormalExpr */
			token.T_21:reduce(9),		/* triggers, reduce: NormalExpr */
        },

	},
	actionRow{ // S47
        canRecover: false,
		actions: map[token.Type]action{ 
			token.EOF:reduce(18),		/* $, reduce: KeyExpr */
			token.T_1:reduce(18),		/* ,, reduce: KeyExpr */
			token.T_2:reduce(18),		/* -<, reduce: KeyExpr */
			token.T_3:reduce(18),		/* ->, reduce: KeyExpr */
			token.T_4:reduce(18),		/* ::, reduce: KeyExpr */
			token.T_5:reduce(18),		/* <-, reduce: KeyExpr */
			token.T_6:reduce(18),		/* >-, reduce: KeyExpr */
			token.T_7:reduce(18),		/* aggregates, reduce: KeyExpr */
			token.T_11:reduce(18),		/* exhibits, reduce: KeyExpr */
			token.T_12:reduce(18),		/* generalizes, reduce: KeyExpr */
			token.T_13:reduce(18),		/* handles, reduce: KeyExpr */
			token.T_14:reduce(18),		/* implements, reduce: KeyExpr */
			token.T_17:reduce(18),		/* owns, reduce: KeyExpr */
			token.T_18:reduce(18),		/* relates, reduce: KeyExpr */
			token.T_19:reduce(18),		/* requires, reduce: KeyExpr */
			token.T_21:reduce(18),		/* triggers, reduce: KeyExpr */
        },

	},
	actionRow{ // S48
        canRecover: false,
		actions: map[token.Type]action{ 
			token.EOF:reduce(16),		/* $, reduce: KeyExpr */
			token.T_1:reduce(16),		/* ,, reduce: KeyExpr */
			token.T_2:reduce(16),		/* -<, reduce: KeyExpr */
			token.T_3:reduce(16),		/* ->, reduce: KeyExpr */
			token.T_4:reduce(16),		/* ::, reduce: KeyExpr */
			token.T_5:reduce(16),		/* <-, reduce: KeyExpr */
			token.T_6:reduce(16),		/* >-, reduce: KeyExpr */
			token.T_7:reduce(16),		/* aggregates, reduce: KeyExpr */
			token.T_11:reduce(16),		/* exhibits, reduce: KeyExpr */
			token.T_12:reduce(16),		/* generalizes, reduce: KeyExpr */
			token.T_13:reduce(16),		/* handles, reduce: KeyExpr */
			token.T_14:reduce(16),		/* implements, reduce: KeyExpr */
			token.T_17:reduce(16),		/* owns, reduce: KeyExpr */
			token.T_18:reduce(16),		/* relates, reduce: KeyExpr */
			token.T_19:reduce(16),		/* requires, reduce: KeyExpr */
			token.T_21:reduce(16),		/* triggers, reduce: KeyExpr */
        },

	},
	actionRow{ // S49
        canRecover: false,
		actions: map[token.Type]action{ 
			token.EOF:reduce(13),		/* $, reduce: NormalExpr */
			token.T_1:shift(19),		/* , */
			token.T_2:reduce(13),		/* -<, reduce: NormalExpr */
			token.T_3:reduce(13),		/* ->, reduce: NormalExpr */
			token.T_4:shift(22),		/* :: */
			token.T_5:reduce(13),		/* <-, reduce: NormalExpr */
			token.T_6:reduce(13),		/* >-, reduce: NormalExpr */
			token.T_7:reduce(13),		/* aggregates, reduce: NormalExpr */
			token.T_11:reduce(13),		/* exhibits, reduce: NormalExpr */
			token.T_12:reduce(13),		/* generalizes, reduce: NormalExpr */
			token.T_13:reduce(13),		/* handles, reduce: NormalExpr */
			token.T_14:reduce(13),		/* implements, reduce: NormalExpr */
			token.T_17:reduce(13),		/* owns, reduce: NormalExpr */
			token.T_18:reduce(13),		/* relates, reduce: NormalExpr */
			token.T_19:reduce(13),		/* requires, reduce: NormalExpr */
			token.T_21:reduce(13),		/* triggers, reduce: NormalExpr */
        },

	},
	actionRow{ // S50
        canRecover: false,
		actions: map[token.Type]action{ 
			token.EOF:reduce(14),		/* $, reduce: NormalExpr */
			token.T_1:shift(19),		/* , */
			token.T_2:reduce(14),		/* -<, reduce: NormalExpr */
			token.T_3:reduce(14),		/* ->, reduce: NormalExpr */
			token.T_4:shift(22),		/* :: */
			token.T_5:reduce(14),		/* <-, reduce: NormalExpr */
			token.T_6:reduce(14),		/* >-, reduce: NormalExpr */
			token.T_7:reduce(14),		/* aggregates, reduce: NormalExpr */
			token.T_11:reduce(14),		/* exhibits, reduce: NormalExpr */
			token.T_12:reduce(14),		/* generalizes, reduce: NormalExpr */
			token.T_13:reduce(14),		/* handles, reduce: NormalExpr */
			token.T_14:reduce(14),		/* implements, reduce: NormalExpr */
			token.T_17:reduce(14),		/* owns, reduce: NormalExpr */
			token.T_18:reduce(14),		/* relates, reduce: NormalExpr */
			token.T_19:reduce(14),		/* requires, reduce: NormalExpr */
			token.T_21:reduce(14),		/* triggers, reduce: NormalExpr */
        },

	},
	actionRow{ // S51
        canRecover: false,
		actions: map[token.Type]action{ 
			token.EOF:reduce(2),		/* $, reduce: EndExpr */
        },

	},
	actionRow{ // S52
        canRecover: false,
		actions: map[token.Type]action{ 
			token.EOF:reduce(4),		/* $, reduce: EndExpr */
			token.T_1:shift(19),		/* , */
			token.T_4:shift(22),		/* :: */
        },

	},
	actionRow{ // S53
        canRecover: false,
		actions: map[token.Type]action{ 
			token.EOF:reduce(3),		/* $, reduce: EndExpr */
			token.T_1:shift(19),		/* , */
			token.T_4:shift(22),		/* :: */
        },

	},
	actionRow{ // S54
        canRecover: false,
		actions: map[token.Type]action{ 
			token.EOF:reduce(12),		/* $, reduce: NormalExpr */
			token.T_1:shift(19),		/* , */
			token.T_2:reduce(12),		/* -<, reduce: NormalExpr */
			token.T_3:reduce(12),		/* ->, reduce: NormalExpr */
			token.T_4:shift(22),		/* :: */
			token.T_5:reduce(12),		/* <-, reduce: NormalExpr */
			token.T_6:reduce(12),		/* >-, reduce: NormalExpr */
			token.T_7:reduce(12),		/* aggregates, reduce: NormalExpr */
			token.T_11:reduce(12),		/* exhibits, reduce: NormalExpr */
			token.T_12:reduce(12),		/* generalizes, reduce: NormalExpr */
			token.T_13:reduce(12),		/* handles, reduce: NormalExpr */
			token.T_14:reduce(12),		/* implements, reduce: NormalExpr */
			token.T_17:reduce(12),		/* owns, reduce: NormalExpr */
			token.T_18:reduce(12),		/* relates, reduce: NormalExpr */
			token.T_19:reduce(12),		/* requires, reduce: NormalExpr */
			token.T_21:reduce(12),		/* triggers, reduce: NormalExpr */
        },

	},
	actionRow{ // S55
        canRecover: false,
		actions: map[token.Type]action{ 
			token.EOF:reduce(10),		/* $, reduce: NormalExpr */
			token.T_1:shift(19),		/* , */
			token.T_2:reduce(10),		/* -<, reduce: NormalExpr */
			token.T_3:reduce(10),		/* ->, reduce: NormalExpr */
			token.T_4:shift(22),		/* :: */
			token.T_5:reduce(10),		/* <-, reduce: NormalExpr */
			token.T_6:reduce(10),		/* >-, reduce: NormalExpr */
			token.T_7:reduce(10),		/* aggregates, reduce: NormalExpr */
			token.T_11:reduce(10),		/* exhibits, reduce: NormalExpr */
			token.T_12:reduce(10),		/* generalizes, reduce: NormalExpr */
			token.T_13:reduce(10),		/* handles, reduce: NormalExpr */
			token.T_14:reduce(10),		/* implements, reduce: NormalExpr */
			token.T_17:reduce(10),		/* owns, reduce: NormalExpr */
			token.T_18:reduce(10),		/* relates, reduce: NormalExpr */
			token.T_19:reduce(10),		/* requires, reduce: NormalExpr */
			token.T_21:reduce(10),		/* triggers, reduce: NormalExpr */
        },

	},
	actionRow{ // S56
        canRecover: false,
		actions: map[token.Type]action{ 
			token.EOF:reduce(17),		/* $, reduce: KeyExpr */
			token.T_1:reduce(17),		/* ,, reduce: KeyExpr */
			token.T_2:reduce(17),		/* -<, reduce: KeyExpr */
			token.T_3:reduce(17),		/* ->, reduce: KeyExpr */
			token.T_4:reduce(17),		/* ::, reduce: KeyExpr */
			token.T_5:reduce(17),		/* <-, reduce: KeyExpr */
			token.T_6:reduce(17),		/* >-, reduce: KeyExpr */
			token.T_7:reduce(17),		/* aggregates, reduce: KeyExpr */
			token.T_11:reduce(17),		/* exhibits, reduce: KeyExpr */
			token.T_12:reduce(17),		/* generalizes, reduce: KeyExpr */
			token.T_13:reduce(17),		/* handles, reduce: KeyExpr */
			token.T_14:reduce(17),		/* implements, reduce: KeyExpr */
			token.T_17:reduce(17),		/* owns, reduce: KeyExpr */
			token.T_18:reduce(17),		/* relates, reduce: KeyExpr */
			token.T_19:reduce(17),		/* requires, reduce: KeyExpr */
			token.T_21:reduce(17),		/* triggers, reduce: KeyExpr */
        },

	},
	actionRow{ // S57
        canRecover: false,
		actions: map[token.Type]action{ 
			token.EOF:reduce(15),		/* $, reduce: KeyExpr */
			token.T_1:reduce(15),		/* ,, reduce: KeyExpr */
			token.T_2:reduce(15),		/* -<, reduce: KeyExpr */
			token.T_3:reduce(15),		/* ->, reduce: KeyExpr */
			token.T_4:reduce(15),		/* ::, reduce: KeyExpr */
			token.T_5:reduce(15),		/* <-, reduce: KeyExpr */
			token.T_6:reduce(15),		/* >-, reduce: KeyExpr */
			token.T_7:reduce(15),		/* aggregates, reduce: KeyExpr */
			token.T_11:reduce(15),		/* exhibits, reduce: KeyExpr */
			token.T_12:reduce(15),		/* generalizes, reduce: KeyExpr */
			token.T_13:reduce(15),		/* handles, reduce: KeyExpr */
			token.T_14:reduce(15),		/* implements, reduce: KeyExpr */
			token.T_17:reduce(15),		/* owns, reduce: KeyExpr */
			token.T_18:reduce(15),		/* relates, reduce: KeyExpr */
			token.T_19:reduce(15),		/* requires, reduce: KeyExpr */
			token.T_21:reduce(15),		/* triggers, reduce: KeyExpr */
        },

	},
	actionRow{ // S58
        canRecover: false,
		actions: map[token.Type]action{ 
			token.EOF:reduce(5),		/* $, reduce: EndExpr */
        },

	},
}

