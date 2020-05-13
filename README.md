# Truco [WIP]

![](assets/1ofspades.jpg)

This Repo is a simple implementation of the game [Truco](https://en.wikipedia.org/wiki/Truco) written in
[go](https://golang.org/) and [react](https://reactjs.org).

# How to Run

1. `cd` to the cloned directory
2. Run the command: `go run main.go`
3. Navigate to `localhost:8000`
4. If you see output similar to:

`{"game_id":"game_ffa71ce1-86e1-46b3-870f-b6fc3a019124"}`

It's working!

# How to deploy

You'll first need to set the `$TRUCO_SERVER` environment variable to a given `user@ipaddress`. I'll share the one I use
 with collaborators. Then run:

```bash
chmod +x /scripts/deploy.sh
./scripts/deploy.sh
```

Take a look at [`deploy.sh`](https://github.com/bkach/truco/blob/master/scripts/deploy.sh) for details.

# Endpoints

For a list of endpoints and how to use them, take a look at the [`routes.go`](https://github.com/bkach/truco/blob/master/routers/routes.go) file.

# ToDo

- [ ] Remove the `board` concept and allow players to keep track of which card they play first, second, and third
- [ ] The entire frontend, [I have some inkling for that might look like](https://docs.google.com/drawings/d/1c_wXfl3P--VGM77Gdfu8vPa-mZLwx8Xqx8Tz1zec_vg/edit?usp=sharing) - but it's very much WIP
