package live_websocket

import (
	"net/http"

	"github.com/logrusorgru/aurora"
	"github.com/qnsoft/live_sdk"
	"github.com/qnsoft/live_utils"
)

var config struct {
	ListenAddr    string
	CertFile      string
	KeyFile       string
	ListenAddrTLS string
}

func init() {
	plugin := &live_sdk.PluginConfig{
		Name:   "LiveWs",
		Config: &config,
		Run:    run,
	}
	live_sdk.InstallPlugin(plugin)
}
func run() {
	if config.ListenAddr != "" || config.ListenAddrTLS != "" {
		live_utils.Print(aurora.Green("LiveWs start at"), aurora.BrightBlue(config.ListenAddr), aurora.BrightBlue(config.ListenAddrTLS))
		live_utils.ListenAddrs(config.ListenAddr, config.ListenAddrTLS, config.CertFile, config.KeyFile, http.HandlerFunc(WsHandler))
	} else {
		live_utils.Print(aurora.Green("LiveWs start reuse websocket port"))
		http.HandleFunc("/livews/", WsHandler)
	}
}
