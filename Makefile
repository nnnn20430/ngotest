define PKGS :=
	github.com/nnnn20430/ngotest/hello
	github.com/nnnn20430/ngotest
endef
OBJS := $(addprefix obj/, $(PKGS:%=%.o))
MAIN := ngotest

.NOTPARALLEL:

.PHONY: all clean $(MAIN) FORCE

all: $(MAIN)

clean:
	rm -rf bin obj

$(MAIN): $(OBJS)
	@mkdir -p bin
	gccgo -static -o bin/$@ $^

obj/%.o: src/% FORCE
	@mkdir -p $(dir $@)
	gccgo -I obj -c -o $@ $(wildcard $(<)/*.go)

FORCE:
