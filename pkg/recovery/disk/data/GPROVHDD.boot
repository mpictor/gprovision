timeout 0
default 0
fallback 1

title Normal
find --set-root /{{.Kernel}}
if not %@root:~5,1%==0 && root (hd0,0)
kernel /{{.Kernel}} real_root=UUID={{.Root_uuid}} quiet {{.Extra_opts}}

title Recovery
find --set-root /{{.Kernel}}
if not %@root:~5,1%==0 && root (hd0,0)
kernel /{{.Kernel}} quiet
