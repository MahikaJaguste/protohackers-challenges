FROM golang:1.22.5
WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
ARG FILE_PATH
WORKDIR $FILE_PATH
RUN go build -o server
CMD [ "./server" ]