package v17localresidentbin

import _ "embed"

//go:embed "stable/lrsrv-debian-x64.zip"
var LRSRV_STABLE_ZIPPED EmbedZipFile
