#ifndef PYTORCH_CPP_GO_EXAMPLE_LIBRARY_H
#define PYTORCH_CPP_GO_EXAMPLE_LIBRARY_H
#include <string>
#include <vector>

int predict(std::string model_path, std::vector<float> *prediction);
#endif