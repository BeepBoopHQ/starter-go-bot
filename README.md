# go-slack-bot
A super simple go slack bot

Buiding the bot outside of docker requires [gb](https://getgb.io):

    go get github.com/constabulary/gb/...
    gb build ./...

Running locally:

    SLACK_TOKEN=<YOUR_SLACK_TOKEN> ./bin/bot 

Building with Docker:

    docker build -t beepboophq/go-slack-bot .

Running

    docker run -it --rm -e SLACK_TOKEN=<YOUR_SLACK_TOKEN> beepboophq/go-slack-bot
