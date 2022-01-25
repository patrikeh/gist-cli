COMPLETIONS_PATH=/usr/local/share/zsh/site-functions
SHELL=zsh
OUT_PATH=./out
BIN_PATH=$(shell go env GOPATH)
BIN_NAME=gist

make:
	@go build -o ${OUT_PATH}/${BIN_NAME}
	@mv ${OUT_PATH}/${BIN_NAME} ${BIN_PATH}/${BIN_NAME}
	@make clean
	@echo Binary ${BIN_NAME} available at ${BIN_PATH}.

completions:
	@${BIN_NAME} completion ${SHELL} > ${COMPLETIONS_PATH}/_${BIN_NAME}

gen:
	@cd docs && go run main.go

clean:
	@rm -rf ${OUT_PATH}