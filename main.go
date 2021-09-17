package jessica

import (
	"net/http"

	. "github.com/logrusorgru/aurora"
	. "github.com/qnsoft/live_sdk"
	"github.com/qnsoft/live_utils"
)

var config struct {
	ListenAddr    string
	CertFile      string
	KeyFile       string
	ListenAddrTLS string
}

func init() {
	plugin := &PluginConfig{
		Name:   "LiveWebSocket",
		Config: &config,
		Run:    run,
	}
	InstallPlugin(plugin)
}
func run() {
	if config.ListenAddr != "" || config.ListenAddrTLS != "" {
		utils.Print(Green("LiveWebSocket start at"), BrightBlue(config.ListenAddr), BrightBlue(config.ListenAddrTLS))
		utils.ListenAddrs(config.ListenAddr, config.ListenAddrTLS, config.CertFile, config.KeyFile, http.HandlerFunc(WsHandler))
	} else {
		utils.Print(Green("LiveWebSocket start reuse gateway port"))
		http.HandleFunc("/jessica/", WsHandler)
	}
}
