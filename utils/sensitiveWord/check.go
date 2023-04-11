package sensitiveWord

import (
	"forum/global"
	"forum/model/forum"
	"io/ioutil"
	"os"
	"strings"

	"go.uber.org/zap"
)

const (
	OtherSen      = "./static/dictionary/其他词库.txt"
	Reactionary   = "./static/dictionary/反动词库.txt"
	ViolenceCould = "./static/dictionary/暴恐词库.txt"
	PeoLivelihood = "./static/dictionary/民生词库.txt"
	Porm          = "./static/dictionary/色情词库.txt"
	Corruption    = "./static/dictionary/贪腐词库.txt"
)

var s *SensitiveMap

type SensitiveMap struct {
	sensitiveNode map[string]interface{}
	isEnd         bool
}

type Target struct {
	Indexes []int
	Len     int
}

// Check 检查是否有敏感词汇
func Check(content string) (*[]forum.SensitiveWord, error) {
	var res []forum.SensitiveWord
	sensitiveMap := getMap()
	target := sensitiveMap.FindAllSensitive(content)
	for key, value := range target {
		t := forum.SensitiveWord{
			Word:    key,
			Indexes: value.Indexes,
			Length:  value.Len,
		}
		res = append(res, t)
	}
	return &res, nil
}

// FindAllSensitive 找到所有的敏感词
func (s *SensitiveMap) FindAllSensitive(text string) map[string]*Target {
	content := []rune(text)
	contentLength := len(content)
	result := false
	ta := make(map[string]*Target)
	for index := range content {
		sMapTmp := s
		target := ""
		in := index
		result = false
		for {
			wo := string(content[in])
			target += wo
			if _, ok := sMapTmp.sensitiveNode[wo]; ok {
				if sMapTmp.sensitiveNode[wo].(*SensitiveMap).isEnd {
					result = true
					break
				}
				if in == contentLength-1 {
					break
				}
				sMapTmp = sMapTmp.sensitiveNode[wo].(*SensitiveMap)
				in++
			} else {
				break
			}
		}
		if result {
			if _, targetInTa := ta[target]; targetInTa {
				ta[target].Indexes = append(ta[target].Indexes, index)
			} else {
				ta[target] = &Target{
					Indexes: []int{index},
					Len:     len([]rune(target)),
				}
			}
		}
	}
	return ta
}

func getMap() *SensitiveMap {
	if s == nil {
		var Sen []string
		Sen = append(Sen, OtherSen, Reactionary, ViolenceCould, PeoLivelihood, Porm, Corruption)
		s = InitDictionary(s, Sen)
	}
	return s
}

func InitDictionary(s *SensitiveMap, dictionary []string) *SensitiveMap {
	s = initSensitiveMap()
	var dictionaryContent []string
	var lenI = len(dictionary)
	for i := 0; i < lenI; i++ {
		dictionaryContentTmp := ReadDictionary(dictionary[i])
		dictionaryContent = append(dictionaryContent, dictionaryContentTmp...)
	}
	for _, words := range dictionaryContent {
		sMapTmp := s
		w := []rune(words)
		wordsLength := len(w)
		for i := 0; i < wordsLength; i++ {
			t := string(w[i])
			isEnd := false
			if i == (wordsLength - 1) {
				isEnd = true
			}
			func(tx string) {
				if _, ok := sMapTmp.sensitiveNode[tx]; !ok {
					sMapTemp := new(SensitiveMap)
					sMapTemp.sensitiveNode = make(map[string]interface{})
					sMapTemp.isEnd = isEnd
					sMapTmp.sensitiveNode[tx] = sMapTemp
				}
				sMapTmp = sMapTmp.sensitiveNode[tx].(*SensitiveMap)
				sMapTmp.isEnd = isEnd
			}(t)
		}
	}
	return s
}

func initSensitiveMap() *SensitiveMap {
	return &SensitiveMap{
		sensitiveNode: make(map[string]interface{}),
		isEnd:         false,
	}
}

func ReadDictionary(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		global.GVA_LOG.Error("read file failed:", zap.Error(err))
		return nil
	}
	defer file.Close()
	str, err := ioutil.ReadAll(file)
	dictionary := strings.Fields(string(str))
	return dictionary
}
