package main

import "bytes"

import "errors"
import "io"
import "os"
import "log"
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
    tpl.Funcs(template.FuncMap{
        "split": TplSplitStr,
    })
    _, err := tpl.Parse(tplSource)
    if err != nil {
        log.Fatal(err)
    }
    err = tpl.Execute(outStream, environ)
    if err != nil {
        log.Fatal(err)
    }
}

func TplSplitStr(args ...interface{}) ([]string, error) {
    rawValue := args[0].(string)
    sep := args[1].(string)
    limit := -1
    if len(args) > 2 {
        parsedLimit, ok := args[2].(int)
        if !ok {
            err := errors.New("Limit parameter (3rd)is not integer")
            return nil, err
        }
        limit = parsedLimit
    }
    return strings.SplitN(rawValue, sep, limit), nil
}

func main() {
    WriteTemplateToStream(ReadTemplate(os.Stdin), ReadEnvVars(os.Environ()), os.Stdout)
}
