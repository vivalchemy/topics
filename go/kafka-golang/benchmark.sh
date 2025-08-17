#!/bin/bash

# Number of requests
N=1000

# A list of random comment samples
COMMENTS=(
  "This is awesome!"
  "Great work 👍"
  "Interesting point."
  "I completely agree."
  "Well explained."
  "Thanks for sharing!"
  "This helped me a lot."
  "Can you explain more?"
  "Brilliant idea 💡"
  "Not sure I agree..."
)

# Export function so GNU parallel can use it
post_comment() {
  local text="$1"
  curl -s -X POST \
    -H "Content-Type: application/json" \
    -d "{\"text\":\"$text\"}" \
    http://localhost:3000/api/v1/comments >/dev/null
}
export -f post_comment

# Generate N random comments and send them in parallel
seq $N | parallel -j 20 post_comment \
  "{= 'echo \"${COMMENTS[$RANDOM % ${#COMMENTS[@]}]}\"' =}"
