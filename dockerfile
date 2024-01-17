FROM golang
WORKDIR /app
ENV PORT 8080
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o ./out/dist ./cmd
CMD ./out/dist

