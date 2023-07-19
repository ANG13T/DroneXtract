<p align="center">
  <img align="center" alt="DroneXtract logo" width="600" src="https://github.com/ANG13T/DroneXtract/blob/main/assets/graphic.png">
</p>


## About
DroneXtract is a comprehensive digital forensics suite for DJI drones made with Golang. It can be used to analyze drone sensor values and telemetry data, visualize drone flight maps, audit for criminal activity, and extract pertinent data within multiple file formats. 

## Preview

<img alt="DroneXtract logo" width="600" src="https://github.com/ANG13T/DroneXtract/blob/main/assets/Preview.png">

## Features
DroneXtract features four main suites for drone forensics and auditing. They include the following:

### DJI File Parsing
You can visualize and extract information from DJI file formats such as CSV, KML, and GPX using the parsing tool.
The parsed information can be saved into an alternative file format when inputted an output file path.
The image below includes an example of a parsed file output and the type of data extracted from the file.

<img alt="DroneXtract logo" height="300" src="https://github.com/ANG13T/DroneXtract/blob/main/assets/demo-1.png">


### Steganography
Steganography refers to the process of revealing information stored within files.
The DroneXtract steganography suite allows you to extract telemetry and valuable data from image and video formats.
Additionally, the extracted data can be exported to four different file formats.

<img alt="DroneXtract logo" height="300" src="https://github.com/ANG13T/DroneXtract/blob/main/assets/demo-2.png">

### Telemetry Visualization
The telemetry visualization suite contains a flight path mapping generator and a telemetry graph visualizer.
The flight path mapping generator creates an image of a map indicating the locations the drone traveled to enroute and the path it took.
The telemetry graph visualizer plots a graph for each of the relevant telemetry or sensor values to be used for auditing purposes. 

<img alt="DroneXtract logo" width="600" src="https://github.com/ANG13T/DroneXtract/blob/main/assets/demo-3.png">

### Flight and Integrity Analysis
The flight and integrity analysis tool iterates through all the telemetry values the drone logged during its flight.
Once the values are collected, it calculates the maximum variance assumed by the value and checks for suspicious data gaps.
This tool can be used to check for anomalous data or any file corruption that may have taken place.

<img alt="DroneXtract logo" height="300" src="https://github.com/ANG13T/DroneXtract/blob/main/assets/demo-4.png">

## Usage
To build from source, you will need Go installed.

```bash
$ export GO111MODULE=on
$ go get ./...
$ go run main.go
```

## Airdata Usage
In order to parse out DJI flight .TXT logs, please use [Airdata's Flight Data Analysis tool](https://app.airdata.com/main?a=upload).
- Airdata CSV output file can be used for the CSV parser, flight path map, and telemetry visualizations
- Airdata KML output file can be used for the KML parser
- Airdata GPX output file can be used for GPX parser 

## Configuration
There are a set of environment variables utilized in DroneXtract. In order to tailor the values to your specific drone / investigation scenario, you can go to the `.env` file and adjust the following values:

`TELEMETRY_VIS_DOWNSAMPLE` downsampling number for values to be used for telemetry visualization

`FLIGHT_MAP_DOWNSAMPLE` downsampling number for values to be used for flight path mapping

`ANALYSIS_DOWNSAMPLE` downsampling number for values to be used for integrity analysis

`ANALYSIS_MAX_VARIANCE` maximum variance allowed between max and min value for analysis values


## Learning and Resources
To learn more about DJI drone digital forensics and the features of DroneXtract, refer to [this article](https://medium.com/@angelinatsuboi/a-comprehensive-guide-to-digital-forensics-with-dji-drones-fd7ef5af2891).

## Contributing
DroneXtract is open to any contributions. Please fork the repository and make a pull request with the features or fixes you want to be implemented.

### Testing
DroneXtract contains testing files for all four suites. 
An example command set for testing the steganography suite
```bash
$ cd steganography
$ go test
```
Example files for testing are included in the `test-data` directory.
If a specific test case includes an output such as a file or image, it will be stored within the `output` directory.
There are already some files contained in the `output` for reference purposes.

## Upcoming
- DUML parser for firmware integrity checking
- DJI Flight Log TXT parsing for the parsing suite
- GEOJSON parsing output for SRT files in the steganography suite

## Support
If you enjoyed DroneXtract, please consider [becoming a sponsor](https://github.com/sponsors/ANG13T) in order to fund my future projects. 

To check out my other works, visit my [GitHub profile](https://github.com/ANG13T).
