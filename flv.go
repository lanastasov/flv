package main

import (
	"fmt"
	"strings"
	"path/filepath"
	"os/exec"
)

func main() {
	files, _ := filepath.Glob("*")
	for i := range files {
		fmt.Println("  <-  ", files[i])
		extensions := []string { ".flv", ".wmv" , ".FLV", ".WMV", ".avi", ".AVI", ".mkv", ".MKV", ".ts", ".TS", ".mpg", ".MPG", ".f4v", ".m4v" }
		for j := range extensions {
			if strings.Contains(files[i], extensions[j]) {
				cmd, err := exec.Command("ffmpeg", "-i", fmt.Sprintf(files[i]), "-vcodec", "libx264", strings.Replace(fmt.Sprintf(files[i]),extensions[j], ".mp4", 1)).Output()
				if err != nil {
					fmt.Println(err)
					return
				}
			fmt.Printf("Command ran: %s\n", string(cmd))
			}	
		}
	}
}
