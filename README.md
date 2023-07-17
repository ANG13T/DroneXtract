<p align="center">
  <img align="center" alt="DroneXtract logo" width="600" src="https://github.com/ANG13T/DroneXtract/blob/main/assets/logo.png">
</p>


## About

## Preview

## Features
DroneXtract features for main suites for drone forensics and auditing. They include the following:

#### DJI File Parsing

#### Steganography

#### Telemetry Visualization

#### Flight and Integrity Analysis


## Usage
To build from source, you will need Go installed.

```bash
$ export GO111MODULE=on
$ go get ./...
$ go run main.go
```

## Configuration
There are a set of environment variables utilized in DroneXtract. In order to tailor the values to your specfic drone / investigation scenario, you can go to the `env.txt` file and adjust the following values:

- downsampling number for graph visualizations
- downsampling number for flight path markers
- downsampling for analysis
- variance values for analysis

## Learning and Resources

## Contributing
### Testing
An example command set for testing the steganography suite
```bash
$ cd steganography
$ go test
```
Example files and output directory

## Upcoming
- DUML parser for firmware integrity checking
- FLIGHT LOG TXT parsing
- GEOJSON parsing output for SRT files

## Support
- donations
- github
