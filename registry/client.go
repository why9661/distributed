package registry

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func RegisterService(r Registration) error {
	buffer := new(bytes.Buffer)
	enc := json.NewEncoder(buffer)
	if err := enc.Encode(&r); err != nil {
		return err
	}

	res, err := http.Post(fmt.Sprintf("http://%s:%s/services", RegistryHost, RegistryPort), "application/json", buffer)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to register service with status code: %v", res.StatusCode)
	}

	return nil
}
