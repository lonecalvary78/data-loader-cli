package dataloader_test

import (
	"testing"

	"github.com/lonecalvary78/data-loader-cli/dataloader"
	"github.com/lonecalvary78/data-loader-cli/helper/batchprofilereader"
	"github.com/stretchr/testify/assert"
)

func FuzzLoadAndWriteToOtherFile(f *testing.F) {
	f.Add("{root}/config/config.yaml", "Batch-01")
	f.Fuzz(func(t *testing.T, batchProfilePath string, batchJobName string) {
		batchjob, _ := batchprofilereader.GetProfileFor(batchProfilePath, batchJobName)
		status := dataloader.LoadAndWriteToOtherFile(batchjob.SourceFile, batchjob.TargetFile, batchjob.Mappings)
		assert.Equal(t, "success", status)
	})
}

func BenchmarkLoadAndWriteOtherFile(b *testing.B) {
	for iterationcounter := 0; iterationcounter < b.N; iterationcounter++ {
		batchjob, _ := batchprofilereader.GetProfileFor("{root}/config/config.yaml", "Batch-01")
		dataloader.LoadAndWriteToOtherFile(batchjob.SourceFile, batchjob.TargetFile, batchjob.Mappings)
	}
}
