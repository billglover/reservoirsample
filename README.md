# reservoirsample
Use Reservoir Sampling to efficiently sample a text file

## Background

 - https://gregable.com/2007/10/reservoir-sampling.html
 - https://en.wikipedia.org/wiki/Reservoir_sampling

## Usage

```
▸ ./reservoirsample -h
Usage of ./reservoirsample:
  -i string
    	input file (default stdin)
  -n int
    	number of samples to select (default 1) (default 1)
  -o string
    	output file (default stdout)
```

```
▸ ./reservoirsample -i in.txt -n 3
N
L
O
```

## Implementation

[![Tech Bytes : Reservoir Sampling](https://img.youtube.com/vi/BzGSTknzp8c/0.jpg)](https://www.youtube.com/watch?v=BzGSTknzp8c)
