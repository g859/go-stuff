# dns-convert
Simple Go program to convert a list of DNS entries to their respective IPV4/IPV6 addresses.

## Usage

1. `make-build-local` -or- `make docker-build-local`
2. Modify `entries.txt` with approrpiate website(s)/entries
3. `./main` -or- `docker run --rm dns-convert:latest`
4. Enjoy list of sorted IPs

> The Dockerfile requires `make build` to be run instead since that will statically compile the program with the required libraries built in.

> The Dockerfile can easily be pushed to Dockerhub with `make docker-build-push`

## To-do

1. Goroutines 
2. Differentiate IPV4 vs IPV6 addresses
3. Write to CSV file for use elsewhere
4. Upload said CSV to an S3 bucket somewhere
5. Use in a K8s Cronjob to run at a consistent time pending new entries are added to the list