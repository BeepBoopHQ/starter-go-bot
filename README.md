# starter-go-bot

## Overview
A simple starting point for creating a Beep Boop hostable, Go based Slack bot.

Visit [Beep Boop](https://beepboophq.com/docs/article/overview) to get the scoop on the the Beep Boop hosting platform. The Slack API documentation can be found [here](https://api.slack.com/).

## Assumptions
* You have already signed up with [Beep Boop](https://beepboophq.com) and have a local fork of this project.
* You have sufficient rights in your Slack team to configure a bot and generate/access a Slack API token.
* Go 1.5 or greater

## Usage

### Run locally
	go get github.com/constabulary/gb/...
	gb build all
	SLACK_TOKEN=<YOUR_SLACK_TOKEN> ./bin/bot

Things are looking good if the console prints something like:

	Connected <your bot name> to <your slack team> team at https://<your slack team>.slack.com.

### Run locally in Docker
	docker build -t beepboophq/go-slack-bot .
	docker run -it --rm -e SLACK_TOKEN=<YOUR_SLACK_TOKEN> beepboophq/go-slack-bot

### Run in BeepBoop
If you have linked your local repo with the Beep Boop service (check [here](https://beepboophq.com/0_o/my-projects)), changes pushed to the remote master branch will automatically deploy.

## Acknowledgements

This code utilizes the super handy https://github.com/nlopes/slack project by [@nlopes](https://github.com/nlopes).

## License

See the [LICENSE](LICENSE.md) file for license rights and limitations (MIT).
