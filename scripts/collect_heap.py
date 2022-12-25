from os import system
from time import sleep

DATA_DIR = "memtest/data"
HEAP_API_URL = "http://localhost:8080/debug/pprof/heap"

def dump_heap(path: str):
    """Dumps app heap usage to file"""
    system(f"curl -s http://localhost:8080/debug/pprof/heap > {path}")

def collect(interval=30):
    """Collects heap usage"""
    count = 0
    while True:
        dump_heap(f"{DATA_DIR}/heap.{count}.pprof")
        count += 1
        if count == 0:
            print("  Writed first dump")
        print(f"  Writed {count} dumps", end='\r')
        sleep(interval)

if __name__ == '__main__':
    collect()
