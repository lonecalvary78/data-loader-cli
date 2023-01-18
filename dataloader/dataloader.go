package dataloader

import (
	"context"
	"strings"
	"time"

	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/core/graph/window"
	_ "github.com/apache/beam/sdks/v2/go/pkg/beam/io/filesystem/local"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/io/textio"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/runners/direct"
	"github.com/lonecalvary78/data-loader-cli/models"
)

const (
	SUCCESS = "success"
	FAILED  = "failed"
)

// To load data from source file and store it into the output file
func LoadAndWriteToOtherFile(sourceFilePath string, targetFilePath string, fieldMappings []models.FieldMapping) string {
	beam.Init()
	pipeline, scope := beam.NewPipelineWithRoot()
	dataRowsFromSourceFile := textio.Read(scope, sourceFilePath)
	windowedRows := beam.WindowInto(scope, window.NewFixedWindows(60*time.Second), dataRowsFromSourceFile)
	transformedRows := beam.ParDo(scope, func(line string, emit func(string)) {
		var builder strings.Builder
		for index := 0; index < len(fieldMappings); index++ {
			builder.WriteString(strings.Trim(string(line[fieldMappings[index].StartPosition:fieldMappings[index].EndPosition]), " "))
			builder.WriteString(",")
		}
		content := builder.String()
		emit(content[0 : len(content)-1])
	}, windowedRows)
	mergedRows := beam.WindowInto(scope, window.NewGlobalWindows(), transformedRows)
	textio.Write(scope, targetFilePath, mergedRows)
	if _, errorOnExecution := direct.Execute(context.Background(), pipeline); errorOnExecution != nil {
		return FAILED
	} else {
		return SUCCESS
	}
}
