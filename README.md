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

## Running Test Coverage

from the terminal, within the base directory, type the following command:

```
go test --cover janio.asia/ticking_clock/pkg/util/time 
```

it should print something as follows:
```
ok      janio.asia/ticking_clock/pkg/util/time  4.018s  coverage: 100.0% of statements
```