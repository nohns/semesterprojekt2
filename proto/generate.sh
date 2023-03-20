#!/bin/bash
#Collection of scripts to generate the protobuf files

echo "Generating protobuf code for ./cloud/app..."
cd ./cloud/app && npx buf generate || { echo "Failed to generate protobuf code for ./cloud/app"; exit 1; }

echo "Changing directory to ../../..."
cd ../../ || { echo "Failed to change directory to ../../"; exit 1; }

echo "Generating protobuf code for ./cloud/bridge..."
cd ./cloud/bridge && buf generate || { echo "Failed to generate protobuf code for ./cloud/bridge"; exit 1; }

echo "Changing directory to ../../..."
cd ../../ || { echo "Failed to change directory to ../../"; exit 1; }

echo "Generating protobuf code for ./cloud/events..."
cd ./cloud/events && buf generate || { echo "Failed to generate protobuf code for ./cloud/events"; exit 1; }

echo "Done."
