NAME="pwg"

.PHONY: $(NAME)

$(NAME):
	go build -o bin/pwg main.go

install:
	mv bin/pwg /usr/bin/pwg
