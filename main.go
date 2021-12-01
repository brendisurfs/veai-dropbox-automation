package main

import (
	"brendisurfs/go-dropbox-processor/enhancer"
	"time"
)

var finished = make(chan bool)

func main() {
	// var proccessedList map[int]string
	// watch the folder
	//  if a new file pops in, check the extension.
	//  - is it: mp4, mov, something?
	// if it is, run the video enhance cli.
	//  - for running the cli, check the file name. This should parse the process to run.
	//  EXAMPLE: if the name has _noisy_ in it, run the denoiser more.
	//  if it has _noscale_ in it, dont upscale.
	//  if it has _low_, _med_, _high_: use the correct res setting for it.
	//  if it has: _cg_ in it, run the Gaia Computer Graphics setting.

	// procFolder, err := ioutil.ReadDir(services.PROC_FOLDER)
	// if err != nil {
	// 	panic(err)
	// }

	fname := enhancer.CreateTempFile("processing video ...")

	time.Sleep(time.Second * 5)

	enhancer.DeleteAlert(fname)

	// for i, v := range procFolder {
	// 	if v.Name() != proccessedList[i] {
	// 		// parses the filename as the first input in the slice, and the seperate args second.
	// 		// might have to do some gnarly regex.
	// 		args, err := enhancer.ParseFilename(v.Name())
	// 		if err != nil {
	// 			panic(err)
	// 		}

	// 		// wait for it to process...
	// 		go enhancer.RunAndAlert(args, finished)

	// 		// once its done
	// 		fmt.Println("done!")

	// 		// delete the standin alert file it creates.
	// 		enhancer.DeleteAlert()
	// 	}
	// }
}
