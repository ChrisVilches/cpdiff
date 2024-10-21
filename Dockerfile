FROM golang:1.23.2-alpine3.20

RUN echo 1.000000001 > in
RUN echo 1.0000002 > ans

RUN go install github.com/ChrisVilches/cpdiff@latest

ENV FORCE_COLOR=1

ENTRYPOINT ["cpdiff", "in", "ans"]
