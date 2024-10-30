## The BTBench Benchmark Suite(Under Updating)

A Benchmark for Comprehensive Binary Translation Performance Evaluation.

## Building

```shell
# install prerequests
./install.sh
# compile workloads
./btbench.py -a install
```

## Running

To run a benchmark with your callback, run:
```bash
./btbench.py [-h] [-a {install,run,uninstall}] [-b BENCHMARK] [-n ITERATIONS] [-c CMD_PREFIX] [-l]
```

./btbench.py usage:
```shell
usage: btbench.py [-h] [-a {install,run,uninstall}] [-b BENCHMARK] [-n ITERATIONS] [-c CMD_PREFIX] [-l]

Run BTBench benchmarks

options:
  -h, --help            show this help message and exit
  -a {install,run,uninstall}, --action {install,run,uninstall}
                        action
  -b BENCHMARK, --benchmark BENCHMARK
                        comma separated workloads name, or c/cpp/go/php_jit/js/php/python/rust/jit/no_jit
  -n ITERATIONS, --iterations ITERATIONS
                        Run each benchmark N times.
  -c CMD_PREFIX, --cmd_prefix CMD_PREFIX
                        cmd prefix before real cmd, %s for output, eg: -c "perf stat -o %s "
  -l, --loose           ignore errors
```

## TODO

* Building all workloads from sources.
* Upload detailed workloads description.
* Upload some test results.
* Creating multi-arch docker images.

