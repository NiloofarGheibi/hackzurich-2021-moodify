import React from "react"

export default function TrackSearchResult({ track, chooseTrack, playingTrack, openRadio }) {
    function handlePlay() {
        chooseTrack(track)
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
            <h1 className="btn rounded-circle mr-2" style={{backgroundColor: "#A9A9A9", color: "white",  marginRight: "10px"}} onClick={openRadio}>+</h1>
        </div>
    )
}