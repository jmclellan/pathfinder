import React, { useState } from 'react';
import logo from './logo.svg';
import './App.css';
import $ from "jquery";
import ReactDOM from 'react-dom';

import Jumbotron from 'react-bootstrap/Jumbotron';
import Container from 'react-bootstrap/Container';
import ListGroup from 'react-bootstrap/ListGroup';
import Collapse from 'react-bootstrap/Collapse';
import Form from 'react-bootstrap/Form';
import Button from 'react-bootstrap/Button';
import 'bootstrap/dist/css/bootstrap.min.css';

class App extends React.Component {
    render() {
        return (
            <Container>
                <Jumbotron>
                    <CoordConsole />
                </Jumbotron>
            </Container>)
    }
}

export default App;


class CoordConsole extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            // here is where we will store the componenats
            coordinates: []
        }
    }


    addCoordiate = (lat, lon) => {
        // debugger;
        let coordinate = { lat: lat, lon: lon }
        console.log(coordinate)
        this.setState({
            coordinates: [coordinate, ...this.state.coordinates]
        })
    }

    submitCoordiates = () => {
        // ajax request
        console.log("SUBMIT THE COORDINATES!")
    }

    render() {
        return (
            <div className="coordinate-console">
                <Form className="user-inputs">
                    <label>Enter your Coordinates one at a time!</label>
                    <Form.Group>
                        <Form.Label>Longitude</Form.Label>
                        <Form.Control id="formLon" type="number" placeholder="" required inputRef={(ref) => { this.input.lon = ref }} />
                        <Form.Control.Feedback type="invalid">
                            Please provide a valid longitude.
                        </Form.Control.Feedback>

                        <Form.Label>Latitude</Form.Label>
                        <Form.Control id="formLat" type="number" required inputRef={(ref) => { this.input.lat = ref }} />
                        <Form.Control.Feedback type="invalid">
                            Please provide a valid Latitude.
                    </Form.Control.Feedback>

                        <Button type="button" onClick={(data) => this.addCoordiate(document.getElementById("formLat").value,
                            document.getElementById("formLon").value)}>
                            enter these coordinates
                            </ Button>
                    </Form.Group>

                </Form>

                <div className="data-container">
                    <ListGroup>
                        {
                            this.state.coordinates.map((coordinateData, index) => <CoordinateContainer key={index}
                                locationName={"point " + (index + 1)}
                                lat={coordinateData.lat}
                                lon={coordinateData.lon}
                                removeElement={() => this.setState({
                                    coordinates: this.state.coordinates.filter((ele) =>
                                        !(ele.lat === coordinateData.lat) && ele.lon === coordinateData.lon)
                                })}
                            />
                            )
                        }
                    </ListGroup>
                </div>

            </div>
        )
    }
}

function CoordinateContainer(props) {
    const [open, setOpen] = useState(false)
    // make function called
    // setButtonText to adjust button name

    return <ListGroup.Item>
        <h1>
            {props.locationName}
        </h1>
        <Button onClick={() => props.removeElement()}>Remove point</Button>
        <Button
            aria-controls="example-collapse-text"
            aria-expanded={open}
            onClick={() => setOpen(!open)}>
            toggle details
        </Button>
        <Collapse in={open}>
            <div id="example-collapse-text">
                <h3>
                    Lat is <i>{props.lat}</i>
                </h3>
                <h3>
                    Lon is <i>{props.lon}</i>
                </h3>
            // add a delete button
            </div>
        </Collapse>
    </ListGroup.Item>
}
