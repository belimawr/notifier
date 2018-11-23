# Notifier
Notifier is a small CLI application to automate notifying application's deployments
on multiple channels. It achieves that by parsing `git log` output and rendering
simple message templates then sending the same message across multiple channels.

## Work in Progress
This is a work in progress, therefore breaking changes will happen very often

## Usage
```sh
APP="my-beautiful-microservice" CHANNEL="channel-where-I-want-to-send-the-notification" TOKEN="bot-authorization-token" go run cli/main.go
```

## License
GPLv3
