# bscdiff

bscdiff compares bsc, issue, fate (it's a SUSE thing) and CVE numbers from a source changelog, to a target changelog. Missing numbers are then printed with their occurrence in the source changelog.


## Usage

```
brejoc@alpha ~> ./bscdiff source.changes target.changes
508: bsc#1098394 -> - Fix file.get_diff regression on 2018.3 (bsc#1098394)
525: bsc#1098394 -> - Fix file.managed binary file utf8 error (bsc#1098394)
4092: bsc#565656565 -> - uploaded to salt 1.12.0 (bsc#565656565, bsc#676767676)
4092: bsc#676767676 -> - uploaded to salt 1.12.0 (bsc#565656565, bsc#676767676)
```

Output is structure like this:

\<line in source.changes\>: \<bsc missing in target.changes\> -> \<line from source.changes\>

# Patterns

bscdiff looks for the following patterns:

* bsc#12345
* CVE-2019-12356
* U#1234
* fate#12345

## Building bscdiff

Since no external dependency was used, you can just do a `go build bscdiff.go`.

# Issues and Contributions

If you've got issues or questions please don't hesitate to open an issue. If you'd like to improve something or found a bug and you want to fix it, just open a pull request.

## License

MIT