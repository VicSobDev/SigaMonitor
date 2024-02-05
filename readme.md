# SIGAMonitor

## About the Project

SIGAMonitor is a tool developed to assist individuals in searching for available appointment times for residence permit renewal on the SIGA platform. Knowing the difficulty of finding available times due to high demand, SIGAMonitor automates the monitoring process, offering a practical and efficient solution for users.

## How It Works

The application, written in Go (Golang), constantly monitors the SIGA platform, specifically the area for residence permit renewal. When a new appointment time becomes available, SIGAMonitor immediately sends a notification through a user-defined channel on Telegram. This allows users to react quickly and secure their appointment.

## Prerequisites

To use SIGAMonitor, you need to have:

- Go installed on your machine.
- A Telegram bot created through BotFather, to send notifications.
- Internet access for the application to monitor the SIGA platform.

## Configuration

### Setting Up the Bot on Telegram

1. Create a new bot on Telegram using BotFather and note the bot token.
2. Create a Telegram channel where notifications will be sent.
3. Add the created bot to the channel as an administrator.

### Preparing the SIGAMonitor Environment

Clone the repository and navigate to the project directory:

```bash
git clone https://github.com/VicSobDev/SigaMonitor
cd SigaMonitor
```

Create a .env file in the project's root directory and add the following variables with their respective values:

```
BOT_TOKEN=your_token_here
CHANNEL_ID=your_channel_id
```

Replace your_token_here and your_channel_id with the actual values obtained during the bot setup on Telegram.

# Usage

To start monitoring, follow the configuration instructions and execute the command:

```
make run 
```

To use Docker for running SIGAMonitor, follow these instructions:

1. Build the Docker image from the project directory:

```bash
docker build -t sigamonitor .
```

2. Once the image is built, you can run the container with the following command:

```bash
docker run -d --env-file .env sigamonitor
```

This command runs the container in detached mode, using the `.env` file you created for environment variables. Ensure that your `.env` file includes the `BOT_TOKEN` and `CHANNEL_ID` values as previously described.

By following these steps, SIGAMonitor will be up and running inside a Docker container, continuously monitoring the SIGA platform for available appointment times and sending notifications to your Telegram channel.