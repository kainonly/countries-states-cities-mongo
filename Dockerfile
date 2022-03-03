FROM alpine:edge

RUN apk add tzdata

COPY dist/event-invoke /

WORKDIR /

CMD [ "/event-invoke" ]