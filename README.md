# telegrambot

A telegram bot library (not a framework) written in go. You can freely combine functions to build your bot.

## Subdirectories

<pre>
- conf                        Config                <a href="https://godoc.org/github.com/laurence6/telegrambot-go/conf">doc</a>
  - conf.go

- handlers                                          <a href="https://godoc.org/github.com/laurence6/telegrambot-go/handlers">doc</a>
  - base.go
  - text_handler.go           TextHandler

- structs                                           <a href="https://godoc.org/github.com/laurence6/telegrambot-go/structs">doc</a>
  - callback_manager.go       CallbackManager
  - queue.go                  MessageQueue

- utils                       Helper functions      <a href="https://godoc.org/github.com/laurence6/telegrambot-go/utils">doc</a>
  - utils.go
</pre>

## License

Copyright (C) 2016-2016  Laurence Liu <liuxy6@gmail.com>

This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with this program.  If not, see <http://www.gnu.org/licenses/>.
