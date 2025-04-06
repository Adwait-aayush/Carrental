FROM golang:1.21-alpine as build
WORKDIR /app
COPY . .
RUN go build -o /carrental .
FROM alpine:latest

COPY --from=build /carrental /carrental
EXPOSE 8080
CMD [ "/carrental" ]