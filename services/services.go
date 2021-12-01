package services

import "os"

// OUTPUT_FOLDER - where the program should output to.
var OUTPUT_FOLDER string = os.Getenv("OUTPUT_FOLDER")

// PROC_FOLDER - folder to ingest footage from.
var PROC_FOLDER string = os.Getenv("PROC_FOLDER")

// BeenProcessed - type for defining whether the file has been processed successfully before.
type BeenProcessed struct {
	Name    string
	Success bool
}

var ProccessedVideos []BeenProcessed
