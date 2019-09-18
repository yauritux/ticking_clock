# ticking_clock
Goes back in time Ticking Clock

## How to Run
1. Build docker image
2. Run the image

### Build Docker Image

```
docker build -t yauritux/ticking-clock:latest .
```

### Run the image

```
docker run yauritux/ticking-clock
```

will run the ticking-clock application with 1 hour as the default parameter

or, you can set the default hour to 3 by using this following command: 

```
docker run -e hour=3 yauritux/ticking-clock
```

or any other arbitrary integer value for hour such as 2, 5, etc.

## Running Test Coverage

from the terminal, within the base directory, type the following command:

```
go test --cover janio.asia/ticking_clock/pkg/util/time 
```

it should print something as follows:
```
ok      janio.asia/ticking_clock/pkg/util/time  4.018s  coverage: 100.0% of statements
```