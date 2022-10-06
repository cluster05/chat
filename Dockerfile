# Alpine is chosen for its small footprint
# compared to Ubuntu
FROM golang:1.16-alpine

# make the 'app' folder the current working directory
WORKDIR /app

# Download necessary Go modules
COPY go.mod ./
COPY go.sum ./

RUN go mod download

# copy remaing files
COPY *.go ./

# build app for production with minification
RUN go build -o /chatapplication

# run server
CMD [ "/chatapplication" ]