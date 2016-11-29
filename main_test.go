package main

import "fmt"
import "bytes"

func ExampleReadEnvVars() {
  samples := []string{"a=A", "b=b=B", "empty=", "spaces=   "}
  envMap := ReadEnvVars(samples)
  iterOrder := []string{"a", "b", "empty", "spaces"}

  for _, k := range iterOrder {
    v := envMap[k]
    fmt.Printf("key '%s' has value '%s'\n", k, v)
  }
  // Output:
  // key 'a' has value 'A'
  // key 'b' has value 'b=B'
  // key 'empty' has value ''
  // key 'spaces' has value '   '
}

func ExampleWriteTemplateToStream_basic() {
  buf := new(bytes.Buffer)
  env := map[string]string{
    "TARGET_PLANET": "earth",
  }
  WriteTemplateToStream("Hello, {{ .TARGET_PLANET }}!\n", env, buf)
  WriteTemplateToStream("{{ if not (exists . \"BEGIN_INVASION\")}}We mean you no{{ else }}Prepare yourselves for{{ end }} harm.", env, buf)
  fmt.Println(buf.String())
  // Output:
  // Hello, earth!
  // We mean you no harm.
}

func ExampleWriteTemplateToStream_strSplit() {
  buf := new(bytes.Buffer)
  env := map[string]string{
    "PARTS": "a,b,c",
  }
  // WriteTemplateToStream("{{ split .PARTS ',' }}", env, buf)
  WriteTemplateToStream("{{ range $i, $v := split .PARTS \",\"   }}List 1 item {{$i}} has value {{$v}}\n{{end}}", env, buf)
  WriteTemplateToStream("{{ range $i, $v := split .PARTS \",\" 2 }}List 2 item {{$i}} has value {{$v}}\n{{end}}", env, buf)
  fmt.Println(buf.String())
  // Output:
  // List 1 item 0 has value a
  // List 1 item 1 has value b
  // List 1 item 2 has value c
  // List 2 item 0 has value a
  // List 2 item 1 has value b,c
}
