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
typedef struct _mvtssr_style_t mvtssr_style_t;
typedef struct _mvtssr_bound_options_t mvtssr_bound_options_t;
typedef struct _mvtssr_resource_t mvtssr_resource_t;
typedef struct _mvtssr_file_source_factory_t mvtssr_file_source_factory_t;
typedef struct _mvtssr_resource_options_t mvtssr_resource_options_t;

typedef struct _mvtssr_headless_frontend_t mvtssr_headless_frontend_t;
typedef struct _mvtssr_map_snapshotter_observer_t
    mvtssr_map_snapshotter_observer_t;
typedef struct _mvtssr_map_snapshotter_t mvtssr_map_snapshotter_t;

typedef struct _mvtssr_offline_region_metadata_t
    mvtssr_offline_region_metadata_t;
typedef struct _mvtssr_offline_region_status_t mvtssr_offline_region_status_t;
typedef struct _mvtssr_offline_region_observer_t
    mvtssr_offline_region_observer_t;
typedef struct _mvtssr_offline_region_t mvtssr_offline_region_t;
typedef struct _mvtssr_offline_region_definition_t
    mvtssr_offline_region_definition_t;

typedef struct _mvtssr_premultiplied_image_t mvtssr_premultiplied_image_t;
typedef struct _mvtssr_unassociatedImage_image_t
    mvtssr_unassociatedImage_image_t;

MVTSSRAPICALL mvtssr_canonical_tileid_t *
new_mvtssr_canonical_tileid(uint8_t z, uint32_t x, uint32_t y);
MVTSSRAPICALL void mvtssr_canonical_tileid_free(mvtssr_canonical_tileid_t *id);
MVTSSRAPICALL mvtssr_canonical_tileid_t *
mvtssr_canonical_tileid_scaled_to(mvtssr_canonical_tileid_t *id, uint8_t z);
MVTSSRAPICALL _Bool mvtssr_canonical_tileid_is_child_of(
    mvtssr_canonical_tileid_t *id, mvtssr_canonical_tileid_t *tar);
MVTSSRAPICALL _Bool mvtssr_canonical_tileid_is_child_less(
    mvtssr_canonical_tileid_t *a, mvtssr_canonical_tileid_t *b);
MVTSSRAPICALL _Bool mvtssr_canonical_tileid_is_child_eq(
    mvtssr_canonical_tileid_t *a, mvtssr_canonical_tileid_t *b);
MVTSSRAPICALL _Bool mvtssr_canonical_tileid_children(
    mvtssr_canonical_tileid_t *id, mvtssr_canonical_tileid_t **childs);

MVTSSRAPICALL mvtssr_latlng_t *new_mvtssr_latlng(double lat, double lon);
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

MVTSSRAPICALL mvtssr_edge_insets_t *new_mvtssr_edge_insets(double t, double l,
                                                           double b, double r);
MVTSSRAPICALL double mvtssr_edge_top(mvtssr_edge_insets_t *edge);
MVTSSRAPICALL double mvtssr_edge_left(mvtssr_edge_insets_t *edge);
MVTSSRAPICALL double mvtssr_edge_bottom(mvtssr_edge_insets_t *edge);
MVTSSRAPICALL double mvtssr_edge_right(mvtssr_edge_insets_t *edge);
MVTSSRAPICALL double mvtssr_edge_is_flush(mvtssr_edge_insets_t *edge);
MVTSSRAPICALL mvtssr_screen_coordinate_t *
mvtssr_edge_get_center(mvtssr_edge_insets_t *edge, uint16_t width,
                       uint16_t height);
MVTSSRAPICALL _Bool mvtssr_edge_eq(mvtssr_edge_insets_t *a,
                                   mvtssr_edge_insets_t *b);
MVTSSRAPICALL void mvtssr_edge_add(mvtssr_edge_insets_t *a,
                                   mvtssr_edge_insets_t *b);

MVTSSRAPICALL mvtssr_screen_coordinate_t *new_screen_coordinate(double x,
                                                                double y);
MVTSSRAPICALL void
mvtssr_screen_coordinate_free(mvtssr_screen_coordinate_t *sc);
MVTSSRAPICALL double mvtssr_screen_coordinate_x(mvtssr_screen_coordinate_t *sc);
MVTSSRAPICALL double mvtssr_screen_coordinate_y(mvtssr_screen_coordinate_t *sc);

MVTSSRAPICALL mvtssr_camera_options_t *new_mvtssr_camera_options();
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
MVTSSRAPICALL mvtssr_latlng_t *
mvtssr_camera_options_get_center(mvtssr_camera_options_t *opt);
MVTSSRAPICALL mvtssr_edge_insets_t *
mvtssr_camera_options_get_padding(mvtssr_camera_options_t *opt);
MVTSSRAPICALL mvtssr_screen_coordinate_t *
mvtssr_camera_options_get_anchor(mvtssr_camera_options_t *opt);
MVTSSRAPICALL double
mvtssr_camera_options_get_zoom(mvtssr_camera_options_t *opt);
MVTSSRAPICALL double
mvtssr_camera_options_get_bearing(mvtssr_camera_options_t *opt);
MVTSSRAPICALL double
mvtssr_camera_options_get_pitch(mvtssr_camera_options_t *opt);

#ifdef __cplusplus
}
#endif

#endif // MVTSSR_C_API_H