package main

import (
	"cd.splunkdev.com/data-availability/lsdc-connector/ms-graph-sharepoint-connector/src/ms"
	"net/url"
	"path"
)

func main() {
	source := "source"
	parentID := "parentID"
	driveItemID := "driveItemID"
	//permissionID:="permissionID"
	u, err := url.Parse(ms.MSGraphAPIRoot)
	if err != nil {
		println(err, "Can't Parse API root url")
	}
	//TestFileFolder := "FolderA"
	//TestFileName := "FileB"
	u.Path = path.Join(u.Path, ms.MSGraphAPIVersion, source, parentID, "drive", "items", driveItemID)

	a := u.String()
	println(a)

	b := assembleUrl(ms.MSGraphAPIVersion, source, parentID, "drive", "items", driveItemID)
	println(b)

	if a == b {
		println("same")
	}
}

func assembleUrl(parts ...string) string {
	u, err := url.Parse(ms.MSGraphAPIRoot)
	if err != nil {
		println(err, "Can't Parse API root url")
	}
	u.Path = path.Join(parts...)
	return u.String()
}
