package utils

type OauthLoginResp struct {
	Success bool   `json:"success"`
	Token   string `json:"token"`
}

type PortalCheckResp struct {
	Success  bool   `json:"success"`
	UserID   string `json:"userId"`
	UserName string `json:"userName"`
	UserType string `json:"userType"`
	Language string `json:"language"`
}

type SimsoCheckResp struct {
	Xm      string `json:"xm"`
	Success bool   `json:"success"`
	Xhzgh   string `json:"xhzgh"`
	Xsmc    string `json:"xsmc"`
	Sfxn    bool   `json:"sfxn"`
	Xsdm    string `json:"xsdm"`
}

type applyHistory struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Row  []struct {
		Sqbh        string `json:"sqbh"`
		Rxjzd       string `json:"rxjzd"`
		Jzdbjyzzj14 string `json:"jzdbjyzzj14"`
		Szxq        string `json:"szxq"`
		Jzdbjjd     string `json:"jzdbjjd"`
		Jzdjwdjsdm  string `json:"jzdjwdjsdm"`
		Sfyxtycj    string `json:"sfyxtycj"`
		Shbz        string `json:"shbz"`
		Zjjl        []struct {
			Formattime string `json:"formattime"`
			Io         string `json:"io"`
			Doorname   string `json:"doorname"`
			ID         int    `json:"id"`
		} `json:"zjjl"`
		Jzdjwdjrq string `json:"jzdjwdjrq"`
		Email     string `json:"email"`
		Rxcs      string `json:"rxcs"`
		Cxxdgj    string `json:"cxxdgj"`
		Tjrq      string `json:"tjrq"`
		Rowno     int    `json:"rowno"`
		Sqrqstr   string `json:"sqrqstr"`
		Sfqdcxrq  string `json:"sfqdcxrq"`
		Xslb      string `json:"xslb"`
		Rxxm      string `json:"rxxm"`
		Jzdbjqx   string `json:"jzdbjqx"`
		Tjr       string `json:"tjr"`
		Cxmdd     string `json:"cxmdd"`
		Rxzjbz    string `json:"rxzjbz"`
		Rxrq      string `json:"rxrq"`
		Czrq      string `json:"czrq"`
		Tjbz      string `json:"tjbz"`
		Cxrq      string `json:"cxrq"`
		Cxzjbz    string `json:"cxzjbz"`
		Jzdbjdjrq string `json:"jzdbjdjrq"`
		Xh        string `json:"xh"`
		Sqlb      string `json:"sqlb"`
		Shrq      string `json:"shrq"`
		Lxdh      string `json:"lxdh"`
		Cxxm      string `json:"cxxm"`
		Sqr       string `json:"sqr"`
		Dfx14Qrbz string `json:"dfx14qrbz"`
		Crxsy     string `json:"crxsy"`
		Cxcs      string `json:"cxcs"`
		Jzdjwgjdq string `json:"jzdjwgjdq"`
		Czr       string `json:"czr"`
		Sfqdhxrq  string `json:"sfqdhxrq"`
		Jzdjwssdm string `json:"jzdjwssdm"`
		Sqrq      string `json:"sqrq"`
	} `json:"row"`
	Success   bool  `json:"success"`
	Timestamp int64 `json:"timestamp"`
}


type SaveSqxxReq struct {
	Sqbh        string `json:"sqbh"`
	Sqlb        string `json:"sqlb"`
	Szxq        string `json:"szxq"`
	Email       string `json:"email"`
	Lxdh        string `json:"lxdh"`
	Crxsy       string `json:"crxsy"`
	Cxrq        string `json:"cxrq"`
	Cxcs        int    `json:"cxcs"`
	Cxxm        string `json:"cxxm"`
	Cxxdgj      string `json:"cxxdgj"`
	Sfqdhxrq    string `json:"sfqdhxrq"`
	Rxrq        string `json:"rxrq"`
	Rxcs        int    `json:"rxcs"`
	Rxxm        string `json:"rxxm"`
	Cxmdd       string `json:"cxmdd"`
	Rxjzd       string `json:"rxjzd"`
	Jzdbjqx     string `json:"jzdbjqx"`
	Jzdbjjd     string `json:"jzdbjjd"`
	Jzdbjyzzj14 string `json:"jzdbjyzzj14"`
	Jzdbjdjrq   string `json:"jzdbjdjrq"`
	Jzdjwgjdq   string `json:"jzdjwgjdq"`
	Jzdjwssdm   string `json:"jzdjwssdm"`
	Jzdjwdjsdm  string `json:"jzdjwdjsdm"`
	Jzdjwdjrq   string `json:"jzdjwdjrq"`
	Sfqdcxrq    string `json:"sfqdcxrq"`
	Dfx14Qrbz   string `json:"dfx14qrbz"`
	Sfyxtycj    string `json:"sfyxtycj"`
	Tjbz        string `json:"tjbz"`
	Shbz        string `json:"shbz"`
	Sfkcx       bool   `json:"sfkcx"`
}

type SaveSqxxResp struct {
	Code      int    `json:"code"`
	Msg       string `json:"msg"`
	Row       string `json:"row"`
	Success   bool   `json:"success"`
	Timestamp int64  `json:"timestamp"`
}

type SimsoLoginResp struct {
	Success   bool   `json:"success"`
	Sessionid string `json:"sessionid"`
	Sid       string `json:"sid"`
}


type SqztResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Row  struct {
		Sfyxsq   string `json:"sfyxsq"`
		LastSqxx struct {
			Sqbh        string `json:"sqbh"`
			Rxjzd       string `json:"rxjzd"`
			Jzdbjyzzj14 string `json:"jzdbjyzzj14"`
			Szxq        string `json:"szxq"`
			Jzdbjjd     string `json:"jzdbjjd"`
			Jzdjwdjsdm  string `json:"jzdjwdjsdm"`
			Sfyxtycj    string `json:"sfyxtycj"`
			Shbz        string `json:"shbz"`
			Jzdjwdjrq   string `json:"jzdjwdjrq"`
			Email       string `json:"email"`
			Rxcs        string `json:"rxcs"`
			Cxxdgj      string `json:"cxxdgj"`
			Tjrq        string `json:"tjrq"`
			Sfqdcxrq    string `json:"sfqdcxrq"`
			Xslb        string `json:"xslb"`
			Rxxm        string `json:"rxxm"`
			Jzdbjqx     string `json:"jzdbjqx"`
			Tjr         string `json:"tjr"`
			Cxmdd       string `json:"cxmdd"`
			Rxzjbz      string `json:"rxzjbz"`
			Rxrq        string `json:"rxrq"`
			Czrq        string `json:"czrq"`
			Tjbz        string `json:"tjbz"`
			Cxrq        string `json:"cxrq"`
			Cxzjbz      string `json:"cxzjbz"`
			Jzdbjdjrq   string `json:"jzdbjdjrq"`
			Xh          string `json:"xh"`
			Sqlb        string `json:"sqlb"`
			Shrq        string `json:"shrq"`
			Lxdh        string `json:"lxdh"`
			Cxxm        string `json:"cxxm"`
			Sqr         string `json:"sqr"`
			Dfx14Qrbz   string `json:"dfx14qrbz"`
			Crxsy       string `json:"crxsy"`
			Cxcs        string `json:"cxcs"`
			Jzdjwgjdq   string `json:"jzdjwgjdq"`
			Czr         string `json:"czr"`
			Sfqdhxrq    string `json:"sfqdhxrq"`
			Jzdjwssdm   string `json:"jzdjwssdm"`
			Sqrq        string `json:"sqrq"`
		} `json:"lastSqxx"`
		Jbxx struct {
			Xh   string `json:"xh"`
			Xm   string `json:"xm"`
			Xsmc string `json:"xsmc"`
			Xslb string `json:"xslb"`
		} `json:"jbxx"`
	} `json:"row"`
	Success   bool  `json:"success"`
	Timestamp int64 `json:"timestamp"`
}
