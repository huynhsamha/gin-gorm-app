This is directory containing HTML files which are share components for each page.

Each file only is pure HTML with data fields `{{ .value }}` and other HTML files use `template` to include.

### Other files using this file

```
{{ include "component_name" }}
```