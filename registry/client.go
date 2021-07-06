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
		return fmt.Errorf("Failed to register service, status code: %v", res.StatusCode)
	}

	return nil
}

func RemoveService(r Registration) error {
	buffer := new(bytes.Buffer)
	enc := json.NewEncoder(buffer)
	if err := enc.Encode(&r); err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("http://%s:%s/services", RegistryHost, RegistryPort), buffer)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to stop service, status code: %v", res.StatusCode)
	}

	return nil
}
