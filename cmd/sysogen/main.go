package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/josephspurrier/goversioninfo"
)

func main() {
	outputFile := flag.String("output", "resource.syso", "Output file to write the .syso to.")
	flag.Parse()

	version := goversioninfo.FileVersion{
		Major: 1,
		Minor: 1,
		Patch: 1,
		Build: 1,
	}

	info := goversioninfo.VersionInfo{
		FixedFileInfo: goversioninfo.FixedFileInfo{
			FileVersion:    version,
			ProductVersion: version,
			FileFlagsMask:  "3f",
			FileFlags:      "00",
			FileOS:         "040004",
			FileType:       "01",
			FileSubType:    "00",
		},
		StringFileInfo: goversioninfo.StringFileInfo{
			CompanyName:      "Example Inc.",
			FileDescription:  "An Example File Description",
			LegalCopyright:   fmt.Sprintf("Copyright (C) %d", time.Now().Year()),
			OriginalFilename: "Example.exe",
			ProductName:      "Example",
			ProductVersion:   formatVersionString(version),
		},
		VarFileInfo: goversioninfo.VarFileInfo{
			Translation: goversioninfo.Translation{
				LangID:    goversioninfo.LngUSEnglish,
				CharsetID: goversioninfo.CsUnicode,
			},
		},
	}

	info.Build()
	info.Walk()

	arch := "386"
	if err := info.WriteSyso(*outputFile, arch); err != nil {
		log.Fatalf("Failed to write the resource.syso file out to %q (arch %q): %v",
			*outputFile, arch, err)
	}
}

func formatVersionString(v goversioninfo.FileVersion) string {
	return fmt.Sprintf("v%d.%d.%d.%d", v.Major, v.Minor, v.Patch, v.Build)
}
