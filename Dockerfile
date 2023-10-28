FROM golang:1.20.1 as builder
WORKDIR /workspace
COPY . /workspace
RUN go build -o server .

FROM debian:11-slim
COPY --from=builder /workspace/server /app/server
EXPOSE 8080
CMD [ "/app/server" ]