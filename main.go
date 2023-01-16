package main

import (
	"flag"
	"log"
	"os"

	"github.com/lonecalvary78/data-loader-cli/dataloader"
	"github.com/lonecalvary78/data-loader-cli/helper/batchprofilereader"
)

var (
	profilerFile = flag.String("profilerFile", "", "profiler file is required")
	batchJob     = flag.String("batchJob", "", "job name is required")
)

func main() {
	flag.Parse()
	batchjob, err := batchprofilereader.GetProfileFor(*profilerFile, *batchJob)
	if err != nil {
		log.Fatalln(err.Error())
		os.Exit(1)
	}
	status := dataloader.LoadAndWriteToOtherFile(batchjob.SourceFile, batchjob.TargetFile, batchjob.Mappings)
	log.Printf("the status of data loading is %s\n", status)
}
