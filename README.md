![Current/Latest Build](https://github.com/lonecalvary78/data-loading-cli/actions/workflows/test-and-test-ci.yml/badge.svg)
## Data Loader CLI

## How to Build
To build the project, you can the normal go build command as you can see on below
```
go build
```

## How to run
To run this CLI, you can follow this command
```
data-loader-cli -profilerFile={the full path of batch profiler file} -batchJob={the target batch job}
```
Once the command was executed, this CLI will generate the a result file inside output directory.
