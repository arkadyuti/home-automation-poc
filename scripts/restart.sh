#!/bin/bash

# Define the file to watch
FILE="relay"
NEW_FILE="newRelay"
BINARY_PROCESS="newRelay"
SLEEP_INTERVAL=2  # Time in seconds to sleep between checks
BINARY_PID=""

# Function to start the binary
start_binary() {
  echo "Starting binary: $BINARY_PROCESS"
  ./$BINARY_PROCESS &
  BINARY_PID=$!
}

# Function to stop the binary
stop_binary() {
  if [[ -n "$BINARY_PID" ]] && ps -p $BINARY_PID > /dev/null; then
    echo "Stopping binary: $BINARY_PROCESS (PID: $BINARY_PID)"
    kill $BINARY_PID
    wait $BINARY_PID 2>/dev/null
    BINARY_PID=""
  fi
}

# Function to check if a file has changed
check_file_change() {
  CURRENT_TIMESTAMP=$(stat -c %Y "$FILE")
  if [[ "$CURRENT_TIMESTAMP" != "$LAST_TIMESTAMP" ]]; then
    return 0  # File has changed
  else
    return 1  # File has not changed
  fi
}

# Initialize the last timestamp
if [[ -e $FILE ]]; then
  LAST_TIMESTAMP=$(stat -c %Y "$FILE")
else
  echo "Error: $FILE does not exist."
  exit 1
fi

# Main loop to monitor the file
while true; do
  if check_file_change; then
    echo "File $FILE has been modified."

    # Stop the old binary if it's running
    stop_binary

    # Rename the file to newRelay
    mv -f $FILE $NEW_FILE
    echo "Renamed $FILE to $NEW_FILE"

    # Start the new binary
    start_binary

    # Update the last modification timestamp
    LAST_TIMESTAMP=$(stat -c %Y "$NEW_FILE")
  fi

  # Sleep for a while before checking again
  sleep $SLEEP_INTERVAL
done
