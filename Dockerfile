FROM alpine:edge

COPY dist/event-invoke /

CMD [ "/event-invoke" ]