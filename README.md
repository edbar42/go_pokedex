# Go Pokédex
Your classic Pokédex app, now on the CLI and powered by Go. Also with an embedded "_catch'em all_" mini-game.

> [!WARNING]
> This project is currently under development.

## 💾 Installation
This project requires `go 1.22` in order to run.

To run the app, clone this repository and `cd` into it. Then, from the project's root directory run:

```sh
go build && ./go_pokedex
```

Alternatively, you may also install it using the `go` toolchain:

```sh
go install github.com/edbar42/go_pokedex@latest
```

If your `$GOPATH` is properly set, you should now be able to start it from the command line using `go-pokedex`.

## 🤝 Contributing

> [!WARNING]
> Please be aware that your suggestions might be implemented in the near future.
> 
> Consider the [TODO](#to-do) section before you start working on changes.
> 
You can contribute by forking this repo and opening pull requests with your suggestion. 

All pull requests should be submitted to the main branch.

## ⚖️ License
This project is licensed under the MIT License - see the LICENSE file for details.

## To-Do
- [x] Implement REPL
- [x] Implement app commands map
- [x] Implement API calls on separate package
- [x] Add application-wide caching
- [x] Test caching expiring timer
- [ ] Allow the user to explore regions via command-line
- [ ] Embed pokemon catching minigame
- [ ] Write Usage section on README
- [ ] Add video of app running to README
