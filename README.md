# Remindme

A Telegram bot to remind you things you don't want to forget

## Why a bot?

I constantly need to remind myself about a weekly thing to do, a daily good habit to follow or a specific event that will occur later on in the year. Sure I could set some alarms, but that does not scale well, or I could put them in the default Mac Notes app. But how do I remember to look through the notes? 

So I thought about creating a Telegram Bot (no fancy thing, it just boils down to a http handler the Telegram API sends the user message to) named `@remindme_later_bot` that I can send a reminder to and that will send it back to me on the specified (recurrent) date and time

## Examples of command

```sh
/remindme start your day with breathing exercises | everyday @ 8:00
/remindme look at the release of this documentary | "2021-12-03" @ 15:00
/remindme checkout the weekly podcast on global warming | each Tuesday @ 18:00
/remindme pay your bills | each 3 @ 8:00
/remindme this is Foo birthday today | each October 19 @ 10:00
```

## WIP

I am currently workong on having a process that checks db in an fix interval of time to get matching reminders and send them to != chats. So basically when sending reminders to the bot at the moment, it won't remind yet (messages are only stored in db)

## Stack stuff

The Golang server and the Postgresql database are running in containers on my Debian server and a Drone instance is used for CI/CD