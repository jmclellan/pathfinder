import React, { useState } from 'react';
import logo from './logo.svg';
import './App.css';
import $ from "jquery";
import ReactDOM from 'react-dom';

// // for attempt at map
// import { Map as LeafletMap, GeoJSON, Marker, Popup } from 'react-leaflet';
// import worldGeoJSON from 'geojson-world-map';

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
            <div>
            <Container>
                <Jumbotron>
                    <CoordConsole />
                </Jumbotron>
            </Container>
            {/* <GeoJsonMap /> */}
            </div>)
            
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


    addCoordiate = (Lat, Lon) => {
        // debugger;
        let coordinate = { Lat: Lat, Lon: Lon }
        console.log(coordinate)
        this.setState({
            coordinates: [coordinate, ...this.state.coordinates]
        })
    }

    submitCoordiates = () => {
        // ajax request
        console.log("SUBMIT THE COORDINATES!")
        let postData = {"Length": this.state.coordinates.length,
                    "Path": this.state.coordinates
                    }
        $.post("/api/optimize_route/", JSON.stringify(postData))
         .then((resp) => {
            let responseObj = JSON.parse(resp)
            this.setState({coordinates: responseObj.path})
            console.log(responseObj)
            })
        console.log("Coordinates submitted!")
    }

    render() {
        return (
            <div className="coordinate-console">
                <Form className="user-inputs">
                    <label>Enter your Coordinates one at a time!</label>
                    <Form.Group>
                        <Form.Label>Longitude</Form.Label>
                        <Form.Control id="formLon" type="number" placeholder="" required inputRef={(ref) => { this.input.Lon = ref }} />
                        <Form.Control.Feedback type="invalid">
                            Please provide a valid longitude.
                        </Form.Control.Feedback>

                        <Form.Label>Latitude</Form.Label>
                        <Form.Control id="formLat" type="number" required inputRef={(ref) => { this.input.Lat = ref }} />
                        <Form.Control.Feedback type="invalid">
                            Please provide a valid Latitude.
                    </Form.Control.Feedback>

                        <Button type="button" onClick={(data) => 
                            this.addCoordiate(Number(document.getElementById("formLat").value),
                                              Number(document.getElementById("formLon").value))}>
                            enter these coordinates!
                        </ Button>
                       
                        <Button type="button" onClick={() => this.submitCoordiates()}> 
                        Submit Coordiates 
                        </Button> 
                    </Form.Group>

                </Form>

                <div className="data-container">
                    <ListGroup>
                        {
                            this.state.coordinates.map((coordinateData, index) => <CoordinateContainer key={index}
                                locationName={"point " + (index + 1)}
                                lat={coordinateData.Lat}
                                lon={coordinateData.Lon}
                                removeElement={() => this.setState({
                                    coordinates: this.state.coordinates.filter((ele) =>
                                        !(ele.Lat === coordinateData.Lat) && ele.Lon === coordinateData.Lon)
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
            </div>
        </Collapse>
    </ListGroup.Item>
}

// class GeoJsonMap extends React.Component {
//     render() {
//       return (
//         <LeafletMap
//           center={[50, 10]}
//           zoom={6}
//           maxZoom={10}
//           attributionControl={true}
//           zoomControl={true}
//           doubleClickZoom={true}
//           scrollWheelZoom={true}
//           dragging={true}
//           animate={true}
//           easeLinearity={0.35}
//         >
//           <GeoJSON
//             data={worldGeoJSON}
//             style={() => ({
//               color: '#4a83ec',
//               weight: 0.5,
//               fillColor: "#1a1d62",
//               fillOpacity: 1,
//             })}
//           />
//           <Marker position={[50, 10]}>
//             <Popup>
//               Popup for any custom information.
//             </Popup>
//           </Marker>
//         </LeafletMap>
//       );
//     }
//   }
  