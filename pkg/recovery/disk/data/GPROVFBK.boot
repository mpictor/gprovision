timeout 0
default 0

title Recovery
errorcheck off
find --set-root /{{.FallFile}}
if not %@root:~5,1%==0 && root (hd0,0)
errorcheck on
kernel /{{.Kernel}} quiet
