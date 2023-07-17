<p align="center">
  <img align="center" alt="DroneXtract logo" width="600" src="https://github.com/ANG13T/DroneXtract/blob/main/assets/logo.png">
</p>


## About
DroneXtract is a comprehensive digital forensics suite for DJI drones made with Golang. It can be used to analyze drone sensor values and telemetry data, visualize drone flight maps, audit for criminal activity, and extract pertinent data within multiple file formats. 

## Preview

<img alt="DroneXtract logo" width="600" src="https://github.com/ANG13T/DroneXtract/blob/main/assets/Preview.png">

## Features
DroneXtract features for main suites for drone forensics and auditing. They include the following:

### DJI File Parsing

### Steganography

### Telemetry Visualization

### Flight and Integrity Analysis


## Usage
To build from source, you will need Go installed.

```bash
$ export GO111MODULE=on
$ go get ./...
$ go run main.go
```

## Configuration
There are a set of environment variables utilized in DroneXtract. In order to tailor the values to your specfic drone / investigation scenario, you can go to the `env.txt` file and adjust the following values:

### Environment Variables
All environment variables can be found and modified in the `.env` file 

`TELEMETRY_VIS_DOWNSAMPLE`

`FLIGHT_MAP_DOWNSAMPLE` 

`ANALYSIS_DOWNSAMPLE` 

`ANALYSIS_MAX_VARIANCE`


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
