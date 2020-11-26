#!/bin/bash
go run ../cmd/stg/main.go \
	        --svg multi-fill-test.svg --out multi-fill-test.gcode \
		--gap 5 --threshold 5 \
		--fill 50,255,0,0 \
		--fill 140,0,255,0 \
		--fill 0,0,0,255 \
		--fill 90,255,0,255 \
		--fill 17,255,255,0
