package mvtssr

type MapModeType uint32

const (
	Continuous MapModeType = 0
	Static     MapModeType = 1
	Tile       MapModeType = 2
)

type ConstrainMode uint32

const (
	None           ConstrainMode = 0
	HeightOnly     ConstrainMode = 1
	WidthAndHeight ConstrainMode = 2
)

type ViewportMode uint32

const (
	Default  ViewportMode = 0
	FlippedY ViewportMode = 1
)

type NorthOrientation uint8

const (
	Upwards   NorthOrientation = 0
	Rightward NorthOrientation = 1
	Downwards NorthOrientation = 2
	Leftwards NorthOrientation = 3
)

type ResourceKind uint8

const (
	Resource_Unknown     ResourceKind = 0
	Resource_Style       ResourceKind = 1
	Resource_Source      ResourceKind = 2
	Resource_Tile        ResourceKind = 3
	Resource_Glyphs      ResourceKind = 4
	Resource_SpriteImage ResourceKind = 5
	Resource_SpriteJSON  ResourceKind = 6
	Resource_Image       ResourceKind = 7
)

type FileType uint8

const (
	Asset          FileType = 0
	Database       FileType = 1
	FileSystem     FileType = 2
	Network        FileType = 3
	ResourceLoader FileType = 4
)

type CameraChangeMode uint32

const (
	Immediate CameraChangeMode = 0
	Animated  CameraChangeMode = 1
)

type RenderMode uint32

const (
	Partial RenderMode = 0
	Full    RenderMode = 1
)

type MapLoadError uint32

const (
	StyleParseError MapLoadError = 0
	StyleLoadError  MapLoadError = 1
	NotFoundError   MapLoadError = 2
	UnknownError    MapLoadError = 3
)

type ReasonError uint8

const (
	Success    ReasonError = 1
	NotFound   ReasonError = 2
	Server     ReasonError = 3
	Connection ReasonError = 4
	RateLimit  ReasonError = 5
	Other      ReasonError = 6
)
