#include "mvtssr_c_api.h"

#ifdef __cplusplus
extern "C" {
#endif

typedef struct _mvtssr_raw_image_t {
  size_t height;
  size_t width;
  uint8_t *data;
} mvtssr_raw_image_t;

typedef struct _mvtssr_snapshot_params_t {
  char *style;
  char *cache_file;
  char *asset_root;
  uint32_t width;
  uint32_t height;
  double ppi_ratio;
  double lat;
  double lng;
  double zoom;
  double pitch;
  double bearing;
} mvtssr_snapshot_params_t;

typedef struct _mvtssr_snapshot_t {
  mvtssr_raw_image_t *image;
  int did_error;
  const char *err;
} mvtssr_snapshot_t;

#ifdef __cplusplus
}
#endif