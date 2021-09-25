import React from "react"

export default function TrackSearchResult({ track, chooseTrack, playingTrack, openRadio }) {
    function handlePlay() {
        chooseTrack(track)
    }
    function SubmitSongAndOpenRadio() {
        const url = 'http://localhost:9091';
        const post = {
            user_id: "c017fb99-067b-481f-979c-6c06e0a45786",
            track_id: track.uri
        };
        const requestPromise = makeRequest(url, post);
        console.log(requestPromise)
        openRadio()
    }
    function makeRequest(url, data) {
        return new Promise((resolve, reject) => {
            let request = new XMLHttpRequest();
            request.open('POST', url + '/radios');
            request.onreadystatechange = () => {
                if (request.readyState === 4) {
                    if (request.status === 200) {
                        resolve(JSON.parse(request.response));
                    } else {
                        reject(JSON.parse(request.response));
                    }
                }
            };
            request.setRequestHeader('Content-Type', 'application/json');
            request.send(JSON.stringify(data));
        });
    }

    return (
        <div
            className={"d-flex m-2 align-items-center justify-content-between bg-light rounded-pill"}
            style={{ cursor: "pointer", padding: "5px"}}
            onClick={handlePlay}
        >
            <div  className="d-flex">
            <img className={"rounded"} src={track.albumUrl} style={{ height: "64px", width: "64px", marginRight: "10px", marginLeft: "20px"}} />
            <div className="ml-3">
                <div>{track.title}</div>
                <div className="text-muted">{track.artist}</div>
            </div>
            </div>
            <h1 className="btn rounded-circle mr-2" style={{backgroundColor: "#A9A9A9", color: "white",  marginRight: "10px"}} onClick={SubmitSongAndOpenRadio}>+</h1>
        </div>
    )
}
