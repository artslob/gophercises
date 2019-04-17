package sitemap

import (
	"bytes"
	"encoding/xml"
)

type url struct {
	XMLName xml.Name `xml:"url"`
	Loc     string   `xml:"loc"`
}

type urlSet struct {
	XMLName   xml.Name `xml:"urlset"`
	Namespace string   `xml:"xmlns,attr"`
	Set       []url
}

func GetXml(siteMap map[string]bool) ([]byte, error) {
	w := &bytes.Buffer{}
	w.WriteString(xml.Header)
	root := urlSet{Namespace: "http://www.sitemaps.org/schemas/sitemap/0.9"}
	for key := range siteMap {
		root.Set = append(root.Set, url{Loc: key})
	}
	marshaled, err := xml.MarshalIndent(root, "", "    ")
	if err != nil {
		return nil, err
	}
	w.Write(marshaled)
	return w.Bytes(), nil
}
