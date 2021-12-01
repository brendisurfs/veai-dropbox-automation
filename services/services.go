package services

// OUTPUT_FOLDER - where the program should output to.
const OUTPUT_FOLDER string = "C:/Users/brend/Dropbox/Enhanced_Video_Exports"

// PROC_FOLDER - folder to ingest footage from.
const PROC_FOLDER string = "/mnt/c/Users/brend/Dropbox/VideoEnhance_ProcFolder/"

// BeenProcessed - type for defining whether the file has been processed successfully before.
type BeenProcessed struct {
	Name    string
	Success bool
}

var ProccessedVideos []BeenProcessed
