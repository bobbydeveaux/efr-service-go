aws dynamodb create-table \
    --table-name Tickets \
    --attribute-definitions \
        AttributeName=TicketID,AttributeType=N \
    --key-schema AttributeName=TicketID,KeyType=HASH   \
    --provisioned-throughput ReadCapacityUnits=1,WriteCapacityUnits=1 \
    --endpoint-url http://localhost:8000