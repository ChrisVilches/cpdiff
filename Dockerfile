FROM golang:1.22-alpine3.19

RUN echo 1.000000001 > in
RUN echo 1.0000002 > ans

RUN go install github.com/ChrisVilches/cpdiff@latest

ENTRYPOINT FORCE_COLOR=1 cpdiff in ans
