# Description

The program accepts an input file of load attempts by the customer.  It will write whether the load attempt is accepted or rejected to an output file.

Rules for rejection: exceeds daily load limit, exceeds weekly load limit, exceeds daily number of loads

If the same event is seen again, the program will ignore it.

Assumptions: load attempts will appear in chronological order in the input file

Requirements: must run dynamoDB locally

# Instructions 

**To run dynamoDB locally:**

In '...\golang assignment\golang-assignment\dynamodb_local_latest', run following:

```
java -Djava.library.path=./DynamoDBLocal_lib -jar DynamoDBLocal.jar -sharedDb
```

[More info] https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/DynamoDBLocal.DownloadingAndRunning.html

**Install AWS CLI:**

[More info] https://docs.aws.amazon.com/cli/latest/userguide/install-cliv2-windows.html#cliv2-windows-install

**Must configure AWS credentials via CLI:**

```
aws configure
AWS Access Key ID [None]: None
AWS Secret Access Key [None]: None
Default region name [None]: None
Default output format [None]: json
```

**To empty the local dynamoDB tables:**

```
aws dynamodb delete-table --table-name EventCorrelationTable --endpoint-url http://localhost:8000

aws dynamodb delete-table --table-name DailyCustomerLoadTable --endpoint-url http://localhost:8000


aws dynamodb create-table --table-name EventCorrelationTable --attribute-definitions AttributeName=CUSTOMER_ID,AttributeType=S AttributeName=EVENT_ID,AttributeType=S --key-schema AttributeName=CUSTOMER_ID,KeyType=HASH AttributeName=EVENT_ID,KeyType=RANGE --provisioned-throughput ReadCapacityUnits=1,WriteCapacityUnits=1 --endpoint-url http://localhost:8000

aws dynamodb create-table --table-name DailyCustomerLoadTable --attribute-definitions AttributeName=CUSTOMER_ID,AttributeType=S AttributeName=DATE,AttributeType=S --key-schema AttributeName=CUSTOMER_ID,KeyType=HASH AttributeName=DATE,KeyType=RANGE --provisioned-throughput ReadCapacityUnits=1,WriteCapacityUnits=1 --endpoint-url http://localhost:8000
```

**External libraries used:**

```
github.com/aws/aws-sdk-go
```

# Improvements

In a real world scenario, these events could be coming in real time.

To improve performance, events of different customers can be processed in parallel.

Since a customer's events need to be processed in some chornological order, we can use some queue that has ordering built in to it(SQS FIFO).

The correlation table could have some TTL on its records (24 hours), to avoid it growing in size. 

A queue would also allow for retrying in cases where I/O calls to the database fail (query, insert).

 


