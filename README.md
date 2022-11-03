# simple-gfs
Simplified simulation of gfs implemented in golang

### Features
* read
* append
* concurrency
* snapshot
* replication
* primary and secondaries
* weak consistency
* simulation of success and failure

### Assumptions
* filename same as chunk handler name and expressed as integer
* each file is unlimited in size, i.e. it can be kept extending, without requirement of further partitions
* strings used rather than files for simplicity
* only read and append operations, no random writes
* read files as a whole rather than a byte range for simplicity
* reading all bits rather than specific bits
* fixed #clients, #servers
* single master
* commit is said to take place if all replicas return success
* else, the same transaction would occur again
* weighted probabilities for success and failure
* p(success): 0.8
* p(failure): 0.2

### To run
```
git clone https://github.com/mgpinf/simple-gfs.git
cd ./simple-gfs/src
go build
./simple-gfs
```
