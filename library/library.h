#ifndef PYTORCH_CPP_GO_EXAMPLE_LIBRARY_H
#define PYTORCH_CPP_GO_EXAMPLE_LIBRARY_H

#include "c_api.h"

PYTORCH_EXAMPLE_C_API
int predict(const char* model_path, float* prediction);

#endif
