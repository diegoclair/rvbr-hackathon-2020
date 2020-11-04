#!/bin/bash

#compile app
go build

#start app after wait the database
sh -c "/wait && ./rvbr-2020-hackathon-time-3-backend"