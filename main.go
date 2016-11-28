package main

import "bytes"
import "io"
import "os"
import "text/template"
import "strings"

func ReadTemplate(instream io.Reader) string {
    buf := new(bytes.Buffer)
    buf.ReadFrom(instream)
    return buf.String()
}

func ReadEnvVars(rawEnv []string) (environ map[string]string) {
    environ = make(map[string]string)
    for _, item := range rawEnv {
        parts := strings.SplitN(item, "=", 2)
        environ[parts[0]] = parts[1]
    }
    return
}

func WriteTemplateToStream(tplSource string, environ map[string]string, outStream io.Writer) {
    tpl := template.New("_root_")
    _, err := tpl.Parse(tplSource)
    if err != nil {
        panic(err)
    }
    tpl.Execute(outStream, environ)
}

func main() {
    WriteTemplateToStream(ReadTemplate(os.Stdin), ReadEnvVars(os.Environ()), os.Stdout)
}
