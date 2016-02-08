package osm

import (
	"encoding/xml"
	"testing"
	"time"

	"github.com/paulmach/go.geo"
)

func TestChangeset(t *testing.T) {
	data := []byte(`
<?xml version="1.0" encoding="UTF-8"?>
<osm version="0.6" generator="replicate_changesets.rb" copyright="OpenStreetMap and contributors" attribution="http://www.openstreetmap.org/copyright" license="http://opendatacommons.org/licenses/odbl/1-0/">
  <changeset id="36947117" created_at="2016-02-01T21:57:17Z" closed_at="2016-02-01T23:05:55Z" open="true" num_changes="86" user="padvinder" uid="978786" min_lat="52.7016394" max_lat="52.7236643" min_lon="5.1545597" max_lon="5.2532961" comments_count="5">
    <tag k="build" v="2.4-16-g0c126d0"/>
    <tag k="created_by" v="Potlatch 2"/>
    <tag k="version" v="2.4"/>
  </changeset>
  <changeset id="36947173" created_at="2016-02-01T22:00:56Z" closed_at="2016-02-01T23:05:06Z" open="false" num_changes="9" user="florijn11" uid="1319603" min_lat="51.5871887" max_lat="51.6032569" min_lon="5.3214071" max_lon="5.33106" comments_count="0">
    <tag k="version" v="2.4"/>
    <tag k="build" v="2.4-16-g0c126d0"/>
    <tag k="comment" v="Fietsdoorsteek aangepast"/>
    <tag k="created_by" v="Potlatch 2"/>
  </changeset>
</osm>`)

	cs := Changesets{}
	err := xml.Unmarshal(data, &cs)
	if err != nil {
		t.Fatalf("unmarshal error: %v", err)
	}

	if l := len(cs.Changesets); l != 2 {
		t.Fatalf("incorrect number of changesets, got %v", l)
	}

	c := cs.Changesets[0]
	if v := c.ID; v != 36947117 {
		t.Errorf("incorrect id, got %v", v)
	}

	if v := c.CreatedAt; v != time.Date(2016, time.February, 1, 21, 57, 17, 0, time.UTC) {
		t.Errorf("incorrect created at, got %v", v)
	}

	if v := c.ClosedAt; v != time.Date(2016, time.February, 1, 23, 05, 55, 0, time.UTC) {
		t.Errorf("incorrect closed at, got %v", v)
	}

	if v := c.ChangesCount; v != 86 {
		t.Errorf("incorrect changes count, got %v", v)
	}

	if v := c.User; v != "padvinder" {
		t.Errorf("incorrect user, got %v", v)
	}

	if v := c.UserID; v != 978786 {
		t.Errorf("incorrect user id, got %v", v)
	}

	if v := c.MinLat; v != 52.7016394 {
		t.Errorf("incorrect min lat, got %v", v)
	}

	if v := c.MaxLat; v != 52.7236643 {
		t.Errorf("incorrect max lat, got %v", v)
	}

	if v := c.MinLng; v != 5.1545597 {
		t.Errorf("incorrect min lng, got %v", v)
	}

	if v := c.MaxLng; v != 5.2532961 {
		t.Errorf("incorrect max lng, got %v", v)
	}

	if v := c.CommentsCount; v != 5 {
		t.Errorf("incorrect comment count, got %v", v)
	}
}

func TestChangesetTags(t *testing.T) {
	data := []byte(`
<changeset id="123123">
  <tag k="comment" v="changeset comment"/>
  <tag k="created_by" v="iD 1.8.3"/>
  <tag k="locale" v="en-US"/>
  <tag k="host" v="http://id.org"/>
  <tag k="imagery_used" v="Bing"/>
  <tag k="source" v="some data"/>
  <tag k="bot" v="yes"/>
</changeset>`)

	var c Changeset
	err := xml.Unmarshal(data, &c)
	if err != nil {
		t.Fatalf("unmarshal error: %v", err)
	}

	if v := c.Comment(); v != "changeset comment" {
		t.Errorf("incorrect comment, got %v", v)
	}

	if v := c.CreatedBy(); v != "iD 1.8.3" {
		t.Errorf("incorrect created by, got %v", v)
	}

	if v := c.Locale(); v != "en-US" {
		t.Errorf("incorrect locale, got %v", v)
	}

	if v := c.Host(); v != "http://id.org" {
		t.Errorf("incorrect host, got %v", v)
	}

	if v := c.ImageryUsed(); v != "Bing" {
		t.Errorf("incorrect imagery used, got %v", v)
	}

	if v := c.Source(); v != "some data" {
		t.Errorf("incorrect source, got %v", v)
	}

	if v := c.Bot(); v != true {
		t.Errorf("incorrect bot, got %v", v)
	}
}

func TestChangesetBound(t *testing.T) {
	data := []byte(`
<changeset id="36947173" created_at="2016-02-01T22:00:56Z" closed_at="2016-02-01T23:05:06Z" open="false" num_changes="9" user="florijn11" uid="1319603" min_lat="51.5871887" max_lat="51.6032569" min_lon="5.3214071" max_lon="5.33106" comments_count="0">
    <tag k="version" v="2.4"/>
    <tag k="build" v="2.4-16-g0c126d0"/>
    <tag k="comment" v="Fietsdoorsteek aangepast"/>
    <tag k="created_by" v="Potlatch 2"/>
</changeset>`)

	var c Changeset
	err := xml.Unmarshal(data, &c)
	if err != nil {
		t.Fatalf("unmarshal error: %v", err)
	}

	b := c.Bound()
	expected := geo.NewBoundFromPoints(
		geo.NewPoint(c.MinLng, c.MinLat),
		geo.NewPoint(c.MaxLng, c.MaxLat),
	)
	if !b.Equals(expected) {
		t.Errorf("incorrect bound, got %v", b)
	}
}