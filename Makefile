
include $(GOROOT)/src/Make.inc

TARG = ansi2utf8
GOFILES = \
	ansi2utf8.go

include $(GOROOT)/src/Make.cmd

run: all
	./ansi2utf8
