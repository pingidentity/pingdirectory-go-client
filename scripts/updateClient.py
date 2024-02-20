#!/usr/bin/env python3
# coding: UTF-8

# Update client.go
updatedClientLines = []
with open("client.go", 'r') as clientFile:
    for line in clientFile:
        # Skip unused imports
        if "\"log\"" in line or "\"net/http/httputil\"" in line:
            continue
        # Remove debug lines for printing request contents
        if "c.cfg.Debug" in line:
            # advance iterator to skip the expected lines
            for i in range(7):
                line = clientFile.readline()
        updatedClientLines.append(line)

with open("client.go", 'w') as clientFile:
    for line in updatedClientLines:
        clientFile.write(line)

