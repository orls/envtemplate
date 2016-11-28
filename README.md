# envtemplate

A super-lightweight tool for templating config files from environment variables – and nothing else.

## Usage

$ envtemplate < my-template-file > my-output-file

Templating is done by the [go template package](https://golang.org/pkg/text/template/), where the only configured inputs are the process's environment variables.

It's kind of a hybrid of [gotpl](https://github.com/tsg/gotpl) and [envtpl](https://github.com/andreasjansson/envtpl). It was built for use in docker containers, where it's useful to:

- have small, easily-installable binary tools
    - ...ruling out `envtpl`; the extra docker image bloat of a python+pip install is.... far from zero
- provide variables directly as env vars
    - ...ruling out `gotpl`, which takes a yml file, needing extra pre-processing if env vars are your only means of configuring.

It shares some spiritual affinity to [confd](https://github.com/kelseyhightower/confd), in the way that a butter knife shares some spiritual affinity to a swiss-army knife.

If you want config pulled from remote datastores, or from yml files, or already have python in the relevant envs and like jinja syntax, then those project may be better fits.
