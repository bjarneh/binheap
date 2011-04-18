include $(GOROOT)/src/Make.inc

TARG=github.com/bjarneh/binheap
GOFILES=\
	binheap.go\
	sort.go

include $(GOROOT)/src/Make.pkg
