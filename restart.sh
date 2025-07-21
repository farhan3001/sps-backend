#!/bin/bash
kill $(ps aux | grep "[m]ain" | awk '{print $2}') 2>/dev/null
nohup ./main >output.log 2>&1 </dev/null &