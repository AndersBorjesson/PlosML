
/*
*/
package parser

const numNTSymbols = 3
type(
	gotoTable [numStates]gotoRow
	gotoRow	[numNTSymbols] int
)

var gotoTab = gotoTable{
	gotoRow{ // S0
		1, // Expr
        2, // Start
        -1, // Stop
        
	},
	gotoRow{ // S1
		-1, // Expr
        -1, // Start
        -1, // Stop
        
	},
	gotoRow{ // S2
		-1, // Expr
        -1, // Start
        -1, // Stop
        
	},
	gotoRow{ // S3
		-1, // Expr
        -1, // Start
        -1, // Stop
        
	},
	gotoRow{ // S4
		-1, // Expr
        -1, // Start
        5, // Stop
        
	},
	gotoRow{ // S5
		-1, // Expr
        -1, // Start
        -1, // Stop
        
	},
	gotoRow{ // S6
		-1, // Expr
        -1, // Start
        -1, // Stop
        
	},
	
}
