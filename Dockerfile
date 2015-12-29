FROM golang:1.5.2

# Copy current project
COPY . /go/src/github.com/mrpatiwi/revel-mongo-api
WORKDIR /go/src/github.com/mrpatiwi/revel-mongo-api

# Install Revel CLI
RUN go get github.com/revel/cmd/revel

# Install project dependencies
RUN go get github.com/tools/godep
RUN godep go install ./app

# Run app
EXPOSE 9000
ENTRYPOINT revel run github.com/mrpatiwi/revel-mongo-api prod 9000
