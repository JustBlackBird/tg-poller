# 🚀 TG-Poller

![Build and test](https://github.com/JustBlackBird/tg-poller/workflows/Build%20and%20test/badge.svg?branch=master)

> Grabs all the updates from Telegram API and streams them to a web-hook.


## Why

When you develop a bot for Telegram you may need to test incoming messages.
The problem is it's not that easy to expose a web-hook from an application
that's running localy to the Internet.

You can use polling in your app but it's not so good in production environment.
You can also support both web-hook and polling in your bot and switch them
depending on the enviroment, but you could end up with "works on my machine"
situation.

**TG-Poller** is a kind of proxy that's polling Telegram API for new updates and
stream them to a web-hook of your local application instance. You do not need to
implement polling in your bot just implement web-hook and use **TG-Poller** to test
your app. When the bot is deployed to the production environment register the web-hook
with Telgram API. You do not need **TG-Poller** anymore.

## Configuration

To configure the poller you have to use enviroment variables:

1. `TELEGRAM_API_HOST` — a host the poller will send requests for updates to. By
    default it's `api.telegram.org` but you can override it to use another host. It
    can be usefull if your environment does not have direct access to telegram API
    (e.g. Telegram is blocked by your organization's firewall).
2. `TELEGRAM_TOKEN` — token which your bot uses to access Telegram bot API. It's
    created for you by botfather when you set up a bot.
3. `UPDATE_HOOK` — URL the poller will stream updates to. Make sure it's available
    from poller's network.

## Usage

The simplest way to use **TG-Poller** is to run a prebuilt docker [image](https://hub.docker.com/repository/docker/justblackbird/tg-poller):

```bash
docker pull justblackbird/tg-poller
```

Below you'll find some guides about how to integrate the poller with your environment.

### Use with docker-compose

Here is an example of `docker-compose.yml` config:

```yml
version: '3.4'

services:
  bot:
    build:
      context: .
      dockerfile: ./Dockerfile
    restart: always

  poller:
    image: justblackbird/tg-poller:latest
    environment:
      TELEGRAM_API_HOST: api.telegram.org
      TELEGRAM_TOKEN: bot123456:ABC-DEF1234ghIkl-zyx57W2v1u123ew11
      UPDATE_HOOK: http://bot:80/api/telegram/update-hook
    restart: always
    depends_on:
      - bot

```

## License

[MIT](http://opensource.org/licenses/MIT) © Dmitry Simushev
