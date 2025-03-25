import (
    "github.com/galeone/tfgo"
    tf "github.com/tensorflow/tensorflow/tensorflow/go"
)

func main() {
    // Load pre-trained model
    model := tfgo.NewModel("path/to/model")
    
    // Example input tensor
    input := tfgo.NewTensor([]float32{1.0, 2.0, 3.0})
    
    // Run the model
    output := model.Exec([]tf.Output{model.Op("output", 0)}, map[tf.Output]*tf.Tensor{
        model.Op("input", 0): input,
    })
    
    // Print the output
    fmt.Println(output[0].Value())
}
