package gostars

import (
	pkggodevclient "github.com/guseggert/pkggodev-client"
	"github.com/pkg/errors"
)

// ============================================================================
//  Type: PkgInfo
// ============================================================================

// PkgInfo holds information about the package from pkg.go.dev.
// It is mainly used to obtain the number of packages using this package.
type PkgInfo struct {
	Name       string `json:"name"`        // Name of the package
	Repository string `json:"repository"`  // Repository URL of the package
	ImportedBy int    `json:"imported_by"` // Number of packages that imports this package
}

// ============================================================================
//  Constructor
// ============================================================================

// NewPkgInfo returns the initialized object of PkgInfo from pkagName.
func NewPkgInfo(pkgName string) (*PkgInfo, error) {
	pkgInfo := &PkgInfo{
		Name: pkgName,
	}

	if err := pkgInfo.Update(); err != nil {
		return nil, err
	}

	return pkgInfo, nil
}

// ============================================================================
//  Methods
// ============================================================================

// Update pulls the package information and sets to the according field.
func (p *PkgInfo) Update() (err error) {
	CoolDown()

	if err = p.UpdateImportedBy(); err == nil {
		err = p.UpdateURLRepository()
	}

	return err
}

func (p *PkgInfo) UpdateImportedBy() error {
	client := pkggodevclient.New()

	// Set ImportedBy
	importedBy, err := client.ImportedBy(pkggodevclient.ImportedByRequest{
		Package: p.Name,
	})
	if err != nil {
		return errors.Wrap(err, "failed to get 'imported by' information")
	}

	p.ImportedBy = len(importedBy.ImportedBy)

	return nil
}

func (p *PkgInfo) UpdateURLRepository() error {
	client := pkggodevclient.New()

	pkgInfo, err := client.DescribePackage(pkggodevclient.DescribePackageRequest{
		Package: p.Name,
	})
	if err != nil {
		return errors.Wrap(err, "failed to get package information")
	}

	p.Repository = "https://" + pkgInfo.Repository

	return nil
}
