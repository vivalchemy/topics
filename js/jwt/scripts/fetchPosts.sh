AUTH_TOKEN=$1

curl -sSL --request GET 'http://localhost:3000/posts' \
--header "Authorization: Bearer $AUTH_TOKEN"
