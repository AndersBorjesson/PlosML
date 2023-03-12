# PlosML - The PloshML System Modelling Language 

This repository contains the PloshML compiler. The compiler transforms an input structure, being the full system representation, to a number of views of the system. The modelling technique is loosly based on Object Process Methodology.

The language and compiler are based on a few basic statements : 

- The language should be fast and efficient to write in order for efficient work. 
- The language input should represent the full structural, logical and functional views of the system. Various views are derived from this one structure.
- The compiler should validate that the model is coherent and consistent. 
- Language simplicity is favoured over exspressiveness.


## Usage 

An input can be compiled via the following command
```bash
./ploshml -infile ploshml_input.opa
```


## Language 