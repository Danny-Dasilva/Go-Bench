./build-docker.sh
./run-docker.sh
./server4benchmarks &  # Inside Docker.
python3 benchmark.py   # Inside Docker.
