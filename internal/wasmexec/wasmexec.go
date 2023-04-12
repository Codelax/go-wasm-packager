package wasmexec

import (
	_ "embed"
	"io"
	"strconv"
	"strings"
)

//go:embed assets/wasm_exec.js
var script string

//go:embed assets/wasm_exec_node.js
var nodescript string

var runner = `
const go = new Go();
go.argv = process.argv.slice(1);
go.env = Object.assign({ TMPDIR: require("os").tmpdir() }, process.env);
go.exit = process.exit;
// fs.readFileSync("main.wasm")
WebAssembly.instantiate(wasmBuffer, go.importObject).then((result) => {
    process.on("exit", (code) => { // Node.js exits if no event handler is pending
        if (code === 0 && !go.exited) {
            // deadlock, make Go print error and stack traces
            go._pendingEvent = { id: 0 };
            go._resume();
        }
    });
    return go.run(result.instance);
}).catch((err) => {
    console.error(err);
    process.exit(1);
});
`

func Script() string {
	return script
}

func NodeScript() string {
	return nodescript
}

func Runner() string {
	return runner
}

func Buffer(reader io.Reader) (string, error) {
	code := "const wasmBuffer = new Int8Array(["
	data, err := io.ReadAll(reader)
	if err != nil {
		return "", err
	}
	ints := make([]string, len(data))
	for i := range data {
		ints[i] = strconv.FormatInt(int64(data[i]), 10)
	}

	code += strings.Join(ints, ",")

	return code + "])", nil
}
