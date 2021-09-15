package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Song struct {
	Artist  string `json:"artist"`
	Channel int32  `json:"channel"`
	Song    string `json:"song"`
	Time    string `json:"time"`
}

type ChristmasRadioResponse struct {
	PageLastPlayed struct {
		Content struct {
			Hero struct {
				Text string `json:"text"`
			} `json:"Hero"`
			RecentlyPlayed struct {
				Songs []Song `json:"songs"`
			} `json:"recently_played"`
		} `json:"Content"`
	} `json:"PageLastPlayed"`
}

func addChristmasRadioLists(listID, sToken string) {
	// loop all interesting lists
	radioUrls := []string{
		"https://jouluradio-wp.production.geniem.io/viimeksi-soitetut/",
		"https://jouluradio-wp.production.geniem.io/viimeksi-soitetut/?kanava=indiejoulu",
		"https://jouluradio-wp.production.geniem.io/viimeksi-soitetut/?kanava=jazzjoulu",
		"https://jouluradio-wp.production.geniem.io/viimeksi-soitetut/?kanava=klassinen-joulu",
	}
	for _, radioURL := range radioUrls {
		fmt.Printf("Adding songs from %s\n", radioURL)
		tracks, err := fetchChristmasRadioSongs(radioURL)
		if err != nil {
			fmt.Printf("Error when fetching songs %s\n", err.Error())
		}
		if err := addSongsFromRadioToPlaylist(tracks, listID, sToken); err != nil {
			fmt.Printf("Error when adding songs %s\n", err.Error())
		}
		fmt.Printf("Sleeping a while...\n")
		time.Sleep(sleepSeconds * time.Second)
	}

	fmt.Println("All done! Merry Christmas! Hyvää joulua!")
}

func fetchChristmasRadioSongs(url string) ([]*Track, error) {
	data, err := doGetRequest(url, "")
	if err != nil {
		return nil, err
	}

	response := ChristmasRadioResponse{}
	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}

	tracks := make([]*Track, 0)
	for _, track := range response.PageLastPlayed.Content.RecentlyPlayed.Songs {
		tracks = append(tracks, &Track{
			Artist: track.Artist,
			Title:  track.Song,
		})
	}

	return tracks, err
}
