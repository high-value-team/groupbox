FROM alpine:3.6
ADD groupbox /go/bin/groupbox
EXPOSE 80

CMD ["/go/bin/groupbox", "--port=80", "--mongodb-url=MONGODB_URL", "--groupbox-root-uri=GROUPBOX_ROOT_URI", "--smtp-username=SMTP_USERNAME", "--smtp-password=SMTP_PASSWORD", "--smtp-no-reply-email=SMTP_NO_REPLY_EMAIL", "--smtp-server-address=SMTP_SERVER_ADDRESS"]
