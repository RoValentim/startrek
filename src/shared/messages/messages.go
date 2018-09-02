package messages

type returnListStruct struct {
        Status  int     `json:"status"`
        Message string  `json:"message"`
}

var ReturnList []returnListStruct

func init() {
        ReturnList = []returnListStruct {
                returnListStruct {
                        Status: 0,
                        Message: "Successful",
                },
                returnListStruct {
                        Status: 1,
                        Message: "", //Error from GoLang
                },
                returnListStruct {
                        Status: 2,
                        Message: "This word could not be translated.",
                },
        }
}
