#!/bin/bash
kill $(ps aux | grep "[m]ain" | awk '{print $2}') 2>/dev/null