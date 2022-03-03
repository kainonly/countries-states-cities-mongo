FROM alpine:edge

RUN apk add tzdata

COPY dist /tmp

WORKDIR /tmp

CMD [ "./main" ]