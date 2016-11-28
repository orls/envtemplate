# envtemplate

A super-lightweight tool for templating config files from environment variables – and nothing else.

## Usage

`$ envtemplate < my-template-file > my-output-file`

Templating is done by the [go template package](https://golang.org/pkg/text/template/), where the only configured variables are the process's environment variables.

For example, a simple template might look like:

```
Hello {{ .USER }}! {{ if .HOME }}Your home dir is {{ .HOME }}.{{ else }}You don't appear to have a home dir set.{{ end }}
```
## Install

### from source

Assuming a working go installation, just  `git clone`, `cd` and `go build`. A binary named `envtemplate` should appear.

### prebuilt binary

A binary is attached to each github release. If you're happy to trust that, just fetch it with curl/similar (being sure to follow redirects):

`curl -L https://github.com/orls/envtemplate/releases/download/0.0.1/envtemplate > /usr/bin/envtemplate && chmod +x /usr/bin/envtemplate`

## Why?

This was borne out of frustration with using regular shell techniques – heredocs, `sed`, and similar – in various docker image-building and container-runtime configuration arrangements; for many config file formats (hi, nginx!) it starts to become unweildy to manage conditional blocks, escaping, etc.

It is a kind of hybrid of [gotpl](https://github.com/tsg/gotpl) and [envtpl](https://github.com/andreasjansson/envtpl). In the target enviroment of docker container management, it's useful to:

- have small, easily-installable binary tools
    - ...ruling out `envtpl`; the extra docker image bloat of a python+pip install is.... far from zero
- provide variables directly as env vars
    - ...ruling out `gotpl`, which takes a yml file, needing extra pre-processing if env vars are your only means of configuring.

It shares some spiritual affinity to [confd](https://github.com/kelseyhightower/confd), in the way that a butter knife shares some spiritual affinity to a swiss-army knife.

If you want config pulled from remote datastores, or from yml files, or already have python in the relevant envs and like jinja syntax, then those project may be better fits.
