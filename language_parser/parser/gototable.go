
/*
*/
package parser

const numNTSymbols = 15
type(
	gotoTable [numStates]gotoRow
	gotoRow	[numNTSymbols] int
)

var gotoTab = gotoTable{
	gotoRow{ // S0
		-1, // Alias
        -1, // Attribute
        1, // Call
        -1, // Dualsidedkey
        2, // EndExpr
        3, // KeyExpr
        4, // NormalExpr
        -1, // Pend
        5, // Pragma
        6, // PragmaExpr
        -1, // Property
        -1, // Pstart
        -1, // Separator
        7, // Structural
        8, // StructuralExpr
        
	},
	gotoRow{ // S1
		-1, // Alias
        -1, // Attribute
        -1, // Call
        -1, // Dualsidedkey
        -1, // EndExpr
        -1, // KeyExpr
        -1, // NormalExpr
        -1, // Pend
        -1, // Pragma
        -1, // PragmaExpr
        -1, // Property
        -1, // Pstart
        -1, // Separator
        -1, // Structural
        -1, // StructuralExpr
        
	},
	gotoRow{ // S2
		-1, // Alias
        -1, // Attribute
        -1, // Call
        -1, // Dualsidedkey
        -1, // EndExpr
        -1, // KeyExpr
        -1, // NormalExpr
        -1, // Pend
        -1, // Pragma
        -1, // PragmaExpr
        -1, // Property
        -1, // Pstart
        -1, // Separator
        -1, // Structural
        -1, // StructuralExpr
        
	},
	gotoRow{ // S3
		-1, // Alias
        -1, // Attribute
        -1, // Call
        16, // Dualsidedkey
        -1, // EndExpr
        -1, // KeyExpr
        -1, // NormalExpr
        -1, // Pend
        -1, // Pragma
        -1, // PragmaExpr
        17, // Property
        -1, // Pstart
        18, // Separator
        -1, // Structural
        -1, // StructuralExpr
        
	},
	gotoRow{ // S4
		-1, // Alias
        -1, // Attribute
        -1, // Call
        34, // Dualsidedkey
        -1, // EndExpr
        -1, // KeyExpr
        -1, // NormalExpr
        -1, // Pend
        -1, // Pragma
        -1, // PragmaExpr
        -1, // Property
        -1, // Pstart
        -1, // Separator
        -1, // Structural
        -1, // StructuralExpr
        
	},
	gotoRow{ // S5
		-1, // Alias
        -1, // Attribute
        -1, // Call
        -1, // Dualsidedkey
        -1, // EndExpr
        -1, // KeyExpr
        -1, // NormalExpr
        -1, // Pend
        -1, // Pragma
        -1, // PragmaExpr
        -1, // Property
        -1, // Pstart
        -1, // Separator
        -1, // Structural
        -1, // StructuralExpr
        
	},
	gotoRow{ // S6
		-1, // Alias
        -1, // Attribute
        -1, // Call
        -1, // Dualsidedkey
        -1, // EndExpr
        -1, // KeyExpr
        -1, // NormalExpr
        -1, // Pend
        -1, // Pragma
        -1, // PragmaExpr
        -1, // Property
        -1, // Pstart
        -1, // Separator
        -1, // Structural
        -1, // StructuralExpr
        
	},
	gotoRow{ // S7
		-1, // Alias
        -1, // Attribute
        -1, // Call
        -1, // Dualsidedkey
        -1, // EndExpr
        -1, // KeyExpr
        -1, // NormalExpr
        -1, // Pend
        -1, // Pragma
        -1, // PragmaExpr
        -1, // Property
        -1, // Pstart
        -1, // Separator
        -1, // Structural
        -1, // StructuralExpr
        
	},
	gotoRow{ // S8
		-1, // Alias
        -1, // Attribute
        -1, // Call
        -1, // Dualsidedkey
        -1, // EndExpr
        -1, // KeyExpr
        -1, // NormalExpr
        -1, // Pend
        -1, // Pragma
        -1, // PragmaExpr
        -1, // Property
        -1, // Pstart
        -1, // Separator
        -1, // Structural
        -1, // StructuralExpr
        
	},
	gotoRow{ // S9
		-1, // Alias
        -1, // Attribute
        -1, // Call
        -1, // Dualsidedkey
        -1, // EndExpr
        -1, // KeyExpr
        -1, // NormalExpr
        -1, // Pend
        -1, // Pragma
        -1, // PragmaExpr
        -1, // Property
        -1, // Pstart
        -1, // Separator
        -1, // Structural
        -1, // StructuralExpr
        
	},
	gotoRow{ // S10
		-1, // Alias
        -1, // Attribute
        -1, // Call
        -1, // Dualsidedkey
        -1, // EndExpr
        -1, // KeyExpr
        -1, // NormalExpr
        -1, // Pend
        -1, // Pragma
        -1, // PragmaExpr
        -1, // Property
        -1, // Pstart
        -1, // Separator
        -1, // Structural
        -1, // StructuralExpr
        
	},
	gotoRow{ // S11
		-1, // Alias
        -1, // Attribute
        -1, // Call
        -1, // Dualsidedkey
        -1, // EndExpr
        -1, // KeyExpr
        -1, // NormalExpr
        -1, // Pend
        -1, // Pragma
        -1, // PragmaExpr
        -1, // Property
        -1, // Pstart
        -1, // Separator
        -1, // Structural
        -1, // StructuralExpr
        
	},
	gotoRow{ // S12
		37, // Alias
        38, // Attribute
        -1, // Call
        39, // Dualsidedkey
        -1, // EndExpr
        -1, // KeyExpr
        -1, // NormalExpr
        -1, // Pend
        -1, // Pragma
        -1, // PragmaExpr
        40, // Property
        -1, // Pstart
        41, // Separator
        -1, // Structural
        -1, // StructuralExpr
        
	},
	gotoRow{ // S13
		-1, // Alias
        -1, // Attribute
        -1, // Call
        -1, // Dualsidedkey
        -1, // EndExpr
        -1, // KeyExpr
        -1, // NormalExpr
        -1, // Pend
        -1, // Pragma
        -1, // PragmaExpr
        -1, // Property
        -1, // Pstart
        -1, // Separator
        -1, // Structural
        -1, // StructuralExpr
        
	},
	gotoRow{ // S14
		-1, // Alias
        -1, // Attribute
        -1, // Call
        -1, // Dualsidedkey
        -1, // EndExpr
        -1, // KeyExpr
        -1, // NormalExpr
        -1, // Pend
        -1, // Pragma
        -1, // PragmaExpr
        -1, // Property
        -1, // Pstart
        -1, // Separator
        -1, // Structural
        -1, // StructuralExpr
        
	},
	gotoRow{ // S15
		44, // Alias
        -1, // Attribute
        -1, // Call
        -1, // Dualsidedkey
        -1, // EndExpr
        -1, // KeyExpr
        -1, // NormalExpr
        -1, // Pend
        -1, // Pragma
        -1, // PragmaExpr
        -1, // Property
        -1, // Pstart
        -1, // Separator
        -1, // Structural
        -1, // StructuralExpr
        
	},
	gotoRow{ // S16
		-1, // Alias
        -1, // Attribute
        -1, // Call
        -1, // Dualsidedkey
        -1, // EndExpr
        45, // KeyExpr
        -1, // NormalExpr
        -1, // Pend
        -1, // Pragma
        -1, // PragmaExpr
        -1, // Property
        -1, // Pstart
        -1, // Separator
        -1, // Structural
        -1, // StructuralExpr
        
	},
	gotoRow{ // S17
		-1, // Alias
        -1, // Attribute
        -1, // Call
        -1, // Dualsidedkey
        -1, // EndExpr
        -1, // KeyExpr
        -1, // NormalExpr
        -1, // Pend
        -1, // Pragma
        -1, // PragmaExpr
        -1, // Property
        -1, // Pstart
        -1, // Separator
        -1, // Structural
        -1, // StructuralExpr
        
	},
	gotoRow{ // S18
		-1, // Alias
        -1, // Attribute
        -1, // Call
        -1, // Dualsidedkey
        -1, // EndExpr
        -1, // KeyExpr
        -1, // NormalExpr
        -1, // Pend
        -1, // Pragma
        -1, // PragmaExpr
        -1, // Property
        -1, // Pstart
        -1, // Separator
        -1, // Structural
        -1, // StructuralExpr
        
	},
	gotoRow{ // S19
		-1, // Alias
        -1, // Attribute
        -1, // Call
        -1, // Dualsidedkey
        -1, // EndExpr
        -1, // KeyExpr
        -1, // NormalExpr
        -1, // Pend
        -1, // Pragma
        -1, // PragmaExpr
        -1, // Property
        -1, // Pstart
        -1, // Separator
        -1, // Structural
        -1, // StructuralExpr
        
	},
	gotoRow{ // S20
		-1, // Alias
        -1, // Attribute
        -1, // Call
        -1, // Dualsidedkey
        -1, // EndExpr
        -1, // KeyExpr
        -1, // NormalExpr
        -1, // Pend
        -1, // Pragma
        -1, // PragmaExpr
        -1, // Property
        -1, // Pstart
        -1, // Separator
        -1, // Structural
        -1, // StructuralExpr
        
	},
	gotoRow{ // S21
		-1, // Alias
        -1, // Attribute
        -1, // Call
        -1, // Dualsidedkey
        -1, // EndExpr
        -1, // KeyExpr
        -1, // NormalExpr
        -1, // Pend
        -1, // Pragma
        -1, // PragmaExpr
        -1, // Property
        -1, // Pstart
        -1, // Separator
        -1, // Structural
        -1, // StructuralExpr
        
	},
	gotoRow{ // S22
		-1, // Alias
        -1, // Attribute
        -1, // Call
        -1, // Dualsidedkey
        -1, // EndExpr
        -1, // KeyExpr
        -1, // NormalExpr
        -1, // Pend
        -1, // Pragma
        -1, // PragmaExpr
        -1, // Property
        -1, // Pstart
        -1, // Separator
        -1, // Structural
        -1, // StructuralExpr
        
	},
	gotoRow{ // S23
		-1, // Alias
        -1, // Attribute
        -1, // Call
        -1, // Dualsidedkey
        -1, // EndExpr
        -1, // KeyExpr
        -1, // NormalExpr
        -1, // Pend
        -1, // Pragma
        -1, // PragmaExpr
        -1, // Property
        -1, // Pstart
        -1, // Separator
        -1, // Structural
        -1, // StructuralExpr
        
	},
	gotoRow{ // S24
		-1, // Alias
        -1, // Attribute
        -1, // Call
        -1, // Dualsidedkey
        -1, // EndExpr
        -1, // KeyExpr
        -1, // NormalExpr
        -1, // Pend
        -1, // Pragma
        -1, // PragmaExpr
        -1, // Property
        -1, // Pstart
        -1, // Separator
        -1, // Structural
        -1, // StructuralExpr
        
	},
	gotoRow{ // S25
		-1, // Alias
        -1, // Attribute
        -1, // Call
        -1, // Dualsidedkey
        -1, // EndExpr
        -1, // KeyExpr
        -1, // NormalExpr
        -1, // Pend
        -1, // Pragma
        -1, // PragmaExpr
        -1, // Property
        -1, // Pstart
        -1, // Separator
        -1, // Structural
        -1, // StructuralExpr
        
	},
	gotoRow{ // S26
		-1, // Alias
        -1, // Attribute
        -1, // Call
        -1, // Dualsidedkey
        -1, // EndExpr
        -1, // KeyExpr
        -1, // NormalExpr
        -1, // Pend
        -1, // Pragma
        -1, // PragmaExpr
        -1, // Property
        -1, // Pstart
        -1, // Separator
        -1, // Structural
        -1, // StructuralExpr
        
	},
	gotoRow{ // S27
		-1, // Alias
        -1, // Attribute
        -1, // Call
        -1, // Dualsidedkey
        -1, // EndExpr
        -1, // KeyExpr
        -1, // NormalExpr
        -1, // Pend
        -1, // Pragma
        -1, // PragmaExpr
        -1, // Property
        -1, // Pstart
        -1, // Separator
        -1, // Structural
        -1, // StructuralExpr
        
	},
	gotoRow{ // S28
		-1, // Alias
        -1, // Attribute
        -1, // Call
        -1, // Dualsidedkey
        -1, // EndExpr
        -1, // KeyExpr
        -1, // NormalExpr
        -1, // Pend
        -1, // Pragma
        -1, // PragmaExpr
        -1, // Property
        -1, // Pstart
        -1, // Separator
        -1, // Structural
        -1, // StructuralExpr
        
	},
	gotoRow{ // S29
		-1, // Alias
        -1, // Attribute
        -1, // Call
        -1, // Dualsidedkey
        -1, // EndExpr
        -1, // KeyExpr
        -1, // NormalExpr
        -1, // Pend
        -1, // Pragma
        -1, // PragmaExpr
        -1, // Property
        -1, // Pstart
        -1, // Separator
        -1, // Structural
        -1, // StructuralExpr
        
	},
	gotoRow{ // S30
		-1, // Alias
        -1, // Attribute
        -1, // Call
        -1, // Dualsidedkey
        -1, // EndExpr
        -1, // KeyExpr
        -1, // NormalExpr
        -1, // Pend
        -1, // Pragma
        -1, // PragmaExpr
        -1, // Property
        -1, // Pstart
        -1, // Separator
        -1, // Structural
        -1, // StructuralExpr
        
	},
	gotoRow{ // S31
		-1, // Alias
        -1, // Attribute
        -1, // Call
        -1, // Dualsidedkey
        -1, // EndExpr
        -1, // KeyExpr
        -1, // NormalExpr
        -1, // Pend
        -1, // Pragma
        -1, // PragmaExpr
        -1, // Property
        -1, // Pstart
        -1, // Separator
        -1, // Structural
        -1, // StructuralExpr
        
	},
	gotoRow{ // S32
		-1, // Alias
        -1, // Attribute
        -1, // Call
        -1, // Dualsidedkey
        -1, // EndExpr
        -1, // KeyExpr
        -1, // NormalExpr
        -1, // Pend
        -1, // Pragma
        -1, // PragmaExpr
        -1, // Property
        -1, // Pstart
        -1, // Separator
        -1, // Structural
        -1, // StructuralExpr
        
	},
	gotoRow{ // S33
		-1, // Alias
        -1, // Attribute
        -1, // Call
        -1, // Dualsidedkey
        -1, // EndExpr
        -1, // KeyExpr
        -1, // NormalExpr
        -1, // Pend
        -1, // Pragma
        -1, // PragmaExpr
        -1, // Property
        -1, // Pstart
        -1, // Separator
        -1, // Structural
        -1, // StructuralExpr
        
	},
	gotoRow{ // S34
		-1, // Alias
        -1, // Attribute
        -1, // Call
        -1, // Dualsidedkey
        -1, // EndExpr
        49, // KeyExpr
        -1, // NormalExpr
        -1, // Pend
        -1, // Pragma
        -1, // PragmaExpr
        -1, // Property
        -1, // Pstart
        -1, // Separator
        -1, // Structural
        -1, // StructuralExpr
        
	},
	gotoRow{ // S35
		-1, // Alias
        -1, // Attribute
        -1, // Call
        -1, // Dualsidedkey
        -1, // EndExpr
        -1, // KeyExpr
        -1, // NormalExpr
        -1, // Pend
        -1, // Pragma
        -1, // PragmaExpr
        -1, // Property
        -1, // Pstart
        -1, // Separator
        -1, // Structural
        -1, // StructuralExpr
        
	},
	gotoRow{ // S36
		-1, // Alias
        -1, // Attribute
        -1, // Call
        -1, // Dualsidedkey
        -1, // EndExpr
        -1, // KeyExpr
        -1, // NormalExpr
        -1, // Pend
        -1, // Pragma
        -1, // PragmaExpr
        -1, // Property
        -1, // Pstart
        -1, // Separator
        -1, // Structural
        -1, // StructuralExpr
        
	},
	gotoRow{ // S37
		-1, // Alias
        -1, // Attribute
        -1, // Call
        -1, // Dualsidedkey
        -1, // EndExpr
        -1, // KeyExpr
        -1, // NormalExpr
        -1, // Pend
        -1, // Pragma
        -1, // PragmaExpr
        -1, // Property
        -1, // Pstart
        -1, // Separator
        -1, // Structural
        -1, // StructuralExpr
        
	},
	gotoRow{ // S38
		-1, // Alias
        -1, // Attribute
        -1, // Call
        -1, // Dualsidedkey
        -1, // EndExpr
        52, // KeyExpr
        -1, // NormalExpr
        -1, // Pend
        -1, // Pragma
        -1, // PragmaExpr
        -1, // Property
        -1, // Pstart
        -1, // Separator
        -1, // Structural
        -1, // StructuralExpr
        
	},
	gotoRow{ // S39
		-1, // Alias
        -1, // Attribute
        -1, // Call
        -1, // Dualsidedkey
        -1, // EndExpr
        54, // KeyExpr
        -1, // NormalExpr
        -1, // Pend
        -1, // Pragma
        -1, // PragmaExpr
        -1, // Property
        -1, // Pstart
        -1, // Separator
        -1, // Structural
        -1, // StructuralExpr
        
	},
	gotoRow{ // S40
		-1, // Alias
        -1, // Attribute
        -1, // Call
        -1, // Dualsidedkey
        -1, // EndExpr
        -1, // KeyExpr
        -1, // NormalExpr
        -1, // Pend
        -1, // Pragma
        -1, // PragmaExpr
        -1, // Property
        -1, // Pstart
        -1, // Separator
        -1, // Structural
        -1, // StructuralExpr
        
	},
	gotoRow{ // S41
		-1, // Alias
        -1, // Attribute
        -1, // Call
        -1, // Dualsidedkey
        -1, // EndExpr
        -1, // KeyExpr
        -1, // NormalExpr
        -1, // Pend
        -1, // Pragma
        -1, // PragmaExpr
        -1, // Property
        -1, // Pstart
        -1, // Separator
        -1, // Structural
        -1, // StructuralExpr
        
	},
	gotoRow{ // S42
		-1, // Alias
        -1, // Attribute
        -1, // Call
        -1, // Dualsidedkey
        -1, // EndExpr
        -1, // KeyExpr
        -1, // NormalExpr
        -1, // Pend
        -1, // Pragma
        -1, // PragmaExpr
        -1, // Property
        -1, // Pstart
        -1, // Separator
        -1, // Structural
        -1, // StructuralExpr
        
	},
	gotoRow{ // S43
		-1, // Alias
        -1, // Attribute
        -1, // Call
        -1, // Dualsidedkey
        -1, // EndExpr
        -1, // KeyExpr
        -1, // NormalExpr
        -1, // Pend
        -1, // Pragma
        -1, // PragmaExpr
        -1, // Property
        -1, // Pstart
        -1, // Separator
        -1, // Structural
        -1, // StructuralExpr
        
	},
	gotoRow{ // S44
		-1, // Alias
        -1, // Attribute
        -1, // Call
        -1, // Dualsidedkey
        -1, // EndExpr
        -1, // KeyExpr
        -1, // NormalExpr
        -1, // Pend
        -1, // Pragma
        -1, // PragmaExpr
        -1, // Property
        -1, // Pstart
        -1, // Separator
        -1, // Structural
        -1, // StructuralExpr
        
	},
	gotoRow{ // S45
		-1, // Alias
        -1, // Attribute
        -1, // Call
        -1, // Dualsidedkey
        -1, // EndExpr
        -1, // KeyExpr
        -1, // NormalExpr
        -1, // Pend
        -1, // Pragma
        -1, // PragmaExpr
        17, // Property
        -1, // Pstart
        18, // Separator
        -1, // Structural
        -1, // StructuralExpr
        
	},
	gotoRow{ // S46
		-1, // Alias
        -1, // Attribute
        -1, // Call
        -1, // Dualsidedkey
        -1, // EndExpr
        -1, // KeyExpr
        -1, // NormalExpr
        -1, // Pend
        -1, // Pragma
        -1, // PragmaExpr
        40, // Property
        -1, // Pstart
        41, // Separator
        -1, // Structural
        -1, // StructuralExpr
        
	},
	gotoRow{ // S47
		-1, // Alias
        -1, // Attribute
        -1, // Call
        -1, // Dualsidedkey
        -1, // EndExpr
        -1, // KeyExpr
        -1, // NormalExpr
        -1, // Pend
        -1, // Pragma
        -1, // PragmaExpr
        -1, // Property
        -1, // Pstart
        -1, // Separator
        -1, // Structural
        -1, // StructuralExpr
        
	},
	gotoRow{ // S48
		-1, // Alias
        -1, // Attribute
        -1, // Call
        -1, // Dualsidedkey
        -1, // EndExpr
        -1, // KeyExpr
        -1, // NormalExpr
        -1, // Pend
        -1, // Pragma
        -1, // PragmaExpr
        -1, // Property
        -1, // Pstart
        -1, // Separator
        -1, // Structural
        -1, // StructuralExpr
        
	},
	gotoRow{ // S49
		-1, // Alias
        -1, // Attribute
        -1, // Call
        -1, // Dualsidedkey
        -1, // EndExpr
        -1, // KeyExpr
        -1, // NormalExpr
        -1, // Pend
        -1, // Pragma
        -1, // PragmaExpr
        17, // Property
        -1, // Pstart
        18, // Separator
        -1, // Structural
        -1, // StructuralExpr
        
	},
	gotoRow{ // S50
		-1, // Alias
        -1, // Attribute
        -1, // Call
        -1, // Dualsidedkey
        -1, // EndExpr
        -1, // KeyExpr
        -1, // NormalExpr
        -1, // Pend
        -1, // Pragma
        -1, // PragmaExpr
        40, // Property
        -1, // Pstart
        41, // Separator
        -1, // Structural
        -1, // StructuralExpr
        
	},
	gotoRow{ // S51
		-1, // Alias
        -1, // Attribute
        -1, // Call
        -1, // Dualsidedkey
        -1, // EndExpr
        -1, // KeyExpr
        -1, // NormalExpr
        -1, // Pend
        -1, // Pragma
        -1, // PragmaExpr
        -1, // Property
        -1, // Pstart
        -1, // Separator
        -1, // Structural
        -1, // StructuralExpr
        
	},
	gotoRow{ // S52
		-1, // Alias
        -1, // Attribute
        -1, // Call
        -1, // Dualsidedkey
        -1, // EndExpr
        -1, // KeyExpr
        -1, // NormalExpr
        -1, // Pend
        -1, // Pragma
        -1, // PragmaExpr
        17, // Property
        -1, // Pstart
        18, // Separator
        -1, // Structural
        -1, // StructuralExpr
        
	},
	gotoRow{ // S53
		-1, // Alias
        -1, // Attribute
        -1, // Call
        -1, // Dualsidedkey
        -1, // EndExpr
        -1, // KeyExpr
        -1, // NormalExpr
        -1, // Pend
        -1, // Pragma
        -1, // PragmaExpr
        40, // Property
        -1, // Pstart
        41, // Separator
        -1, // Structural
        -1, // StructuralExpr
        
	},
	gotoRow{ // S54
		-1, // Alias
        -1, // Attribute
        -1, // Call
        -1, // Dualsidedkey
        -1, // EndExpr
        -1, // KeyExpr
        -1, // NormalExpr
        -1, // Pend
        -1, // Pragma
        -1, // PragmaExpr
        17, // Property
        -1, // Pstart
        18, // Separator
        -1, // Structural
        -1, // StructuralExpr
        
	},
	gotoRow{ // S55
		-1, // Alias
        -1, // Attribute
        -1, // Call
        -1, // Dualsidedkey
        -1, // EndExpr
        -1, // KeyExpr
        -1, // NormalExpr
        -1, // Pend
        -1, // Pragma
        -1, // PragmaExpr
        40, // Property
        -1, // Pstart
        41, // Separator
        -1, // Structural
        -1, // StructuralExpr
        
	},
	gotoRow{ // S56
		-1, // Alias
        -1, // Attribute
        -1, // Call
        -1, // Dualsidedkey
        -1, // EndExpr
        -1, // KeyExpr
        -1, // NormalExpr
        -1, // Pend
        -1, // Pragma
        -1, // PragmaExpr
        -1, // Property
        -1, // Pstart
        -1, // Separator
        -1, // Structural
        -1, // StructuralExpr
        
	},
	gotoRow{ // S57
		-1, // Alias
        -1, // Attribute
        -1, // Call
        -1, // Dualsidedkey
        -1, // EndExpr
        -1, // KeyExpr
        -1, // NormalExpr
        -1, // Pend
        -1, // Pragma
        -1, // PragmaExpr
        -1, // Property
        -1, // Pstart
        -1, // Separator
        -1, // Structural
        -1, // StructuralExpr
        
	},
	gotoRow{ // S58
		-1, // Alias
        -1, // Attribute
        -1, // Call
        -1, // Dualsidedkey
        -1, // EndExpr
        -1, // KeyExpr
        -1, // NormalExpr
        -1, // Pend
        -1, // Pragma
        -1, // PragmaExpr
        -1, // Property
        -1, // Pstart
        -1, // Separator
        -1, // Structural
        -1, // StructuralExpr
        
	},
	
}
