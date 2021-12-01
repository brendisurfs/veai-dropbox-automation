package enhancer

import (
	"brendisurfs/go-dropbox-processor/services"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

type StatusCode int

const (
	OK StatusCode = iota
	FAILED
)

func (sc StatusCode) ToStr() string {
	return []string{"[OK]", "[FAILED]"}[sc]
}

var tempFileName string

// CreateTempFile - creates a custom tempfile, whether thats an updater or something else.
func CreateTempFile(info string) string {
	nameInfo := fmt.Sprintf("%s/%s.txt", services.PROC_FOLDER, info)
	f, err := os.Create(nameInfo)
	if err != nil {
		log.Fatalf("%s: could not create temp file. error: %v", FAILED.ToStr(), err)
	}
	tmpName := f.Name()

	log.Printf("%v: created the temp file.", OK.ToStr())
	tempFileName = tmpName

	return tempFileName
}

// CheckPreviousSuccess - checks if the file being named has been processed before.
// if the filename is already in the list
func CheckPreviousSuccess(filename string) {
	for _, v := range services.ProccessedVideos {
		if filename == v.Name {
			break
		}
	}
}

// ReadName - reads the file name and parses the arguments from it.
func ParseFilename(filename string) ([]string, error) {
	splitExt := strings.Split(filename, ".")

	// remove last element (the file extension)
	splitExt = splitExt[:len(splitExt)-1]
	return splitExt, nil
}

// RunAndAlert - take args, runs the video processor, and exports to the right path.
// uses goroutine to wait until its done to use max gpu power for each vid.
func RunAndAlert(args []string, finished chan bool) error {

	jointArgs := strings.Join(args, " ")
	cmd := exec.Command("videoenhance", jointArgs)

	CreateTempFile("running the program...")

	if err := cmd.Run(); err != nil {
		CreateTempFile("could not run the program!")
		recover()
		return err
		// recover the program.
	}
	log.Println("finished processing video.")
	finished <- true

	return nil
}

// DeleteAlert - deletes the temp file we made to alert the user about the progress (since not terminal program.)
func DeleteAlert(temp string) {
	// check if there is a temp file.
	// if so, delete.

	os.Remove(temp)
	log.Printf("%v : deleted info temp file.", OK.ToStr())
}
