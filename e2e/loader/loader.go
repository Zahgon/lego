package loader

import (
	"bytes"
	"context"
	"io"
	"os/exec"
	"testing"
)

const (
	cmdNamePebble   = "pebble"
	cmdNameChallSrv = "pebble-challtestsrv"
)

const (
	envLegoTests = "LEGO_E2E_TESTS"

	envLegoTestsSkipClean = "LEGO_E2E_TESTS_SKIP_CLEAN"
)

type CmdOption struct {
	HealthCheckURL string
	Args           []string
	Env            []string
	Dir            string
}

type EnvLoader struct {
	PebbleOptions *CmdOption
	LegoOptions   []string
	ChallSrv      *CmdOption
	lego          string
}

func (l *EnvLoader) MainTest(ctx context.Context, m *testing.M) int {
	_ = "STUB: not implemented"
	return 0
}

func (l *EnvLoader) RunLegoCombinedOutput(ctx context.Context, arg ...string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (l *EnvLoader) RunLego(ctx context.Context, arg ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *EnvLoader) RunLegoWithInput(ctx context.Context, stdin io.Reader, arg ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *EnvLoader) launchPebble(ctx context.Context) func() { _ = "STUB: not implemented"; return nil }

func (l *EnvLoader) cmdPebble(ctx context.Context) (*exec.Cmd, *bytes.Buffer) {
	_ = "STUB: not implemented"
	return nil, nil
}

func pebbleHealthCheck(options *CmdOption) { _ = "STUB: not implemented"; return }

func (l *EnvLoader) launchChallSrv(ctx context.Context) func() {
	_ = "STUB: not implemented"
	return nil
}

func (l *EnvLoader) cmdChallSrv(ctx context.Context) (*exec.Cmd, *bytes.Buffer) {
	_ = "STUB: not implemented"
	return nil, nil
}

func buildLego(ctx context.Context) (string, func(), error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func getProjectRoot(ctx context.Context) (string, error) { _ = "STUB: not implemented"; return "", nil }

func build(ctx context.Context, binary string) error { _ = "STUB: not implemented"; return nil }

func goToolPath(ctx context.Context) (string, error) { _ = "STUB: not implemented"; return "", nil }

func goTool(ctx context.Context) (string, error) { _ = "STUB: not implemented"; return "", nil }

func CleanLegoFiles(ctx context.Context) { _ = "STUB: not implemented"; return }
