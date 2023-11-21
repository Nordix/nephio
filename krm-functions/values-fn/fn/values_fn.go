package fn

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/GoogleContainerTools/kpt-functions-sdk/go/fn"
	"gopkg.in/yaml.v2"
)

const (
	Namespace        = "namespace"
	ConfigMapName    = "name-values"
	APIVersion       = "v1"
	KindConfigMap    = "ConfigMap"
	DataLabel        = "data"
	ValuesFileName   = "values"
	ModifiedFileName = "values_modified"
	HelmValuesPath   = "./package/helm/"
)

func Process(rl *fn.ResourceList) (bool, error) {
	var err error
	err = modifyAndSaveValues(HelmValuesPath, "replicaCount", 4)
	if err != nil {
		return false, fmt.Errorf("replicaCount key not found in values.yaml")
	}

	valuesStr, err := extractValuesAsString(HelmValuesPath, ModifiedFileName)
	if err != nil {
		return false, err
	}

	cmko, err := createNewValuesConfigMap(valuesStr)
	if err != nil {
		return false, fmt.Errorf("failed to create object: %v", err)
	}

	err = rl.UpsertObjectToItems(cmko, nil, true)
	if err != nil {
		return false, fmt.Errorf("failed to upsert object: %v", err)
	}

	return true, nil
}

func createNewValuesConfigMap(valuesStr string) (*fn.KubeObject, error) {
	cmko := fn.NewEmptyKubeObject()

	err := cmko.SetAPIVersion(APIVersion)
	if err != nil {
		return nil, err
	}
	err = cmko.SetKind(KindConfigMap)
	if err != nil {
		return nil, err
	}
	_, err = cmko.RemoveNestedField("metadata", "creationTimestamp")
	if err != nil {
		return nil, err
	}
	err = cmko.SetName(ConfigMapName)
	if err != nil {
		return nil, err
	}
	err = cmko.SetNamespace(Namespace)
	if err != nil {
		return nil, err
	}

	err = cmko.SetNestedField(valuesStr, DataLabel)
	if err != nil {
		return nil, err
	}

	return cmko, nil
}
func extractValuesFromValuesDirectly(packagePath string) (map[string]interface{}, error) {
	filePath := filepath.Join(packagePath, ValuesFileName)

	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("cannot open the values.yaml in: %v", filePath)
	}
	defer file.Close()

	var values map[string]interface{}
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&values); err != nil {
		return nil, fmt.Errorf("failed to decode values.yaml: %v", err)
	}

	return values, nil
}
func extractValuesAsString(packagePath string, fileName string) (string, error) {
	filePath := filepath.Join(packagePath, fileName)

	filePath = filepath.Clean(filePath)
	file, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read %s: %v", fileName, err)
	}
	return string(file), nil
}
func modifyAndSaveValues(packagePath string, key string, newValue interface{}) error {
	// Step 1: Read the YAML file into a map
	values, err := extractValuesFromValuesDirectly(packagePath)
	if err != nil {
		return err
	}

	// Step 2: Modify the specific value in the map
	values[key] = newValue

	// Step 3: Encode the modified map back to YAML
	yamlBytes, err := yaml.Marshal(values)
	if err != nil {
		return fmt.Errorf("failed to encode values to YAML: %v", err)
	}

	// Step 4: Save the encoded YAML to the file
	newFilePath := filepath.Join(packagePath, ModifiedFileName)
	err = os.WriteFile(newFilePath, yamlBytes, 0600)
	if err != nil {
		return fmt.Errorf("failed to write modified YAML to file: %v", err)
	}

	return nil
}
