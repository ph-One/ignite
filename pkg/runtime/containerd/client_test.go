package containerd

import (
	"testing"

	meta "github.com/weaveworks/ignite/pkg/apis/meta/v1alpha1"
	"github.com/weaveworks/ignite/pkg/runtime"

	v2shim "github.com/containerd/containerd/runtime/v2/shim"
)

var client runtime.Interface

func init() {
	var clienterr error
	client, clienterr = GetContainerdClient()
	if clienterr != nil {
		panic(clienterr)
	}
}

var imageName, _ = meta.NewOCIImageRef("docker.io/library/busybox:latest")

func TestPull(t *testing.T) {
	t.Error(client.PullImage(imageName))
}

func TestInspect(t *testing.T) {
	t.Error(client.InspectImage(imageName))
}

/*func TestExport(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)
	fmt.Println(tempDir)

	tarCmd := exec.Command("tar", "-x", "-C", tempDir)
	reader, _, err = client.ExportImage(imageName)
	if err != nil {
		t.Fatal("export err:", err)
	}

	tarCmd.Stdin = reader
	if err := tarCmd.Start(); err != nil {
		t.Fatal(err)
	}

	if err := tarCmd.Wait(); err != nil {
		t.Fatal(err)
	}

	if err := reader.Close(); err != nil {
		t.Fatal(err)
	}
	t.Error("done", tempDir)
}*/

func TestRunRemove(t *testing.T) {
	cfg := &runtime.ContainerConfig{
		Cmd: []string{
			"/bin/sh",
			"-c",
			"echo hello && sleep 3600",
		},
		Binds: []*runtime.Bind{
			runtime.BindBoth("/tmp/foo"),
		},
		Devices: []*runtime.Bind{
			runtime.BindBoth("/dev/kvm"),
		},
	}
	t.Error(client.RunContainer(imageName, cfg, "ignite-test-foo2", "test-foo2"))
	t.Error(client.RemoveContainer("ignite-test-foo2"))
}

func TestV2ShimRuntimesHaveBinaryNames(t *testing.T) {
	for _, runtime := range v2ShimRuntimes {
		if v2shim.BinaryName(runtime) == "" {
			t.Errorf("shim binary could not be found -- %q is an invalid runtime/v2/shim", runtime)
		}
	}
}