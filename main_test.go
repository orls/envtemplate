package main

import "fmt"
import "bytes"

func ExampleReadEnvVars() {
  samples := []string{"a=A", "b=b=B", "empty=", "spaces=   "}
  for k, v := range ReadEnvVars(samples) {
    fmt.Printf("key '%s' has value '%s'\n", k, v)
  }
  // Output: key 'a' has value 'A'
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
  WriteTemplateToStream("{{ if not .BEGIN_INVASION}}We mean you no{{ else }}Prepare yourselves for{{ end }} harm.", env, buf)
  fmt.Println(buf.String())
  // Output: Hello, earth!
  // We mean you no harm.
}
