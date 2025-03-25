package main

import (
    "fmt"
    "gorgonia.org/gorgonia"
    "gorgonia.org/tensor"
)

func main() {
    g := gorgonia.NewGraph()

    // Define the shapes
    inputShape := tensor.Shape{batchSize, inputSize}
    latentShape := tensor.Shape{batchSize, latentSize}
    outputShape := tensor.Shape{batchSize, outputSize}

    // Define input nodes
    x := gorgonia.NewTensor(g, tensor.Float32, 2, gorgonia.WithShape(inputShape...), gorgonia.WithName("x"))

    // Define weights and biases for encoder
    wEnc := gorgonia.NewTensor(g, tensor.Float32, 2, gorgonia.WithShape(inputSize, latentSize), gorgonia.WithName("wEnc"), gorgonia.WithInit(gorgonia.GlorotU(1)))
    bEnc := gorgonia.NewTensor(g, tensor.Float32, 1, gorgonia.WithShape(latentSize), gorgonia.WithName("bEnc"), gorgonia.WithInit(gorgonia.Zeroes()))

    // Define encoder operations
    enc := gorgonia.Must(gorgonia.Add(gorgonia.Must(gorgonia.Mul(x, wEnc)), bEnc))
    latent := gorgonia.Must(gorgonia.Rectify(enc))

    // Define weights and biases for decoder
    wDec := gorgonia.NewTensor(g, tensor.Float32, 2, gorgonia.WithShape(latentSize, outputSize), gorgonia.WithName("wDec"), gorgonia.WithInit(gorgonia.GlorotU(1)))
    bDec := gorgonia.NewTensor(g, tensor.Float32, 1, gorgonia.WithShape(outputSize), gorgonia.WithName("bDec"), gorgonia.WithInit(gorgonia.Zeroes()))

    // Define decoder operations
    dec := gorgonia.Must(gorgonia.Add(gorgonia.Must(gorgonia.Mul(latent, wDec)), bDec))
    output := gorgonia.Must(gorgonia.Rectify(dec))

    // Define the loss function
    loss := gorgonia.Must(gorgonia.Mean(gorgonia.Must(gorgonia.Square(gorgonia.Must(gorgonia.Sub(output, x))))))

    // Define the gradient
    grads, err := gorgonia.Gradient(loss, wEnc, bEnc, wDec, bDec)
    if err != nil {
        fmt.Println("Error calculating gradient:", err)
        return
    }

    // Create a VM to run the graph
    vm := gorgonia.NewTapeMachine(g)

    // Training loop
    for i := 0; i < numEpochs; i++ {
        if err := vm.RunAll(); err != nil {
            fmt.Println("Error during training:", err)
            return
        }

        // Update weights using gradients
        gorgonia.WithLearnRate(learningRate).Apply(wEnc, grads[wEnc])
        gorgonia.WithLearnRate(learningRate).Apply(bEnc, grads[bEnc])
        gorgonia.WithLearnRate(learningRate).Apply(wDec, grads[wDec])
        gorgonia.WithLearnRate(learningRate).Apply(bDec, grads[bDec])

        // Reset the VM for the next iteration
        vm.Reset()
    }

    fmt.Println("Training completed!")
}
