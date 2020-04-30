# Truco Backend [WIP]

This Repo is a simple backend for the game [Truco](https://en.wikipedia.org/wiki/Truco) written in [go](https://golang.org/).

## Endpoints and Documentation

### ``GET/newGame`

Starts a new game and builds a deck

### ``POST /addPlayer`

Adds a player to the game

#### Request

```
{
	"name": "my player's name"
}
```

#### Response

```
{
    "PlayerState": {
        "Info": {
            "Name": "boris",
            "ID": "player_boris"
        },
        "Hand": [
            {
                "value": 1,
                "house": "clubs"
            },
            {
                "value": 2,
                "house": "gold"
            },
            {
                "value": 2,
                "house": "cups"
            }
        ]
    },
    "Board": []
}
```

### ``POST /player`

Gets information on a given player

#### Request

```
{
	"playerId": "player_boris"
}
```

#### Response

```
{
    "PlayerState": {
        "Info": {
            "Name": "boris",
            "ID": "player_boris"
        },
        "Hand": [
            {
                "value": 1,
                "house": "cups"
            },
            {
                "value": 2,
                "house": "gold"
            },
            {
                "value": 2,
                "house": "cups"
            },
        ]
    },
    "Board": []
}
```

### ``POST /playCard`

Gets information on a given player

#### Request

```
{
	"ID" : "player_boris",
	"card": {
		"value": 1, "house": "cups"
	}
}
```

#### Response

```
{
    "PlayerState": {
        "Info": {
            "Name": "boris",
            "ID": "player_boris"
        },
        "Hand": [
            {
                "value": 2,
                "house": "gold"
            },
            {
                "value": 2,
                "house": "cups"
            },
        ]
    },
    "Board": [
            {
                "value": 1,
                "house": "cups"
            }
    ]
}
```

## ToDo

Quite a lot. WIP!