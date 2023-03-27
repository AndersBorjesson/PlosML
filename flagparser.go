package main

import (
	"flag"
	"log"
	"os"
	"path"
)

type inflags struct {
	infile  *string
	inpath  *string
	outpath *string
	makegv  *bool
}

func parse_flags() inflags {
	inf := inflags{
		infile:  flag.String("infile", "./ploshml_input_test.opa", "Main input file"),
		outpath: flag.String("outpath", ".", "Path to output"),
		makegv:  flag.Bool("opmstyle", true, "Make OPM-style image"),
	}
	flag.Parse()
	inf = add_defaut_values(inf)

	return inf
}

func add_defaut_values(inf inflags) inflags {
	if *inf.outpath == "." {
		*inf.outpath = path.Dir(*inf.outpath)

	}
	_, err := check_file_exists(*inf.outpath)
	if err != nil {
		log.Println("Output path does not exist")
		os.Exit(1)
	}
	return inf
}
