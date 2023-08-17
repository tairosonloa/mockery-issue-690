# Mockery issue

This repo is a minimal reproduction of the code that is facing an error / bug with mockery.

It's related with https://github.com/vektra/mockery/issues/690

## Steps to reproduce

1. Install mockery with `brew`
2. Run `mockery` in the project root
3. See the error

```
$> mockery

17 Aug 23 12:16 CEST INF Starting mockery dry-run=false version=v2.32.4
17 Aug 23 12:16 CEST INF Using config: /Users/aalonso/workspaces/issue/.mockery.yaml dry-run=false version=v2.32.4
17 Aug 23 12:16 CEST ERR encountered error when loading package error="-: package go-pvpc is not in GOROOT (/opt/homebrew/Cellar/go/1.20.7/libexec/src/go-pvpc)" dry-run=false version=v2.32.4
17 Aug 23 12:16 CEST ERR unable to parse packages error="error occurred when loading packages" dry-run=false version=v2.32.4
error occurred when loading packages
```
