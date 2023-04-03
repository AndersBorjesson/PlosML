A -> B >- C -> D >- E -> F
D >- G -> H 
F, H >- I

lambda.Ingestor handles A 
dynamodb.Database handles C 
beanstalk.Compute1 handles E 
batch.Compute2 handles G 
redshift.Storage handles I
