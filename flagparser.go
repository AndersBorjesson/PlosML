package main

import "flag"

type inflags struct {
	infile  string
	outpath string
	makegv  bool
}

func parse_flags() inflags {
	inf := inflags{
		infile:  *flag.String("infile", "", "Main input file"),
		outpath: *flag.String("outpath", ".", "Path to output"),
		makegv:  *flag.Bool("opmstyle", true, "Make OPM-style image"),
	}
	flag.Parse()
	return inf
}
