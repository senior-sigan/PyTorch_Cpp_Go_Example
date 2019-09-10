import torch
import torchvision


def main():
    model = torchvision.models.resnet18()

    example = torch.rand(1, 3, 224, 224)

    traced_script_module = torch.jit.trace(model, example)

    traced_script_module.save("traced_resnet_model.pt")


if __name__ == "__main__":
    main()
