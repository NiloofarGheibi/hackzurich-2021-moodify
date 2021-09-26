import React from "react"
import { useState, useEffect } from "react"
import SpotifyWebApi from "spotify-web-api-node"
import useAuth from "./useAuth";
import { Container, Form } from "react-bootstrap"
import TrackSearchResult from "./TrackSearchResult"
import Player from "./Player"

const spotifyApi = new SpotifyWebApi({
    clientId: <CLIENT_ID>,
})

export default function Dashboard({code}){
    const accessToken = useAuth(code)
    const [search, setSearch] = useState("")
    const [searchResults, setSearchResults] = useState([])
    const [playingTrack, setPlayingTrack] = useState()
    const [trackSelected, setTrackSelected] = useState(false)


    function chooseTrack(track) {
        setPlayingTrack(track)
        //setSearch("")
    }

    function openRadio() {
        setTrackSelected(true)
    }

    useEffect(() => {
        if (!accessToken) return
        spotifyApi.setAccessToken(accessToken)
    }, [accessToken])


    useEffect(() => {
        if (!search) return setSearchResults([])
        if (!accessToken) return

        let cancel = false
        spotifyApi.searchTracks(search).then(res => {
            if (cancel) return
            setSearchResults(
                res.body.tracks.items.map(track => {
                    const smallestAlbumImage = track.album.images.reduce(
                        (smallest, image) => {
                            if (image.height < smallest.height) return image
                            return smallest
                        },
                        track.album.images[0]
                    )

                    return {
                        artist: track.artists[0].name,
                        title: track.name,
                        uri: track.uri,
                        albumUrl: smallestAlbumImage.url,
                    }
                })
            )
        })

        return () => cancel = true
    }, [search, accessToken])


    return (
        trackSelected ? (
                <Container className="d-flex align-items-center flex-column" style={{ height: "100vh", marginTop: "20%" }}>
                    <h1 style={{ marginBottom: "40px" }}>Your Team Radio 📡</h1>
                    <Player accessToken={accessToken} trackUri={playingTrack?.uri} />
                </Container>
            )
            : <Container className="d-flex flex-column py-2" style={{ height: "100vh" }}>
                <h1 className="d-flex justify-content-center">Hey you :)</h1>
                <h1 className="d-flex justify-content-center"> </h1>
                <h1 className="d-flex justify-content-center mb-2">Which music is in your mind?</h1>
                <h1 className="d-flex justify-content-center"> </h1>
                <Form.Control
                            type="search"
                            placeholder="Search Songs/Artists"
                            value={search}
                            onChange={e => setSearch(e.target.value)}
                        />
                        <div className="flex-grow-1 my-2" style={{overflowY: "auto"}}>
                            {searchResults.map(track => (
                                <TrackSearchResult
                                track={track}
                                key={track.uri}
                                chooseTrack={chooseTrack}
                                 playingTrack={playingTrack}
                                openRadio={openRadio}
                                />
                            ))}
                        </div>
                        {searchResults.length != 0 && (
                            <div>
                                <Player accessToken={accessToken} trackUri={playingTrack?.uri} />
                            </div>
                        )}
            </Container>
    )
}
