# telegrambot

A telegram bot library written in go

## Intro

This is a telegram bot library (not a framework). You can freely combine functions to build your bot.

## Subdirectories

    conf                        Config                      [doc][conf]
        conf.go

    handlers                                                [doc][handlers]
        base.go
        text_handler.go         TextHandler

    structs                                                 [doc][structs]
        callback_manager.go     CallbackManager
        queue.go                MessageQueue

    utils                       Helper functions            [doc][utils]
        utils.go

[conf]: https://godoc.org/github.com/laurence6/telegrambot-go/conf
[handlers]: https://godoc.org/github.com/laurence6/telegrambot-go/handlers
[structs]: https://godoc.org/github.com/laurence6/telegrambot-go/structs
[utils]: https://godoc.org/github.com/laurence6/telegrambot-go/utils

## License

Copyright (C) 2016-2016  Laurence Liu <liuxy6@gmail.com>

This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with this program.  If not, see <http://www.gnu.org/licenses/>.
