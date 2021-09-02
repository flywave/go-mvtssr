#ifndef MVTSSR_C_API_H
#define MVTSSR_C_API_H

#if defined(WIN32) || defined(WINDOWS) || defined(_WIN32) || defined(_WINDOWS)
#define MVTSSRAPICALL __declspec(dllexport)
#else
#define MVTSSRAPICALL
#endif

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

typedef struct _mvtssr_canonical_tileid_t mvtssr_canonical_tileid_t;
typedef struct _mvtssr_latlng_t mvtssr_latlng_t;
typedef struct _mvtssr_latlng_bounds_t mvtssr_latlng_bounds_t;
typedef struct _mvtssr_edge_insets_t mvtssr_edge_insets_t;
typedef struct _mvtssr_screen_coordinate_t mvtssr_screen_coordinate_t;
typedef struct _mvtssr_camera_options_t mvtssr_camera_options_t;
typedef struct _mvtssr_runloop_t mvtssr_runloop_t;
typedef struct _mvtssr_size_t mvtssr_size_t;
typedef struct _mvtssr_map_observer_t mvtssr_map_observer_t;
typedef struct _mvtssr_map_options_t mvtssr_map_options_t;
typedef struct _mvtssr_file_source_t mvtssr_file_source_t;
typedef struct _mvtssr_file_source_manager_t mvtssr_file_source_manager_t;
typedef struct _mvtssr_map_t mvtssr_map_t;
typedef struct _mvtssr_style_t mvtssr_style_t;
typedef struct _mvtssr_bound_options_t mvtssr_bound_options_t;
typedef struct _mvtssr_resource_t mvtssr_resource_t;
typedef struct _mvtssr_file_source_factory_t mvtssr_file_source_factory_t;
typedef struct _mvtssr_resource_options_t mvtssr_resource_options_t;
typedef struct _mvtssr_headless_frontend_t mvtssr_headless_frontend_t;
typedef struct _mvtssr_map_snapshotter_observer_t
    mvtssr_map_snapshotter_observer_t;
typedef struct _mvtssr_map_snapshotter_t mvtssr_map_snapshotter_t;
typedef struct _mvtssr_map_snapshotter_result_t mvtssr_map_snapshotter_result_t;
typedef struct _mvtssr_premultiplied_image_t mvtssr_premultiplied_image_t;

MVTSSRAPICALL mvtssr_canonical_tileid_t *
new_mvtssr_canonical_tileid(uint8_t z, uint32_t x, uint32_t y);
MVTSSRAPICALL void mvtssr_canonical_tileid_free(mvtssr_canonical_tileid_t *id);
MVTSSRAPICALL mvtssr_canonical_tileid_t *
mvtssr_canonical_tileid_scaled_to(mvtssr_canonical_tileid_t *id, uint8_t z);
MVTSSRAPICALL _Bool mvtssr_canonical_tileid_is_child_of(
    mvtssr_canonical_tileid_t *id, mvtssr_canonical_tileid_t *tar);
MVTSSRAPICALL _Bool mvtssr_canonical_tileid_is_child_less(
    mvtssr_canonical_tileid_t *a, mvtssr_canonical_tileid_t *b);
MVTSSRAPICALL _Bool mvtssr_canonical_tileid_eq(mvtssr_canonical_tileid_t *a,
                                               mvtssr_canonical_tileid_t *b);
MVTSSRAPICALL _Bool mvtssr_canonical_tileid_children(
    mvtssr_canonical_tileid_t *id, mvtssr_canonical_tileid_t **childs);

MVTSSRAPICALL mvtssr_latlng_t *mvtssr_new_latlng(double lat, double lon);
MVTSSRAPICALL void mvtssr_mvtssr_latlng_free(mvtssr_latlng_t *ll);
MVTSSRAPICALL mvtssr_latlng_t *
new_mvtssr_latlng_with_id(mvtssr_canonical_tileid_t *id);
MVTSSRAPICALL double mvtssr_latlng_latitude(mvtssr_latlng_t *latlon);
MVTSSRAPICALL double mvtssr_latlng_longitude(mvtssr_latlng_t *latlon);

MVTSSRAPICALL mvtssr_latlng_bounds_t *latlng_bounds_world();
MVTSSRAPICALL mvtssr_latlng_bounds_t *
new_latlng_bounds(mvtssr_latlng_t *latlon);
MVTSSRAPICALL mvtssr_latlng_bounds_t *hull_latlng_bounds(mvtssr_latlng_t *a,
                                                         mvtssr_latlng_t *b);
MVTSSRAPICALL mvtssr_latlng_bounds_t *latlng_bounds_empty();
MVTSSRAPICALL mvtssr_latlng_bounds_t *
new_latlng_bounds_with_id(mvtssr_canonical_tileid_t *id);
MVTSSRAPICALL void mvtssr_latlng_bounds_free(mvtssr_latlng_bounds_t *llb);
MVTSSRAPICALL _Bool latlng_bounds_valid(mvtssr_latlng_bounds_t *bounds);
MVTSSRAPICALL double latlng_bounds_south(mvtssr_latlng_bounds_t *bounds);
MVTSSRAPICALL double latlng_bounds_west(mvtssr_latlng_bounds_t *bounds);
MVTSSRAPICALL double latlng_bounds_north(mvtssr_latlng_bounds_t *bounds);
MVTSSRAPICALL double latlng_bounds_east(mvtssr_latlng_bounds_t *bounds);
MVTSSRAPICALL mvtssr_latlng_t *
latlng_bounds_southwest(mvtssr_latlng_bounds_t *bounds);
MVTSSRAPICALL mvtssr_latlng_t *
latlng_bounds_northeast(mvtssr_latlng_bounds_t *bounds);
MVTSSRAPICALL mvtssr_latlng_t *
latlng_bounds_southeast(mvtssr_latlng_bounds_t *bounds);
MVTSSRAPICALL mvtssr_latlng_t *
latlng_bounds_northwest(mvtssr_latlng_bounds_t *bounds);
MVTSSRAPICALL mvtssr_latlng_t *
latlng_bounds_center(mvtssr_latlng_bounds_t *bounds);
MVTSSRAPICALL mvtssr_latlng_t *
latlng_bounds_constrain(mvtssr_latlng_bounds_t *bounds, mvtssr_latlng_t *p);
MVTSSRAPICALL void latlng_bounds_extend(mvtssr_latlng_bounds_t *bounds,
                                        mvtssr_latlng_t *p);
MVTSSRAPICALL void
latlng_bounds_extend_bounds(mvtssr_latlng_bounds_t *bounds,
                            mvtssr_latlng_bounds_t *ext_bounds);
MVTSSRAPICALL _Bool latlng_bounds_is_empty(mvtssr_latlng_bounds_t *bounds);
MVTSSRAPICALL _Bool latlng_bounds_contains_id(mvtssr_latlng_bounds_t *bounds,
                                              mvtssr_canonical_tileid_t *id);
MVTSSRAPICALL _Bool latlng_bounds_contains_point(mvtssr_latlng_bounds_t *bounds,
                                                 mvtssr_latlng_t *point);
MVTSSRAPICALL _Bool latlng_bounds_contains_bounds(
    mvtssr_latlng_bounds_t *bounds, mvtssr_latlng_bounds_t *area);
MVTSSRAPICALL _Bool latlng_bounds_contains_intersects(
    mvtssr_latlng_bounds_t *bounds, mvtssr_latlng_bounds_t *area);

MVTSSRAPICALL mvtssr_edge_insets_t *mvtssr_new_edge_insets(double t, double l,
                                                           double b, double r);
MVTSSRAPICALL void mvtssr_edge_free(mvtssr_edge_insets_t *edge);
MVTSSRAPICALL double mvtssr_edge_top(mvtssr_edge_insets_t *edge);
MVTSSRAPICALL double mvtssr_edge_left(mvtssr_edge_insets_t *edge);
MVTSSRAPICALL double mvtssr_edge_bottom(mvtssr_edge_insets_t *edge);
MVTSSRAPICALL double mvtssr_edge_right(mvtssr_edge_insets_t *edge);
MVTSSRAPICALL _Bool mvtssr_edge_is_flush(mvtssr_edge_insets_t *edge);
MVTSSRAPICALL mvtssr_screen_coordinate_t *
mvtssr_edge_get_center(mvtssr_edge_insets_t *edge, uint16_t width,
                       uint16_t height);
MVTSSRAPICALL _Bool mvtssr_edge_eq(mvtssr_edge_insets_t *a,
                                   mvtssr_edge_insets_t *b);
MVTSSRAPICALL void mvtssr_edge_add(mvtssr_edge_insets_t *a,
                                   mvtssr_edge_insets_t *b);

MVTSSRAPICALL mvtssr_screen_coordinate_t *
mvtssr_new_screen_coordinate(double x, double y);
MVTSSRAPICALL void
mvtssr_screen_coordinate_free(mvtssr_screen_coordinate_t *sc);
MVTSSRAPICALL double mvtssr_screen_coordinate_x(mvtssr_screen_coordinate_t *sc);
MVTSSRAPICALL double mvtssr_screen_coordinate_y(mvtssr_screen_coordinate_t *sc);

MVTSSRAPICALL mvtssr_camera_options_t *mvtssr_new_camera_options();
MVTSSRAPICALL void mvtssr_camera_options_free(mvtssr_camera_options_t *opt);
MVTSSRAPICALL void
mvtssr_camera_options_set_center(mvtssr_camera_options_t *opt,
                                 mvtssr_latlng_t *point);
MVTSSRAPICALL void
mvtssr_camera_options_set_padding(mvtssr_camera_options_t *opt,
                                  mvtssr_edge_insets_t *edge);
MVTSSRAPICALL void
mvtssr_camera_options_set_anchor(mvtssr_camera_options_t *opt,
                                 mvtssr_screen_coordinate_t *sc);
MVTSSRAPICALL void mvtssr_camera_options_set_zoom(mvtssr_camera_options_t *opt,
                                                  double zoom);
MVTSSRAPICALL void
mvtssr_camera_options_set_bearing(mvtssr_camera_options_t *opt, double bearing);
MVTSSRAPICALL void mvtssr_camera_options_set_pitch(mvtssr_camera_options_t *opt,
                                                   double pitch);

MVTSSRAPICALL
mvtssr_size_t *mvtssr_new_size(uint32_t width, uint32_t height);
MVTSSRAPICALL void mvtssr_size_free(mvtssr_size_t *si);
MVTSSRAPICALL uint32_t mvtssr_size_area(mvtssr_size_t *si);
MVTSSRAPICALL float mvtssr_size_aspect_ratio(mvtssr_size_t *si);
MVTSSRAPICALL _Bool mvtssr_size_is_empty(mvtssr_size_t *si);

MVTSSRAPICALL
mvtssr_runloop_t *mvtssr_new_runloop();
MVTSSRAPICALL void mvtssr_runloop_free(mvtssr_runloop_t *loop);

MVTSSRAPICALL
mvtssr_map_observer_t *mvtssr_null_map_observer();
MVTSSRAPICALL
mvtssr_map_observer_t *mvtssr_new_map_observer(void *ctx);
MVTSSRAPICALL void mvtssr_map_observer_free(mvtssr_map_observer_t *ob);

MVTSSRAPICALL mvtssr_map_options_t *mvtssr_new_map_options();
MVTSSRAPICALL void mvtssr_map_options_free(mvtssr_map_options_t *op);
MVTSSRAPICALL void mvtssr_map_options_set_map_mode(mvtssr_map_options_t *opt,
                                                   uint32_t mode);
MVTSSRAPICALL void
mvtssr_map_options_set_constrain_mode(mvtssr_map_options_t *opt, uint32_t mode);
MVTSSRAPICALL void
mvtssr_map_options_set_viewport_mode(mvtssr_map_options_t *opt, uint32_t mode);
MVTSSRAPICALL void
mvtssr_map_options_set_cross_source_collisions(mvtssr_map_options_t *opt,
                                               _Bool sc);
MVTSSRAPICALL void
mvtssr_map_options_set_north_orientation(mvtssr_map_options_t *opt,
                                         uint8_t ori);
MVTSSRAPICALL void mvtssr_map_options_set_size(mvtssr_map_options_t *opt,
                                               mvtssr_size_t *si);
MVTSSRAPICALL void mvtssr_map_options_set_pixel_ratio(mvtssr_map_options_t *opt,
                                                      float ratio);

MVTSSRAPICALL
mvtssr_file_source_t *mvtssr_new_file_source(void *ctx);
MVTSSRAPICALL void mvtssr_file_source_free(mvtssr_file_source_t *s);

MVTSSRAPICALL
mvtssr_file_source_manager_t *mvtssr_get_file_source_manager();
MVTSSRAPICALL void
mvtssr_file_source_manager_free(mvtssr_file_source_manager_t *s);
MVTSSRAPICALL void mvtssr_file_source_manager_register_file_source_factory(
    mvtssr_file_source_manager_t *s, mvtssr_file_source_factory_t *factory);
MVTSSRAPICALL void mvtssr_file_source_manager_unregister_file_source_factory(
    mvtssr_file_source_manager_t *s, mvtssr_file_source_factory_t *factory);

MVTSSRAPICALL
mvtssr_file_source_factory_t *mvtssr_new_file_source_factory(uint8_t file_type,
                                                             void *ctx);
MVTSSRAPICALL void
mvtssr_file_source_factory_free(mvtssr_file_source_factory_t *s);

MVTSSRAPICALL mvtssr_resource_options_t *mvtssr_new_resource_options();
MVTSSRAPICALL void mvtssr_resource_options_free(mvtssr_resource_options_t *op);
MVTSSRAPICALL void
mvtssr_resource_options_set_access_token(mvtssr_resource_options_t *opt,
                                         const char *token);
MVTSSRAPICALL void
mvtssr_resource_options_set_base_url(mvtssr_resource_options_t *opt,
                                     const char *url);
MVTSSRAPICALL void
mvtssr_resource_options_set_asset_path(mvtssr_resource_options_t *opt,
                                       const char *path);
MVTSSRAPICALL void
mvtssr_resource_options_set_maximum_cache_size(mvtssr_resource_options_t *opt,
                                               uint64_t size);

MVTSSRAPICALL mvtssr_headless_frontend_t *
mvtssr_new_headless_frontend(mvtssr_size_t *size, float pixelRatio);
MVTSSRAPICALL void
mvtssr_headless_frontend_free(mvtssr_headless_frontend_t *op);
MVTSSRAPICALL void
mvtssr_headless_frontend_reset(mvtssr_headless_frontend_t *op);
MVTSSRAPICALL mvtssr_screen_coordinate_t *
mvtssr_headless_frontend_pixel_for_latlng(mvtssr_headless_frontend_t *op,
                                          mvtssr_latlng_t *latlon);
MVTSSRAPICALL mvtssr_latlng_t *
mvtssr_headless_frontend_latlng_for_pixel(mvtssr_headless_frontend_t *op,
                                          mvtssr_screen_coordinate_t *coord);
MVTSSRAPICALL void
mvtssr_headless_frontend_set_size(mvtssr_headless_frontend_t *op,
                                  mvtssr_size_t *size);
MVTSSRAPICALL mvtssr_size_t *
mvtssr_headless_frontend_get_size(mvtssr_headless_frontend_t *op);
MVTSSRAPICALL mvtssr_premultiplied_image_t *
mvtssr_headless_frontend_render(mvtssr_headless_frontend_t *op,
                                mvtssr_map_t *map);

MVTSSRAPICALL
mvtssr_map_snapshotter_observer_t *mvtssr_null_map_snapshotter_observer();
MVTSSRAPICALL
mvtssr_map_snapshotter_observer_t *
mvtssr_new_map_snapshotter_observer(void *ctx);
MVTSSRAPICALL void
mvtssr_map_snapshotter_observer_free(mvtssr_map_snapshotter_observer_t *op);

MVTSSRAPICALL mvtssr_map_snapshotter_t *
mvtssr_new_map_snapshotter(mvtssr_size_t *size, float pixelRatio,
                           mvtssr_resource_options_t *opts,
                           mvtssr_map_snapshotter_observer_t *obser);
MVTSSRAPICALL void mvtssr_map_snapshotter_free(mvtssr_map_snapshotter_t *snap);
MVTSSRAPICALL void
mvtssr_map_snapshotter_set_style_url(mvtssr_map_snapshotter_t *snap,
                                     const char *url);
MVTSSRAPICALL char *
mvtssr_map_snapshotter_get_style_url(mvtssr_map_snapshotter_t *snap);

MVTSSRAPICALL void
mvtssr_map_snapshotter_set_style(mvtssr_map_snapshotter_t *snap,
                                 const char *style);
MVTSSRAPICALL char *
mvtssr_map_snapshotter_get_style(mvtssr_map_snapshotter_t *snap);
MVTSSRAPICALL void
mvtssr_map_snapshotter_set_size(mvtssr_map_snapshotter_t *snap,
                                mvtssr_size_t *size);
MVTSSRAPICALL mvtssr_size_t *
mvtssr_map_snapshotter_get_size(mvtssr_map_snapshotter_t *snap);

MVTSSRAPICALL void
mvtssr_map_snapshotter_set_camera_options(mvtssr_map_snapshotter_t *snap,
                                          mvtssr_camera_options_t *opts);

MVTSSRAPICALL mvtssr_camera_options_t *
mvtssr_map_snapshotter_get_camera_options(mvtssr_map_snapshotter_t *snap);
MVTSSRAPICALL void
mvtssr_map_snapshotter_set_region(mvtssr_map_snapshotter_t *snap,
                                  mvtssr_latlng_bounds_t *bounds);
MVTSSRAPICALL mvtssr_latlng_bounds_t *
mvtssr_map_snapshotter_get_region(mvtssr_map_snapshotter_t *snap);
MVTSSRAPICALL void
mvtssr_map_snapshotter_cancel(mvtssr_map_snapshotter_t *snap);
MVTSSRAPICALL void mvtssr_map_snapshotter_snapshot(
    mvtssr_map_snapshotter_t *snap, mvtssr_map_snapshotter_result_t *result);

MVTSSRAPICALL
mvtssr_map_snapshotter_result_t *mvtssr_new_map_snapshotter_result(void *ctx);
MVTSSRAPICALL void
mvtssr_map_snapshotter_result_free(mvtssr_map_snapshotter_result_t *op);
MVTSSRAPICALL mvtssr_premultiplied_image_t *
mvtssr_map_snapshotter_result_get_image(mvtssr_map_snapshotter_result_t *op);
MVTSSRAPICALL char *
mvtssr_map_snapshotter_result_get_error(mvtssr_map_snapshotter_result_t *op);
MVTSSRAPICALL mvtssr_screen_coordinate_t *
mvtssr_map_snapshotter_result_pixel_for_latlng(
    mvtssr_map_snapshotter_result_t *op, mvtssr_latlng_t *latlon);
MVTSSRAPICALL mvtssr_latlng_t *mvtssr_map_snapshotter_result_latlng_for_pixel(
    mvtssr_map_snapshotter_result_t *op, mvtssr_screen_coordinate_t *coord);

MVTSSRAPICALL
mvtssr_style_t *mvtssr_new_style(mvtssr_file_source_t *source,
                                 float pixelRatio);
MVTSSRAPICALL void mvtssr_style_free(mvtssr_style_t *m);
MVTSSRAPICALL char *mvtssr_style_get_json(mvtssr_style_t *m);
MVTSSRAPICALL char *mvtssr_style_get_url(mvtssr_style_t *m);

MVTSSRAPICALL
mvtssr_bound_options_t *mvtssr_new_bound_options();
MVTSSRAPICALL void mvtssr_bound_options_free(mvtssr_bound_options_t *m);
MVTSSRAPICALL void
mvtssr_bound_options_set_bounds(mvtssr_bound_options_t *opt,
                                mvtssr_latlng_bounds_t *bounds);
MVTSSRAPICALL void mvtssr_bound_options_set_min_zoom(mvtssr_bound_options_t *m,
                                                     double z);
MVTSSRAPICALL void mvtssr_bound_options_set_max_zoom(mvtssr_bound_options_t *m,
                                                     double z);
MVTSSRAPICALL void mvtssr_bound_options_set_min_pitch(mvtssr_bound_options_t *m,
                                                      double z);
MVTSSRAPICALL void mvtssr_bound_options_set_max_pitch(mvtssr_bound_options_t *m,
                                                      double z);

MVTSSRAPICALL
mvtssr_map_t *mvtssr_new_map(mvtssr_headless_frontend_t *fr,
                             mvtssr_map_observer_t *obser,
                             mvtssr_map_options_t *opts,
                             mvtssr_resource_options_t *ropts);
MVTSSRAPICALL void mvtssr_map_free(mvtssr_map_t *m);
MVTSSRAPICALL void mvtssr_map_set_style(mvtssr_map_t *m, mvtssr_style_t *s);
MVTSSRAPICALL mvtssr_camera_options_t *
mvtssr_map_camera_options(mvtssr_map_t *m, mvtssr_edge_insets_t *e);
MVTSSRAPICALL void mvtssr_map_jump_to(mvtssr_map_t *m,
                                      mvtssr_camera_options_t *opt);
MVTSSRAPICALL mvtssr_camera_options_t *mvtssr_map_camera_for_latlng_bounds(
    mvtssr_map_t *m, mvtssr_latlng_bounds_t *bounds, mvtssr_edge_insets_t *e,
    double *bearing, double *pitch);
MVTSSRAPICALL mvtssr_latlng_bounds_t *
mvtssr_map_camera_latlng_bounds_for_camera(mvtssr_map_t *m,
                                           mvtssr_camera_options_t *opt);
MVTSSRAPICALL mvtssr_latlng_bounds_t *
mvtssr_map_camera_latlng_bounds_for_camera_unwrapped(
    mvtssr_map_t *m, mvtssr_camera_options_t *opt);
MVTSSRAPICALL void mvtssr_map_set_bounds(mvtssr_map_t *m,
                                         mvtssr_bound_options_t *opts);
MVTSSRAPICALL mvtssr_bound_options_t *mvtssr_map_get_bounds(mvtssr_map_t *m);
MVTSSRAPICALL void mvtssr_map_set_north_orientation(mvtssr_map_t *m,
                                                    uint32_t ori);
MVTSSRAPICALL void mvtssr_map_set_constrain_mode(mvtssr_map_t *m,
                                                 uint32_t mode);
MVTSSRAPICALL void mvtssr_map_set_viewport_mode(mvtssr_map_t *m, uint32_t mode);
MVTSSRAPICALL void mvtssr_map_set_size(mvtssr_map_t *m, mvtssr_size_t *si);
MVTSSRAPICALL mvtssr_map_options_t *mvtssr_map_get_map_options(mvtssr_map_t *m);
MVTSSRAPICALL mvtssr_screen_coordinate_t *
mvtssr_map_pixel_for_latlng(mvtssr_map_t *m, mvtssr_latlng_t *ll);
MVTSSRAPICALL mvtssr_latlng_t *
mvtssr_map_latlng_for_pixel(mvtssr_map_t *m, mvtssr_screen_coordinate_t *coord);

MVTSSRAPICALL
mvtssr_resource_t *mvtssr_new_resource_style(const char *url);
MVTSSRAPICALL
mvtssr_resource_t *mvtssr_new_resource_source(const char *url);
MVTSSRAPICALL
mvtssr_resource_t *mvtssr_new_resource_tile(const char *urltpl,
                                            float pixelRatio, int32_t x,
                                            int32_t y, int8_t z, _Bool isTms);
MVTSSRAPICALL
mvtssr_resource_t *mvtssr_new_resource_glyphs(const char *urltpl,
                                              const char *fontStack,
                                              uint16_t start, uint16_t end);
MVTSSRAPICALL
mvtssr_resource_t *mvtssr_new_resource_sprite_image(const char *base,
                                                    float pixelRatio);
MVTSSRAPICALL
mvtssr_resource_t *mvtssr_new_resource_sprite_json(const char *base,
                                                   float pixelRatio);
MVTSSRAPICALL
mvtssr_resource_t *mvtssr_new_resource_image(const char *url);
MVTSSRAPICALL void mvtssr_resource_free(mvtssr_resource_t *r);
MVTSSRAPICALL uint8_t mvtssr_resource_get_kind(mvtssr_resource_t *r);

MVTSSRAPICALL
mvtssr_premultiplied_image_t *mvtssr_empty_premultiplied_image();
MVTSSRAPICALL
mvtssr_premultiplied_image_t *
mvtssr_new_premultiplied_image(mvtssr_size_t *size);
MVTSSRAPICALL
mvtssr_premultiplied_image_t *
mvtssr_new_premultiplied_image_with_data(mvtssr_size_t *size,
                                         const uint8_t *data, size_t length);
MVTSSRAPICALL void
mvtssr_premultiplied_image_free(mvtssr_premultiplied_image_t *r);
MVTSSRAPICALL _Bool
mvtssr_premultiplied_image_valid(mvtssr_premultiplied_image_t *r);
MVTSSRAPICALL size_t
mvtssr_premultiplied_image_stride(mvtssr_premultiplied_image_t *r);
MVTSSRAPICALL size_t
mvtssr_premultiplied_image_bytes(mvtssr_premultiplied_image_t *r);
MVTSSRAPICALL uint8_t *
mvtssr_premultiplied_image_data(mvtssr_premultiplied_image_t *r);
MVTSSRAPICALL void
mvtssr_premultiplied_image_size(mvtssr_premultiplied_image_t *r,
                                uint32_t *width, uint32_t *height);

#ifdef __cplusplus
}
#endif

#endif // MVTSSR_C_API_H