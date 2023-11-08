SRC_DIR = ./cmd
BIN_DIR = ./bin
RUN_DIR = ./runtime

INSTALL_DIR ?= ~/.local/share/lspcli

all: lspcli

clean:
	rm -rf bin

bin:
	@mkdir -p ${BIN_DIR}

lspcli: bin
	go build -o ${BIN_DIR}/lspcli ${SRC_DIR}/lspcli/main.go

install: lspcli
	cp -rf ${RUN_DIR} ${INSTALL_DIR}
	cp -rf ${BIN_DIR}/lspcli ${INSTALL_DIR}/bin/lspcli

uninstall:
	@rm -rf ${INSTALL_DIR}

.PHONY: all install clean


