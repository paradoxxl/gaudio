//+build windows

package main

import(

)
import (
	"os/exec"
	//"time"
	"os"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	//ffmpeg  -f dshow -i audio="Microphone (Realtek High Definition Audio)" yo.mp3

	inputDeviceName := "Microphone (Realtek High Definition Audio)"
	rand.Seed(time.Now().UTC().UnixNano())
	outputFile := strconv.Itoa(rand.Int()) +".mp3"//time.Now().String() + ".mp3"

	cmd := exec.Command("C:/Dev/ffmpeg/bin/ffmpeg.exe", "-f","dshow", "-i","audio=" + inputDeviceName, outputFile)
	//cmd := exec.Command("C:/Dev/ffmpeg/bin/ffmpeg.exe")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err := cmd.Start()
	if err != nil {
		println(err.Error())
	}



	time.AfterFunc(time.Second*5,func(){
		cmd.Process.Kill()
	})


	err = cmd.Wait()
	if err != nil {
		println(err.Error())
	}

}
