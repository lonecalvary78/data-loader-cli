package batchprofilereader_test

import (
	"testing"

	"github.com/lonecalvary78/data-loader-cli/helper/batchprofilereader"
	"github.com/stretchr/testify/assert"
)

func TestReadBatchProfile(t *testing.T) {
	batchjob, _ := batchprofilereader.GetProfileFor("{root}/config/config.yaml", "Batch-01")
	assert.Equal(t, "Batch-01", batchjob.Name)
	assert.Equal(t, "firstname", batchjob.Mappings[0].FieldName)
	assert.Equal(t, 0, batchjob.Mappings[0].StartPosition)
	assert.Equal(t, 13, batchjob.Mappings[0].EndPosition)
	assert.Equal(t, "lastname", batchjob.Mappings[1].FieldName)
	assert.Equal(t, 14, batchjob.Mappings[1].StartPosition)
	assert.Equal(t, 29, batchjob.Mappings[1].EndPosition)
}

func BenchmarkReadBatchProfile(b *testing.B) {
	for counter := 0; counter < b.N; counter++ {
		batchprofilereader.GetProfileFor("{root}/config/config.yaml", "Batch-01")
	}
}
