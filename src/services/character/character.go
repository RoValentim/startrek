package character

import (
        "fmt"
        "strings"

        "shared/defines"
        "shared/logger"
        "shared/messages"
        "shared/urls"
)

func GetSpecie(ch chan bool, name string, specie *string) {
        logger.Log("2", "Getting Specie for " + name, logger.FATAL)
        var aux string

        data := `name=` + name
        form := strings.NewReader(data)
	url  := fmt.Sprintf("%s/search", defines.StapiRESTCharacter)
	logger.Log("2", url, logger.DEBUG)

        headers := make(map[string]string)
	headers["Content-Type"]  = "application/x-www-form-urlencoded"

        ret, err := urls.SendRequest("2", "POST", url, form, headers, defines.Timeout)
	if err != nil {
                logger.Log("2", err.Error(), logger.DEBUG)
                ch <- false
                return
	}

        if ret["characters"] == nil {
                logger.Log("2", messages.ReturnList[3].Message, logger.DEBUG)
                ch <- false
                return
        }

        for _, v := range ret["characters"].([]interface{}) {
                for key, value := range v.(map[string]interface{}) {
                        if key == "uid" {
                                aux = getFullCharacter(value.(string))
                                if aux != "" {
                                        *specie = aux
                                        ch <- true
                                        return
                                }
                        }
                }
        }

        logger.Log("2", messages.ReturnList[3].Message, logger.DEBUG)
        ch <- false
}

func getFullCharacter(uid string) string {
        logger.Log("2", "Getting Full Character for " + uid, logger.FATAL)

	url := fmt.Sprintf("%s?uid=%s", defines.StapiRESTCharacter, uid)
	logger.Log("2", url, logger.DEBUG)

        ret, err := urls.SendRequest("2", "GET", url, nil, nil, defines.Timeout)
	if err != nil {
		return ""
	}

        if ret["character"] == nil {
                return ""
        }

        for k, v := range ret["character"].(map[string]interface{}) {
                if k == "characterSpecies" {
                        for _, array := range v.([]interface{}) {
                                for key, value := range array.(map[string]interface{}) {
                                        if key == "name" {
                                                return value.(string)
                                        }
                                }
                        }
                }
        }

        return ""
}
