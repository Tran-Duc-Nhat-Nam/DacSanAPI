FROM golang
WORKDIR /app
EXPOSE 8080
ENV PORT 8080
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o ./out/dist ./cmd
CMD ./out/dist

