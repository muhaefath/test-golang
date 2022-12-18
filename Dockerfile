FROM golang:1.18

RUN mkdir -p /go/src/github.com && cd /go/src/github.com && mkdir bee && cd bee &&  git clone https://github.com/beego/bee.git && cd bee && git checkout 1ed5c7108761c0be6cf && go install
WORKDIR /go/src/netpro/test-golang

RUN apt-get update && apt-get install -y poppler-utils wv unrtf tidy

RUN ["go", "install", "github.com/githubnemo/CompileDaemon@latest"]
CMD CompileDaemon -log-prefix=false -build="go build" -command="./test-golang"

# CMD /go/bin/bee run

