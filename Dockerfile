FROM golang:1.9.2-alpine

ARG app_env
ENV APP_ENV $app_env

COPY . /go/src/app
WORKDIR /go/src/app

RUN apk add --no-cache git mercurial
RUN go get ./
RUN go build

# if dev setting will use pilu/fresh for code reloading via docker-compose volume sharing with local machine
# if production setting will build binary
CMD if [ ${APP_ENV} = production ]; \
	then \
	./app; \
	else \
	go get github.com/pilu/fresh && \
	fresh; \
	fi

EXPOSE 8080
