import React, { Component } from "react";
import {render} from "react-dom";
import $ from "jquery";
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

//   class Home extends React.Component {
//     render() {
//       return (
//         <div className="container">
//           <div className="col-xs-8 col-xs-offset-2 jumbotron text-center">
//             <h1>Jokeish</h1>
//             <p>A load of Dad jokes XD</p>
//             <p>Sign in to get access </p>
//             <a onClick={this.authenticate} className="btn btn-primary btn-lg btn-login btn-block">Sign In</a>
//           </div>
//         </div>
//       )
//     }
//   }


//   class LoggedIn extends React.Component {
//     constructor(props) {
//       super(props);
//       this.state = {
//         jokes: []
//       }
//     }
      
//     render() {
//       return (
//         <div className="container">
//           <div className="col-lg-12">
//             <br />
//             <span className="pull-right"><a onClick={this.logout}>Log out</a></span>
//             <h2>Jokeish</h2>
//             <p>Let's feed you with some funny Jokes!!!</p>
//             <div className="row">
//               {this.state.jokes.map(function(joke, i){
//                 return (<Joke key={i} joke={joke} />);
//               })}
//             </div>
//           </div>
//         </div>
//       )
//     }
//   }

ReactDOM.render(<App />, document.getElementById('app'));