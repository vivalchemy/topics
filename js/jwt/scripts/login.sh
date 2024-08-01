USERNAME=$1

curl --location --request POST 'http://localhost:4000/login' \
--header 'Content-Type: application/json' \
--data '{"username": "'"$USERNAME"'"}'
