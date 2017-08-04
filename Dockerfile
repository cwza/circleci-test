FROM alpine:3.2
USER root

COPY ./test.txt /

CMD ["tail", "-f", "/test.txt"]