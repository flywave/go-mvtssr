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

class GoMapObserver : public mbgl::MapObserver {
public:
  GoMapObserver(void *ctx_) : ctx(ctx_) {}

  virtual void onCameraWillChange(CameraChangeMode) override {}
  virtual void onCameraIsChanging() override {}
  virtual void onCameraDidChange(CameraChangeMode) override {}
  virtual void onWillStartLoadingMap() override {}
  virtual void onDidFinishLoadingMap() override {}
  virtual void onDidFailLoadingMap(mbgl::MapLoadError,
                                   const std::string &) override {}
  virtual void onWillStartRenderingFrame() override {}
  virtual void onDidFinishRenderingFrame(RenderFrameStatus) override {}
  virtual void onWillStartRenderingMap() override {}
  virtual void onDidFinishRenderingMap(RenderMode) override {}
  virtual void onDidFinishLoadingStyle() override {}
  virtual void onSourceChanged(mbgl::style::Source &) override {}
  virtual void onDidBecomeIdle() override {}
  virtual void onStyleImageMissing(const std::string &) override {}

  void *ctx;
};

class GoFileSource : public mbgl::FileSource {
public:
  GoFileSource(void *ctx_) : ctx(ctx_) {}

  virtual std::unique_ptr<mbgl::AsyncRequest> request(const mbgl::Resource &,
                                                      Callback) override {}

  virtual bool canRequest(const mbgl::Resource &) const override {}

  void *ctx;
};

struct GoFileSourceFactory {
  GoFileSourceFactory(void *ctx_) : ctx(ctx_) {}

  std::unique_ptr<mbgl::FileSource> operator()(const mbgl::ResourceOptions &) {
    return nullptr;
  }

  void *ctx;
};

class GoMapSnapshotterObserver : public mbgl::MapSnapshotterObserver {
public:
  GoMapSnapshotterObserver(void *ctx_) : ctx(ctx_) {}

  virtual void onDidFailLoadingStyle(const std::string &) override {}
  virtual void onDidFinishLoadingStyle() override {}
  virtual void onStyleImageMissing(const std::string &) override {}

  void *ctx;
};

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
  std::unique_ptr<mbgl::MapObserver> obser;
};

struct _mvtssr_map_options_t {
  mbgl::MapOptions opt;
};

struct _mvtssr_file_source_t {
  std::shared_ptr<mbgl::FileSource> src;
};

struct _mvtssr_file_source_manager_t {
  mbgl::FileSourceManager *mag;
};

struct _mvtssr_file_source_factory_t {
  mbgl::FileSourceType type;
  GoFileSourceFactory factory;
};

struct _mvtssr_resource_options_t {
  mbgl::ResourceOptions opt;
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

struct _mvtssr_map_t {
  std::unique_ptr<mbgl::Map> map;
};

struct _mvtssr_style_t {
  std::unique_ptr<mbgl::style::Style> st;
};

struct _mvtssr_map_snapshotter_t {
  mbgl::MapSnapshotter snap;
};

struct _mvtssr_map_snapshotter_result_t {
  void *ctx;
  std::exception_ptr err;
  mbgl::PremultipliedImage img;
  std::vector<std::string> attr;
  std::function<mbgl::ScreenCoordinate(const mbgl::LatLng &)> point_for;
  std::function<mbgl::LatLng(const mbgl::ScreenCoordinate &)> latLng_for;
};

struct _mvtssr_premultiplied_image_t {
  mbgl::PremultipliedImage img;
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

MVTSSRAPICALL _Bool mvtssr_canonical_tileid_eq(mvtssr_canonical_tileid_t *a,
                                               mvtssr_canonical_tileid_t *b) {
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

MVTSSRAPICALL mvtssr_latlng_t *mvtssr_new_latlng(double lat, double lon) {
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

MVTSSRAPICALL mvtssr_edge_insets_t *mvtssr_new_edge_insets(double t, double l,
                                                           double b, double r) {
  return new mvtssr_edge_insets_t{mbgl::EdgeInsets(t, l, b, r)};
}

MVTSSRAPICALL void mvtssr_edge_free(mvtssr_edge_insets_t *edge) { delete edge; }

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

MVTSSRAPICALL _Bool mvtssr_edge_is_flush(mvtssr_edge_insets_t *edge) {
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

MVTSSRAPICALL mvtssr_screen_coordinate_t *
mvtssr_new_screen_coordinate(double x, double y) {
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

MVTSSRAPICALL mvtssr_camera_options_t *mvtssr_new_camera_options() {
  return new mvtssr_camera_options_t{};
}

MVTSSRAPICALL void mvtssr_camera_options_free(mvtssr_camera_options_t *opt) {
  delete opt;
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

MVTSSRAPICALL
mvtssr_size_t *mvtssr_new_size(uint32_t width, uint32_t height) {
  return new mvtssr_size_t{mbgl::Size(width, height)};
}

MVTSSRAPICALL void mvtssr_size_free(mvtssr_size_t *si) { delete si; }

MVTSSRAPICALL uint32_t mvtssr_size_area(mvtssr_size_t *si) {
  return si->si.area();
}

MVTSSRAPICALL float mvtssr_size_aspect_ratio(mvtssr_size_t *si) {
  return si->si.aspectRatio();
}

MVTSSRAPICALL _Bool mvtssr_size_is_empty(mvtssr_size_t *si) {
  return si->si.isEmpty();
}

MVTSSRAPICALL
mvtssr_runloop_t *mvtssr_new_runloop() { return new mvtssr_runloop_t{}; }

MVTSSRAPICALL void mvtssr_runloop_free(mvtssr_runloop_t *loop) { delete loop; }

MVTSSRAPICALL
mvtssr_map_observer_t *mvtssr_null_map_observer() {
  return new mvtssr_map_observer_t{std::make_unique<mbgl::MapObserver>()};
}

MVTSSRAPICALL
mvtssr_map_observer_t *mvtssr_new_map_observer(void *ctx) {
  return new mvtssr_map_observer_t{
      std::unique_ptr<mbgl::MapObserver>(new GoMapObserver{ctx})};
}

MVTSSRAPICALL void mvtssr_map_observer_free(mvtssr_map_observer_t *ob) {
  delete ob;
}

MVTSSRAPICALL mvtssr_map_options_t *mvtssr_new_map_options() {
  return new mvtssr_map_options_t{};
}

MVTSSRAPICALL void mvtssr_map_options_free(mvtssr_map_options_t *op) {
  delete op;
}

MVTSSRAPICALL void mvtssr_map_options_set_map_mode(mvtssr_map_options_t *opt,
                                                   uint32_t mode) {
  opt->opt.withMapMode(static_cast<mbgl::MapMode>(mode));
}

MVTSSRAPICALL void
mvtssr_map_options_set_constrain_mode(mvtssr_map_options_t *opt,
                                      uint32_t mode) {
  opt->opt.withConstrainMode(static_cast<mbgl::ConstrainMode>(mode));
}

MVTSSRAPICALL void
mvtssr_map_options_set_viewport_mode(mvtssr_map_options_t *opt, uint32_t mode) {
  opt->opt.withViewportMode(static_cast<mbgl::ViewportMode>(mode));
}

MVTSSRAPICALL void
mvtssr_map_options_set_cross_source_collisions(mvtssr_map_options_t *opt,
                                               _Bool sc) {
  opt->opt.withCrossSourceCollisions(sc);
}

MVTSSRAPICALL void
mvtssr_map_options_set_north_orientation(mvtssr_map_options_t *opt,
                                         uint8_t ori) {
  opt->opt.withNorthOrientation(static_cast<mbgl::NorthOrientation>(ori));
}

MVTSSRAPICALL void mvtssr_map_options_set_size(mvtssr_map_options_t *opt,
                                               mvtssr_size_t *si) {
  opt->opt.withSize(si->si);
}

MVTSSRAPICALL void mvtssr_map_options_set_pixel_ratio(mvtssr_map_options_t *opt,
                                                      float ratio) {
  opt->opt.withPixelRatio(ratio);
}

MVTSSRAPICALL
mvtssr_file_source_t *mvtssr_new_file_source(void *ctx) {
  return new mvtssr_file_source_t{std::make_shared<GoFileSource>(ctx)};
}

MVTSSRAPICALL void mvtssr_file_source_free(mvtssr_file_source_t *s) {
  delete s;
}

MVTSSRAPICALL mvtssr_file_source_manager_t *mvtssr_get_file_source_manager() {
  return new mvtssr_file_source_manager_t{mbgl::FileSourceManager::get()};
}

MVTSSRAPICALL void
mvtssr_file_source_manager_free(mvtssr_file_source_manager_t *s) {
  delete s;
}

MVTSSRAPICALL void mvtssr_file_source_manager_register_file_source_factory(
    mvtssr_file_source_manager_t *s, mvtssr_file_source_factory_t *factory) {
  s->mag->registerFileSourceFactory(factory->type, factory->factory);
}

MVTSSRAPICALL void mvtssr_file_source_manager_unregister_file_source_factory(
    mvtssr_file_source_manager_t *s, mvtssr_file_source_factory_t *factory) {
  s->mag->unRegisterFileSourceFactory(factory->type);
}

MVTSSRAPICALL
mvtssr_file_source_factory_t *mvtssr_new_file_source_factory(uint8_t file_type,
                                                             void *ctx) {
  return new mvtssr_file_source_factory_t{
      static_cast<mbgl::FileSourceType>(file_type), GoFileSourceFactory(ctx)};
}

MVTSSRAPICALL void
mvtssr_file_source_factory_free(mvtssr_file_source_factory_t *s) {
  delete s;
}

MVTSSRAPICALL mvtssr_resource_options_t *mvtssr_new_resource_options() {
  return new mvtssr_resource_options_t{};
}

MVTSSRAPICALL void mvtssr_resource_options_free(mvtssr_resource_options_t *op) {
  delete op;
}

MVTSSRAPICALL void
mvtssr_resource_options_set_access_token(mvtssr_resource_options_t *opt,
                                         const char *token) {
  opt->opt.withAccessToken(token);
}

MVTSSRAPICALL void
mvtssr_resource_options_set_base_url(mvtssr_resource_options_t *opt,
                                     const char *url) {
  opt->opt.withBaseURL(std::string(url));
}

MVTSSRAPICALL void
mvtssr_resource_options_set_asset_path(mvtssr_resource_options_t *opt,
                                       const char *path) {
  opt->opt.withAssetPath(path);
}

MVTSSRAPICALL void
mvtssr_resource_options_set_maximum_cache_size(mvtssr_resource_options_t *opt,
                                               uint64_t size) {
  opt->opt.withMaximumCacheSize(size);
}

MVTSSRAPICALL mvtssr_headless_frontend_t *
mvtssr_new_headless_frontend(mvtssr_size_t *size, float pixelRatio) {
  return new mvtssr_headless_frontend_t{{size->si, pixelRatio}};
}

MVTSSRAPICALL void
mvtssr_headless_frontend_free(mvtssr_headless_frontend_t *op) {
  delete op;
}

MVTSSRAPICALL void
mvtssr_headless_frontend_reset(mvtssr_headless_frontend_t *op) {
  op->frontend.reset();
}

MVTSSRAPICALL mvtssr_screen_coordinate_t *
mvtssr_headless_frontend_pixel_for_latlng(mvtssr_headless_frontend_t *op,
                                          mvtssr_latlng_t *latlon) {
  return new mvtssr_screen_coordinate_t{
      op->frontend.pixelForLatLng(latlon->ll)};
}

MVTSSRAPICALL mvtssr_latlng_t *
mvtssr_headless_frontend_latlng_for_pixel(mvtssr_headless_frontend_t *op,
                                          mvtssr_screen_coordinate_t *coord) {
  return new mvtssr_latlng_t{op->frontend.latLngForPixel(coord->sc)};
}

MVTSSRAPICALL void
mvtssr_headless_frontend_set_size(mvtssr_headless_frontend_t *op,
                                  mvtssr_size_t *size) {
  op->frontend.setSize(size->si);
}

MVTSSRAPICALL mvtssr_size_t *
mvtssr_headless_frontend_get_size(mvtssr_headless_frontend_t *op) {
  return new mvtssr_size_t{op->frontend.getSize()};
}

MVTSSRAPICALL mvtssr_premultiplied_image_t *
mvtssr_headless_frontend_render(mvtssr_headless_frontend_t *op,
                                mvtssr_map_t *map) {
  auto res = op->frontend.render(*map->map);
  return new mvtssr_premultiplied_image_t{std::move(res.image)};
}

MVTSSRAPICALL
mvtssr_map_snapshotter_observer_t *mvtssr_null_map_snapshotter_observer() {
  return new mvtssr_map_snapshotter_observer_t{};
}

MVTSSRAPICALL
mvtssr_map_snapshotter_observer_t *
mvtssr_new_map_snapshotter_observer(void *ctx) {
  return new mvtssr_map_snapshotter_observer_t{GoMapSnapshotterObserver(ctx)};
}

MVTSSRAPICALL void
mvtssr_map_snapshotter_observer_free(mvtssr_map_snapshotter_observer_t *op) {
  delete op;
}

MVTSSRAPICALL mvtssr_map_snapshotter_t *
mvtssr_new_map_snapshotter(mvtssr_size_t *size, float pixelRatio,
                           mvtssr_resource_options_t *opts,
                           mvtssr_map_snapshotter_observer_t *obser) {
  return new mvtssr_map_snapshotter_t{
      {size->si, pixelRatio, opts->opt, obser->obser}};
}

MVTSSRAPICALL void mvtssr_map_snapshotter_free(mvtssr_map_snapshotter_t *snap) {
  delete snap;
}

MVTSSRAPICALL void
mvtssr_map_snapshotter_set_style_url(mvtssr_map_snapshotter_t *snap,
                                     const char *url) {
  snap->snap.setStyleURL(url);
}

MVTSSRAPICALL char *
mvtssr_map_snapshotter_get_style_url(mvtssr_map_snapshotter_t *snap) {
  return strdup(snap->snap.getStyleURL().c_str());
}

MVTSSRAPICALL void
mvtssr_map_snapshotter_set_style(mvtssr_map_snapshotter_t *snap,
                                 const char *style) {
  snap->snap.setStyleJSON(style);
}

MVTSSRAPICALL char *
mvtssr_map_snapshotter_get_style(mvtssr_map_snapshotter_t *snap) {
  return strdup(snap->snap.getStyleJSON().c_str());
}

MVTSSRAPICALL void
mvtssr_map_snapshotter_set_size(mvtssr_map_snapshotter_t *snap,
                                mvtssr_size_t *size) {
  snap->snap.setSize(size->si);
}

MVTSSRAPICALL mvtssr_size_t *
mvtssr_map_snapshotter_get_size(mvtssr_map_snapshotter_t *snap) {
  return new mvtssr_size_t{snap->snap.getSize()};
}

MVTSSRAPICALL void
mvtssr_map_snapshotter_set_camera_options(mvtssr_map_snapshotter_t *snap,
                                          mvtssr_camera_options_t *opts) {
  snap->snap.setCameraOptions(opts->opt);
}

MVTSSRAPICALL mvtssr_camera_options_t *
mvtssr_map_snapshotter_get_camera_options(mvtssr_map_snapshotter_t *snap) {
  return new mvtssr_camera_options_t{snap->snap.getCameraOptions()};
}

MVTSSRAPICALL void
mvtssr_map_snapshotter_set_region(mvtssr_map_snapshotter_t *snap,
                                  mvtssr_latlng_bounds_t *bounds) {
  snap->snap.setRegion(bounds->bounds);
}

MVTSSRAPICALL mvtssr_latlng_bounds_t *
mvtssr_map_snapshotter_get_region(mvtssr_map_snapshotter_t *snap) {
  return new mvtssr_latlng_bounds_t{snap->snap.getRegion()};
}

MVTSSRAPICALL void
mvtssr_map_snapshotter_cancel(mvtssr_map_snapshotter_t *snap) {
  snap->snap.cancel();
}

void mvtssr_map_snapshotter_result_finish(
    mvtssr_map_snapshotter_result_t *result) {
  // TODO
}

MVTSSRAPICALL void mvtssr_map_snapshotter_snapshot(
    mvtssr_map_snapshotter_t *snap, mvtssr_map_snapshotter_result_t *result) {
  snap->snap.snapshot([result](std::exception_ptr e,
                               mbgl::PremultipliedImage img,
                               mbgl::MapSnapshotter::Attributions attr,
                               mbgl::MapSnapshotter::PointForFn p,
                               mbgl::MapSnapshotter::LatLngForFn l) {
    result->err = e;
    result->img = std::move(img);
    result->attr = attr;
    result->point_for = p;
    result->latLng_for = l;
    mvtssr_map_snapshotter_result_finish(result);
  });
}

MVTSSRAPICALL
mvtssr_map_snapshotter_result_t *mvtssr_new_map_snapshotter_result(void *ctx) {
  return new mvtssr_map_snapshotter_result_t{ctx};
}

MVTSSRAPICALL void
mvtssr_map_snapshotter_result_free(mvtssr_map_snapshotter_result_t *op) {
  delete op;
}

MVTSSRAPICALL mvtssr_premultiplied_image_t *
mvtssr_map_snapshotter_result_get_image(mvtssr_map_snapshotter_result_t *op) {
  return new mvtssr_premultiplied_image_t{std::move(op->img)};
}

MVTSSRAPICALL char *
mvtssr_map_snapshotter_result_get_error(mvtssr_map_snapshotter_result_t *op) {
  if (op->err != nullptr) {
    try {
      std::rethrow_exception(op->err);
    } catch (std::exception &e) {
      return strdup(e.what());
    }
  }
  return nullptr;
}

MVTSSRAPICALL mvtssr_screen_coordinate_t *
mvtssr_map_snapshotter_result_pixel_for_latlng(
    mvtssr_map_snapshotter_result_t *op, mvtssr_latlng_t *latlon) {
  return new mvtssr_screen_coordinate_t{op->point_for(latlon->ll)};
}

MVTSSRAPICALL mvtssr_latlng_t *mvtssr_map_snapshotter_result_latlng_for_pixel(
    mvtssr_map_snapshotter_result_t *op, mvtssr_screen_coordinate_t *coord) {
  return new mvtssr_latlng_t{op->latLng_for(coord->sc)};
}

MVTSSRAPICALL
mvtssr_style_t *mvtssr_new_style(mvtssr_file_source_t *source,
                                 float pixelRatio) {
  return new mvtssr_style_t{
      std::make_unique<mbgl::style::Style>(source->src, pixelRatio)};
}

MVTSSRAPICALL void mvtssr_style_free(mvtssr_style_t *m) { delete m; }

MVTSSRAPICALL char *mvtssr_style_get_json(mvtssr_style_t *m) {
  return strdup(m->st->getJSON().c_str());
}

MVTSSRAPICALL char *mvtssr_style_get_url(mvtssr_style_t *m) {
  return strdup(m->st->getURL().c_str());
}

MVTSSRAPICALL
mvtssr_bound_options_t *mvtssr_new_bound_options() {
  return new mvtssr_bound_options_t{};
}

MVTSSRAPICALL void mvtssr_bound_options_free(mvtssr_bound_options_t *m) {
  delete m;
}

MVTSSRAPICALL void
mvtssr_bound_options_set_bounds(mvtssr_bound_options_t *opt,
                                mvtssr_latlng_bounds_t *bounds) {
  opt->opt.bounds = bounds->bounds;
}

MVTSSRAPICALL void mvtssr_bound_options_set_min_zoom(mvtssr_bound_options_t *m,
                                                     double z) {
  m->opt.minZoom = z;
}

MVTSSRAPICALL void mvtssr_bound_options_set_max_zoom(mvtssr_bound_options_t *m,
                                                     double z) {
  m->opt.maxZoom = z;
}

MVTSSRAPICALL void mvtssr_bound_options_set_min_pitch(mvtssr_bound_options_t *m,
                                                      double z) {
  m->opt.minPitch = z;
}

MVTSSRAPICALL void mvtssr_bound_options_set_max_pitch(mvtssr_bound_options_t *m,
                                                      double z) {
  m->opt.maxPitch = z;
}

MVTSSRAPICALL
mvtssr_map_t *mvtssr_new_map(mvtssr_headless_frontend_t *fr,
                             mvtssr_map_observer_t *obser,
                             mvtssr_map_options_t *opts,
                             mvtssr_resource_options_t *ropts) {
  return new mvtssr_map_t{std::make_unique<mbgl::Map>(
      fr->frontend, *obser->obser, opts->opt, ropts->opt)};
}

MVTSSRAPICALL void mvtssr_map_free(mvtssr_map_t *m) { delete m; }

MVTSSRAPICALL void mvtssr_map_set_style(mvtssr_map_t *m, mvtssr_style_t *s) {
  m->map->setStyle(std::move(s->st));
}

MVTSSRAPICALL mvtssr_camera_options_t *
mvtssr_map_camera_options(mvtssr_map_t *m, mvtssr_edge_insets_t *e) {
  return new mvtssr_camera_options_t{m->map->getCameraOptions()};
}

MVTSSRAPICALL void mvtssr_map_jump_to(mvtssr_map_t *m,
                                      mvtssr_camera_options_t *opt) {
  m->map->jumpTo(opt->opt);
}

MVTSSRAPICALL mvtssr_camera_options_t *mvtssr_map_camera_for_latlng_bounds(
    mvtssr_map_t *m, mvtssr_latlng_bounds_t *bounds, mvtssr_edge_insets_t *e,
    double *bearing, double *pitch) {
  return new mvtssr_camera_options_t{m->map->cameraForLatLngBounds(
      bounds->bounds, e->edge,
      bearing != nullptr ? *bearing : mbgl::optional<double>{},
      pitch != nullptr ? *pitch : mbgl::optional<double>{})};
}

MVTSSRAPICALL mvtssr_latlng_bounds_t *
mvtssr_map_camera_latlng_bounds_for_camera(mvtssr_map_t *m,
                                           mvtssr_camera_options_t *opt) {
  m->map->latLngBoundsForCamera(opt->opt);
}

MVTSSRAPICALL mvtssr_latlng_bounds_t *
mvtssr_map_camera_latlng_bounds_for_camera_unwrapped(
    mvtssr_map_t *m, mvtssr_camera_options_t *opt) {
  m->map->latLngBoundsForCameraUnwrapped(opt->opt);
}

MVTSSRAPICALL void mvtssr_map_set_bounds(mvtssr_map_t *m,
                                         mvtssr_bound_options_t *opts) {
  m->map->setBounds(opts->opt);
}

MVTSSRAPICALL mvtssr_bound_options_t *mvtssr_map_get_bounds(mvtssr_map_t *m) {
  return new mvtssr_bound_options_t{m->map->getBounds()};
}

MVTSSRAPICALL void mvtssr_map_set_north_orientation(mvtssr_map_t *m,
                                                    uint32_t ori) {
  m->map->setNorthOrientation(static_cast<mbgl::NorthOrientation>(ori));
}

MVTSSRAPICALL void mvtssr_map_set_constrain_mode(mvtssr_map_t *m,
                                                 uint32_t mode) {
  m->map->setConstrainMode(static_cast<mbgl::ConstrainMode>(mode));
}

MVTSSRAPICALL void mvtssr_map_set_viewport_mode(mvtssr_map_t *m,
                                                uint32_t mode) {
  m->map->setViewportMode(static_cast<mbgl::ViewportMode>(mode));
}

MVTSSRAPICALL void mvtssr_map_set_size(mvtssr_map_t *m, mvtssr_size_t *si) {
  m->map->setSize(si->si);
}

MVTSSRAPICALL mvtssr_map_options_t *
mvtssr_map_get_map_options(mvtssr_map_t *m) {
  return new mvtssr_map_options_t{m->map->getMapOptions()};
}

MVTSSRAPICALL mvtssr_screen_coordinate_t *
mvtssr_map_pixel_for_latlng(mvtssr_map_t *m, mvtssr_latlng_t *ll) {
  return new mvtssr_screen_coordinate_t{m->map->pixelForLatLng(ll->ll)};
}

MVTSSRAPICALL mvtssr_latlng_t *
mvtssr_map_latlng_for_pixel(mvtssr_map_t *m,
                            mvtssr_screen_coordinate_t *coord) {
  return new mvtssr_latlng_t{m->map->latLngForPixel(coord->sc)};
}

MVTSSRAPICALL
mvtssr_resource_t *mvtssr_new_resource_style(const char *url) {
  return new mvtssr_resource_t{mbgl::Resource::style(url)};
}

MVTSSRAPICALL
mvtssr_resource_t *mvtssr_new_resource_source(const char *url) {
  return new mvtssr_resource_t{mbgl::Resource::source(url)};
}

MVTSSRAPICALL
mvtssr_resource_t *mvtssr_new_resource_tile(const char *urltpl,
                                            float pixelRatio, int32_t x,
                                            int32_t y, int8_t z, _Bool isTms) {
  return new mvtssr_resource_t{mbgl::Resource::tile(
      urltpl, pixelRatio, x, y, z, static_cast<mbgl::Tileset::Scheme>(isTms))};
}

MVTSSRAPICALL
mvtssr_resource_t *mvtssr_new_resource_glyphs(const char *urltpl,
                                              const char *fontStack,
                                              uint16_t start, uint16_t end) {
  mbgl::FontStack st{fontStack};
  return new mvtssr_resource_t{
      mbgl::Resource::glyphs(urltpl, st, std::make_pair(start, end))};
}

MVTSSRAPICALL
mvtssr_resource_t *mvtssr_new_resource_sprite_image(const char *base,
                                                    float pixelRatio) {
  return new mvtssr_resource_t{mbgl::Resource::spriteImage(base, pixelRatio)};
}

MVTSSRAPICALL
mvtssr_resource_t *mvtssr_new_resource_sprite_json(const char *base,
                                                   float pixelRatio) {
  return new mvtssr_resource_t{mbgl::Resource::spriteJSON(base, pixelRatio)};
}

MVTSSRAPICALL
mvtssr_resource_t *mvtssr_new_resource_image(const char *url) {
  return new mvtssr_resource_t{mbgl::Resource::image(url)};
}

MVTSSRAPICALL void mvtssr_resource_free(mvtssr_resource_t *r) { delete r; }

MVTSSRAPICALL uint8_t mvtssr_resource_get_kind(mvtssr_resource_t *r) {
  return static_cast<uint8_t>(r->res.kind);
}

MVTSSRAPICALL
mvtssr_premultiplied_image_t *mvtssr_empty_premultiplied_image() {
  return new mvtssr_premultiplied_image_t{};
}

MVTSSRAPICALL
mvtssr_premultiplied_image_t *
mvtssr_new_premultiplied_image(mvtssr_size_t *size) {
  return new mvtssr_premultiplied_image_t{{size->si}};
}

MVTSSRAPICALL
mvtssr_premultiplied_image_t *
mvtssr_new_premultiplied_image_with_data(mvtssr_size_t *size,
                                         const uint8_t *data, size_t length) {
  return new mvtssr_premultiplied_image_t{{size->si, data, length}};
}

MVTSSRAPICALL void
mvtssr_premultiplied_image_free(mvtssr_premultiplied_image_t *r) {
  delete r;
}

MVTSSRAPICALL _Bool
mvtssr_premultiplied_image_valid(mvtssr_premultiplied_image_t *r) {
  return r->img.valid();
}

MVTSSRAPICALL size_t
mvtssr_premultiplied_image_stride(mvtssr_premultiplied_image_t *r) {
  return r->img.stride();
}

MVTSSRAPICALL size_t
mvtssr_premultiplied_image_bytes(mvtssr_premultiplied_image_t *r) {
  return r->img.bytes();
}

MVTSSRAPICALL uint8_t *
mvtssr_premultiplied_image_data(mvtssr_premultiplied_image_t *r) {
  return r->img.data.get();
}

MVTSSRAPICALL void
mvtssr_premultiplied_image_size(mvtssr_premultiplied_image_t *r,
                                uint32_t *width, uint32_t *height) {
  *width = r->img.size.width;
  *height = r->img.size.height;
}

#ifdef __cplusplus
}
#endif