#include "mvtssr_c_api.h"

#include <limits>

#include <mbgl/gfx/headless_frontend.hpp>
#include <mbgl/map/bound_options.hpp>
#include <mbgl/map/camera.hpp>
#include <mbgl/map/map.hpp>
#include <mbgl/map/map_snapshotter.hpp>
#include <mbgl/storage/file_source.hpp>
#include <mbgl/storage/file_source_manager.hpp>
#include <mbgl/storage/offline.hpp>
#include <mbgl/storage/resource.hpp>
#include <mbgl/style/style.hpp>
#include <mbgl/util/geo.hpp>
#include <mbgl/util/image.hpp>
#include <mbgl/util/run_loop.hpp>

#ifdef __cplusplus
extern "C" {
#endif

struct _mvtssr_canonical_tileid_t {
  mbgl::CanonicalTileID id;
};

struct _mvtssr_latlng_t {
  mbgl::LatLng ll;
};

struct _mvtssr_latlng_bounds_t {
  mbgl::LatLngBounds bounds;
};

struct _mvtssr_edge_insets_t {
  mbgl::EdgeInsets edge;
};

struct _mvtssr_screen_coordinate_t {
  mbgl::ScreenCoordinate sc;
};

struct _mvtssr_camera_options_t {
  mbgl::CameraOptions opt;
};

struct _mvtssr_runloop_t {
  mbgl::util::RunLoop *loop;
};

struct _mvtssr_size_t {
  mbgl::Size si;
};

struct _mvtssr_headless_frontend_t {
  mbgl::HeadlessFrontend frontend;
};

struct _mvtssr_map_observer_t {
  mbgl::MapObserver obser;
};

struct _mvtssr_map_options_t {
  mbgl::MapOptions opt;
};

struct _mvtssr_file_source_t {
  mbgl::PassRefPtr<mbgl::FileSource> src;
};

struct _mvtssr_file_source_manager_t {
  mbgl::FileSourceManager *mag;
};

struct _mvtssr_file_source_factory_t {
  std::function<std::unique_ptr<mbgl::FileSource>(
      const mbgl::ResourceOptions &)>
      func;
};

struct _mvtssr_resource_options_t {
  mbgl::ResourceOptions opt;
};

struct _mvtssr_style_t {
  mbgl::style::Style sty;
};

struct _mvtssr_bound_options_t {
  mbgl::BoundOptions opt;
};

struct _mvtssr_resource_t {
  mbgl::Resource res;
};

struct _mvtssr_map_snapshotter_observer_t {
  mbgl::MapSnapshotterObserver obser;
};

struct _mvtssr_map_snapshotter_t {
  mbgl::MapSnapshotter snap;
};

struct _mvtssr_offline_region_metadata_t {
  mbgl::OfflineRegionMetadata md;
};

struct _mvtssr_offline_region_status_t {
  mbgl::OfflineRegionStatus status;
};

struct _mvtssr_offline_region_observer_t {
  mbgl::OfflineRegionObserver obser;
};

struct _mvtssr_offline_region_definition_t {
  mbgl::OfflineRegionDefinition def;
};

struct _mvtssr_offline_region_t {
  mbgl::OfflineRegion region;
};

struct _mvtssr_premultiplied_image_t {
  mbgl::PremultipliedImage img;
};

struct _mvtssr_unassociatedImage_image_t {
  mbgl::UnassociatedImage img;
};

MVTSSRAPICALL mvtssr_canonical_tileid_t *
new_mvtssr_canonical_tileid(uint8_t z, uint32_t x, uint32_t y) {
  return new mvtssr_canonical_tileid_t{mbgl::CanonicalTileID(z, x, y)};
}

MVTSSRAPICALL void mvtssr_canonical_tileid_free(mvtssr_canonical_tileid_t *id) {
  delete id;
}

MVTSSRAPICALL mvtssr_canonical_tileid_t *
mvtssr_canonical_tileid_scaled_to(mvtssr_canonical_tileid_t *id, uint8_t z) {
  return new mvtssr_canonical_tileid_t{id->id.scaledTo(z)};
}

MVTSSRAPICALL _Bool mvtssr_canonical_tileid_is_child_of(
    mvtssr_canonical_tileid_t *id, mvtssr_canonical_tileid_t *tar) {
  return id->id.isChildOf(tar->id);
}

MVTSSRAPICALL _Bool mvtssr_canonical_tileid_is_child_less(
    mvtssr_canonical_tileid_t *a, mvtssr_canonical_tileid_t *b) {
  return a->id < b->id;
}

MVTSSRAPICALL _Bool mvtssr_canonical_tileid_is_child_eq(
    mvtssr_canonical_tileid_t *a, mvtssr_canonical_tileid_t *b) {
  return a->id == b->id;
}

MVTSSRAPICALL _Bool mvtssr_canonical_tileid_children(
    mvtssr_canonical_tileid_t *id, mvtssr_canonical_tileid_t **childs) {
  std::array<mbgl::CanonicalTileID, 4> childrens = id->id.children();
  childs[0] = new mvtssr_canonical_tileid_t{childrens[0]};
  childs[1] = new mvtssr_canonical_tileid_t{childrens[1]};
  childs[2] = new mvtssr_canonical_tileid_t{childrens[2]};
  childs[3] = new mvtssr_canonical_tileid_t{childrens[3]};
  return true;
}

MVTSSRAPICALL mvtssr_latlng_t *new_mvtssr_latlng(double lat, double lon) {
  return new mvtssr_latlng_t{mbgl::LatLng(lat, lon)};
}

MVTSSRAPICALL void mvtssr_mvtssr_latlng_free(mvtssr_latlng_t *ll) { delete ll; }

MVTSSRAPICALL mvtssr_latlng_t *
new_mvtssr_latlng_with_id(mvtssr_canonical_tileid_t *id) {
  return new mvtssr_latlng_t{mbgl::LatLng(id->id)};
}

MVTSSRAPICALL double mvtssr_latlng_latitude(mvtssr_latlng_t *latlon) {
  return latlon->ll.latitude();
}

MVTSSRAPICALL double mvtssr_latlng_longitude(mvtssr_latlng_t *latlon) {
  return latlon->ll.longitude();
}

MVTSSRAPICALL mvtssr_latlng_bounds_t *latlng_bounds_world() {
  return new mvtssr_latlng_bounds_t{mbgl::LatLngBounds::world()};
}

MVTSSRAPICALL mvtssr_latlng_bounds_t *
new_latlng_bounds(mvtssr_latlng_t *latlon) {
  return new mvtssr_latlng_bounds_t{mbgl::LatLngBounds::singleton(latlon->ll)};
}

MVTSSRAPICALL mvtssr_latlng_bounds_t *hull_latlng_bounds(mvtssr_latlng_t *a,
                                                         mvtssr_latlng_t *b) {
  return new mvtssr_latlng_bounds_t{mbgl::LatLngBounds::hull(a->ll, b->ll)};
}

MVTSSRAPICALL mvtssr_latlng_bounds_t *latlng_bounds_empty() {
  return new mvtssr_latlng_bounds_t{mbgl::LatLngBounds::empty()};
}

MVTSSRAPICALL mvtssr_latlng_bounds_t *
new_latlng_bounds_with_id(mvtssr_canonical_tileid_t *id) {
  return new mvtssr_latlng_bounds_t{id->id};
}

MVTSSRAPICALL _Bool latlng_bounds_valid(mvtssr_latlng_bounds_t *bounds) {
  return bounds->bounds.valid();
}

MVTSSRAPICALL double latlng_bounds_south(mvtssr_latlng_bounds_t *bounds) {
  return bounds->bounds.south();
}

MVTSSRAPICALL double latlng_bounds_west(mvtssr_latlng_bounds_t *bounds) {
  return bounds->bounds.west();
}

MVTSSRAPICALL double latlng_bounds_north(mvtssr_latlng_bounds_t *bounds) {
  return bounds->bounds.north();
}

MVTSSRAPICALL double latlng_bounds_east(mvtssr_latlng_bounds_t *bounds) {
  return bounds->bounds.east();
}

MVTSSRAPICALL mvtssr_latlng_t *
latlng_bounds_southwest(mvtssr_latlng_bounds_t *bounds) {
  return new mvtssr_latlng_t{bounds->bounds.southwest()};
}

MVTSSRAPICALL mvtssr_latlng_t *
latlng_bounds_northeast(mvtssr_latlng_bounds_t *bounds) {
  return new mvtssr_latlng_t{bounds->bounds.northeast()};
}

MVTSSRAPICALL mvtssr_latlng_t *
latlng_bounds_southeast(mvtssr_latlng_bounds_t *bounds) {
  return new mvtssr_latlng_t{bounds->bounds.southeast()};
}

MVTSSRAPICALL mvtssr_latlng_t *
latlng_bounds_northwest(mvtssr_latlng_bounds_t *bounds) {
  return new mvtssr_latlng_t{bounds->bounds.northwest()};
}

MVTSSRAPICALL mvtssr_latlng_t *
latlng_bounds_center(mvtssr_latlng_bounds_t *bounds) {
  return new mvtssr_latlng_t{bounds->bounds.center()};
}

MVTSSRAPICALL mvtssr_latlng_t *
latlng_bounds_constrain(mvtssr_latlng_bounds_t *bounds, mvtssr_latlng_t *p) {
  return new mvtssr_latlng_t{bounds->bounds.constrain(p->ll)};
}

MVTSSRAPICALL void latlng_bounds_extend(mvtssr_latlng_bounds_t *bounds,
                                        mvtssr_latlng_t *p) {
  bounds->bounds.extend(p->ll);
}

MVTSSRAPICALL void
latlng_bounds_extend_bounds(mvtssr_latlng_bounds_t *bounds,
                            mvtssr_latlng_bounds_t *ext_bounds) {
  bounds->bounds.extend(ext_bounds->bounds);
}

MVTSSRAPICALL _Bool latlng_bounds_is_empty(mvtssr_latlng_bounds_t *bounds) {
  return bounds->bounds.isEmpty();
}

MVTSSRAPICALL void mvtssr_latlng_bounds_free(mvtssr_latlng_bounds_t *llb) {
  delete llb;
}

MVTSSRAPICALL _Bool latlng_bounds_contains_id(mvtssr_latlng_bounds_t *bounds,
                                              mvtssr_canonical_tileid_t *id) {
  return bounds->bounds.contains(id->id);
}

MVTSSRAPICALL _Bool latlng_bounds_contains_point(mvtssr_latlng_bounds_t *bounds,
                                                 mvtssr_latlng_t *point) {
  return bounds->bounds.contains(point->ll);
}

MVTSSRAPICALL _Bool latlng_bounds_contains_bounds(
    mvtssr_latlng_bounds_t *bounds, mvtssr_latlng_bounds_t *area) {
  return bounds->bounds.contains(area->bounds);
}

MVTSSRAPICALL _Bool latlng_bounds_contains_intersects(
    mvtssr_latlng_bounds_t *bounds, mvtssr_latlng_bounds_t *area) {
  return bounds->bounds.intersects(area->bounds);
}

MVTSSRAPICALL mvtssr_edge_insets_t *new_mvtssr_edge_insets(double t, double l,
                                                           double b, double r) {
  return new mvtssr_edge_insets_t{mbgl::EdgeInsets(t, l, b, r)};
}

MVTSSRAPICALL double mvtssr_edge_top(mvtssr_edge_insets_t *edge) {
  return edge->edge.top();
}

MVTSSRAPICALL double mvtssr_edge_left(mvtssr_edge_insets_t *edge) {
  return edge->edge.left();
}

MVTSSRAPICALL double mvtssr_edge_bottom(mvtssr_edge_insets_t *edge) {
  return edge->edge.bottom();
}

MVTSSRAPICALL double mvtssr_edge_right(mvtssr_edge_insets_t *edge) {
  return edge->edge.right();
}

MVTSSRAPICALL double mvtssr_edge_is_flush(mvtssr_edge_insets_t *edge) {
  return edge->edge.isFlush();
}

MVTSSRAPICALL mvtssr_screen_coordinate_t *
mvtssr_edge_get_center(mvtssr_edge_insets_t *edge, uint16_t width,
                       uint16_t height) {
  return new mvtssr_screen_coordinate_t{edge->edge.getCenter(width, height)};
}

MVTSSRAPICALL _Bool mvtssr_edge_eq(mvtssr_edge_insets_t *a,
                                   mvtssr_edge_insets_t *b) {
  return a->edge == b->edge;
}

MVTSSRAPICALL void mvtssr_edge_add(mvtssr_edge_insets_t *a,
                                   mvtssr_edge_insets_t *b) {
  a->edge += b->edge;
}

MVTSSRAPICALL mvtssr_screen_coordinate_t *new_screen_coordinate(double x,
                                                                double y) {
  return new mvtssr_screen_coordinate_t{mbgl::ScreenCoordinate(x, y)};
}

MVTSSRAPICALL void
mvtssr_screen_coordinate_free(mvtssr_screen_coordinate_t *sc) {
  delete sc;
}

MVTSSRAPICALL double
mvtssr_screen_coordinate_x(mvtssr_screen_coordinate_t *sc) {
  return sc->sc.x;
}

MVTSSRAPICALL double
mvtssr_screen_coordinate_y(mvtssr_screen_coordinate_t *sc) {
  return sc->sc.y;
}

MVTSSRAPICALL mvtssr_camera_options_t *new_mvtssr_camera_options() {
  return new mvtssr_camera_options_t{};
}

MVTSSRAPICALL void
mvtssr_camera_options_set_center(mvtssr_camera_options_t *opt,
                                 mvtssr_latlng_t *point) {
  opt->opt.center = point->ll;
}

MVTSSRAPICALL void
mvtssr_camera_options_set_padding(mvtssr_camera_options_t *opt,
                                  mvtssr_edge_insets_t *edge) {
  opt->opt.padding = edge->edge;
}

MVTSSRAPICALL void
mvtssr_camera_options_set_anchor(mvtssr_camera_options_t *opt,
                                 mvtssr_screen_coordinate_t *sc) {
  opt->opt.anchor = sc->sc;
}

MVTSSRAPICALL void mvtssr_camera_options_set_zoom(mvtssr_camera_options_t *opt,
                                                  double zoom) {
  opt->opt.zoom = zoom;
}

MVTSSRAPICALL void
mvtssr_camera_options_set_bearing(mvtssr_camera_options_t *opt,
                                  double bearing) {
  opt->opt.bearing = bearing;
}

MVTSSRAPICALL void mvtssr_camera_options_set_pitch(mvtssr_camera_options_t *opt,
                                                   double pitch) {
  opt->opt.pitch = pitch;
}

MVTSSRAPICALL mvtssr_latlng_t *
mvtssr_camera_options_get_center(mvtssr_camera_options_t *opt) {
  if (opt->opt.center) {
    return new mvtssr_latlng_t{*opt->opt.center};
  }
  return nullptr;
}

MVTSSRAPICALL mvtssr_edge_insets_t *
mvtssr_camera_options_get_padding(mvtssr_camera_options_t *opt) {
  if (opt->opt.padding) {
    return new mvtssr_edge_insets_t{*opt->opt.padding};
  }
  return nullptr;
}

MVTSSRAPICALL mvtssr_screen_coordinate_t *
mvtssr_camera_options_get_anchor(mvtssr_camera_options_t *opt) {
  if (opt->opt.padding) {
    return new mvtssr_screen_coordinate_t{*opt->opt.anchor};
  }
  return nullptr;
}

MVTSSRAPICALL double
mvtssr_camera_options_get_zoom(mvtssr_camera_options_t *opt) {
  if (opt->opt.zoom) {
    return *opt->opt.zoom;
  }
  return std::numeric_limits<double>::infinity();
}

MVTSSRAPICALL
double mvtssr_camera_options_get_bearing(mvtssr_camera_options_t *opt) {
  if (opt->opt.bearing) {
    return *opt->opt.bearing;
  }
  return std::numeric_limits<double>::infinity();
}

MVTSSRAPICALL double
mvtssr_camera_options_get_pitch(mvtssr_camera_options_t *opt) {
  if (opt->opt.pitch) {
    return *opt->opt.pitch;
  }
  return std::numeric_limits<double>::infinity();
}

#ifdef __cplusplus
}
#endif