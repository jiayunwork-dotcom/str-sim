# str-sim examples

Offline usage examples (no network required).

Build first:

```bash
export GOTOOLCHAIN=local CGO_ENABLED=0
go build -o /tmp/str-sim .
```

Compute a similarity score:

```bash
/tmp/str-sim levenshtein kitten sitting
/tmp/str-sim jaro-winkler prefix prefixation -p 4
```

Threshold match:

```bash
/tmp/str-sim match jaro-winkler prefix prefixation 0.9
```

Unknown algorithm name exits non-zero without panicking:

```bash
/tmp/str-sim frobnicate a b; echo $?
```
