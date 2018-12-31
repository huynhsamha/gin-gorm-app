This is directory containing HTML files which are partials for each page.

Each file `define` new components and other HTML files use `template` to include.

### File in this directory

```
{{ define "component_name" }}

...


{{ end }}
```

### Other files using this file

```
{{ template "component_name" . }}

```