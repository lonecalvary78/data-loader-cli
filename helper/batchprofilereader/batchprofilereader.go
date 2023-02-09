package batchprofilereader

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/goccy/go-yaml"
	"github.com/lonecalvary78/data-loading-cli/models"
)

// Get the batch job from the external batch profiling file
func GetProfileFor(profilerFilePath string, batchJobName string) (models.BatchJob, error) {
	var batchProfile models.BatchProfile
	contents, errorOnLoad := os.ReadFile(profilerFilePath)
	if errorOnLoad != nil {
		return models.BatchJob{}, fmt.Errorf("unable to read the configuration file since it does not exists [file: %s]", profilerFilePath)
	}
	yaml.Unmarshal(contents, &batchProfile)

	if len(batchProfile.BatchJobs) > 0 {
		for _, batchjob := range batchProfile.BatchJobs {
			if strings.Contains(batchjob.Name, batchJobName) {
				return batchjob, nil
			}
		}
	}
	return models.BatchJob{}, errors.New("no batch job that match with the given job name")
}
