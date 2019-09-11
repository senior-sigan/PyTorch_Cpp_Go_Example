package main

// #cgo LDFLAGS: -L${SRCDIR}/library/libs/osx/libtorch/lib -L${SRCDIR}/library/cmake-build-debug -lstdc++ -lpytorch_cpp_go_example -ltorch -lc10
// #include <stdlib.h>
// int predict(const char *model_path, float *prediction);
import "C"
import (
	"fmt"
	"unsafe"
)

func predict(modelPath string) {
	cModelPath := C.CString(modelPath)
	defer C.free(unsafe.Pointer(cModelPath))
	prediction := make([]float32, 1000)
	cPrediction := (*C.float)(unsafe.Pointer(&prediction[0]))

	C.predict(cModelPath, cPrediction)
	for i := 0;i < 1000; i++ {
		fmt.Printf("%f ", prediction[i])
	}
}

func main() {
	fmt.Println("Hello world")
	modelPath := "model/traced_resnet_model.pt"
	predict(modelPath)
}
