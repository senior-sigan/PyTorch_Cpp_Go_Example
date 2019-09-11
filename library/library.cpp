#include "library.h"
#include <torch/script.h>

#include <iostream>
#include <string>
#include <vector>

int predict(const char* model_path, float* prediction) {
  torch::jit::script::Module module;
  try {
    module = torch::jit::load(model_path);
    std::cout << "Model loaded" << std::endl;
  } catch (const c10::Error& e) {
    std::cerr << "error loading the model" << std::endl;
    return -1;
  }

  std::vector<torch::jit::IValue> inputs;
  inputs.emplace_back(torch::ones({1, 3, 224, 224}));

  at::Tensor output = module.forward(inputs).toTensor();

  auto a = output.accessor<float, 2>();
  if (a.size(1) != 1000) {
    return -2;
  }
  if (a.size(0) != 1) {
    return -3;
  }
  for (int64_t i = 0; i < a.size(0); i++) {
    for (int64_t j = 0; j < a.size(1); j++) {
      prediction[j] = a[i][j];
    }
  }

  return 0;
}
