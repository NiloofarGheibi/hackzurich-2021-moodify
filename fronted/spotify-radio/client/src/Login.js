import React from "react";
import {Container} from "react-bootstrap";

const AUTH_URL="https://accounts.spotify.com/authorize?client_id=4dce3921b23d43bda057e0cf0f5d7808&response_type=code&redirect_uri=http://localhost:3000&scope=streaming%20user-read-email%20user-read-private%20user-library-read%20user-library-modify%20user-read-playback-state%20user-modify-playback-state";
;

export default function Login(){
    return (
        <Container className="d-flex justify-content-center align-items-center" style={{minHeight: "100vh"}}>
            <a className="btn btn-success btn-lg rounded-pill" style={{backgroundColor: "#1ed760", borderColor: "#1ed760", fontWeight: "500"}} href={AUTH_URL}>LOGIN WITH SPOTIFY</a>
        </Container>
    )
}