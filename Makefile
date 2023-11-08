SRC_DIR := cmd

# Define the output folder for binary files
BIN_DIR := bin

# Find all subdirectories in the SRC_DIR
SUBDIRS := $(wildcard $(SRC_DIR)/*)

# Generate the binary file names based on the subdirectory names
BINARIES := $(patsubst $(SRC_DIR)/%,$(BIN_DIR)/%,$(SUBDIRS))

# The default target builds all binaries
all: $(BINARIES)

# Define a pattern rule to build each binary from its corresponding source directory
$(BIN_DIR)/%: $(SRC_DIR)/%
	@mkdir -p $(BIN_DIR)
	go build -o $@ $</main.go

clean:
	rm -rf $(BINARIES)

.PHONY: all clean


