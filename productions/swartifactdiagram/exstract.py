#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
Created on Mon Feb 13 12:42:11 2023

@author: aborjes3
"""


def readfile(filename):
    with open(filename, "r") as f:
        d = f.read()
        
    d = d.split("\n")
    return d

def writefile(data, filename):
    with open(filename, "w") as f:
        for l1 in data:
            f.write(l1 + "\n")
            
def clean_up(d):
    candidates  =[]
    for l1 in d:
        if "Container" in l1:
            if "*diagram.Node" in l1:
                candidates.append(l1)
        else:
            pass
    return candidates


def str_replace(st):
    out = st
    if st == "r":
        out = "rlang"
    if st == "c":
        out = "clang"
    return out

def derive_funcs(c):
    funcs  = {}
    for l1 in c:
        package = l1.split("/")[0]
        name = l1.split("/")[1].split(".")[0]
        obj = l1.split("ontainer)")[1].split("(opts")[0].replace(" ", "")
        U = name[0].upper()
        name  = U + name[1:]
        key = str_replace(obj.lower())
        funcs[key] = package + "." + name+ "."  + obj
    return funcs

def generate_code(funcs):
    d = []
    d.append("package swartifactdiagram")
    d.append("func artifacts()map[string]func()*diagram.Node{")
    d.append("afs := make(map[string]func() *diagram.Node)")
    for _, l1 in enumerate(funcs):
        st = 'afs["' + l1 + '"] = func() *diagram.Node { return '
        st = st + funcs[l1] + "()}"
        
        d.append(st)
        
    d.append("return afs")
    d.append("}")
    return d

if __name__ == "__main__":
    filename_in = "/home/aborjes3/tmp.txt"
    filename_out = "./artifacts.go"
    raw = readfile(filename_in)
    candidates = clean_up(raw)
    funcs = derive_funcs(candidates)
    code = generate_code(funcs)
    writefile(code, filename_out)