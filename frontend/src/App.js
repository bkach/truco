import React from "react";
import CenteredCard from "./CenteredCard";
import ItemList from "./ItemList";
import ItemListItem from "./ItemListItem";
import Cookies from "js-cookie";

class GameListComponent extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            gameItems: [],
            open: false
        }
    }

    updateGames() {
        fetch("api/games", {headers: {"accepts": "application/json"}})
            .then(res => res.json())
            .then(
                (result) => {
                    if (result != null) {
                        const gameItems = result.map((game) => new ItemListItem(game.id, game.name))
                        this.setState({
                            gameItems: gameItems
                        })
                    }
                    console.log(this.state)
                },
                (error) => {

                }
            )
    }

    handleAdd() {
        console.log("handle add")
        this.setState({
            open: true
        })
    }

    componentDidMount() {
        this.updateGames()
        this.timer = setInterval(() => {
            this.updateGames()
        }, 2000)
    }

    componentWillUnmount() {
        clearInterval(this.timer)
        this.timer = null
    }

    render() {

        return (
            <CenteredCard>
                <ItemList items={this.state.gameItems} onAdd={this.handleAdd}/>
            </CenteredCard>
        )
    }
}

function selectedGame() {
    return Cookies.get('selectedGame')
}

export default () => {
    return (
        <>
            <GameListComponent />
        </>
    );
}
