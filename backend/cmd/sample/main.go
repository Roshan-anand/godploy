package main

import "fmt"

func main() {
	envs := []string{"PROT=8080", "HOST=localhost", "DEBUG=true"}

	resolver := []struct {
		Key   string
		Value interface{}
	}{
		{"PROT", 8080},
		{"HOST", "localhost"},
		{"DEBUG", true},
	}

	for _, env := range resolver {
		val, ok := env.Value.(string)
		if ok {
			envs = append(envs, env.Key+"="+val)
		} else {
			envs = append(envs, env.Key+"="+fmt.Sprintf("%v", env.Value))
		}
	}

	fmt.Println(envs)
}
