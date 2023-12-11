package main

import "fmt"

type Streamer interface {
	play()
	pause()
	stop()
}

type MusicPlayer struct {
	name string
}

func (mp MusicPlayer) play() {
	fmt.Printf("%s is playing a song\n", mp.name)
}

func (mp MusicPlayer) pause() {
	fmt.Printf("%s is paused playing the song\n", mp.name)
}

func (mp MusicPlayer) stop() {
	fmt.Printf("%s is stopped playing the song\n", mp.name)
}

func main() {
	mp := MusicPlayer{name: "MP3"}
	mp.play()
	mp.pause()
	mp.stop()
}
