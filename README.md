# kube-resource-unit-converter

Go utility for normalizing k8s resource quantities.

## Usage

```
$ go build ./...

$ # memory/space
$ echo '300Mi' | ./kube-resource-unit-converter
314572800

$ echo '3Gi' | ./kube-resource-unit-converter
3221225472

$ # for cores, you may want to pass -m to get the millivalue back
$ echo '500' | ./kube-resource-unit-converter -m
500
$ echo '1' | ./kube-resource-unit-converter -m
1000
```
