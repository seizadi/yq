# yq
I have a probelm of using kubectl and in some cases it just displays yaml without option to
output json so I can use 'jq', so I wanted a command to convert yaml to json so I can use 
jq. I tried to use [this python tool](https://kislyuk.github.io/yq/).
I tried to use it on my use case and it failed :(
```bash
❯ k describe ts podinfo
Name:         podinfo
Namespace:    test
....
Spec:
  Backends:
    Service:  podinfo-canary
    Weight:   0
    Service:  podinfo-primary
    Weight:   100
  Service:    podinfo
Events:       <none>

k describe ts podinfo | yq .Spec
{
  "Backends": {
    "Service": "podinfo-primary",
    "Weight": 100
  },
  "Service": "podinfo"
}
```

To use it initially you would just run
```bash
k describe ts podinfo | yq | jq .Spec
```

Later I will implement it so it will call jq for you.

Looks like my convert routine needs more help since the current implementation suffers the same
fate and will not print out the missing elements:

```json
{
  "Namespace": "test",
  "Spec": {
    "Backends": {
      "Service": "podinfo-primary",
      "Weight": 100
    },
    "Service": "podinfo"
  }
}
```
You see the YAML input in yp.go, I think that YAML is not formed properly and that is why
you don't have an arry but just a single element.
