package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/go-yaml/yaml"
	"log"
)
const s = `Name:         podinfo
Namespace:    test
Labels:       <none>
Annotations:  <none>
API Version:  split.smi-spec.io/v1alpha1
Kind:         TrafficSplit
Metadata:
  Creation Timestamp:  2020-06-13T17:24:11Z
  Generation:          23
  Managed Fields:
    API Version:  split.smi-spec.io/v1alpha1
    Fields Type:  FieldsV1
    fieldsV1:
      f:metadata:
        f:ownerReferences:
      f:spec:
        .:
        f:backends:
        f:service:
    Manager:    flagger
    Operation:  Update
    Time:       2020-06-14T04:45:11Z
  Owner References:
    API Version:           flagger.app/v1beta1
    Block Owner Deletion:  true
    Controller:            true
    Kind:                  Canary
    Name:                  podinfo
    UID:                   1735f61e-26a8-4a81-be8a-3392642160b3
  Resource Version:        115164
  Self Link:               /apis/split.smi-spec.io/v1alpha1/namespaces/test/trafficsplits/podinfo
  UID:                     88d7682f-e18a-43f4-ba55-ac7018731eb4
Spec:
  Backends:
    Service:  podinfo-canary
    Weight:   0
    Service:  podinfo-primary
    Weight:   100
  Service:    podinfo
Events:       <none>
`

func yp(args []string) error {
	bytes := []byte(s)
	//bytes, err := ioutil.ReadAll(os.Stdin)
	//if err != nil {
	//	log.Fatal("Could not read from Stdin", err)
	//}

	var body interface{}
	if err := yaml.Unmarshal(bytes, &body); err != nil {
		log.Fatal("Could not Unmarshal YAML", err)
	}

	body = convert(body)

	b, err := json.Marshal(body)
	if err != nil {
		log.Fatal("Could not Marshal to JSON", err)
	}

	fmt.Printf("Output: %s\n", b)
	return nil
}

func convert(i interface{}) interface{} {
	switch x := i.(type) {
	case map[interface{}]interface{}:
		m2 := map[string]interface{}{}
		for k, v := range x {
			m2[k.(string)] = convert(v)
		}
		return m2
	case []interface{}:
		for i, v := range x {
			x[i] = convert(v)
		}
	}
	return i
}
