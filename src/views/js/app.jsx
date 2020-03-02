// import React, { Component } from "react";
// import {render} from "react-dom";
// import $ from "jquery";

// this will run everything else and as such will be the root of deciding if user is logged in or not
// and how that affects what we show
class App extends React.Component {
    render() {
        CoordConsole()
  }
}
// we want a coodConsole componant - which will house 
class CoordConsole extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            // here is where we will store the componenats
            coordinates: []
        }
    }


    addCoordiate = () => {
        this.setState({
            coordinates: [{Lat: $(".lat-input").val(), Lon: $(".lon-input").val()}, ...this.state.coordinates]
        })
        $(".lat-input").val(0);
        $(".lon-input").val(0)
    }

    submitCoordiates = () => {
        // ajax request
        console.log("SUBMIT THE COORDINATES!")
    }

    render() {
        return (
            <div className="coordinate-console">
                <div className="user-inputs">
                    <label>Enter your Coordinates one at a time!</label>
                    <input type="number" id="lat-input"></input>
                    <input type="number" id="lon-input"></input>
                    <button type="button" onClick="">enter these coordinates</button>
                </div>
                
                <div className="data-container">
                    <ol>
                        {this.state.coordinates.map((coordinateData) => coordinateContainer(coordinateData))}
                    </ol>
                </div>

            </div>
        )
    }
}

function coordinateContainer(props) {
    return <li>
        <h1>
            Lat is <i>{props.Lat}</i>
        </h1>
        <h3>
            Lon is <i>{props.Lon}</i>
        </h3>
        // add a delete button
    </li>
}


ReactDOM.render(<App />, document.getElementById('app'));