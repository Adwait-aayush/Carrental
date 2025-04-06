FROM golang:1.23.4-alpine as build
WORKDIR /app
COPY . .
RUN go build -o /carental .
FROM scratch

COPY --from=build /carental /carental
EXPOSE 8080
CMD [ "/carental" ]