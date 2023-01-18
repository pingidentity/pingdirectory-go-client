#!/usr/bin/env python3
# coding: UTF-8
import glob

# Find all files in the repo starting with "model_enum"
enumModelFiles = glob.glob('model_enum*.go')

# Search in each of those files for the line corresponding to the enum name
enumNames = []
for enumFile in enumModelFiles:
    with open(enumFile) as file:
        for line in file:
            if line.startswith("type ") and line.rstrip().endswith(" string"):
                enumNames.append(line[5:-8])

# Create enum_stringer_implementation.go
with open("enum_stringer_implementation.go", 'w') as stringerFile:
    stringerFile.write("package openapi\n")
    stringerFile.write("\n")
    stringerFile.write("// Generated by scripts/generateStringerImplementation.py. Implements Stringer interface for each enum in the client.\n")
    for enum in enumNames:
        stringerFile.write("\n")
        stringerFile.write(f'func (e {enum}) String() string {{\n')
        stringerFile.write('    return string(e)\n')
        stringerFile.write('}\n')