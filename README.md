# rJob
rJob is a remote job service for Linux systems.

## Server
The rjob server runs on the system you'd like to run jobs on.
```
Synopsis
  rjob start
```
  

## Remote Client
The rjob command line client is used to connect to the rJob service and start, stop, list, and query jobs.
```
Synopsis
  rclient start [OPTION…] COMMAND
  rclient stop [OPTION…]
  rclient list
  rclient monitor
  rclient status

Options
  Start Options
    --cpu
      CPU limit in percent (0-100)
    --memory
      Memory limit in KB
     --io
      IO limit in percent (0-100)
  General Options
    --jobid
      The ID of the job to interact with
    --target
      Hostname and port of server
```

## Requirements
* The server application must be run as root.
* Cgroups must be enabled along with support for the io, cpu, and memory
  subsystems.
* All certs and keys are currently hardcoded. Place them in `/tmp/rjob/ssl` before
  using the server or client.

## Quick setup
1. Copy the example `ssl` directory to `/tmp/rjob`
2. Run `make build`
3. Start the server `sudo in/rjob start`
4. Start a job `bin/rclient --target localhost:9080 start ip link`
5. Check output of the job `bin/rclient --target localhost:9080 monitor --jobid <JOBID>`
