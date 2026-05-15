package verify

import (
	"log/slog"

	"github.com/adr/ad-guidance-tool/internal/ade/application/shared"
	"github.com/adr/ad-guidance-tool/internal/ade/domain"
	"github.com/adr/ad-guidance-tool/internal/ade/rule"
)

type VerifyInfo struct {
	InputFile  string
	PluginName string
	RootDir    string
}

func Verify(info VerifyInfo) {
	slog.Debug("starting verify", "file", info.InputFile)

	ir, err := shared.CompileSpec(info.InputFile)
	domain.CheckFatalError(err, "loading spec")

	if info.RootDir != "" {
		ir.OutputDir = info.RootDir
	}
	ir.Mode = rule.InvocationMode_MODE_VERIFY

	slog.Debug("executing plugin", "plugin", info.PluginName)

	err = shared.RunPlugin(info.PluginName, ir)
	domain.CheckFatalError(err, "running plugin")

	slog.Debug("plugin finished", "plugin", info.PluginName)
}
