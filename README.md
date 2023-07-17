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
You can visualize and extract information from DJI file formats such as CSV, KML, and GPX using the parsing tool.
The parsed information can be saved into an alternative file format when inputted an output file path.
The image below includes an example of a parsed file output and the type of data extracted from the file.


### Steganography
Steganography refers to the process of revealing information stored within files.
The DroneXtract steganography suite allows you to extract telemetry and valuable data from image and video formats.
Additionally, the extracted data can be exported to four different file formats.


### Telemetry Visualization
The telemetry visualization suite contains a flight path mapping generator and a telemetry graph visualizer.
The flight path mapping generator creates an image of a map indicating the locations the drone traveled to enroute and the path it took.
The telemetry graph visualizer plots a graph for each of the relevant telemetry or sensor values to be used for auditing purposes. 


### Flight and Integrity Analysis
The flight and integrity analysis tool iterates through all the telemetry values the drone logged during its flight.
Once the values are collected, it calculates the maximum variance assumed by the value and checks for suspicious data gaps.
This tool can be used to check for anomalous data or any file corruption that may have taken place/


## Usage
To build from source, you will need Go installed.

```bash
$ export GO111MODULE=on
$ go get ./...
$ go run main.go
```

## Configuration
There are a set of environment variables utilized in DroneXtract. In order to tailor the values to your specific drone / investigation scenario, you can go to the `env.txt` file and adjust the following values:

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
