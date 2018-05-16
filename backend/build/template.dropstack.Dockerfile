FROM alpine:3.6

ADD groupbox-backend /go/bin/groupbox-backend

EXPOSE 80

RUN apk add --no-cache ca-certificates

CMD ["/go/bin/groupbox-backend", "--port=80", "--mongodb-url=${MONGODB_URL}", "--smtp-username=${SMTP_USERNAME}", "--smtp-password=${SMTP_PASSWORD}", "--smtp-no-reply-email=${SMTP_NO_REPLY_EMAIL}", "--smtp-server-address=${SMTP_SERVER_ADDRESS}"]
