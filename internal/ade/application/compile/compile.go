package compile

import (
	"log/slog"

	"github.com/adr/ad-guidance-tool/internal/ade/application/shared"
	"github.com/adr/ad-guidance-tool/internal/ade/domain"
	"github.com/adr/ad-guidance-tool/internal/ade/rule"
)

type CompileInfo struct {
	InputFile  string
	OutputDir  string
	PluginName string
}

func Compile(info CompileInfo) {
	slog.Debug("starting compilation", "file", info.InputFile)

	ir, err := shared.CompileSpec(info.InputFile)
	domain.CheckFatalError(err, "compiling spec")

	ir.OutputDir = info.OutputDir
	ir.Mode = rule.InvocationMode_MODE_COMPILE

	slog.Debug("executing plugin", "plugin", info.PluginName)

	err = shared.RunPlugin(info.PluginName, ir)
	domain.CheckFatalError(err, "running plugin")

	slog.Debug("plugin finished", "plugin", info.PluginName)
}
