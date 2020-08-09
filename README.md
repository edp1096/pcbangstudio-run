# PC-BangStudio Launcher

For using [`pcbangstudio`](https://github.com/edp1096/pcbangstudio)

## Build
```powershell
# go get github.com/akavel/rsrc
bin/rsrc -arch="amd64" -manifest run.manifest -ico="icons/browser.ico" -o rsrc.syso

go build -ldflags="-H windowsgui"
```
