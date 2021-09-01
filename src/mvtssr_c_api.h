#ifndef PRONTO_C_API_H
#define PRONTO_C_API_H

#if defined(WIN32) || defined(WINDOWS) || defined(_WIN32) || defined(_WINDOWS)
#define MVTSSRAPICALL __declspec(dllexport)
#else
#define MVTSSRAPICALL
#endif

#include <stdbool.h>

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
}
#endif

#endif // PRONTO_C_API_H