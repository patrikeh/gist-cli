COMPLETIONS_PATH=/usr/local/share/zsh/site-functions
SHELL=zsh
OUT_PATH=./out
BIN_PATH=${GOPATH}/bin
BIN_NAME=gist

make:
	@go build -o ${OUT_PATH}/${BIN_NAME}
	@mv ${OUT_PATH}/${BIN_NAME} ${BIN_PATH}/${BIN_NAME}
	@make completions
	@make clean
	@echo Binary ${BIN_NAME} available at ${BIN_PATH}.

completions:
	@${BIN_NAME} completion ${SHELL} > ${COMPLETIONS_PATH}/_${BIN_NAME}

clean:
	@rm -rf ${OUT_PATH}