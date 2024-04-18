package v17localresidentbin

import _ "embed"

//go:embed "stable/lrsrv-windows-x64.exe.zip"
var LRSRV_STABLE_ZIPPED EmbedZipFile
