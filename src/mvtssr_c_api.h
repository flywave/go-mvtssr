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

typedef struct _mvtssr_latlng_t mvtssr_latlng_t;
typedef struct _mvtssr_latlng_bounds_t mvtssr_latlng_bounds_t;
typedef struct _mvtssr_latlng_altitude_t mvtssr_latlng_altitude_t;
typedef struct _mvtssr_edge_insets_t mvtssr_edge_insets_t;
typedef struct _mvtssr_screen_coordinate_t mvtssr_screen_coordinate_t;
typedef struct _mvtssr_camera_options_t mvtssr_camera_options_t;
typedef struct _mvtssr_thread_pool_t mvtssr_thread_pool_t;
typedef struct _mvtssr_runloop_t mvtssr_runloop_t;
typedef struct _mvtssr_size_t mvtssr_size_t;
typedef struct _mvtssr_scheduler_t mvtssr_scheduler_t;
typedef struct _mvtssr_headless_frontend_t mvtssr_headless_frontend_t;
typedef struct _mvtssr_map_observer_t mvtssr_map_observer_t;
typedef struct _mvtssr_map_options_t mvtssr_map_options_t;
typedef struct _mvtssr_file_source_t mvtssr_file_source_t;
typedef struct _mvtssr_file_source_manager_t mvtssr_file_source_manager_t;
typedef struct _mvtssr_style_t mvtssr_style_t;
typedef struct _mvtssr_bound_options_t mvtssr_bound_options_t;
typedef struct _mvtssr_resource_t mvtssr_resource_t;

typedef struct _mvtssr_map_snapshotter_observer_t
    mvtssr_map_snapshotter_observer_t;
typedef struct _mvtssr_map_snapshotter_t mvtssr_map_snapshotter_t;

typedef struct _mvtssr_offline_region_metadata_t
    mvtssr_offline_region_metadata_t;
typedef struct _mvtssr_offline_region_status_t mvtssr_offline_region_status_t;
typedef struct _mvtssr_offline_region_observer_t
    mvtssr_offline_region_observer_t;
typedef struct _mvtssr_offline_region_t mvtssr_offline_region_t;

typedef struct _mvtssr_premultiplied_image_t mvtssr_premultiplied_image_t;

#ifdef __cplusplus
}
#endif

#endif // MVTSSR_C_API_H