# Sam Knows Metric Analyser 

## Documentation

### Testing

To run the tests use the test make command.
The make command first build the application so that end to end test can be run along with Unit tests.

```
make test
```

### Build

To build the script use the build make command

```
make build
```

### Run

The script (`./sam`) requires you to provide to arguments to run an input file location and an output file location
The file at the input file location must exist.
The location for the output file location cannot already exist.
```
./sam ./inputs/1.json ./report
```

## Assumptions

  - That the report will only need to be rendered in Megabits per Second
  - That the will only been a single period of under performance per data set
  - That under performance is defined as a when the metricValue is less that the average metricValue minus half the minimum metricValue 
  
## Requirements


<div align="center">
  <img src="https://samknows.com/img/sk-logo.svg" align="center" width="60">
  <h1 align="center">SamKnows Backend Engineering Test</h1>
</div>

### Summary

The scope of the test is to generate the expected output files (in the `outputs` folder) given the input files (in the `inputs` folder).

The application should:
1. Display the min, max, median and average for a data set.
2. Discover under-performing periods of download performance.

If you do submit it without doing everything you'd like to do, then add a TODO file in root with the changes
you'd like to make and document any assumptions made during the implementation.

### Data
In the data provided, `dtime` represents the date of the measurement and `metricValue` represents 
the measurement in bytes per second.

### What we're looking for

We would like you to write a small application in one of the following languages/frameworks:
- PHP (Symfony 4+)
- Java (Spring)
- Go

The application should:

- Be well-structured.
- Be covered by automated tests.
- Include clear and concise commit messages.
- Include relevant documentation and/or comments.


### When you're finished

Either share your repo or send over your code in a compressed format.


Thank you and good luck!
