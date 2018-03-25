package nginxconf

import (
	"fmt"
	"io"
)

func PrintConfiguration(out io.Writer, config *SiteConfiguration) error {
	switch config.Template {
	case StaticSite, ProxySite, RedirectSite:
		// pass
	default:
		return fmt.Errorf("invalid config type: %s", config.Template)
	}

	if err := siteConfigTmpl.Execute(out, config); err != nil {
		fmt.Fprintf(out, "\n\nerror: %+v\n\n", err)
		return err
	}
	fmt.Fprint(out, "\n")

	return nil
}
