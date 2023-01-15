package models

type BatchProfile struct {
	BatchJobs []BatchJob `yaml:"jobs"`
}

type BatchJob struct {
	Name       string         `yaml:"name"`
	SourceFile string         `yaml:"sourceFile"`
	TargetFile string         `yaml:"targetFile"`
	Mappings   []FieldMapping `yaml:"mappings"`
}

type FieldMapping struct {
	FieldName     string `yaml:"name"`
	StartPosition int    `yaml:"start_position"`
	EndPosition   int    `yaml:"end_position"`
}
