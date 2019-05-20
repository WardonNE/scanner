package scanner

import (
	"github.com/json-iterator/go"
	"github.com/sipt/GoJsoner"

	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var invalidUrlRulesFilePath string

var invalidUrlRulesFileName string = "conf/invalid_url_rules.json"

func init() {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	reader := strings.NewReader(loadConfFile())
	decoder := json.NewDecoder(reader)
	jsonObject := make(map[string][]string)
	if err := decoder.Decode(&jsonObject); err == nil {
		setRules(jsonObject)
	} else {
		panic(err)
	}
}

func loadConfFile() string {
	if binRunPath, err := os.Executable(); err == nil {
		binRunDir := filepath.Dir(binRunPath)
		invalidUrlRulesFilePath = binRunDir + "/../" + invalidUrlRulesFileName
		if content, err := ioutil.ReadFile(invalidUrlRulesFilePath); err == nil {
			if jsonstr, err := GoJsoner.Discard(string(content)); err == nil {
				return jsonstr
			} else {
				panic(err)
			}
		} else {
			panic(err)
		}
	} else {
		panic(err)
	}
}

func setRules(rules map[string][]string) {
	for k, r := range rules {
		addNewRule(k, r)
	}
}
