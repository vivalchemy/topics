REFRESH_TOKEN=$1

curl --location --request DELETE 'http://localhost:4000/logout' \
--header 'Content-Type: application/json' \
--data '{"token": "'"$REFRESH_TOKEN"'"}'
