# syntax=docker/dockerfile:1

FROM golang:1.18.1-windowsservercore-1809 AS build

WORKDIR /app

COPY go.mod ./
# COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /checkSocketFile.exe

## CMD [ "/checkSocketFile.exe" ]

FROM mcr.microsoft.com/windows/servercore:ltsc2019

WORKDIR /

COPY --from=build /checkSocketFile.exe /checkSocketFile.exe

## USER nonroot:nonroot

ENTRYPOINT ["/checkSocketFile.exe"]