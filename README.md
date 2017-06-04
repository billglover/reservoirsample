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

For example:

```
▸ ./reservoirsample -i in.txt -n 3
N
L
O
```

## Implementation

<div style="position:relative;height:0;padding-bottom:56.25%"><iframe src="https://www.youtube.com/embed/BzGSTknzp8c?ecver=2" width="640" height="360" frameborder="0" style="position:absolute;width:100%;height:100%;left:0" allowfullscreen></iframe></div>
