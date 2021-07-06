package config

// Config - LINE botの設定
type Config struct {
	Host                      string // LINEのホスト(スキームも含む)
	TalkServicePath           string // TalkServiceのパス
	TalkServicePathForPolling string // TalkServiceのパス(Polling用)
	QrPath                    string // QRログインのパス
	PermitNoticePath          string // クライアントへログイン処理させるためのパス
	UserAgent                 string // User-Agentヘッダ
	LINEApp                   string // X-Line-Applicationヘッダ
}

// NewConfig - Configのコンストラクタ
func NewConfig() Config {
	return Config{
		Host:                      "https://legy-jp.line.naver.jp",
		TalkServicePath:           "/S4",
		TalkServicePathForPolling: "/P4",
		QrPath:                    "/acct/lgn/sq/v1",
		PermitNoticePath:          "/acct/lp/lgn/sq/v1",
		UserAgent:                 "LLA/2.13.2 Nexus Mozilla/5.0 (Linux; Android 6.0; CPH1609 Build/MRA58K; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/55.0.2883.91 Mobile Safari/537.36",
		LINEApp:                   "ANDROID\t10.20.1\tAndroid OS\t10;SECONDARY",
	}
}
