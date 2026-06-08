package mvtssr

import (
	"image/color"
	"testing"
)

func TestTileID(t *testing.T) {
	id := NewTileID(5, 10, 20)
	if id == nil {
		t.Fatal("NewTileID returned nil")
	}

	if !id.Eq(NewTileID(5, 10, 20)) {
		t.Error("equal tile IDs should be equal")
	}
	if id.Eq(NewTileID(5, 10, 21)) {
		t.Error("different tile IDs should not be equal")
	}

	scaled := id.ScaledTo(10)
	if scaled == nil {
		t.Fatal("ScaledTo returned nil")
	}

	if !scaled.IsChildOf(id) {
		t.Error("zoom-10 scaled tile should be child of zoom-5 tile")
	}

	if !id.Less(NewTileID(5, 10, 21)) {
		t.Error("id (5,10,20) should be less than (5,10,21) — same z, same x, smaller y")
	}

	children := id.Children()
	if len(children) != 4 {
		t.Fatalf("expected 4 children, got %d", len(children))
	}
	if !children[0].IsChildOf(id) || !children[3].IsChildOf(id) {
		t.Error("all children should be children of parent")
	}
}

func TestLatLng(t *testing.T) {
	ll := NewLatLng(39.9, 116.4)
	if ll == nil {
		t.Fatal("NewLatLng returned nil")
	}

	lat, lon := ll.GetData()
	if lat != 39.9 || lon != 116.4 {
		t.Errorf("got (%.1f, %.1f), want (39.9, 116.4)", lat, lon)
	}
}

func TestLatLngWithTileID(t *testing.T) {
	id := NewTileID(0, 0, 0)
	ll := NewLatLngWithID(id)
	if ll == nil {
		t.Fatal("NewLatLngWithID returned nil")
	}
	lat, lon := ll.GetData()
	if lat == 0 && lon == 0 {
		t.Log("tile 0/0/0 center is at origin")
	}
}

func TestLatLngBounds(t *testing.T) {
	world := LatLngBoundsWorld()
	if world == nil {
		t.Fatal("LatLngBoundsWorld returned nil")
	}
	if !world.Valid() {
		t.Error("world bounds should be valid")
	}
	if world.South() >= world.North() {
		t.Error("south should be less than north")
	}
	if world.West() >= world.East() {
		t.Error("west should be less than east")
	}

	empty := LatLngBoundsEmpty()
	if empty == nil {
		t.Fatal("LatLngBoundsEmpty returned nil")
	}
	if !empty.Empty() {
		t.Error("empty bounds should be empty")
	}

	sw := NewLatLng(-10, -20)
	ne := NewLatLng(10, 20)
	hull := HullLatLngBounds(sw, ne)
	if hull == nil {
		t.Fatal("HullLatLngBounds returned nil")
	}
	if !hull.Valid() {
		t.Error("hull should be valid")
	}
	if hull.South() != -10 || hull.North() != 10 {
		t.Error("hull latitude bounds mismatch")
	}
	if hull.West() != -20 || hull.East() != 20 {
		t.Error("hull longitude bounds mismatch")
	}

	center := hull.Center()
	if center == nil {
		t.Fatal("Center returned nil")
	}
	cLat, cLon := center.GetData()
	if cLat != 0 || cLon != 0 {
		t.Errorf("center should be at (0,0), got (%.1f, %.1f)", cLat, cLon)
	}

	p := NewLatLng(0, 0)
	if !hull.ContainsPoint(p) {
		t.Error("hull should contain origin")
	}

	hull.Extend(NewLatLng(20, 30))
	if hull.North() != 20 || hull.East() != 30 {
		t.Error("extend should expand bounds")
	}

	southWest := hull.SouthWest()
	if southWest == nil {
		t.Fatal("SouthWest returned nil")
	}
	if lat, lon := southWest.GetData(); lat != -10 || lon != -20 {
		t.Errorf("SW should be (-10,-20), got (%.1f,%.1f)", lat, lon)
	}

	northEast := hull.NorthEast()
	if northEast == nil {
		t.Fatal("NorthEast returned nil")
	}
	if lat, lon := northEast.GetData(); lat != 20 || lon != 30 {
		t.Errorf("NE should be (20,30), got (%.1f,%.1f)", lat, lon)
	}

	boundsWithID := NewLatLngBoundsWithID(NewTileID(0, 0, 0))
	if boundsWithID == nil {
		t.Fatal("NewLatLngBoundsWithID returned nil")
	}
	if !boundsWithID.ContainsTile(NewTileID(0, 0, 0)) {
		t.Error("bounds should contain its own tile")
	}

	singleton := NewLatLngBounds(NewLatLng(5, 10))
	if singleton == nil {
		t.Fatal("NewLatLngBounds returned nil")
	}
	if !singleton.Valid() {
		t.Error("singleton bounds should be valid")
	}

	p2 := NewLatLng(15, 25)
	constrained := hull.Constrain(p2)
	if constrained == nil {
		t.Fatal("Constrain returned nil")
	}

	if hull.ContainsBounds(hull) != true {
		t.Error("bounds should contain itself")
	}

	if hull.Intersects(hull) != true {
		t.Error("bounds should intersect itself")
	}
}

func TestEdgeInsets(t *testing.T) {
	e := NewEdgeInsets(1, 2, 3, 4)
	if e == nil {
		t.Fatal("NewEdgeInsets returned nil")
	}
	if e.Top() != 1 || e.Left() != 2 || e.Bottom() != 3 || e.Right() != 4 {
		t.Error("edge inset values mismatch")
	}
	if e.IsFlush() {
		t.Error("(1,2,3,4) should not be flush")
	}

	flush := NewEdgeInsets(0, 0, 0, 0)
	if !flush.IsFlush() {
		t.Error("(0,0,0,0) should be flush")
	}

	if !e.Eq(NewEdgeInsets(1, 2, 3, 4)) {
		t.Error("equal insets should be equal")
	}

	center := e.Center(100, 200)
	if center == nil {
		t.Fatal("Center returned nil")
	}
	x, y := center.Get()
	if x <= 0 || y <= 0 {
		t.Errorf("center should be positive, got (%.1f, %.1f)", x, y)
	}
}

func TestScreenCoordinate(t *testing.T) {
	sc := NewScreenCoordinate(123.4, 567.8)
	if sc == nil {
		t.Fatal("NewScreenCoordinate returned nil")
	}
	x, y := sc.Get()
	if x != 123.4 || y != 567.8 {
		t.Errorf("got (%.1f, %.1f), want (123.4, 567.8)", x, y)
	}
}

func TestSize(t *testing.T) {
	s := NewSize(200, 100)
	if s == nil {
		t.Fatal("NewSize returned nil")
	}
	if s.Area() != 20000 {
		t.Errorf("area = %d, want 20000", s.Area())
	}
	if s.AspectRatio() != 2.0 {
		t.Errorf("aspect ratio = %.2f, want 2.0", s.AspectRatio())
	}
	if s.Empty() {
		t.Error("(200,100) should not be empty")
	}

	empty := NewSize(0, 0)
	if !empty.Empty() {
		t.Error("(0,0) should be empty")
	}
}

func TestImage(t *testing.T) {
	empty := NewEmptyImage()
	if empty == nil {
		t.Fatal("NewEmptyImage returned nil")
	}
	if empty.Valid() {
		t.Error("empty image should not be valid")
	}

	si := NewSize(10, 10)
	img := NewImage(si)
	if img == nil {
		t.Fatal("NewImage returned nil")
	}
	if !img.Valid() {
		t.Error("new image should be valid")
	}
	w, h := img.Size()
	if w != 10 || h != 10 {
		t.Errorf("size = (%d,%d), want (10,10)", w, h)
	}
	stride := img.Stride()
	if stride <= 0 {
		t.Errorf("stride = %d, want > 0", stride)
	}
	bytes := img.Bytes()
	if bytes <= 0 {
		t.Errorf("bytes = %d, want > 0", bytes)
	}

	data := img.Data()
	if len(data) <= 0 {
		t.Error("image data should not be empty")
	}

	goImg := img.Image()
	if goImg == nil {
		t.Fatal("Image() returned nil")
	}
	bounds := goImg.Bounds()
	if bounds.Dx() != 10 || bounds.Dy() != 10 {
		t.Errorf("go image bounds = %dx%d, want 10x10", bounds.Dx(), bounds.Dy())
	}

	data2 := make([]byte, 10*10*4)
	for i := range data2 {
		data2[i] = 0x80
	}
	img2 := NewImageWithData(NewSize(10, 10), data2)
	if img2 == nil {
		t.Fatal("NewImageWithData returned nil")
	}
	if !img2.Valid() {
		t.Error("image with data should be valid")
	}
	pixel := img2.Image().At(0, 0)
	r, g, b, a := pixel.RGBA()
	if r>>8 != 0x80 || g>>8 != 0x80 || b>>8 != 0x80 || a>>8 != 0x80 {
		t.Errorf("pixel RGBA = (%d,%d,%d,%d), want (128,128,128,128)",
			r>>8, g>>8, b>>8, a>>8)
	}
}

func TestResourceOptions(t *testing.T) {
	opts := NewResourceOptions()
	if opts == nil {
		t.Fatal("NewResourceOptions returned nil")
	}
	opts.SetAccessToken("test-token")
	opts.SetBaseURL("https://example.com")
	opts.SetAssetPath("/assets")
	opts.SetMaximumCacheSize(1024 * 1024)
}

func TestMapOptions(t *testing.T) {
	opts := NewMapOptions()
	if opts == nil {
		t.Fatal("NewMapOptions returned nil")
	}
	opts.SetMapMode(Static)
	opts.SetConstrainMode(HeightOnly)
	opts.SetViewportMode(FlippedY)
	opts.SetCrossSourceCollisions(false)
	opts.SetNorthOrientation(Upwards)
	opts.SetSize(NewSize(512, 256))
	opts.SetPixelRatio(2.0)
}

func TestCameraOptions(t *testing.T) {
	opts := NewCameraOptions()
	if opts == nil {
		t.Fatal("NewCameraOptions returned nil")
	}
	opts.SetCenter(NewLatLng(10, 20))
	opts.SetZoom(5.5)
	opts.SetBearing(45)
	opts.SetPitch(30)
	opts.SetPadding(NewEdgeInsets(1, 1, 1, 1))
	opts.SetAnchor(NewScreenCoordinate(100, 200))
}

type testFileLoader struct{}

func (f *testFileLoader) LoadAsync(req *FileSourceRequest, res *Resource) {
	resp := NewErrorFileSourceResponse(NotFound, "not found")
	req.SetResponse(resp)
}
func (f *testFileLoader) Pause()  {}
func (f *testFileLoader) Resume() {}

func TestFileSourceResponse(t *testing.T) {
	resp := NewFileSourceResponse([]byte("hello"))
	if resp == nil {
		t.Fatal("NewFileSourceResponse returned nil")
	}

	errResp := NewErrorFileSourceResponse(NotFound, "not found")
	if errResp == nil {
		t.Fatal("NewErrorFileSourceResponse returned nil")
	}

	emptyResp := NewFileSourceResponse([]byte{})
	if emptyResp == nil {
		t.Fatal("NewFileSourceResponse([]byte{}) returned nil")
	}
}

func TestResource(t *testing.T) {
	r := NewStyleResource("https://example.com/style.json")
	if r == nil || r.GetKind() != Resource_Style {
		t.Error("expected Style resource")
	}

	r = NewSourceResource("https://example.com/source.json")
	if r == nil || r.GetKind() != Resource_Source {
		t.Error("expected Source resource")
	}

	r = NewTileResource("https://example.com/{z}/{x}/{y}.pbf", 1.0, 0, 0, 0, false, LM_All)
	if r == nil || r.GetKind() != Resource_Tile {
		t.Error("expected Tile resource")
	}

	r = NewGlyphsResource("https://example.com/{fontstack}/{start}-{end}.pbf", "Arial", 0, 255)
	if r == nil || r.GetKind() != Resource_Glyphs {
		t.Error("expected Glyphs resource")
	}

	r = NewSpriteImageResource("https://example.com/sprite", 1.0)
	if r == nil || r.GetKind() != Resource_SpriteImage {
		t.Error("expected SpriteImage resource")
	}

	r = NewSpriteJSONResource("https://example.com/sprite", 1.0)
	if r == nil || r.GetKind() != Resource_SpriteJSON {
		t.Error("expected SpriteJSON resource")
	}

	r = NewImageResource("https://example.com/image.png")
	if r == nil || r.GetKind() != Resource_Image {
		t.Error("expected Image resource")
	}
}

func TestResourceUsage(t *testing.T) {
	r := NewStyleResource("https://example.com/style.json")
	r.SetUsage(Offline)
	if r.GetUsage() != Offline {
		t.Error("expected Offline usage")
	}
	r.SetUsage(Online)
	if r.GetUsage() != Online {
		t.Error("expected Online usage")
	}
}

func TestBoundOptions(t *testing.T) {
	opts := NewBoundOptions()
	if opts == nil {
		t.Fatal("NewBoundOptions returned nil")
	}
	opts.SetCenter(LatLngBoundsWorld())
	opts.SetMinZoom(0)
	opts.SetMaxZoom(20)
	opts.SetMinPitch(0)
	opts.SetMaxPitch(60)
}

func TestRunLoop(t *testing.T) {
	loop := NewRunLoop()
	if loop == nil {
		t.Fatal("NewRunLoop returned nil")
	}
	loop.Stop()
}

func TestMapSnapshotterSetGet(t *testing.T) {
	// NOTE: MapSnapshotter creation requires a working OpenGL context.
	// On headless/CI systems this may crash with SIGSEGV inside mbgl.
	// Only test property getters/setters on an already-created snapshotter.
	// Full integration tests (Snapshot()) need a display server.
	t.Skip("requires OpenGL context — run on a Mac with a display")
}

func TestColorConversion(t *testing.T) {
	si := NewSize(2, 2)
	data := []byte{
		255, 0, 0, 255,
		0, 255, 0, 255,
		0, 0, 255, 255,
		128, 128, 128, 128,
	}
	img := NewImageWithData(si, data)
	if img == nil {
		t.Fatal("NewImageWithData returned nil")
	}
	goImg := img.Image()
	if goImg == nil {
		t.Fatal("Image() returned nil")
	}
	pixel := color.RGBAModel.Convert(goImg.At(0, 0)).(color.RGBA)
	expected := color.RGBA{R: 255, G: 0, B: 0, A: 255}
	if pixel != expected {
		t.Errorf("pixel (0,0) = %+v, want %+v", pixel, expected)
	}
}

func TestConstants(t *testing.T) {
	tests := []struct {
		name string
		got  interface{}
		want interface{}
	}{
		{"Continuous", MapModeType(Continuous), MapModeType(0)},
		{"Static", MapModeType(Static), MapModeType(1)},
		{"Tile", MapModeType(Tile), MapModeType(2)},
		{"None", ConstrainMode(None), ConstrainMode(0)},
		{"HeightOnly", ConstrainMode(HeightOnly), ConstrainMode(1)},
		{"WidthAndHeight", ConstrainMode(WidthAndHeight), ConstrainMode(2)},
		{"Default_Viewport", ViewportMode(Default), ViewportMode(0)},
		{"FlippedY", ViewportMode(FlippedY), ViewportMode(1)},
		{"Upwards", NorthOrientation(Upwards), NorthOrientation(0)},
		{"Resource_Unknown", Resource_Unknown, ResourceKind(0)},
		{"Resource_Style", Resource_Style, ResourceKind(1)},
		{"Resource_Tile", Resource_Tile, ResourceKind(3)},
		{"Asset", FileSourceType(Asset), FileSourceType(0)},
		{"Network", FileSourceType(Network), FileSourceType(3)},
		{"Immediate", CameraChangeMode(Immediate), CameraChangeMode(0)},
		{"Partial", RenderMode(Partial), RenderMode(0)},
		{"Full", RenderMode(Full), RenderMode(1)},
		{"StyleParseError", MapLoadError(StyleParseError), MapLoadError(0)},
		{"StyleLoadError", MapLoadError(StyleLoadError), MapLoadError(1)},
		{"Success", ReasonError(Success), ReasonError(1)},
		{"NotFound_Reason", ReasonError(NotFound), ReasonError(2)},
		{"Server", ReasonError(Server), ReasonError(3)},
		{"NoDebug", MapDebugOptions(NoDebug), MapDebugOptions(0)},
		{"TileBorders", MapDebugOptions(TileBorders), MapDebugOptions(2)},
		{"Timestamps", MapDebugOptions(Timestamps), MapDebugOptions(8)},
		{"Online", ResourceUsage(Online), ResourceUsage(false)},
		{"Offline", ResourceUsage(Offline), ResourceUsage(true)},
		{"LM_None", LoadingMethod(LM_None), LoadingMethod(0)},
		{"LM_Cache", LoadingMethod(LM_Cache), LoadingMethod(1)},
		{"LM_Network", LoadingMethod(LM_Network), LoadingMethod(2)},
		{"LM_All", LoadingMethod(LM_All), LoadingMethod(3)},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.got != tc.want {
				t.Errorf("got %v, want %v", tc.got, tc.want)
			}
		})
	}
}

func TestStyleResource(t *testing.T) {
	// This just tests that style construction functions don't panic
	si := NewSize(256, 256)
	_ = si

	ropts := NewResourceOptions()
	_ = ropts
}

func TestTileData(t *testing.T) {
	pngHeader := []byte{0x89, 'P', 'N', 'G', 0x0D, 0x0A, 0x1A, 0x0A}
	d := &TileData{Image: pngHeader, X: 1, Y: 2, Z: 3}
	if d.IsEmpty() {
		t.Error("PNG tile should not be empty")
	}
	if d.Size() != 8 {
		t.Errorf("size = %d, want 8", d.Size())
	}
	if d.Format() != "png" {
		t.Errorf("format = %q, want png", d.Format())
	}
	if d.X != 1 || d.Y != 2 || d.Z != 3 {
		t.Errorf("coords = (%d,%d,%d), want (1,2,3)", d.X, d.Y, d.Z)
	}

	jpegHeader := []byte{0xFF, 0xD8, 0xFF, 0xE0}
	d2 := &TileData{Image: jpegHeader}
	if d2.Format() != "jpeg" {
		t.Errorf("format = %q, want jpeg", d2.Format())
	}

	unknown := &TileData{Image: []byte{0, 1, 2, 3}}
	if unknown.Format() != "" {
		t.Errorf("format = %q, want empty", unknown.Format())
	}

	empty := &TileData{Image: []byte{}}
	if !empty.IsEmpty() {
		t.Error("empty data should be empty")
	}
	if empty.Size() != 0 {
		t.Errorf("size = %d, want 0", empty.Size())
	}
	if empty.Format() != "" {
		t.Errorf("format = %q, want empty", empty.Format())
	}
}

func TestRasterTile(t *testing.T) {
	td := &TileData{Image: []byte("test"), X: 5, Y: 10, Z: 15}
	rt := &RasterTile{data: []*TileData{td}}

	if rt.Count() != 1 {
		t.Errorf("Count = %d, want 1", rt.Count())
	}
	all := rt.GetAllData()
	if len(all) != 1 || all[0] != td {
		t.Error("GetAllData mismatch")
	}
	got, err := rt.GetData(0)
	if err != nil {
		t.Fatalf("GetData(0) error: %v", err)
	}
	if got != td {
		t.Error("GetData(0) mismatch")
	}
	_, err = rt.GetData(1)
	if err == nil {
		t.Error("GetData(1) should error")
	}
}

func TestTilerGridGlobalGeodetic(t *testing.T) {
	grid := NewGlobalGeodeticGrid(256, 1)
	if grid.TileSize() != 256 {
		t.Errorf("TileSize = %d, want 256", grid.TileSize())
	}

	// Zoom 0, tile (0,0) is the whole world
	mlon, mlat, xlon, xlat := grid.TileBounds(0, 0, 0)
	if mlon != -180 || xlon != 180 || mlat != -90 || xlat != 90 {
		t.Errorf("z0 bounds = (%.1f,%.1f,%.1f,%.1f), want (-180,-90,180,90)",
			mlon, mlat, xlon, xlat)
	}

	// Zoom 1, tile (0,0) is NW quadrant
	mlon, mlat, xlon, xlat = grid.TileBounds(0, 0, 1)
	if mlon != -180 || xlon != 0 || mlat != 0 || xlat != 90 {
		t.Errorf("z1(0,0) bounds = (%.1f,%.1f,%.1f,%.1f), want (-180,0,0,90)",
			mlon, mlat, xlon, xlat)
	}

	// Zoom 1, tile (1,1) is SE quadrant
	mlon, mlat, xlon, xlat = grid.TileBounds(1, 1, 1)
	if mlon != 0 || xlon != 180 || mlat != -90 || xlat != 0 {
		t.Errorf("z1(1,1) bounds = (%.1f,%.1f,%.1f,%.1f), want (0,-90,180,0)",
			mlon, mlat, xlon, xlat)
	}

	// Extent at zoom 0
	minX, minY, maxX, maxY := grid.Extent(0)
	if minX != 0 || minY != 0 || maxX != 0 || maxY != 0 {
		t.Errorf("z0 extent = (%d,%d,%d,%d), want (0,0,0,0)", minX, minY, maxX, maxY)
	}

	// Extent at zoom 1
	minX, minY, maxX, maxY = grid.Extent(1)
	if minX != 0 || minY != 0 || maxX != 1 || maxY != 1 {
		t.Errorf("z1 extent = (%d,%d,%d,%d), want (0,0,1,1)", minX, minY, maxX, maxY)
	}

	// Tile count
	count := grid.TileCount(0, 1)
	if count != 5 { // 1 + 4 = 5
		t.Errorf("TileCount(0,1) = %d, want 5", count)
	}
}

func TestTilerGridGlobalMercator(t *testing.T) {
	grid := NewGlobalMercatorGrid(256, 1)
	if grid.TileSize() != 256 {
		t.Errorf("TileSize = %d, want 256", grid.TileSize())
	}

	// Zoom 0, tile (0,0) is the whole world
	mlon, mlat, xlon, xlat := grid.TileBounds(0, 0, 0)
	if mlon != -180 || xlon != 180 {
		t.Errorf("z0 lon bounds = (%.1f,%.1f), want (-180,180)", mlon, xlon)
	}
	if mlat != -85.05112877980659 || xlat != 85.05112877980659 {
		t.Errorf("z0 lat bounds = (%.14f,%.14f), want (-85.05112877980659,85.05112877980659)", mlat, xlat)
	}

	// Extent at zoom 1
	_, _, maxX, maxY := grid.Extent(1)
	if maxX != 1 || maxY != 1 {
		t.Errorf("z1 extent max = (%d,%d), want (1,1)", maxX, maxY)
	}
}

func TestTilerGridWithRootTiles(t *testing.T) {
	grid := NewGlobalGeodeticGrid(256, 2)
	if grid.TileSize() != 256 {
		t.Errorf("TileSize = %d, want 256", grid.TileSize())
	}

	// With rootTiles=2, zoom 0 has 4 tiles
	minX, minY, maxX, maxY := grid.Extent(0)
	if minX != 0 || minY != 0 || maxX != 1 || maxY != 1 {
		t.Errorf("z0 extent = (%d,%d,%d,%d), want (0,0,1,1)", minX, minY, maxX, maxY)
	}

	// Tile (0,0) at zoom 0 with rootTiles=2
	mlon, mlat, xlon, xlat := grid.TileBounds(0, 0, 0)
	if mlon != -180 || xlon != 0 || mlat != 0 || xlat != 90 {
		t.Errorf("z0(0,0) bounds = (%.1f,%.1f,%.1f,%.1f), want (-180,0,0,90)",
			mlon, mlat, xlon, xlat)
	}

	// Extent at zoom 1: 2^1 * 2 = 4 tiles per dimension
	minX, minY, maxX, maxY = grid.Extent(1)
	if maxX != 3 || maxY != 3 {
		t.Errorf("z1 extent = (%d,%d), want (3,3)", maxX, maxY)
	}

	// Tile count: z0 = 4, z1 = 16
	count := grid.TileCount(0, 1)
	if count != 20 {
		t.Errorf("TileCount(0,1) = %d, want 20", count)
	}
}

func TestTilerGridLocal(t *testing.T) {
	box := [4]float64{-10, -5, 10, 5}
	grid := NewGlobalGeodeticLocalGrid(256, box, 1)
	if grid.TileSize() != 256 {
		t.Errorf("TileSize = %d, want 256", grid.TileSize())
	}
}

func TestGridIterator(t *testing.T) {
	grid := NewGlobalGeodeticGrid(256, 1)
	iter := NewGridIterator(grid, 0)

	count := 0
	iter.TraveGridIterator(func(x, y uint32, minLon, minLat, maxLon, maxLat float64) {
		count++
		if x != 0 || y != 0 {
			t.Errorf("expected tile (0,0), got (%d,%d)", x, y)
		}
		if minLon != -180 || maxLon != 180 || minLat != -90 || maxLat != 90 {
			t.Errorf("z0 bounds = (%.1f,%.1f,%.1f,%.1f)", minLon, minLat, maxLon, maxLat)
		}
	})
	if count != 1 {
		t.Errorf("expected 1 tile at zoom 0, got %d", count)
	}

	iter2 := NewGridIterator(grid, 1)
	count2 := 0
	iter2.TraveGridIterator(func(x, y uint32, _, _, _, _ float64) {
		count2++
	})
	if count2 != 4 {
		t.Errorf("expected 4 tiles at zoom 1, got %d", count2)
	}
}
