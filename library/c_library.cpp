#include "c_library.h"
#include "library.h"
#include <vector>

int predict_c(const char *model_path, float *prediction) {
  std::vector<float> prediction_vec;
  auto ret = predict(model_path, &prediction_vec);
  if (ret != 0) {
    return ret;
  }
  if (prediction_vec.size() != 1000) {
    return -2;
  }
  for (int64_t i =0; i < 1000; i++) {
    prediction[i] = prediction_vec[i];
  }
  return 0;
}