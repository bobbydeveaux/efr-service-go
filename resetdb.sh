aws dynamodb delete-table --region eu-west-1 --table-name Tickets --endpoint-url http://localhost:8000         
aws dynamodb delete-table --region eu-west-1 --table-name Winners --endpoint-url http://localhost:8000         

aws dynamodb create-table \
    --table-name Tickets \
    --attribute-definitions \
        AttributeName=TicketID,AttributeType=S \
    --key-schema AttributeName=TicketID,KeyType=HASH   \
    --provisioned-throughput ReadCapacityUnits=1,WriteCapacityUnits=1 \
    --region eu-west-1 \
    --endpoint-url http://localhost:8000


 aws dynamodb create-table \
    --table-name Winners \
    --attribute-definitions \
        AttributeName=WinnerID,AttributeType=N \
    --key-schema AttributeName=WinnerID,KeyType=HASH   \
    --provisioned-throughput ReadCapacityUnits=1,WriteCapacityUnits=1 \
    --region eu-west-1 \
    --endpoint-url http://localhost:8000