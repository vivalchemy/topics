
REFRESH_TOKEN=$1

curl --location --request POST 'http://localhost:4000/token' \
--header 'Content-Type: application/json' \
--data '{"token": "'"$REFRESH_TOKEN"'"}'
