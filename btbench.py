#!/bin/env python3
from pathlib import Path
import os
import sys
from datetime import datetime
import time
import argparse
import math

parser = argparse.ArgumentParser(description =
"""
Run BTBench benchmarks
""", formatter_class=argparse.RawTextHelpFormatter)
parser.add_argument('-a', '--action', default="run", choices=['install', 'run', 'uninstall'], help="action")
parser.add_argument('-b', '--benchmark', default="all", help="comma separated workloads name, or c/cpp/go/php_jit/js/php/python/rust/jit/no_jit")
parser.add_argument('-n', '--iterations', type=int, default=1, help="Run each benchmark N times.")
parser.add_argument('-c', '--cmd_prefix', default='', help=r'cmd prefix before real cmd, %%s for output, eg: -c "perf stat -o %%s "')
parser.add_argument('-l', '--loose', action='store_true', help="ignore errors")
# parser.add_argument('-v', '--verbose', action='store_true', help="Print cmd before exec cmd")
# parser.add_argument('--title', default="test_title")
# parser.add_argument('--stamp', default="time")
# parser.add_argument('--result_dir', default=os.path.expanduser('~') + '/runspec_result', help="location of cmd_prefix logs, defaults to ~/runspec_result")
args = parser.parse_args()

# set some environment variables to get more stable results
os.environ["PYTHONDONTWRITEBYTECODE"] = "1"
os.environ["GOMAXPROCS"] = "1"
os.environ["UV_THREADPOOL_SIZE"] = "1"

benchmark_info = {
    "gcc"           : {"dir" : "c_gcc",                "lan" : "c",          "jit" : False, },
    "x264"          : {"dir" : "c_x264",               "lan" : "c",          "jit" : False, },
    "xz"            : {"dir" : "c_xz",                 "lan" : "c",          "jit" : False, },
    "nlohmann_json" : {"dir" : "cpp_nlohmann_json",    "lan" : "cpp",        "jit" : False, },
    "xalan"         : {"dir" : "cpp_xalan",            "lan" : "cpp",        "jit" : False, },
    "buntdb"        : {"dir" : "go_buntdb",            "lan" : "go",         "jit" : False, },
    "gse"           : {"dir" : "go_gse",               "lan" : "go",         "jit" : False, },
    "jit_zend"      : {"dir" : "jit_php_zend",         "lan" : "php_jit",    "jit" : True , },
    "jit_adatron"   : {"dir" : "jit_pypy_adatron",     "lan" : "python_jit", "jit" : True , },
    "jit_aes"       : {"dir" : "jit_pypy_aes",         "lan" : "python_jit", "jit" : True , },
    "navier_stokes" : {"dir" : "octane_NavierStokes",  "lan" : "js",         "jit" : True , },
    "pdfjs"         : {"dir" : "octane_PdfJS",         "lan" : "js",         "jit" : True , },
    "script_compile": {"dir" : "octane_Typescript",    "lan" : "js",         "jit" : True , },
    "zlib"          : {"dir" : "octane_zlib",          "lan" : "js",         "jit" : True , },
    "phpbench"      : {"dir" : "php_phpbench",         "lan" : "php",        "jit" : False, },
    "zend"          : {"dir" : "php_zend",             "lan" : "php",        "jit" : False, },
    "adatron"       : {"dir" : "python_adatron",       "lan" : "python",     "jit" : False, },
    "aes"           : {"dir" : "python_aes",           "lan" : "python",     "jit" : False, },
    "mdbook"        : {"dir" : "rust_mdbook",          "lan" : "rust",       "jit" : False, },
    "rust_bio"      : {"dir" : "rust_rust-bio",        "lan" : "rust",       "jit" : False, },
    "svgbob"        : {"dir" : "rust_svgbob",          "lan" : "rust",       "jit" : False, },
}


all_benchmark = list(benchmark_info.keys())

benchmarks = []
barg = args.benchmark.strip(",").split(",")
for bselect in barg:
    if bselect == "all":
        benchmarks = all_benchmark
    elif bselect == "jit":
        benchmarks.extend([k for k,v in benchmark_info.items() if v['jit']])
    elif bselect == "no_jit":
        benchmarks.extend([k for k,v in benchmark_info.items() if not v['jit']])
    elif bselect in [ "c", "cpp", "go", "php_jit", "js", "php", "python", "rust",]:
        benchmarks.extend([k for k,v in benchmark_info.items() if v['lan'] == bselect])
    elif bselect in benchmark_info:
        benchmarks.append(bselect)
    else:
        print("can not find workload", bselect)
        exit(1)

benchmarks = sorted(list(set(benchmarks)))
benchmarks_score = {}

assert Path("bench").exists()

print("selected benchmarks:", benchmarks)

cmd_prefix = args.cmd_prefix
log_dir = Path("log_" + datetime.now().strftime("%Y_%m_%d_%H_%M_%S"))
if "%s" in cmd_prefix:
    os.makedirs(log_dir)

log_dir = log_dir.absolute()

def system_run(dir, cmd):
    shell_cmd = "cd %s && %s" % (dir, cmd)
    # print(shell_cmd)
    begin = time.time()
    r = os.system(shell_cmd)
    duration = time.time() - begin
    if r:
        print("error", shell_cmd)
        return -1
    return duration

def run_one(benchmark):
    print("running", benchmark)
    work_dir = Path("bench") / benchmark_info[benchmark]["dir"]
    init_sh = work_dir / "init.sh"
    install_sh = work_dir / "install.sh"
    uninstall_sh = work_dir / "uninstall.sh"
    if args.action == "run":
        if install_sh.exists():
            if system_run(work_dir, "./install.sh") < 0:
                return False
        if init_sh.exists():
            if system_run(work_dir, "./init.sh") < 0:
                return False
        dt = []
        for it in range(args.iterations):
            if cmd_prefix:
                if "%s" in cmd_prefix:
                    os.environ["CMD_PREFIX"] = cmd_prefix % (str(log_dir / benchmark) + "_"  + str(it))
                else:
                    os.environ["CMD_PREFIX"] = cmd_prefix

                print(os.environ["CMD_PREFIX"])

            duration = system_run(work_dir, "./run.sh")
            if duration < 0:
                print("FAILED:", benchmark)
                if not args.loose:
                    exit(1)
            else:
                dt.append(duration)
                print("SUCESS:", benchmark, "round %d" % it ,duration)
        if dt:
            if len(dt) < 3:
                dt = sum(dt) / len(dt)
            else:
                dt = sorted(dt)[len(dt) // 2]
            benchmarks_score[benchmark] = dt

    elif args.action == "install":
        if install_sh.exists():
            system_run(work_dir, "./install.sh")
    elif args.action == "uninstall":
        if uninstall_sh.exists():
            system_run(work_dir, "./uninstall.sh")
        system_run(".", "rm -rf bench/*/stdout.txt bench/*/*/stdout.txt bench/*/stderr.txt bench/*/*/stderr.txt")
    return True


for benchmark in benchmarks:
    if not run_one(benchmark):
        print("FAILED:", benchmark)

def geomean(xs):
    return math.exp(math.fsum(math.log(x) for x in xs) / len(xs))

if benchmarks_score:
    for k,v in benchmarks_score.items():
        print("%s:%s" % (k, v))

    print("score:%f" % geomean(benchmarks_score.values()))
