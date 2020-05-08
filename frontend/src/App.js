import React from "react";

class TestComponent extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            gameCreated: false
        }
    }

    componentDidMount() {
        fetch("api/createGame", {headers:{"accepts":"application/json"}})
            .then(
                (result) => {
                    console.log(result)
                    this.setState({
                        gameCreated: true
                    })
                },
                (error) => {}
            )
    }

    render () {
        return (<h1>Game Created: {this.state.gameCreated.toString()}</h1>)
    }
}

export default () => (
  <>
    <h1>Truco Test</h1>
      <TestComponent />
  </>
);
