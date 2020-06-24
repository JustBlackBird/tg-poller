# TG-Poller

> Grabbs all the updates from Telegram API and streams them to a web-hook.

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


## License

[MIT](http://opensource.org/licenses/MIT) © Dmitry Simushev
