from os import system

DATA_DIR = "memtest/data"

def dump_path(index: int):
    return f"{DATA_DIR}/heap.{index}.pprof"

def compare(base: int, target: int):
    cmd = "go tool pprof -http 0.0.0.0:2020 "
    cmd += f"-diff_base={dump_path(base)} "
    cmd += dump_path(target)
    system(cmd)

if __name__ == '__main__':
    compare(1, 2)
# system(f"go tool pprof -http 0.0.0.0:2020 -diff_base=memtest/data/heap.1.pprof memtest/data/heap.2.pprof")
