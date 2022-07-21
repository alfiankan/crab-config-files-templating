# Crab 🦀
Dynamic configuration file templating tool for kubernetes manifest or general configuration files

[![codecov](https://codecov.io/gh/alfiankan/crab-config-files-templating/branch/master/graph/badge.svg?token=DJHuZmbYeU)](https://codecov.io/gh/alfiankan/crab-config-files-templating)
[![Go Reference](https://pkg.go.dev/badge/github.com/alfiankan/crab-config-files-templating.svg)](https://pkg.go.dev/github.com/alfiankan/crab-config-files-templating)
[![Go Report Card](https://goreportcard.com/badge/github.com/alfiankan/crab-config-files-templating)](https://goreportcard.com/report/github.com/alfiankan/crab-config-files-templating)
[![Generate release-artifacts](https://github.com/alfiankan/crab-config-files-templating/actions/workflows/go.yml/badge.svg?branch=v1.0.5)](https://github.com/alfiankan/crab-config-files-templating/actions/workflows/go.yml)
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)  


## How to install

Download according to your computer architecture at release page [go to release page](https://github.com/alfiankan/crab-config-files-templating/releases)

| OS        | Platform           | Status  |
| ------------- |:-------------:| -----:|
| Ubuntu, Debian | amd64, arm64, 386| available |
| MacOS, Darwin  | amd64, arm64, 386| available |
| Windows | amd64, arm64, 386| available |

Extract using command like :
```bash
tar -xzf crab-cli-v0.1.0-darwin-amd64.tar.gz -C crab
```

Move to /usr/local/bin (Optional) :
```bash
cp crab/crab /usr/local/bin
```

Then check crab cli working properly :
<image cli>

## Build From Source
- clone this repository
- make sure you have go version 18
- run `make build`

## How to use
basic command fromat :
```bash
  crab -f <fileinputpath> -r <keyvalue> -o <outputfilepath>
```
Flags :
| Flag       | Description           | Required  | default |
| :------------- |:-------------| :-----| :-----|
| -f | file input path location eg. ./manifest/nginx.yaml (accept any config file ext, .env, json, yml, yaml etc.) | true | - |
| -o  | write output file location eg. ./manifest/nginx-result.yaml | false  | same as input file path (overwrite) |
| -r | key value replacable | true | - |
| -q | key value replacable with quotes | false | - |
| -v | verbose process | false | false |

### Case 1 Standart replace
First add {{replacableName}} to your file :
```yaml
apiVersion: v1
kind: Service
metadata:
    name: nginx
    namespace: {{namespace}}
spec:
    ports:
      - targetPort: 80
        name: nginxhttp
        port: {{exposePort}}
```
crab cli will replace {{replacableName}} as the key, based on the example above :
- {{namespace}}
- {{exposePort}}

then run crab command :
```bash
  crab -f inputfile.yaml -r namespace=production -r exposePort=8081
```
yes you can override multiple key values  ✔️
  
write output to another file add -o flag:
```bash
  crab -f inputfile.yaml -r namespace=production -r exposePort=8081 -o result.yaml
```
  
to make verbose add -v flag:
```bash
  crab -f inputfile.yaml -r namespace=production -r exposePort=8081 -o result.yaml -v
```
verbose output :
```bash
[REPLACED] from namespace to production
[REPLACED] from exposePort to 8081
[DONE] Crab output result at result.yaml
```

the result will be :
```yaml
apiVersion: v1
kind: Service
metadata:
    name: nginx
    namespace: production
spec:
    ports:
      - targetPort: 80
        name: nginxhttp
        port: 8081
```
Click to watch the demo :
[![asciicast](https://asciinema.org/a/E76kRn2G8uoIXBm6KZPtPn27x.svg)](https://asciinema.org/a/E76kRn2G8uoIXBm6KZPtPn27x)

Youtube demo video :
- [go to youtube](https://youtu.be/HdRQ3mz64us)

  
### Case 2 Quotes replace
sometimes your config file needs string quotes, (like for env vars or connection string):
  ```sh
    namespace="production"
  ```
template file example :
```yaml
apiVersion: v1
kind: Service
metadata:
    name: nginx
    namespace: {{namespace}}
spec:
    ports:
      - targetPort: 80
        name: nginxhttp
        port: {{exposePort}}
```
let's say we need a quote in the namespace :
```bash
  crab -f inputfile.yaml -q namespace=production -r exposePort=8081 -o result.yaml -v
```
yass you can simultaneously Replace multiple key values with quotes or not  ✔️

result will be :
```yaml
apiVersion: v1
kind: Service
metadata:
    name: nginx
    namespace: "production"
spec:
    ports:
      - targetPort: 80
        name: nginxhttp
        port: 8081
```

## Sample Github Action

```yaml
name: Demo deployment using crab

on:
  workflow_dispatch:

jobs:
  crab-manifest:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - name: mock job
        run: pwd

  deploy-test:
    runs-on: ubuntu-latest
    needs: crab-manifest
    steps:
      - uses: actions/checkout@v3
      - name: install crab cli
        run: wget -c https://github.com/alfiankan/crab-config-files-templating/releases/download/v1.0.5/crab-v1.0.5-linux-amd64.tar.gz -O - | tar -xz crab
      - name: recreate deployment manifest on test
        run: |
          ./crab -f example/manifest/nginx.yaml \
          -r namespace=test \
          -r publishPort=8000 \
          -q portName=test-server \
          -o nginx-test.yaml \
          -v
      - name: view manifest
        run: cat nginx-test.yaml
                    
  
  deploy-production:
    runs-on: ubuntu-latest
    needs: crab-manifest
    steps:
      - uses: actions/checkout@v3
      - name: install crab cli
        run: wget -c https://github.com/alfiankan/crab-config-files-templating/releases/download/v1.0.5/crab-v1.0.5-linux-amd64.tar.gz -O - | tar -xz crab
      - name: recreate deployment manifest on production
        run: |
          ./crab -f example/manifest/nginx.yaml \
          -r namespace=production \
          -r publishPort=80 \
          -q portName=prod-server \
          -o nginx-prod.yaml \
          -v
      - name: view manifest
        run: cat nginx-prod.yaml
```

## Related article
- Replacing kubernetes manifest value dynamicly (coming soon on medium)

