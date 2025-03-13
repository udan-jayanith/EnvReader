package envReader

import (
	"bufio"
	"fmt"
	"os"
)

type Env struct {
	envMap map[string]string
}

func (env *Env) setKeyValue(key, value string) {
	if env.envMap == nil {
		env.envMap = map[string]string{}
	}

	env.envMap[key] = value
}

var EnvFile Env

func LoadEnv(){
	file, err := os.Open(".env")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		envFileLine := scanner.Bytes()
		key, value := "", ""
		isKey := true
		for _, v := range envFileLine {
			// 32 means space. A tab is equal 4 or 3 spaces(32)
			if v == 32 {
				continue
			} else if v == 35 { // 35 means #(comment)
				break
			} else if isKey {
				// 61 means =
				if v == 61 {
					isKey = false
					continue
				}
				key += string(v)
			} else {
				// 34 means "
				if v == 34 {
					continue
				}
				value += string(v)
			}
		}

		if key == "" || key == " " {
			continue
		}
		EnvFile.setKeyValue(key, value)
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println(err)
	}
}

func (env *Env) get(value string) string {
	if env.envMap == nil {
		LoadEnv()
	}
	return env.envMap[value]
}

func Get(value string) string{
	return EnvFile.get(value)
}