package builder

var tmpl = `FROM $base AS builder

ADD . .

WORKDIR $GOPATH/$path

$add-on
COPY . ./
RUN go build ./...
RUN rm -rf $GOPATH/bin/dep

FROM scratch
ENTRYPOINT ["./app"]
`

var shTmpl = `#!/bin/sh

docker build -t $projectName:$projectTag .

dangling_docker=$(docker images -f 'dangling=true' -q)
if [ -z "$dangling_docker" ]; then
    exit 1
fi

docker rmi $dangling_docker --force
`