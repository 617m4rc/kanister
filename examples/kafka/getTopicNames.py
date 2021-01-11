import os 
import boto3
# Create an S3 client
s3 = boto3.client('s3', aws_access_key_id = os.environ.get('AWS_ACCESS_KEY'), aws_secret_access_key = os.environ.get('AWS_SECRET_KEY'), region_name = os.environ.get('region'))
bucket = os.environ.get('bucket')
prefix = 'topics/'  

result = s3.list_objects(Bucket=bucket, Prefix=prefix, Delimiter='/')
for o in result.get('CommonPrefixes'):
    print(o.get('Prefix'))
    