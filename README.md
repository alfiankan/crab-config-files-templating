# Crab 🦀
Dynamic configuration file templating tool for kubernetes manifest or general configuration files

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
```
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
crab cli will replace {{replacableName}} as key, base on example above :
- {{namespace}}
- {{exposePort}}

then run crab command :
```bash
  crab -f inputfile.yaml -r namespace=production -r exposePort=8081
```
yass you can replace multiple key value ✔️
  
write output to another file add -o flag:
```bash
  crab -f inputfile.yaml -r namespace=production -r exposePort=8081 -o result.yaml
```
  
to make verbose add -v flag:
```bash
  crab -f inputfile.yaml -r namespace=production -r exposePort=8081 -o result.yaml -v
```

the result will be :

  
### Case 2 Quotes replace
sometimes your config file needs string quotes :  
  ```sh
    baseUrl="http://domain.com"
  ```
template file example :
```
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
let says wi need quotes on namespace :
```bash
  crab -f inputfile.yaml -q namespace=production -r exposePort=8081 -o result.yaml -v
```
yass you can silmutanly replacing multiple key value with quotes or not ✔️

result will be

## Related article
- Replacing kubernetes manifest value dynamicly (coming soon)

