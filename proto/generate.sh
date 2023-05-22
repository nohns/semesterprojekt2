#!/bin/bash
#Collection of scripts to generate the protobuf files

lock="./lock"

#lock
echo "Generating protobuf code for $lock..."
cd $lock && npx buf generate || { echo "Failed to generate protobuf code for $lock"; exit 1; }
echo "Changing directory to ../../..."
cd ../ || { echo "Failed to change directory to ../"; exit 1; }

#create string variabel called pairing
pairing="./pairing"

