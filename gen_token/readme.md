From : (https://docker.atlassian.net/wiki/display/OR/Generating+request+for+CVE+db+URLs)[https://docker.atlassian.net/wiki/display/OR/Generating+request+for+CVE+db+URLs]

docker build -t clemenko/docker_scanning_database .


docker run --rm -it -v /Users/clemenko/Desktop/:/data/ clemenko/docker_scanning_database 
