module github.com/venkat0724/numtowords

go 1.26.4

// Discovered unneeded library references in go.mod and go.sum files. The following retractions have been added to the go.mod file to prevent the use of these versions in future builds.
retract v1.1.1
